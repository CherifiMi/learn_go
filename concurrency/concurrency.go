package main

type WebCheker func(string) bool
type result struct {
	string
	bool
}

func CheckWeb(wc WebCheker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChanel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChanel
		results[r.string] = r.bool
	}

	return results
}
