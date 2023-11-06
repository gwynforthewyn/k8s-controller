package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"time"
)
import "sigs.k8s.io/controller-runtime"
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
	opts := manager.Options{}

	ctrlManager, err := controllerruntime.NewManager(cfg, opts)
	if err != nil {
		panic(err)
	}

	reconciler := HelloReconciler{
		Client: ctrlManager.GetClient(),
		Scheme: ctrlManager.GetScheme(),
	}

	//What does the controller do here?
	ctrl, err := controller.New("gwyn-controller", ctrlManager, controller.Options{
		Reconciler: &reconciler,
	})

	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(time.Now())
		fmt.Println(ctrlManager)
		time.Sleep(3000000000)
	}
}
