package main

import (
	"os"
	"path/filepath"

	"fmt"

	corev1 "k8s.io/api/core/v1"
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

		secret := &corev1.Secret{
			
			ObjectMeta: metav1.ObjectMeta {
				Name: "k8s-talk-createsecret",
				Namespace: "k8s-talk-createsecret",
				Labels: map[string]string {
					"app": "k8s-talk-createsecret",
				},
			},

			StringData: map[string]string {
				"username": "simon",
				"password": "secretpw123",
			},
		}

		res, err := clientset.CoreV1().Secrets(secret.ObjectMeta.Namespace).Create(secret)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Successfully created secret '%s' in namespace '%s'\n", res.ObjectMeta.Name, res.ObjectMeta.Namespace)
}