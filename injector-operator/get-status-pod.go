package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/etc/kubernetes/admin.conf", "absolute path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	// Replace "capi-quickstart" and "default" with your namespace if needed
	clusterName := "my-nginx-54b466b7b4-28qbw"
	namespace := "default"

	// Get information about Pod
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), clusterName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting pod: %v\n", err)
		os.Exit(1)
	}

	// Print information about Pod
	fmt.Printf("Name: %s\n", pod.Name)
	fmt.Printf("Phase: %s\n", pod.Status.Phase)
	fmt.Printf("Age: %v\n", pod.ObjectMeta.CreationTimestamp.Time)
	fmt.Printf("Version: %s\n", pod.Status.ContainerStatuses[0].Image)
}
