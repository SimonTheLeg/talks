package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	var config *rest.Config

	config, err := rest.InClusterConfig()

	// Please handle errors more gracefully in your code ;)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		panic(err)
	}

	log.Printf("it works: Clientset is %v", clientset)
	log.Printf("Going into endless loop now, so container doesn't die :)")
	for {
	}
}
