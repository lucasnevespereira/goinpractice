## Waitgroups

Sometimes you want to start multiple goroutines but not continue working until they are finished.
Go wait groups are a simple way to do that.
To illustrate we have here a gzip compression script.

Example run command
```
go run main.go exampleData/*
```