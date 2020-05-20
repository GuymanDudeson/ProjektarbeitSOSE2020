package Projektarbeit_Aufgabe4

type Optional struct {
	b bool
	val string
}

func newEmptyOptional() *Optional {
	return &Optional{
		b:   false,
		val: "",
	}
}

func newOptional(val string) *Optional{
	return &Optional{
		b:   true,
		val: val,
	}
}

func isJust(o *Optional) bool{
	return o.b
}

func isNothing(o *Optional) bool{
	return !o.b
}

func fromJust(o *Optional) string{
	return o.val
}

func nothing() *Optional{
	return newEmptyOptional()
}

func just(val string) *Optional{
	return newOptional(val)
}



