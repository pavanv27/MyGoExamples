Important commands:

***  Test cover:  ***
go test -cover
    PASS
    coverage: 33.3% of statements

go test -coverprofile c.out
    PASS
    coverage: 28.6% of statements

go tool cover -html=c.out
    [shows the untittesting uncovered and covered code]

go tool cover -html=c.out -o coverage.html
    [writes the test coverage info of all files to a html file]


*** Benchmark ***
go test -bench .
    goos: windows
    goarch: amd64
    pkg: Myworkspace/ninja13_testing/1
    BenchmarkSum-4          278886609                4.06 ns/op
    PASS



*** Godoc ***
godoc -http=:8080
    [brings up the godoc page in localhost:8000 page]