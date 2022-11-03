package build

import (
	"context"

	"dagger.io/dagger"

	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

const (
	serverImage = "kylepenfound/pyserver:latest"
)

func Build(ctx context.Context, client *dagger.Client, os, arch string) *dagger.Directory {
	directory := client.Host().Workdir()
	serverPy := directory.File("services/py-server/server.py")

	// py-server image requires go-uname from tools
	uname := gouname.Build(ctx, client, os, arch)
	unameBinary := uname.File("/go-uname")

	buildOutputs := client.Directory().
		WithFile("/server.py", serverPy).
		WithFile("/go-uname", unameBinary)

	return buildOutputs
}

func Image(ctx context.Context, client *dagger.Client, build *dagger.Directory) *dagger.Container {
	// Load python alpine image
	py := client.Container().From("python:3.11-alpine")
	// Mount build to container
	py = py.WithMountedDirectory("/build", build)

	// Copy build artifacts off of mounted directory
	py = py.Exec(dagger.ContainerExecOpts{
		Args: []string{"cp", "/build/go-uname", "/usr/bin/go-uname"},
	})
	py = py.Exec(dagger.ContainerExecOpts{
		Args: []string{"cp", "/build/server.py", "/server.py"},
	})

	// Setup environment
	py = py.WithEnvVariable("PYTHONUNBUFFERED", "1")
	// Set container entrypiont
	py = py.WithEntrypoint([]string{"python", "/server.py"})

	return py
}

func Push(ctx context.Context, img *dagger.Container) (string, error) {
	return img.Publish(ctx, serverImage)
}
