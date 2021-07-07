# Go scaffold template for hexagonal microservices using DDD

## Architectural principles

For each domain, one folder is created with 4 sub-folders:

1. **ports/**: The main entry-point to the application. It can be a rest endpoints, a GrapQL server, a pub/sub listener, gRPC server, etc...
2. **app/**: A thin layer gluing the other components. Think about it like the orchestrator of the other components.
3. **domain/**: Where the business objects and logic live. It must be isolated and know nothing about other components.
4. **adapters/**: The layer where the application communicates with the outer world. It can be a database, pub/sub publisher, gRPC clint, etc...

Remeber to **Accept interfaces, return structs**, this will maintain code testable and extensible.  
Under the example folder you can see a project following these principles created with the same template and the `values_example.json` file.

## Required dependencies

- [Scaffold](https://pkg.go.dev/github.com/metakeule/scaffold)
- [Make](https://www.gnu.org/software/make/manual/make.html) (comes by default on most unix machines, check installation with `make --help`)
- [Mockery](https://github.com/vektra/mockery)

## Scaffold project

1. create a values.json file with the desired output (values_example.json was used to crete the example)
2. Run `scaffold -t=template.templ < values.json`
3. Run `make deps` inside the generated project
4. the project is ready to use

## Commands

All common commands are available on the generated Makefile.  
To run a command run:

```bash
   make [command]
```

## Tests mocks

To generate mocks run `make gen-mocks`, this will generate mocks from all the exported interfaces and place them under the `mocks/` folder. Now you just need to import them from this directory and use them for your unit tests.  
This uses the library [Testsify](https://github.com/stretchr/testify)

