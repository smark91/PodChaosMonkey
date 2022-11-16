package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {

	deleteDelay := os.Getenv("DELETE_DELAY")
	if len(deleteDelay) == 0 {
		panic("DELETE_DELAY missing!")
	}
	deleteNamespace := os.Getenv("DELETE_NAMESPACE")
	if len(deleteNamespace) == 0 {
		panic("DELETE_NAMESPACE missing!")
	}

	fmt.Println("DELETE_DELAY:", deleteDelay)
	fmt.Println("DELETE_NAMESPACE:", deleteNamespace)

	deleteDelayInt, err := strconv.Atoi(deleteDelay)
	if err != nil {
		panic(err.Error())
	}

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(time.Now().Unix())

	for {
		pods, err := clientset.CoreV1().Pods(deleteNamespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\nThere are %d pods in the \"%s\" namespace:\n", len(pods.Items), deleteNamespace)
		for _, pod := range pods.Items {
			fmt.Printf("- %s\n", pod.Name)
		}

		nextDeletePod := pods.Items[rand.Intn(len(pods.Items))]
		fmt.Printf("\nSelected randomly %s pod for deletion...\n", nextDeletePod.Name)
		err = clientset.CoreV1().Pods(deleteNamespace).Delete(context.TODO(), nextDeletePod.Name, metav1.DeleteOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%s pod deleted!\n", nextDeletePod.Name)

		fmt.Printf("\nSleeping %d seconds...\n", deleteDelayInt)
		time.Sleep(time.Duration(deleteDelayInt) * time.Second)
	}
}
