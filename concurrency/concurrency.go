package concurrency

type WebsiteCheker func(string) bool

func CheckWebsites(wc WebsiteCheker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}
	return results
}
