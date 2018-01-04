package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/docker/distribution"
	"github.com/docker/distribution/configuration"
	"github.com/docker/distribution/context"
	"github.com/docker/distribution/digest"
	"github.com/docker/distribution/registry/handlers"
	_ "github.com/docker/distribution/registry/storage/driver/inmemory"

	imageapiv1 "github.com/openshift/api/image/v1"
	registryclient "github.com/openshift/image-registry/pkg/dockerregistry/server/client"
	registrytest "github.com/openshift/image-registry/pkg/dockerregistry/testutil"
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
)

func createTestRegistryServer(t *testing.T, ctx context.Context) *httptest.Server {
	// pullthrough middleware will attempt to pull from this registry instance
	remoteRegistryApp := handlers.NewApp(ctx, &configuration.Configuration{
		Loglevel: "debug",
		Storage: configuration.Storage{
			"inmemory": configuration.Parameters{},
			"cache": configuration.Parameters{
				"blobdescriptor": "inmemory",
			},
			"delete": configuration.Parameters{
				"enabled": true,
			},
			"maintenance": configuration.Parameters{
				"uploadpurging": map[interface{}]interface{}{
					"enabled": false,
				},
			},
		},
	})
	return httptest.NewServer(remoteRegistryApp)
}

func TestPullthroughManifests(t *testing.T) {
	namespace := "fuser"
	repo := "zapp"
	repoName := fmt.Sprintf("%s/%s", namespace, repo)
	tag := "latest"

	backgroundCtx := context.Background()
	backgroundCtx = registrytest.WithTestLogger(backgroundCtx, t)

	backgroundCtx = withAppMiddleware(backgroundCtx, &appMiddlewareChain{
		&fakeAccessControllerMiddleware{t: t},
		&fakeBlobDescriptorServiceMiddleware{t: t, respectPassthrough: true},
	})

	remoteRegistryServer := createTestRegistryServer(t, backgroundCtx)
	defer remoteRegistryServer.Close()

	serverURL, err := url.Parse(remoteRegistryServer.URL)
	if err != nil {
		t.Fatalf("error parsing server url: %v", err)
	}

	ms1dgst, ms1canonical, _, ms1manifest, err := registrytest.CreateAndUploadTestManifest(
		backgroundCtx, registrytest.ManifestSchema1, 2, serverURL, nil, repoName, "schema1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, ms1payload, err := ms1manifest.Payload()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Logf("ms1dgst=%s, ms1manifest: %s", ms1dgst, ms1canonical)

	image, err := registrytest.NewImageForManifest(repoName, string(ms1payload), "", false)
	if err != nil {
		t.Fatal(err)
	}
	image.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", serverURL.Host, namespace, repo, image.Name)
	image.DockerImageManifest = ""

	fos, imageClient := registrytest.NewFakeOpenShiftWithClient(backgroundCtx)
	registrytest.AddImageStream(t, fos, namespace, repo, map[string]string{
		imageapi.InsecureRepositoryAnnotation: "true",
	})
	registrytest.AddImage(t, fos, image, namespace, repo, tag)

	for _, tc := range []struct {
		name                  string
		manifestDigest        digest.Digest
		localData             map[digest.Digest]distribution.Manifest
		expectedLocalCalls    map[string]int
		expectedError         bool
		expectedNotFoundError bool
	}{
		{
			name:           "manifest digest from local store",
			manifestDigest: ms1dgst,
			localData: map[digest.Digest]distribution.Manifest{
				ms1dgst: ms1manifest,
			},
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},
		{
			name:           "manifest served from remote repository",
			manifestDigest: digest.Digest(image.Name),
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},
		{
			name:                  "unknown manifest digest",
			manifestDigest:        unknownBlobDigest,
			expectedNotFoundError: true,
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},
	} {
		localManifestService := newTestManifestService(repoName, tc.localData)

		repo := newTestRepository(backgroundCtx, t, namespace, repo, testRepositoryOptions{
			client:            registryclient.NewFakeRegistryAPIClient(nil, imageClient),
			enablePullThrough: true,
		})

		ptms := &pullthroughManifestService{
			ManifestService: localManifestService,
			repo:            repo,
		}

		ctx := WithTestPassthroughToUpstream(backgroundCtx, false)
		manifestResult, err := ptms.Get(ctx, tc.manifestDigest)
		switch err.(type) {
		case distribution.ErrManifestUnknownRevision:
			if !tc.expectedNotFoundError {
				t.Fatalf("[%s] unexpected error: %#+v", tc.name, err)
			}
		case nil:
			if tc.expectedError || tc.expectedNotFoundError {
				t.Fatalf("[%s] unexpected successful response", tc.name)
			}
		default:
			if tc.expectedError {
				break
			}
			t.Fatalf("[%s] unexpected error: %#+v", tc.name, err)
		}

		if tc.localData != nil {
			if manifestResult != nil && manifestResult != tc.localData[tc.manifestDigest] {
				t.Fatalf("[%s] unexpected result returned", tc.name)
			}
		}

		for name, count := range localManifestService.calls {
			expectCount, exists := tc.expectedLocalCalls[name]
			if !exists {
				t.Errorf("[%s] expected no calls to method %s of local manifest service, got %d", tc.name, name, count)
			}
			if count != expectCount {
				t.Errorf("[%s] unexpected number of calls to method %s of local manifest service, got %d", tc.name, name, count)
			}
		}
	}
}

func TestPullthroughManifestInsecure(t *testing.T) {
	namespace := "fuser"
	repo := "zapp"
	repoName := fmt.Sprintf("%s/%s", namespace, repo)

	backgroundCtx := context.Background()
	backgroundCtx = registrytest.WithTestLogger(backgroundCtx, t)

	backgroundCtx = withAppMiddleware(backgroundCtx, &appMiddlewareChain{
		&fakeAccessControllerMiddleware{t: t},
		&fakeBlobDescriptorServiceMiddleware{t: t, respectPassthrough: true},
	})

	remoteRegistryServer := createTestRegistryServer(t, backgroundCtx)
	defer remoteRegistryServer.Close()

	serverURL, err := url.Parse(remoteRegistryServer.URL)
	if err != nil {
		t.Fatalf("error parsing server url: %v", err)
	}

	ms1dgst, ms1canonical, _, ms1manifest, err := registrytest.CreateAndUploadTestManifest(
		backgroundCtx, registrytest.ManifestSchema1, 2, serverURL, nil, repoName, "schema1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, ms1payload, err := ms1manifest.Payload()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Logf("ms1dgst=%s, ms1manifest: %s", ms1dgst, ms1canonical)
	ms2dgst, ms2canonical, ms2config, ms2manifest, err := registrytest.CreateAndUploadTestManifest(
		backgroundCtx, registrytest.ManifestSchema2, 2, serverURL, nil, repoName, "schema2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, ms2payload, err := ms2manifest.Payload()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Logf("ms2dgst=%s, ms2manifest: %s", ms2dgst, ms2canonical)

	ms1img, err := registrytest.NewImageForManifest(repoName, string(ms1payload), "", false)
	if err != nil {
		t.Fatal(err)
	}
	ms1img.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", serverURL.Host, namespace, repo, ms1img.Name)
	ms1img.DockerImageManifest = ""
	ms2img, err := registrytest.NewImageForManifest(repoName, string(ms2payload), ms2config, false)
	if err != nil {
		t.Fatal(err)
	}
	ms2img.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", serverURL.Host, namespace, repo, ms2img.Name)
	ms2img.DockerImageManifest = ""

	for _, tc := range []struct {
		name                string
		manifestDigest      digest.Digest
		localData           map[digest.Digest]distribution.Manifest
		fakeOpenShiftInit   func(fos *registrytest.FakeOpenShift)
		expectedManifest    distribution.Manifest
		expectedLocalCalls  map[string]int
		expectedErrorString string
	}{

		{
			name:           "fetch schema 1 with allowed insecure",
			manifestDigest: ms1dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				registrytest.AddImageStream(t, fos, namespace, repo, map[string]string{
					imageapi.InsecureRepositoryAnnotation: "true",
				})
				registrytest.AddImage(t, fos, ms1img, namespace, repo, "schema1")
			},
			expectedManifest: ms1manifest,
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},

		{
			name:           "fetch schema 2 with allowed insecure",
			manifestDigest: ms2dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				registrytest.AddImageStream(t, fos, namespace, repo, map[string]string{
					imageapi.InsecureRepositoryAnnotation: "true",
				})
				registrytest.AddImage(t, fos, ms2img, namespace, repo, "schema2")
			},
			expectedManifest: ms2manifest,
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},

		{
			name:           "explicit forbid insecure",
			manifestDigest: ms1dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				registrytest.AddImageStream(t, fos, namespace, repo, map[string]string{
					imageapi.InsecureRepositoryAnnotation: "false",
				})
				registrytest.AddImage(t, fos, ms1img, namespace, repo, "schema1")
			},
			expectedErrorString: "server gave HTTP response to HTTPS client",
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},

		{
			name:           "implicit forbid insecure",
			manifestDigest: ms1dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				registrytest.AddImageStream(t, fos, namespace, repo, nil)
				registrytest.AddImage(t, fos, ms1img, namespace, repo, "schema1")
			},
			expectedErrorString: "server gave HTTP response to HTTPS client",
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},

		{
			name:           "pullthrough from insecure tag",
			manifestDigest: ms1dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				image, err := registrytest.NewImageForManifest(repoName, string(ms1payload), "", false)
				if err != nil {
					t.Fatal(err)
				}
				image.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", serverURL.Host, namespace, repo, ms1dgst)
				image.DockerImageManifest = ""

				registrytest.AddUntaggedImage(t, fos, image)
				registrytest.AddImageStream(t, fos, namespace, repo, nil)
				registrytest.AddImageStreamTag(t, fos, ms1img, namespace, repo, &imageapiv1.TagReference{
					Name:         "schema1",
					ImportPolicy: imageapiv1.TagImportPolicy{Insecure: true},
				})
			},
			expectedManifest: ms1manifest,
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},

		{
			name:           "pull insecure if either image stream is insecure or the tag",
			manifestDigest: ms2dgst,
			fakeOpenShiftInit: func(fos *registrytest.FakeOpenShift) {
				image, err := registrytest.NewImageForManifest(repoName, string(ms2payload), ms2config, false)
				if err != nil {
					t.Fatal(err)
				}
				image.DockerImageReference = fmt.Sprintf("%s/%s/%s@%s", serverURL.Host, namespace, repo, image.Name)
				image.DockerImageManifest = ""

				registrytest.AddUntaggedImage(t, fos, image)
				registrytest.AddImageStream(t, fos, namespace, repo, map[string]string{
					imageapi.InsecureRepositoryAnnotation: "true",
				})
				registrytest.AddImageStreamTag(t, fos, image, namespace, repo, &imageapiv1.TagReference{
					Name: "schema2",
					// the value doesn't override is annotation because we cannot determine whether the
					// value is explicit or just the default
					ImportPolicy: imageapiv1.TagImportPolicy{Insecure: false},
				})
			},
			expectedManifest: ms2manifest,
			expectedLocalCalls: map[string]int{
				"Get": 1,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fos, imageClient := registrytest.NewFakeOpenShiftWithClient(backgroundCtx)

			tc.fakeOpenShiftInit(fos)

			localManifestService := newTestManifestService(repoName, tc.localData)

			ctx := WithTestPassthroughToUpstream(backgroundCtx, false)
			repo := newTestRepository(ctx, t, namespace, repo, testRepositoryOptions{
				client:            registryclient.NewFakeRegistryAPIClient(nil, imageClient),
				enablePullThrough: true,
			})

			ptms := &pullthroughManifestService{
				ManifestService: localManifestService,
				repo:            repo,
			}

			manifestResult, err := ptms.Get(ctx, tc.manifestDigest)
			switch err.(type) {
			case nil:
				if len(tc.expectedErrorString) > 0 {
					t.Fatalf("unexpected successful response")
				}
			default:
				if len(tc.expectedErrorString) > 0 {
					if !strings.Contains(err.Error(), tc.expectedErrorString) {
						t.Fatalf("expected error string %q, got instead: %s (%#+v)", tc.expectedErrorString, err.Error(), err)
					}
					break
				}
				t.Fatalf("unexpected error: %#+v (%s)", err, err.Error())
			}

			if tc.localData != nil {
				if manifestResult != nil && manifestResult != tc.localData[tc.manifestDigest] {
					t.Fatalf("unexpected result returned")
				}
			}

			registrytest.AssertManifestsEqual(t, tc.name, manifestResult, tc.expectedManifest)

			for name, count := range localManifestService.calls {
				expectCount, exists := tc.expectedLocalCalls[name]
				if !exists {
					t.Errorf("expected no calls to method %s of local manifest service, got %d", name, count)
				}
				if count != expectCount {
					t.Errorf("unexpected number of calls to method %s of local manifest service, got %d", name, count)
				}
			}
		})
	}
}

