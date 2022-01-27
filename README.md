## Wordle Solver

A helper utility for solving [Wordle](https://www.powerlanguage.co.uk/wordle/)

## Installation

- Run `go mod download` to fetch all dependencies
- `make run` to run without generating a binary
- `make build` to build the executable

## Usage

Modify the following in `main.go` file to arrive at the answer

```go
// Add characters that you know exist in the word to `whitelist`
// eg -- []string{"a"}
whitelist := []string{}

// Add characters that you don't exist in the world to `blacklist`
// eg -- []string{"z"}
blacklist := []string{}

// Add character that exists and it's position in the word is known to `mappings`
// eg -- map[int]string{"0": "a"}
mappings := map[int]string{}
```