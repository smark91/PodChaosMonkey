package main

import (
	"strings"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func newTestSimpleK8s(object ...runtime.Object) *k8s {
	client := k8s{}
	client.Clientset = fake.NewSimpleClientset(object...)
	return &client
}

func TestListPods(t *testing.T) {

	obj := []runtime.Object{
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod1", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod2", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod3", Namespace: "test-namespace"}},
	}

	fakeClient := newTestSimpleK8s(obj...)
	pods, err := listK8sPodsInNamespace(fakeClient, "test-namespace")
	if err != nil {
		t.Errorf(err.Error())
	}

	if len(pods.Items) != 3 {
		t.Errorf("Expected 3 pods, find %d", len(pods.Items))
	}
}

func TestPickRandomPod(t *testing.T) {

	obj := []runtime.Object{
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod1", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod2", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod3", Namespace: "test-namespace"}},
	}

	fakeClient := newTestSimpleK8s(obj...)
	pods, err := listK8sPodsInNamespace(fakeClient, "test-namespace")
	if err != nil {
		t.Errorf(err.Error())
	}

	pod := getRandomK8sPodFromList(fakeClient, "test-namespace", pods)
	if !(strings.HasPrefix(pod.Name, "test-pod")) {
		t.Errorf("Expected pod name starting with test-pod, got %s", pod.Name)
	}
}

func TestDeletePod(t *testing.T) {
	// creates the fake clientset
	obj := []runtime.Object{
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod1", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod2", Namespace: "test-namespace"}},
		&v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "test-pod3", Namespace: "test-namespace"}},
	}

	fakeClient := newTestSimpleK8s(obj...)
	err := deleteK8sPodInNamespace(fakeClient, "test-namespace", "test-pod1")
	if err != nil {
		t.Errorf(err.Error())
	}
}
