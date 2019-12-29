package sitemap

import "net/url"

var ChangeFreqs = []string{"always", "hourly", "daily", "weekly", "monthly", "yearly", "never"}

func validateURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func validateChangeFreq(changeFreq string) bool {
	for _, b := range ChangeFreqs {
		if b == changeFreq {
			return true
		}
	}
	return false
}