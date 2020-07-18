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

func (l *AsnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())

}

func main() {

	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := asn.NewASNLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := asn.NewASNParser(stream)

	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.AssignmentList()
	var listener AsnListener

	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)

}
