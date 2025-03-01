package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) { // u is a copy of url
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for range urls {
		r := <-resultChannel // this loop reads the data and what happens is this loop will be blocked for next iteration and wait for the upper loops goroutine to write to the channel
		results[r.string] = r.bool
	}

	return results
}
