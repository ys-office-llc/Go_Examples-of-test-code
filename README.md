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
go test -v -coverprofile=coverage.out ./...
=== RUN   TestGetUsers
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /users                    --> github.com/ys-office-llc/Go_Examples-of-test-code/api.GetUsers (3 handlers)
[GIN] 2023/04/09 - 14:02:53 | 200 |      28.853µs |                 | GET      "/users"
--- PASS: TestGetUsers (0.00s)
PASS
        github.com/ys-office-llc/Go_Examples-of-test-code/api   coverage: 25.0% of statements
ok      github.com/ys-office-llc/Go_Examples-of-test-code/api   0.005s  coverage: 25.0% of statements
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
ok      github.com/ys-office-llc/Go_Examples-of-test-code/calc  0.002s  coverage: 100.0% of statements
```

## Bibliography
- [テスト（go test/testing）](https://www.twihike.dev/docs/golang-primer/testing)
- [Effective Go](https://go.dev/doc/effective_go)
