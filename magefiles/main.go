//go:build mage

package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"dagger.io/dagger"
	goserver "github.com/kpenfound/hello-monorepo/services/go-server/build"
	pyserver "github.com/kpenfound/hello-monorepo/services/py-server/build"
	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

func GoUname(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	os, arch := getOsArch()

	uname := gouname.Build(ctx, client, os, arch)

	_, err = uname.Export(ctx, ".")
	if err != nil {
		panic(err)
	}
}

func GoServer(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	os, arch := getOsArch()

	server := goserver.Build(ctx, client, os, arch)

	_, err = server.Export(ctx, ".")
	if err != nil {
		panic(err)
	}
}

func PyServerPush(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	build := pyserver.Build(ctx, client, "linux", "amd64")
	image := pyserver.Image(ctx, client, build)
	addr, err := pyserver.Push(ctx, image)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pushed py-server to %s\n", addr)
}

func getOsArch() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
