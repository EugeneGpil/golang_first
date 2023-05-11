Init app
```
go mod init example/hello
```

Run the app
```
go run .
```

Find all imports and install
```
go mod tidy
```

Replace package with local package
```
go mod edit -replace greetings/greetings=../greetings
go mod edit -replace <package name>     =<local path>
```

Run tests
```
go test
```

Run tests with explanation
```
go test -v
```