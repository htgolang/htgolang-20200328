package main

import (
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func editReplicas() {
	configPath := "etc/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"

	name := "nginx"

	scale, err := clientset.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metaV1.GetOptions{})

	scale.Spec.Replicas = 2
	scale, err = clientset.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), name, scale, metaV1.UpdateOptions{})
	fmt.Println(err, scale)
}
