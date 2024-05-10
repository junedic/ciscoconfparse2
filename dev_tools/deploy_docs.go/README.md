
# Overview

This is just a quick Go tool to upload my ccp markdown docs to my webserver with private key auth.

It uses:
- [github melbahja/goph](https://github.com/melbahja/goph) for ssh automation
- [github gleich/logoru](https://github.com/gleich/logoru) for terminal logging

To build this tool, `make build` and it will get all the required dependencies and compile them.

Run with `make build`

By default, you should see output like this...

```shell
$ ./deploy_docs
2023-06-23T09:13:24.844-0500    INFO    deploy_docs/deploy_docs.go:16   Starting ciscoconfparse2 new documentation upload
2023-06-23T09:13:24.844-0500    INFO    deploy_docs/deploy_docs.go:18       Initialize ssh key-auth
2023-06-23T09:13:24.844-0500    INFO    deploy_docs/deploy_docs.go:25       ssh into remote webhost
2023-06-23T09:13:25.477-0500    INFO    deploy_docs/deploy_docs.go:31       Remove old documentation files
2023-06-23T09:13:25.937-0500    INFO    deploy_docs/deploy_docs.go:37       Start copying doc tarball to remote ssh server
2023-06-23T09:13:28.659-0500    INFO    deploy_docs/deploy_docs.go:42       Finished copying tarball to remote ssh server
2023-06-23T09:13:28.659-0500    INFO    deploy_docs/deploy_docs.go:44       extract CiscoConfParse documentation tarball
$
```
