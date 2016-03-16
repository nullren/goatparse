goatparse.go: goatparse.y
	go tool yacc -p Goatparse -o goatparse.y.go goatparse.y

EBNF2Y:=$(GOPATH)/bin/ebnf2y

goatparse.y: goatparse.ebnf $(EBNF2Y)
	$(EBNF2Y) -pkg gen -start Goatparse goatparse.ebnf > goatparse.y

$(EBNF2Y):
	go get github.com/cznic/ebnf2y
