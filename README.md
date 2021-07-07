# Analysis of Software Engineering Automation Tools for Go

## Introduction

Go is a modern programming language designed by Ken Thompson, Rob Pike, and Robert Greissmer at Google, releasing its first version in 2012 [1]. Since its release as an open-source project, it has been one of the fastest-growing programming languages for building scalable and distributed systems thanks to its concurrency model and its simplicity. In this study we will analyze different automation tools for software engineering tasks for the Go programming language and evaluate the cognitive load reduction on the developers while writing good quality software.  

In this repository you will find code snippets of the tests carried out on industrial projects. On each folder you will find a "Makefile" with the commands to run each tool.

## Tools

### Code Generation Tools

- [Mockgen](https://github.com/golang/mock)  
The Mockgen tool generates mocks  from analyzing defined interfaces on a Go source code file. These mocks are then printed to the standard output or a file. These mocks can be used in unit tests and the desired calls to them can be defined.  

- [Mockery](https://github.com/vektra/mockery)  
This tool has the same purpose as the previous tool. The main difference is that it creates mocks for the Testify testing framework.  

- [Scaffold](https://pkg.go.dev/github.com/metakeule/scaffold)  
The Scaffold tool is a CLI implemented in Go that leverages the generate package from the standard library to scaffold new projects.

- [Gqlgen](https://github.com/99designs/gqlgen)  
The Gqlgen library leverages code generation strategies to create graphQL servers without worrying about the implementation of the transport layer details. The library scaffolds the project where the developer defines the grapQL schema, then executes the command and the server is generated with some method stubs where the developer needs to implement the functionality to have a fully functional server.  

- [gRPC](https://grpc.io/docs/languages/go/)  
gRPC is a high-performance remote procedure call framework using HTTP/2 and protocol buffers. The way this framework works with Go is by generating all the transport layer code by reading from the protocol’s buffer definition file. The produced code is generated into a single file which can be accessed from other packages.  

### Static Analysis Tools

- [Check](https://gitlab.com/opennota/check)  
The Check tool analyses code to detect inefficiently packed structs and tells the developers which ones could be declared more efficiently.  

In Go structs are packed into memory in the order their fields are declared, meaning that depending on the types of the fields some padding may be added to have the fields in the same order on the machine’s memory. This may lead to structs with a higher memory footprint than the one they could have if they are efficiently declared.  

- [Dupl](https://github.com/mibk/dupl)  
The Dupl tool analyses code to detect duplicated code. Duplicated code could lead to maintainability issues because this could mean that a change has to be done in various parts of the codebase.  

- [Errcheck](https://github.com/kisielk/errcheck)  
The Errcheck tools analyses code to detect not handled errors. In Go errors are returned as the last value of a function, this means that there will not be errors that stop the program if they are not checked, but they should be checked to have a correct working program.  

- [Gocyclo](https://github.com/fzipp/gocyclo)  
The Gocyclo tool calculates the cyclomatic complexity  of all the declared functions in the program.  

- [Gosec](https://github.com/securego/gosec)  
The Gosec tool analyses code to detect possible security vulnerabilities on the program’s code.  

- [Prealloc](https://github.com/alexkohler/prealloc)  
The Prealloc tool analyses code to find slices that can be replaced by arrays to improve the program’s memory consumption and increase its efficiency while not having to redeclare a backing array when the slice’s size increases.  

- [Safesql](https://github.com/stripe/safesql)  
The Safesql tool analyses code to find possible security vulnerabilities when performing SQL queries into a relational database.  

### Dynamic Analysis Tools

- [Race-detector](https://blog.golang.org/race-detector)  
The Race Detector tool analyses code to find possible race conditions. A race condition occurs when two concurrent processes access the same position in memory, causing incorrect functionality on the program’s behavior.  

- [Benchark-tests](https://golang.org/pkg/testing/)  
The Benchmark Tests analyses running code to get metrics on the performance of the tested functions. The main purpose of these metrics is to compare the performance of each implementation.  

### Debuggers

- [Delve](https://github.com/go-delve/delve)  
The Delve tool is a debugger that gives the developers tools to run the program under a certain environment and provides information about the execution of it.  

### Formatters

- [Gofmt](https://golang.org/cmd/gofmt/)  
The Gofmt tool is the Go’s official formatting tool. It is the standard tool to format Go code into the official best practices of the language.  

### Documentation Generation Tools

- [Godoc](https://blog.golang.org/godoc)  
The Godoc tool is the official documentation generation tool for Go.  



## References

[1] C. Ou, “Go: A Documentary.” https://golang.design/history/