package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1Types "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// API client for managing secrets
var secretsClient coreV1Types.SecretInterface

func initClient() {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secretsClient = clientset.CoreV1().Secrets("default")
}

func main() {
	initClient()
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "my-secret",
		},
		Data: map[string][]byte{
			"secret-data": []byte("secret-value-1"),
		},
	}
	_, err := secretsClient.Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	// read the secret
	_, err = secretsClient.Get(context.TODO(), metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
}

