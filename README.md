# Go examples of test code

## To run a test

```shell
make test
```

## To create a coverage file

```shell
make cover
```

## To display coverage

```shell
make open
```

## Execution sample

```shell
$ make test
go test -coverprofile=coverage.out ./...
ok      github.com/ys-office-llc/Go_Examples-of-test-code/calc  0.001s  coverage: 100.0% of statements
```
