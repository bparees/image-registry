package testframework

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	authorizationapiv1 "github.com/openshift/api/authorization/v1"
	projectapiv1 "github.com/openshift/api/project/v1"
	authorizationv1 "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	projectv1 "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
)

func CreateProject(clientConfig *rest.Config, namespace string, adminUser string) (*projectapiv1.Project, error) {
	projectClient := projectv1.NewForConfigOrDie(clientConfig)
	project, err := projectClient.ProjectRequests().Create(&projectapiv1.ProjectRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	})
	if err != nil {
		return project, err
	}

	authorizationClient := authorizationv1.NewForConfigOrDie(clientConfig)
	_, err = authorizationClient.RoleBindings(namespace).Update(&authorizationapiv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "admin",
		},
		UserNames: []string{adminUser},
		RoleRef: corev1.ObjectReference{
			Name: "admin",
		},
	})
	return project, err
}
