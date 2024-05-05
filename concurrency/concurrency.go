package main

type WebCheker func(string) bool

func CheckWeb(wc WebCheker, urls []string) map[string]bool {
	result := make(map[string]bool)
	for _, url := range urls {
		result[url] = wc(url)
	}
	return result
}
