package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"time"
)
import "sigs.k8s.io/controller-runtime/pkg/reconcile"

// ./k8s-controller
// result ->
// >>> "hello, world"

type HelloReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// The reconcile method will contain the logic to see whether a playtechnique/labelme label exists and add
// a playtechnique/labelled label.
func (r *HelloReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	fmt.Println("Hello, world!")

	return reconcile.Result{}, nil
}

func main() {
	cfg := config.GetConfigOrDie()
	//Is opts the right type to pass in to the Pods filter?
	podOpts := v1.ListOptions{}

	clientset, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		panic(err)
	}

	//Pods.List() takes a context.TODO() and a list of options that can contain a metav1.ListOptions
	//cf. /Users/gwyn/go/pkg/mod/k8s.io/apimachinery@v0.28.3/pkg/apis/meta/v1/types.go:322
	// This can contain a label to select pods!
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), podOpts)

	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(time.Now())
		fmt.Println(pods)
		time.Sleep(3000000000)
	}
}