func TestPullthroughManifestDockerReference(t *testing.T) {
	namespace := "user"
	repo1 := "repo1"
	repo2 := "repo2"
	tag := "latest"

	type testServer struct {
		*httptest.Server
		name    string
		touched bool
	}

	startServer := func(name string) *testServer {
		s := &testServer{
			name: name,
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.touched = true
			http.Error(w, "dummy implementation", http.StatusInternalServerError)
		})

		s.Server = httptest.NewServer(handler)
		return s
	}

	dockerImageReference := func(s *testServer, rest string) string {
		serverURL, err := url.Parse(s.Server.URL)
		if err != nil {
			t.Fatal(err)
		}
		return fmt.Sprintf("%s/%s", serverURL.Host, rest)
	}

	server1 := startServer("server1")
	defer server1.Close()

	server2 := startServer("server2")
	defer server2.Close()

	img, err := registrytest.CreateRandomImage(namespace, "dummy")
	if err != nil {
		t.Fatal(err)
	}
	img.DockerImageManifest = ""

	image1 := *img
	image1.DockerImageReference = dockerImageReference(server1, "repo/name")

	image2 := *img
	image2.DockerImageReference = dockerImageReference(server2, "foo/bar")

	ctx := context.Background()
	ctx = registrytest.WithTestLogger(ctx, t)

	fos, imageClient := registrytest.NewFakeOpenShiftWithClient(ctx)
	registrytest.AddImageStream(t, fos, namespace, repo1, map[string]string{
		imageapi.InsecureRepositoryAnnotation: "true",
	})
	registrytest.AddImageStream(t, fos, namespace, repo2, map[string]string{
		imageapi.InsecureRepositoryAnnotation: "true",
	})
	registrytest.AddImage(t, fos, &image1, namespace, repo1, tag)
	registrytest.AddImage(t, fos, &image2, namespace, repo2, tag)

	for _, tc := range []struct {
		name             string
		repoName         string
		touchedServers   []*testServer
		untouchedServers []*testServer
	}{
		{
			name:             "server 1",
			repoName:         repo1,
			touchedServers:   []*testServer{server1},
			untouchedServers: []*testServer{server2},
		},
		{
			name:             "server 2",
			repoName:         repo2,
			touchedServers:   []*testServer{server2},
			untouchedServers: []*testServer{server1},
		},
	} {
		for _, s := range append(tc.touchedServers, tc.untouchedServers...) {
			s.touched = false
		}

		r := newTestRepository(ctx, t, namespace, tc.repoName, testRepositoryOptions{
			client:            registryclient.NewFakeRegistryAPIClient(nil, imageClient),
			enablePullThrough: true,
		})

		ptms := &pullthroughManifestService{
			ManifestService: newTestManifestService(tc.repoName, nil),
			repo:            r,
		}

		ptms.Get(ctx, digest.Digest(img.Name))

		for _, s := range tc.touchedServers {
			if !s.touched {
				t.Errorf("[%s] %s not touched", tc.name, s.name)
			}
		}

		for _, s := range tc.untouchedServers {
			if s.touched {
				t.Errorf("[%s] %s touched", tc.name, s.name)
			}
		}
	}
}

