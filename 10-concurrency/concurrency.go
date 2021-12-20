package concurrency

type WebsiteChecker func(url string) bool
type Result struct {
	url    string
	result bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	ch := make(chan Result)
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			ch <- Result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-ch
		results[r.url] = r.result
	}

	return results
}
