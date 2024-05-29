package main

type WebChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWeb(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChanal := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChanal <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChanal
		results[r.string] = r.bool
	}

	return results
}
