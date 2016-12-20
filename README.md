# go_sandbox
My GO code sandbox

## Install:
[Download](https://golang.org/dl/) the installer.
Continue -> Continue -> Continue -> Install -> bla bla next done.
```
$ cd $HOME/hd2/code
$ git clone GIT_REPO_NAME
$ export GOPATH=$HOME/hd2/code/go_sandbox
$ cd GIT_REPO_NAME
```

### I used simple one-liner's here for copy paste ease. so the semi-colon
### might be misleading for new users. have a look at the file after gofmt
### and goimport with less.

## Hello World Example: 
```
$ mkdir src/hello
$ touch src/hello/main.go
$ echo 'package main; import "fmt"; func main() { fmt.Println("Hello World!") }' > src/hello/main.go
$ gofmt -w src/hello/main.go
$ less src/hello/main.go
$ go run src/hello/main.go
```

## Command Line Args:
```
$ mkdir src/cmdline
$ touch src/cmdline/main.go
$ echo 'package main;import("fmt";"os");func main(){if len(os.Args)>1{fmt.Println(os.Args[1])}else{fmt.Println("Hello World!")}}' > src/cmdline/main.go
$ gofmt -w src/cmdline/main.go
$ less src/cmdline/main.go
$ go run src/cmdline/main.go
$ go run src/cmdline/main.go "Hello me"
```

## Add missing imports:
### Install [it](https://godoc.org/golang.org/x/tools/cmd/goimports):
```
$ go get golang.org/x/tools/cmd/goimports
```
Try it:
```
$ mkdir src/miss
$ touch src/miss/main.go
$ echo 'package main;func main(){if len(os.Args)>1{fmt.Println(os.Args[1])}else{fmt.Println("Hello World!")}}' > src/miss/main.go
$ go run src/miss/main.go
```
*Witness import's are missing
```
$ ./bin/goimports -w src/miss/main.go
$ go run src/miss/main.go
$ go run src/cmdline/main.go "I fixed the imports!"
```