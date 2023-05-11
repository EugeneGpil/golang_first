Compile the code into the executable
```
go build
```

Get install path
```
go list -f '{{.Target}}'

go list -f '{{.Target}}'
/home/app/go/bin/hello
```

Add Go install directory to PATH
```
export PATH=$PATH:/path/to/your/install/directory

export PATH=$PATH:/home/app/go/bin
```

Install
```
go install
```
