# go-sudoku
- [Knuth's Algorithm X](https://en.wikipedia.org/wiki/Knuth%27s_Algorithm_X) based sudoku puzzle solver and generator.
- Terminal data: `Json`
- Edges: `1 * 1`, `2 * 2`, `3 * 3`, ..., `n * n`
- Block: `square`, `random`, `irregular`


## Usage

#### Solver

```go
func SolveString(raw string) string
func SolveByte(raw []byte) []byte
func SolveTerminalJson(t *TerminalJson) *TerminalJson
```

#### Generator

```go
func GenString(edge int, mode GeneratorMode, minSubGiven int, minTotalGiven int) string
func GenByte(edge int, mode GeneratorMode, minSubGiven int, minTotalGiven int) []byte
```

#### Mobile
The [Go mobile](https://github.com/golang/go/wiki/Mobile) subrepository adds support for mobile platforms (Android and iOS).

## Dependencies
- https://github.com/jiangyang5157/go-dlx
- https://github.com/jiangyang5157/go-graph/tree/master/graph
- https://github.com/jiangyang5157/golang-start/tree/master/data/stack
