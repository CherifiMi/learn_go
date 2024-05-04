package maps

type Dic map[string]string

func (d Dic) Search(s string) string {
	return d[s]
}
