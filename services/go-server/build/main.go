package build

import (
	"context"

	"dagger.io/dagger"
	"github.com/kpenfound/hello-monorepo/daggerutils"

	goping "github.com/kpenfound/hello-monorepo/tools/go-ping/build"
	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

func Build(ctx context.Context, client *dagger.Client, os, arch string) (*dagger.Directory, error) {
	directory, err := client.Host().Workdir().Read().ID(ctx)
	if err != nil {
		return nil, err
	}

	build := daggerutils.GoBuild(daggerutils.GoBuildInput{
		Client:    client,
		Os:        os,
		Arch:      arch,
		Ctx:       ctx,
		Directory: directory,
		Workdir:   "services/go-server",
	})
	buildBinary, err := build.File("/go-server").ID(ctx)
	if err != nil {
		return nil, err
	}

	// go-server image requires go-ping and go-uname from tools
	ping, err := goping.Build(ctx, client, os, arch)
	if err != nil {
		return nil, err
	}
	pingBinary, err := ping.File("/go-ping").ID(ctx)
	if err != nil {
		return nil, err
	}

	uname, err := gouname.Build(ctx, client, os, arch)
	if err != nil {
		return nil, err
	}
	unameBinary, err := uname.File("/go-uname").ID(ctx)
	if err != nil {
		return nil, err
	}

	buildOutputs := client.Directory().
		WithCopiedFile("/go-server", buildBinary).
		WithCopiedFile("/go-ping", pingBinary).
		WithCopiedFile("/go-uname", unameBinary)

	return buildOutputs, nil
}
