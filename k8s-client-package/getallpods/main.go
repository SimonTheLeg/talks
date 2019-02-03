package getallpods

import (
	"os"
	"path/filepath"

	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// Fetch a list of all Pods in all namespaces
	allPods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range allPods.Items {
		fmt.Println(pod.ObjectMeta.Name)
	}
}
