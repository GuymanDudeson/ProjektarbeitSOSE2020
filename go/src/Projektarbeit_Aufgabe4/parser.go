package Projektarbeit_Aufgabe4

type Parser struct {
	t Tokenizer
}

func newParser(s string) Parser{
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
	if p.t.token == PLUS {
		p.t.nextToken()

		var right = p.parseT()
		if right.isNothing() {
			return right
		}

		return p.parseE2(newPlusExp(left, right.fromJust()))
	}

	return newOptional(left)
}

func(p Parser) parseT() Optional{
	var f = p.parseF()
	if f.isNothing() {
		return f
	}

	return p.parseT2(f.fromJust())
}

func(p Parser) parseT2(left Exp) Optional{
	if p.t.token == MULT {
		p.t.nextToken()

		var right = p.parseF()
		if right.isNothing() {
			return right
		}
		return p.parseT2(newMultExp(left, right.fromJust()))
	}

	return newOptional(left)
}

func(p Parser) parseF() Optional{
	switch p.t.token {
	case ZERO:
		p.t.nextToken()
		return newOptional(newIntExp(0))
	case ONE:
		p.t.nextToken()
		return newOptional(newIntExp(1))
	case TWO:
		p.t.nextToken()
		return newOptional(newIntExp(2))
	case OPEN:
		p.t.nextToken()

		var e = p.parseE()
		if e.isNothing() {
			return e
		}

		if p.t.token != CLOSE {
			return newEmptyOptional()
		}

		p.t.nextToken()
		return e
	default:
		return newEmptyOptional()
	}
}
