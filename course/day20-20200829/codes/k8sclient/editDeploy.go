package main

import (
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func editDeploy() {
	configPath := "etc/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"

	var replicas int32 = 1

	name := "nginx"

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})

	deployment.Spec.Replicas = &replicas
	deployment.Spec.Template.Spec.Containers[0].Image = "nginx:1.14"

	deployment, err = clientset.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	fmt.Println(err, deployment)
}
