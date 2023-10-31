package main

import (
	"context"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"time"
)
import "sigs.k8s.io/controller-runtime"
import "sigs.k8s.io/controller-runtime/pkg/reconcile"

// ./k8s-controller
// result ->
// >>> "hello, world"

type HelloReconciler struct {
}

func (r *HelloReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	fmt.Println("Hello, world!")

	return reconcile.Result{}, nil
}

//type hello struct{}

func main() {
	cfg := config.GetConfigOrDie()
	opts := manager.Options{}

	ctrlManager, err := controllerruntime.NewManager(cfg, opts)
	if err != nil {
		panic(err)
	}

	//builder := controllerruntime.NewControllerManagedBy(ctrlManager)
	//err = builder.For(&hello{}).Complete(&HelloReconciler{})

	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(time.Now())
		fmt.Println(ctrlManager)
		time.Sleep(3000000000)
	}
}
