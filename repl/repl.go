package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/zeddee/learn-write-interpreter-in-go/lexer"
	"github.com/zeddee/learn-write-interpreter-in-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			// if no input,
			// then break for loop
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
