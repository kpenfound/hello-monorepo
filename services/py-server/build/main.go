package build

import (
	"context"

	"dagger.io/dagger"

	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

const (
	serverImage = "kylepenfound/pyserver:latest"
)

func Build(ctx context.Context, client *dagger.Client, os, arch string) (*dagger.Directory, error) {
	directory := client.Host().Workdir().Read()

	serverPy, err := directory.File("services/py-server/server.py").ID(ctx)
	if err != nil {
		return nil, err
	}

	// py-server image requires go-uname from tools
	uname, err := gouname.Build(ctx, client, os, arch)
	if err != nil {
		return nil, err
	}
	unameBinary, err := uname.File("/go-uname").ID(ctx)
	if err != nil {
		return nil, err
	}

	buildOutputs := client.Directory().
		WithCopiedFile("/server.py", serverPy).
		WithCopiedFile("/go-uname", unameBinary)

	return buildOutputs, nil
}

func Image(ctx context.Context, client *dagger.Client, build *dagger.Directory) (*dagger.Container, error) {
	// Get build DirectoryID
	buildId, err := build.ID(ctx)
	if err != nil {
		return nil, err
	}

	// Load python alpine image
	py := client.Container().From("python:3.11-alpine")
	// Mount build to container
	py = py.WithMountedDirectory("/build", buildId)

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

	return py, nil
}

func Push(ctx context.Context, img *dagger.Container) (string, error) {
	return img.Publish(ctx, serverImage)
}
