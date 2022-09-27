## DB Browser for SQLite
- https://sqlitebrowser.org/

### Linux installation
```
sudo snap install sqlitebrowser
```

## Run
```
go test --bench=. -count 3
```

## Results
| Test                     | Execution count | Execution time |
|--------------------------|-----------------|----------------|
| BenchmarkInsertRowsCGO-8 | 409             | 3034936 ns/op  |
| BenchmarkInsertRowsCGO-8 | 379             | 3360784 ns/op  |
| BenchmarkInsertRowsCGO-8 | 352             | 3268604 ns/op  |
| BenchmarkSelectRowsCGO-8 | 12066           | 93303 ns/op    |
| BenchmarkSelectRowsCGO-8 | 11474           | 90582 ns/op    |
| BenchmarkSelectRowsCGO-8 | 11580           | 93422 ns/op    |
| BenchmarkInsertRowsGO-8  | 291             | 4172648 ns/op  |
| BenchmarkInsertRowsGO-8  | 277             | 4204259 ns/op  |
| BenchmarkInsertRowsGO-8  | 284             | 4185075 ns/op  |
| BenchmarkSelectRowsGO-8  | 8052            | 129715 ns/op   |
| BenchmarkSelectRowsGO-8  | 8624            | 116418 ns/op   |
| BenchmarkSelectRowsGO-8  | 8526            | 131107 ns/op   |




| Test                   | Time average            | Difference |
|------------------------|-------------------------|------------|
| BenchmarkInsertRowsCGO | 3221441,33 ns/op        | -23,06%    |
| BenchmarkInsertRowsGO  | 4187327,33 ns/op        |            |


| Test                   | Time average            | Difference |
|------------------------|-------------------------|------------|
| BenchmarkSelectRowsCGO | 277307 ns/op            | -26,49%    |
| BenchmarkSelectRowsGO  | 377240 ns/op            |            |
