package parser

import (
	"Projektarbeit_Aufgabe4/src/ast"
	"Projektarbeit_Aufgabe4/src/tokenizer"
	"Projektarbeit_Aufgabe4/src/utility"
)


type Parser struct {
	t tokenizer.Tokenizer
}

func NewParser(s string) Parser{
	return Parser{t: tokenizer.NewTokenizer(s)}
}

func(p Parser) Parse() utility.Optional{
	return p.ParseE()
}

func(p Parser) ParseE() utility.Optional{
	var t = p.ParseT()

	if t.IsNothing(){
		return t
	}

	return p.ParseE2(t.FromJust())
}

func(p Parser) ParseE2(left ast.Exp) utility.Optional{
	if p.t.Token == tokenizer.PLUS {
		p.t.NextToken()

		var right = p.ParseT()
		if right.IsNothing() {
			return right
		}

		return p.ParseE2(ast.NewPlusExp(left, right.FromJust()))
	}

	return utility.NewOptional(left)
}

func(p Parser) ParseT() utility.Optional{
	var f = p.ParseF()
	if f.IsNothing() {
		return f
	}

	return p.ParseT2(f.FromJust())
}

func(p Parser) ParseT2(left ast.Exp) utility.Optional{
	if p.t.Token == tokenizer.MULT {
		p.t.NextToken()

		var right = p.ParseF()
		if right.IsNothing() {
			return right
		}
		return p.ParseT2(ast.NewMultExp(left, right.FromJust()))
	}

	return utility.NewOptional(left)
}

func(p Parser) ParseF() utility.Optional{
	switch p.t.Token {
	case tokenizer.ZERO:
		p.t.NextToken()
		return utility.NewOptional(ast.NewIntExp(0))
	case tokenizer.ONE:
		p.t.NextToken()
		return utility.NewOptional(ast.NewIntExp(1))
	case tokenizer.TWO:
		p.t.NextToken()
		return utility.NewOptional(ast.NewIntExp(2))
	case tokenizer.OPEN:
		p.t.NextToken()

		var e = p.ParseE()
		if e.IsNothing() {
			return e
		}

		if p.t.Token != tokenizer.CLOSE {
			return utility.NewEmptyOptional()
		}

		p.t.NextToken()
		return e
	default:
		return utility.NewEmptyOptional()
	}
}
