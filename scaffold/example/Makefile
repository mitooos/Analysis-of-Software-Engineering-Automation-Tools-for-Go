

# run the project
run:
	go run src/main.go

# build the project
build:
	go build src/main.go

# tidies dependencies
tidy-deps:
	go mod tidy

# vendors dependencies
deps:
	go mod tidy
	go mod vendor

# run all tests with coverage
test-cover:
	go test ./... -cover -coverprofile=coverage.out -race

# launches browser with coverage report from last test run
watch-coverage:
	go tool cover -html=coverage.out

# serves project's documentation on port 6060
docs:
	godoc -http=:6060

# formats project with go's style guidelines
fmt:
	gofmt -w -s ./src

# checks the project for possible errors
check:
	staticcheck ./...

# generates mocks for interfaces defined in the project
gen-mocks:
	mockery -r --srcpkg=./src/ --all

