package main

import (
	"context"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	configPath := "etc/kube.conf"
	config, _ := clientcmd.BuildConfigFromFlags("", configPath)
	clientset, _ := kubernetes.NewForConfig(config)

	namespace := "default"

	name, serviceName := "nginx", "nginx-service"

	clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})

	clientset.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metaV1.DeleteOptions{})

	// clientset.ExtensionsV1beta1().Ingresses()
}
