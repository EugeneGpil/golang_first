# How to publish and fetch new version of package

In github add new release

In browser
```
                                         only lowercase username
                                         VVVVVVVVVV
https://sum.golang.org/lookup/github.com/eugenegpil/greetings@v0.0.1
                                                    ^^^^^^^^^
                                                    only lowercase package name
```

in `go.mod`
```
module github.com/EugeneGpil/go_getting_started

go 1.20

//                                      update version manually
//                                      VVVVVV
require github.com/EugeneGpil/greetings v0.0.1
```

in terminal
```
go mod tidy
```
