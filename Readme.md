## Required commands

- Create Go file and run

```
mkdir test
cd test
go mod init test
ni main.go
go run main.go
```

- Build app in Go

```
cd test
go build
```
or 
```
cd test
GOOS="windows" go build
```