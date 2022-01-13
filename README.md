# gqlgen-error

gqlgen error MVP

## How to Start

1. install [task command](https://taskfile.dev) (Optional if you want by yourself.)
2. `task down` to install `gqlgen v0.10.2`.
3. `task` to generate with gqlgen.
4. `task up` to install `gqlgen@master`
5. `task` to generate with gqlgen.

## The outputs

This code is works fine with `v0.10.2`, but can not used by `v0.11.0+`:

```
$ task   
task: [default] go generate ./...
task: [default] go run github.com/99designs/gqlgen --verbose
merging type systems failed: unable to bind to interface: github.com/ipfans/gqlgen-error/business/entity.City does not satisfy the interface gitlab.com/hookactions/gqlgen-relay/relay.Node
exit status 1
```
