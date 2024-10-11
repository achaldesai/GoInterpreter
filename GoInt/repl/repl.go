package repl

import (
	"GoInt/evaluator"
	"GoInt/lexer"
	"GoInt/parser"
	"bufio"
	"fmt"
	"io"
)

const (
	PROMPT      = ">> "
	errorHeader = "---------------------Whoops that's an error alright !!---------------------" + "\n"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, errorHeader)
		io.WriteString(out, "\t"+msg+"\n")
	}
}
