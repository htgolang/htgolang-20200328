package main

import (
	"context"
	"fmt"
	"log"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func list() {
	configPath := "etc/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	nodeList, _ := clientset.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})

	fmt.Println("node:")
	for _, node := range nodeList.Items {
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\n%s\n",
			node.Name,
			node.Status.Phase,
			node.Status.Addresses,
			node.Status.NodeInfo.OSImage,
			node.Status.NodeInfo.KubeletVersion,
			node.Status.NodeInfo.OperatingSystem,
			node.Status.NodeInfo.Architecture,
			node.CreationTimestamp,
		)
	}

	namespaceList, _ := clientset.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})

	fmt.Println("namespace:")
	for _, namespace := range namespaceList.Items {
		fmt.Println(namespace.Name, namespace.CreationTimestamp, namespace.Status.Phase)
	}

	servicesList, _ := clientset.CoreV1().Services("").List(context.TODO(), metaV1.ListOptions{})
	fmt.Println("service:")

	for _, service := range servicesList.Items {
		fmt.Println(service.Namespace, service.Name, service.Spec.Type, service.CreationTimestamp, service.Spec.ClusterIP, service.Spec.Ports)
	}

	deploymentList, _ := clientset.AppsV1().Deployments("").List(context.TODO(), metaV1.ListOptions{})

	fmt.Println("deployment:")
	for _, deployment := range deploymentList.Items {
		fmt.Println(deployment.Namespace, deployment.Name, deployment.Labels, deployment.CreationTimestamp, deployment.Spec.Selector.MatchLabels, deployment.Status.Replicas, deployment.Status.AvailableReplicas)
	}
}
