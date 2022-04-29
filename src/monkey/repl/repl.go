package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
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

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			{
				_, err := io.WriteString(out, evaluated.Inspect())
				if err != nil {
					panic(err)
				}
			}
			{
				_, err := io.WriteString(out, "\n")
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	{
		_, err := io.WriteString(out, KABAO)
		if err != nil {
			panic(err)
		}
	}
	{
		_, err := io.WriteString(out,
			"すごい一体感を感じる。今までにない何か熱い一体感を。\n")
		if err != nil {
			panic(err)
		}
	}
	{
		_, err := io.WriteString(out, "\n")
		if err != nil {
			panic(err)
		}
	}
	{
		_, err := io.WriteString(out, "parser errors: \n ")
		if err != nil {
			panic(err)
		}
	}
	for _, msg := range errors {
		_, err := io.WriteString(out, fmt.Sprintf("\t%s\n", msg))
		if err != nil {
			panic(err)
		}
	}
}