type testManifestService struct {
	name  string
	data  map[digest.Digest]distribution.Manifest
	calls map[string]int
}

var _ distribution.ManifestService = &testManifestService{}

func newTestManifestService(name string, data map[digest.Digest]distribution.Manifest) *testManifestService {
	b := make(map[digest.Digest]distribution.Manifest)
	for d, content := range data {
		b[d] = content
	}
	return &testManifestService{
		name:  name,
		data:  b,
		calls: make(map[string]int),
	}
}

func (t *testManifestService) Exists(ctx context.Context, dgst digest.Digest) (bool, error) {
	t.calls["Exists"]++
	_, exists := t.data[dgst]
	return exists, nil
}

func (t *testManifestService) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	t.calls["Get"]++
	content, exists := t.data[dgst]
	if !exists {
		return nil, distribution.ErrManifestUnknownRevision{
			Name:     t.name,
			Revision: dgst,
		}
	}
	return content, nil
}

func (t *testManifestService) Put(ctx context.Context, manifest distribution.Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	t.calls["Put"]++
	_, payload, err := manifest.Payload()
	if err != nil {
		return "", err
	}
	dgst := digest.FromBytes(payload)
	t.data[dgst] = manifest
	return dgst, nil
}

func (t *testManifestService) Delete(ctx context.Context, dgst digest.Digest) error {
	t.calls["Delete"]++
	return fmt.Errorf("method not implemented")
}

const etcdDigest = "sha256:958608f8ecc1dc62c93b6c610f3a834dae4220c9642e6e8b4e0f2b3ad7cbd238"
