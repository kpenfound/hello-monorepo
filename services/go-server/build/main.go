package build

import (
	"context"

	"dagger.io/dagger"
	"github.com/kpenfound/hello-monorepo/daggerutils"

	goping "github.com/kpenfound/hello-monorepo/tools/go-ping/build"
	gouname "github.com/kpenfound/hello-monorepo/tools/go-uname/build"
)

func Build(ctx context.Context, client *dagger.Client, os, arch string) *dagger.Directory {
	directory := client.Host().Workdir()

	build := daggerutils.GoBuild(daggerutils.GoBuildInput{
		Client:    client,
		Os:        os,
		Arch:      arch,
		Ctx:       ctx,
		Directory: directory,
		Workdir:   "services/go-server",
	})

	// go-server image requires go-ping and go-uname from tools
	ping := goping.Build(ctx, client, os, arch)
	uname := gouname.Build(ctx, client, os, arch)

	buildOutputs := client.Directory().
		WithFile("/go-server", build.File("/go-server")).
		WithFile("/go-ping", ping.File("/go-ping")).
		WithFile("/go-uname", uname.File("/go-uname"))

	return buildOutputs
}
