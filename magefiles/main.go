//go:build mage

package main

import (
	"context"
	"os"
	"runtime"

	"dagger.io/dagger"
	goserver "github.com/kpenfound/hello-monorepo/services/go-server/build"
	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

func GoUname(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	os, arch := getOsArch()

	uname, err := gouname.Build(ctx, client, os, arch)
	if err != nil {
		panic(err)
	}
	unameID, err := uname.ID(ctx)
	if err != nil {
		panic(err)
	}

	workdir := client.Host().Workdir()

	_, err = workdir.Write(ctx, unameID, dagger.HostDirectoryWriteOpts{})
}

func GoServer(ctx context.Context) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	os, arch := getOsArch()

	server, err := goserver.Build(ctx, client, os, arch)
	if err != nil {
		panic(err)
	}
	serverID, err := server.ID(ctx)
	if err != nil {
		panic(err)
	}

	workdir := client.Host().Workdir()

	_, err = workdir.Write(ctx, serverID, dagger.HostDirectoryWriteOpts{})
}

func getOsArch() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
