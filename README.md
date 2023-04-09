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
go test -v -coverprofile=coverage.out ./...
ここに前処理を書く
=== RUN   TestMaxTableDrivenTests
=== RUN   TestMaxTableDrivenTests/正の数
=== RUN   TestMaxTableDrivenTests/負の数
=== RUN   TestMaxTableDrivenTests/両方
--- PASS: TestMaxTableDrivenTests (0.00s)
    --- PASS: TestMaxTableDrivenTests/正の数 (0.00s)
    --- PASS: TestMaxTableDrivenTests/負の数 (0.00s)
    --- PASS: TestMaxTableDrivenTests/両方 (0.00s)
=== RUN   ExampleMax
--- PASS: ExampleMax (0.00s)
PASS
        github.com/ys-office-llc/Go_Examples-of-test-code/calc  coverage: 100.0% of statements
ここに後処理を書く
ok      github.com/ys-office-llc/Go_Examples-of-test-code/calc  0.001s  coverage: 100.0% of statements
```

## Bibliography
- [テスト（go test/testing）](https://www.twihike.dev/docs/golang-primer/testing)
- [Effective Go](https://go.dev/doc/effective_go)
