# module_testing

1. go mod init creates a new module, initializing the go.mod file that describes it.
2. go build, go test, and other package-building commands add new dependencies to go.mod as needed.
3. go list -m all prints the current module’s dependencies.
4. go get changes the required version of a dependency (or adds a new dependency).
5. go mod tidy removes unused dependencies.