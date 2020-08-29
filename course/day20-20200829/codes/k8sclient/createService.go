package main

import (
	"context"
	"fmt"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func createService() {
	configPath := "etc/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "nginx-service",
			Labels: map[string]string{
				"env": "dev",
			},
		},
		Spec: coreV1.ServiceSpec{
			Type: coreV1.ServiceTypeNodePort,
			Selector: map[string]string{
				"env": "dev",
				"app": "nginx",
			},
			Ports: []coreV1.ServicePort{
				{
					Name:     "http",
					Port:     80,
					Protocol: coreV1.ProtocolTCP,
				},
			},
		},
	}
	service, err := clientset.CoreV1().Services(namespace).Create(context.TODO(), service, metaV1.CreateOptions{})

	fmt.Println(err, service)
}
