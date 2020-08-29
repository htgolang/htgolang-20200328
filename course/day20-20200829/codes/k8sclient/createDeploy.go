package main

import (
	"context"
	"fmt"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createDeploy() {
	configPath := "etc/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"
	var replicas int32 = 3

	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				"env": "dev",
			},
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
					"env": "dev",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
						"env": "dev",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.16.1",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), deployment, metaV1.CreateOptions{})

	fmt.Println(err, deployment)
}
