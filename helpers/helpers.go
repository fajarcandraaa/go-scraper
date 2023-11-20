package helpers

import (
	"regexp"

	"github.com/gocolly/colly/v2"
)

func SetAgent(r *colly.Request) {
	r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
}

func FormatCurrency(input string) string {
	re := regexp.MustCompile(`(?i)(Rp)(\d[\d.]*)`)
	return re.ReplaceAllString(input, "$1 $2")
}
