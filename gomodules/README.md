# Creating Go modules 

## Steps


### Create module a `greeting` module (to be imported by others)

```bash
cd greetings/
# make this directory a module
go mod init src/greetings
```

### Create module that will import `greeting` module

```bash
cd hello/
# make this directory a module
go mod init src/hello
# point module to local directory
go mod edit -replace src/greetings=../greetings
# make go recognize module and assign version
go mod tidy
```
