package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">> "

const KABAO = `      ∩＿∩
   ／ ＼ ／＼
  |（ﾟ）=（ﾟ）|
  |   ●_●  |
 /           ヽ
 | 〃 ------ ヾ
 ＼＿＿二＿＿ノ
`

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, KABAO)
	io.WriteString(out, "すごい一体感を感じる。今までにない何か熱い一体感を。\n")
	io.WriteString(out, "parser errors: \n ")
	for _, msg := range errors {
		io.WriteString(out, fmt.Sprintf("\t%s\n", msg))
	}
}
