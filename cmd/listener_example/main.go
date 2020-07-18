// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/adibrastegarnia/asn1toproto/pkg/asn"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// AsnListener ...
type AsnListener struct {
	*asn.BaseASNListener
}

// EnterEnumeratedValue ...
func (l *AsnListener) EnterEnumeratedValue(ctx *asn.EnumeratedValueContext) {
	fmt.Println("hello")
	fmt.Println(ctx.GetText())
}

// ExitEveryRule ...
func (l *AsnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("Exit:", ctx.GetText())
}

// VisitTerminal ...
func (l *AsnListener) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("Terminal:", node.GetText())
}

// EnterEveryRule ...
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
