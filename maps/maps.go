package maps

import "errors"

type Dic map[string]string

func (d Dic) Search(s string) (string, error) {
	r, ok := d[s]

	if !ok {
		return "", errors.New("could not find the word")
	} else {
		return r, nil
	}
}

func (d Dic) Add(key, value string) {
	d[key] = value
}
