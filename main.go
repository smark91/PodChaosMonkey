package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	deleteNamespace, deleteDelayInt, err := getEnvVars()
	if err != nil {
		panic(err.Error())
	}

	client, err := getK8sClientset()
	if err != nil {
		panic(err.Error())
	}

	for {
		// Sleep
		fmt.Printf("\nSleeping %d seconds...\n", deleteDelayInt)
		time.Sleep(time.Duration(deleteDelayInt) * time.Second)

		// Get pods list
		pods, err := listK8sPodsInNamespace(client, deleteNamespace)
		if err != nil {
			panic(err.Error())
		} else if len(pods.Items) < 1 {
			fmt.Printf("\nThere are no pods in the \"%s\" namespace\n", deleteNamespace)
			continue
		}
		fmt.Printf("\nThere are %d pods in the \"%s\" namespace:\n", len(pods.Items), deleteNamespace)
		for _, pod := range pods.Items {
			fmt.Printf("\t- %s\n", pod.Name)
		}

		// Select random pod for deletion
		nextDeletePod := getRandomK8sPodFromList(client, deleteNamespace, pods)
		fmt.Printf("\nSelected randomly %s pod for deletion...\n", nextDeletePod.Name)

		// Delete selected pod
		err = deleteK8sPodInNamespace(client, deleteNamespace, nextDeletePod.Name)
		if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("%s pod deleted!\n", nextDeletePod.Name)
		}
	}
}

func getEnvVars() (string, int, error) {
	deleteDelay := os.Getenv("DELETE_DELAY")
	if len(deleteDelay) == 0 {
		return "", 0, errors.New("DELETE_DELAY missing!")
	}
	deleteNamespace := os.Getenv("DELETE_NAMESPACE")
	if len(deleteNamespace) == 0 {
		return "", 0, errors.New("DELETE_DELAY missing!")
	}

	fmt.Println("DELETE_DELAY:", deleteDelay)
	fmt.Println("DELETE_NAMESPACE:", deleteNamespace)

	deleteDelayInt, err := strconv.Atoi(deleteDelay)
	if err != nil {
		panic(err.Error())
	}
	return deleteNamespace, deleteDelayInt, err
}
