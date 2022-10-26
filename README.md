# hello-monorepo

A simple monorepo with the following structure:
```
├── services
│   ├── go-server
│   ├── go-worker
│   └── py-server
└── tools
    ├── go-ping
    └── go-uname
```

## Dagger

CLI tasks can be performed with [Mage](https://magefile.org/)

Try tasks like:
- `mage GoServer`
- `mage GoUname`

These will build the subproject and put the build output in your local directory.

Each project has it's own `build` submodule which defines it's build instructions for Dagger to use. The Mage tasks will handle connecting to the dagger engine and executing a build of a specified project.

In this example, the `services/go-server` project depends on `tools/go-uname` and `tools/go-ping`, so you can see how it imports the build instructions for those projects.
