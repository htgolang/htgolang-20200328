package services

import (
	"cmdb/forms"
	"context"

	"github.com/astaxie/beego"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sService struct {
}

func (s *k8sService) GetClient() (*kubernetes.Clientset, error) {
	path := beego.AppConfig.DefaultString("k8s::path", "conf/k8s/kube.conf")
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

type deploymentService struct {
	k8sService
}

func (s *deploymentService) Query() []appsV1.Deployment {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return []appsV1.Deployment{}
	}
	deploymentList, err := clientset.AppsV1().Deployments("").List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []appsV1.Deployment{}
	}
	return deploymentList.Items
}

func (s *deploymentService) Delete(name string, namespace string) {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return
	}
	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (s *deploymentService) Create(form *forms.DeploymentCreateForm) {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return
	}
	var replicas int32 = 1
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name:   form.Name,
			Labels: form.GetLabels(),
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: form.GetSelectors(),
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name:   form.Name,
					Labels: form.GetSelectors(),
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  form.GetImageName(),
							Image: form.Image,
							Ports: form.GetPorts(),
						},
					},
				},
			},
		},
	}
	clientset.AppsV1().Deployments(form.Namespace).Create(context.TODO(), deployment, metaV1.CreateOptions{})
}

type namespaceService struct {
	k8sService
}

func (s *namespaceService) Query() []coreV1.Namespace {
	clientset, err := s.GetClient()
	if err != nil {
		beego.Error(err)
		return []coreV1.Namespace{}
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		beego.Error(err)
		return []coreV1.Namespace{}
	}
	return namespaceList.Items
}

var DeploymentService = new(deploymentService)
var NamespaceService = new(namespaceService)
