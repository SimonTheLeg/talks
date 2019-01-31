package main

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"path/filepath"
	"os"
	"fmt"
)

func main() {
	var config *rest.Config

	kubecfgpath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubecfgpath)

	clientset, err := kubernetes.NewForConfig(config)

	// Please handle errors more gracefully in your code ;)
	if err != nil {
		panic(err)
	}

	fmt.Println(clientset)

}
