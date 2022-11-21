package main

import (
	"context"
	"math/rand"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type k8s struct {
	Clientset kubernetes.Interface
}

func getK8sClientset() (*k8s, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client := k8s{}
	client.Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &client, err
}

func listK8sPodsInNamespace(client *k8s, namespace string) (*v1.PodList, error) {
	pods, err := client.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	return pods, err
}

func getRandomK8sPodFromList(client *k8s, namespace string, pods *v1.PodList) v1.Pod {
	rand.Seed(time.Now().Unix())
	selectedPod := pods.Items[rand.Intn(len(pods.Items))]
	return selectedPod
}

func deleteK8sPodInNamespace(client *k8s, namespace string, podName string) error {
	err := client.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	return err
}
