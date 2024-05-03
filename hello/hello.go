package hello

const (
	english_pre = "hello, "
	spanish_pre = "hola, "
	franch_pre  = "oui, "

	franch  = "franch"
	spanish = "spanish"
)

func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return prefix(lang) + name
}

func prefix(lang string) (pre string) {
	switch lang {
	case spanish:
		pre = spanish_pre
	case franch:
		pre = franch_pre
	default:
		pre = english_pre
	}
	return
}
