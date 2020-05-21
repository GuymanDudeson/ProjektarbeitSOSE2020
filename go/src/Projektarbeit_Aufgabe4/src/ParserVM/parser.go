package ParserVM

type Parser struct {
	t Tokenizer
}

func NewParser(s string) Parser {
	return Parser{t: newTokenizer(s)}
}

func(p Parser) parse() Optional{
	return p.parseE()
}

func(p Parser) parseE() Optional{
	var t = p.parseT()

	if t.isNothing(){
		return t
	}

	return p.parseE2(t.fromJust())
}

func(p Parser) parseE2(left Exp) Optional{
	if p.t.Token == PLUS {
		p.t.nextToken()

		var right = p.parseT()
		if right.isNothing() {
			return right
		}

		return p.parseE2(NewPlusExp(left, right.fromJust()))
	}

	return NewOptional(left)
}

func(p Parser) parseT() Optional{
	var f = p.parseF()
	if f.isNothing() {
		return f
	}

	return p.parseT2(f.fromJust())
}

func(p Parser) parseT2(left Exp) Optional{
	if p.t.Token == MULT {
		p.t.nextToken()

		var right = p.parseF()
		if right.isNothing() {
			return right
		}
		return p.parseT2(NewMultExp(left, right.fromJust()))
	}

	return NewOptional(left)
}

func(p Parser) parseF() Optional{
	switch p.t.Token {
	case ZERO:
		p.t.nextToken()
		return NewOptional(NewIntExp(0))
	case ONE:
		p.t.nextToken()
		return NewOptional(NewIntExp(1))
	case TWO:
		p.t.nextToken()
		return NewOptional(NewIntExp(2))
	case OPEN:
		p.t.nextToken()

		var e = p.parseE()
		if e.isNothing() {
			return e
		}

		if p.t.Token != CLOSE {
			return NewEmptyOptional()
		}

		p.t.nextToken()
		return e
	default:
		return NewEmptyOptional()
	}
}
