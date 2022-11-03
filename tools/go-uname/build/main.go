package build

import (
	"context"

	"dagger.io/dagger"
	"github.com/kpenfound/hello-monorepo/daggerutils"
)

func Build(ctx context.Context, client *dagger.Client, os, arch string) *dagger.Directory {
	directory := client.Host().Workdir()
	return daggerutils.GoBuild(daggerutils.GoBuildInput{
		Client:    client,
		Os:        os,
		Arch:      arch,
		Ctx:       ctx,
		Directory: directory,
		Workdir:   "tools/go-uname",
	})
}
