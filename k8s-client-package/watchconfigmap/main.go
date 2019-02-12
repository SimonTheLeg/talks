package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Connect to the k8s cluster
	var config *rest.Config

	kubecfgpath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubecfgpath)

	clientset, err := kubernetes.NewForConfig(config)

	// Please handle errors more gracefully in your code ;)
	if err != nil {
		panic(err)
	}
	cmns, cmname := "simons-ns", "simons-configmap"
	listOptions := metav1.ListOptions{
		FieldSelector: fmt.Sprintf("metadata.name==%s", cmname),
	}

	log.Printf("Establishing watch on configmap '%s/%s'", cmns, cmname)
	watcher, err := clientset.CoreV1().ConfigMaps(cmns).Watch(listOptions)

	if err != nil {
		fmt.Printf("Could not establish watch on '%s/%s': %v", cmns, cmname, err)
	}

	for event := range watcher.ResultChan() {
		switch event.Type {
			
		case watch.Added:
			log.Printf("configmap '%s/%s' was created", cmns, cmname)
		case watch.Modified:
			log.Printf("configmap '%s/%s' was modified", cmns, cmname)
		case watch.Deleted:
			log.Printf("configmap '%s/%s' was deleted :(", cmns, cmname)
		}
	}
}
