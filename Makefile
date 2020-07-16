

parser:
	go run golang.org/x/tools/cmd/goyacc -l -p parser -o pkg/parser/parser.go pkg/parser/asn1.y