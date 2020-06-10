# Sudoku Go

A mini app to solve Sudoku puzzles.

## How to run with Go

0. Install and configure Golang on your machine [https://golang.org/dl/](https://golang.org/dl/)
1. Add the puzzle into the puzzle.txt file; A sample file is provided as in this repo.
2. Run `make` OR `go run main.go puzzle.txt` and it will output the result.

## How to go without Go

If you don't want to bother about Golang installation, you can compile and run the app using the following Docker command:

```
docker run -it --rm -v $PWD:/go/src/app -w /go/src/app golang:1.14 make
```

## Benchmark the performance of Go functions

Run `make benchmark` OR `go test -bench=.` if you have Golang installed;

Otherwise, run

```
docker run -it --rm -v $PWD:/go/src/app -w /go/src/app golang:1.14 make benchmark
```

Here is a sample output from my environment

```
goos: linux
goarch: amd64
pkg: github.com/jiaqi-yin/sudoku-go
BenchmarkSolvePuzzle-4      2697090        429 ns/op
PASS
ok      github.com/jiaqi-yin/sudoku-go      1.836s
```