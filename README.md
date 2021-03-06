# ![grapi](./grapi.png)
[![Build Status](https://travis-ci.org/izumin5210/grapi.svg?branch=master)](https://travis-ci.org/izumin5210/grapi)
[![GoDoc](https://godoc.org/github.com/izumin5210/grapi/pkg/grapiserver?status.svg)](https://godoc.org/github.com/izumin5210/grapi/pkg/grapiserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/izumin5210/grapi)](https://goreportcard.com/report/github.com/izumin5210/grapi)
[![Go project version](https://badge.fury.io/go/github.com%2Fizumin5210%2Fgrapi.svg)](https://badge.fury.io/go/github.com%2Fizumin5210%2Fgrapi)
[![license](https://img.shields.io/github/license/izumin5210/grapi.svg)](./LICENSE)

:open_mouth: A surprisingly easy API server in gRPC and Go

## Features
- You can develop and deploy API servers blazing fastly :zap:
- Easily code generator
	- application  (inspired by `rails new` and `create-react-app`)
	- gRPC services and their implementations (inspired by `rails g (scaffold_)controller`)
- User-friendly `protoc` wrapper (inspired by [protoeasy](https://github.com/peter-edge/protoeasy-go))
- Provides gRPC and HTTP JSON API  with single implementation, by using [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- Generates codes followed [google's API design guideline](https://cloud.google.com/apis/design/)

[![asciicast](https://asciinema.org/a/172436.png)](https://asciinema.org/a/172436)

## Getting Started
### Installation

```
$ brew install protobuf
$ brew install izumin5210/tools/grapi
```

### Creates a new application
```
$ grapi init awesome-app
```

### Creates a new service
```
$ grapi g service books
```

Or if you need full [standard methods](https://cloud.google.com/apis/design/standard_methods), you can get them with following command:

```
$ grapi g scaffold-service books
```

And you should register generated services to the `grapiserver.Engine` instance:

```diff
 // app/run.go
 
 // Run starts the grapiserver.
 func Run() error {
 	s := grapiserver.New(
 		grapiserver.WithDefaultLogger(),
 		grapiserver.WithServers(
+			server.NewBookServiceServer(),
-		// TODO
 		),
 	)
 	return s.Serve()
 }
```

If you updated service definition, you can re-generate `.pb.go` and `.pb.gw.go` with the following command:

```
$ grapi protoc
```

### Starts server

```
$ grapi server
```

### User-defined commands

```
$ grapi g command import-books
$ vim cmd/import-books/run.go  # implements the command
$ grapi import-books  # run the command
```

### Build commands (including server)

```
$ grapi build
```
