package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/kgosse/monkeygo/token"

	"github.com/kgosse/monkeygo/lexer"
)

// PROMPT contains the starting message of the prompt
const PROMPT = "42sh> "

// Start starts the repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
