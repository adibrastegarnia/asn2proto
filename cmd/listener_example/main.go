package main

import (
	"fmt"
	"os"

	"github.com/adibrastegarnia/asn1toproto/pkg/asn"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type AsnListener struct {
	*asn.BaseASNListener
}

func NewAsnListener() *AsnListener {
	return new(AsnListener)
}

func (l *AsnListener) EnterEnumeratedValue(ctx *asn.EnumeratedValueContext) {
	fmt.Println("hello")
	fmt.Println(ctx.GetText())
}

func (l *AsnListener) EnterObjectClass(ctx asn.ObjectClassContext) {

}

func (l *AsnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("Exit:", ctx.GetText())
}

func (l *AsnListener) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("Terminal:", node.GetText())
}

func (l *AsnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("Entry:", ctx.GetText())
	//fmt.Println()

}

func main() {

	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := asn.NewASNLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := asn.NewASNParser(stream)

	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.Assignment()
	var listener AsnListener

	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)

}
