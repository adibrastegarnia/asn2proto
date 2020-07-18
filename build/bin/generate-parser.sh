#!/bin/sh

GRAMMAR_PATH=${GOPATH}/go/src/github.com/adibrastegarnia/asn1toproto/pkg/grammar/ASN.g4
GENERATED_PARSER_PATH=${GOPATH}/go/src/github.com/adibrastegarnia/asn1toproto/pkg/asn

java  -jar /usr/local/lib/antlr-4.8-complete.jar -Dlanguage=Go -o ${GENERATED_PARSER_PATH}  -package asn -visitor ${GRAMMER_PATH}