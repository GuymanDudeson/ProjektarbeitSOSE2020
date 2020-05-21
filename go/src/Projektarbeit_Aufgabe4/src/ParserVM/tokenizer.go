package ParserVM

type Token int
const (
	EOS   Token = 0
	ZERO  Token = 1
	ONE   Token = 2
	TWO   Token = 3
	OPEN  Token = 4
	CLOSE Token = 5
	PLUS  Token = 6
	MULT  Token = 7
)

func showTok(t Token) string {
	switch t {
		case EOS: return "EOS"
		case ZERO: return "ZERO"
		case ONE: return "ONE"
		case TWO: return "TWO"
		case OPEN: return "OPEN"
		case CLOSE: return "CLOSE"
		case PLUS: return "PLUS"
		case MULT: return "MULT"
	}
	return "Error"
}

type Tokenize struct {
	s string
	pos int
}

func newTokenize(_s string) Tokenize {
	var t = Tokenize{
		s:   _s,
		pos: 0,
	}
	return t
}

func (t Tokenize) Next() Token {
	if len(t.s) <= t.pos {
		return EOS
	}

	for true {
		if len(t.s) <= t.pos {
			return EOS
		}

		switch t.s[t.pos] {
			case '0':
				t.pos++
				return ZERO
			case '1':
				t.pos++
				return ONE
			case '2':
				t.pos++
				return TWO
			case '(':
				t.pos++
				return OPEN
			case ')':
				t.pos++
				return CLOSE
			case '+':
				t.pos++
				return PLUS
			case '*':
				t.pos++
				return MULT
		default:
			t.pos++
			break
		}
	}
	return EOS
}

func (t Tokenize) scan() []Token {
	var v []Token
	var tempToken Token

	for {
		tempToken = t.Next()
		v = append(v, tempToken)
		if tempToken == EOS {
			break
		}
	}
	return v
}

func (t Tokenize) show() string {
	var v = t.scan()
	var s string

	for i, currentToken := range v {
		s += showTok(Token(currentToken))

		if i+1 < len(v) {
			s += ";"
		}
	}
	return s
}

type Tokenizer struct {
	Tokenize Tokenize
	Token    Token
}

func newTokenizer(s string) Tokenizer {
	var tokenize = newTokenize(s)
	var token = tokenize.Next()
	return Tokenizer{tokenize, token}
}

func (t Tokenizer) nextToken() {
	t.Token = t.Tokenize.Next()
}

