package model

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var Client = GetKubernetesClient()

func GetKubernetesClient() *kubernetes.Clientset {
	kubeconfig := "/root/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Println(err)
	}
	Clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
	}
	return Clientset
}
