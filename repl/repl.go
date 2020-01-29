package repl

import (
	"github.com/Delta456/interpreter_go/lexer"
	"github.com/Delta456/interpreter_go/token"

	"bufio"
	"fmt"
	"io"
)

// PROMPT is the prompt for the REPL
const PROMPT = ">> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		if scanned := scanner.Scan(); !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		if line == "exit" {
			break
		}
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
