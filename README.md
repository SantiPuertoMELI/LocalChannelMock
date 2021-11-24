# Local Channel Mock

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This is application is intended to mock the ISO channel during the development of a new integration.

## Exposed endpoints

By default, the application is listening requests at port 8080 but it can be changed.

### Purchase operation

In the path `/purchase` you can send a **GET** request that will return a channel response with data from a successful
purchase.

## First steps

### Go Runtime Version

Specify the Go runtime version tag you desire in your `Dockerfile`. If in doubt, it's completely safe to always use the
latest one given the [Go 1 compatibility guarantees](https://golang.org/doc/go1compat).

```docker
FROM hub.furycloud.io/mercadolibre/go:1.15-mini
```

> You can find all available image tags for your Dockerfile
> [here](https://github.com/mercadolibre/fury_go-mini#supported-tags).

### Release Process Support

Set the application module name in the first line of the `go.mod` file with your application's GitHub repository URL.
Avoiding to do so will result in the CI process, and go command failing.

The file should start with:

```
module github.com/mercadolibre/<fury_fury-app-name>

[...]
```

For more information about all the features provided by Release Process refer to the docker image documentation
[here](https://github.com/mercadolibre/fury_go-mini#testing-support).

### Dependency management

This image has native support for Go Modules and requires the use of it as the dependency management tool.

For more information refer to the
[`go-mini` docs](https://github.com/mercadolibre/fury_go-mini#dependency-management-support).

## Questions

* [Fury Issue Tracker](https://github.com/mercadolibre/fury/issues)