run: build
	@./sudoku-go puzzle.txt

build:
	@go build .

benchmark:
	@go test -bench=.