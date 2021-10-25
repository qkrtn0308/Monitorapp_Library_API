package crawler

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

const (
	yes24 = "http://www.yes24.com/searchcorner/Search?keywordAd=&keyword=&domain=ALL&qdomain=%C0%FC%C3%BC&query="
)

func Crawling(c echo.Context) error {
	q := c.QueryParam("query")
	url := yes24 + url.QueryEscape(q)
	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(doc)

	for i := 1; i <= 10; i++ {
		div_three, _ := doc.Find("tr").Eq(i).Html()
		log.Println(div_three)
		encoded, _ := URLDecode(div_three)
		log.Println(encoded)
	}

	return c.NoContent(200)
}
