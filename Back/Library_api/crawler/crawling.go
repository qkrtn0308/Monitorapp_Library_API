package crawler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/labstack/echo"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

const (
	yes24 = "http://www.yes24.com/searchcorner/Search?keywordAd=&keyword=&domain=ALL&qdomain=%C0%FC%C3%BC&query="
	kyobo = "https://search.kyobobook.co.kr/web/search?vPstrKeyWord="
)

var selector string

var data = make([]Yes24Model, 20)

func Crawling(c echo.Context) error {
	//# 사이트와 검색어 쿼리로 받음
	s := c.QueryParam("site")
	q := c.QueryParam("query")
	log.Println(q)
	//# 사이트 선택함
	if s == "yes24" {
		selector = yes24
		log.Println(selector)
		//# yes24 키워드 깨짐 방지로 인코딩 (교보는 돌리면 안됨)
		q, _ = encodeToEUCKR(q)
	} else if s == "kyobo" {
		selector = kyobo
		log.Println(selector)
	} else {
		log.Println("site 선택 오류")
		return c.NoContent(404)
	}
	url := selector + url.QueryEscape(q)
	log.Println(url)

	resp, err := soup.Get(url)
	if err != nil {
		log.Println("<===== Check out your url =====>")
	}

	doc := soup.HTMLParse(resp)

	//! yes24 태그 전용!

	var j int = 0

	trs := doc.Find("div", "class", "goodsList").Find("tbody").FindAll("tr")

	for i, tr := range trs {
		if i%2 != 0 {
			log.Print("<-- -- -->")
		} else {
			cnt := tr.Find("span", "class", "goods_cnt")
			image := tr.Find("a", "class", "img_bdr").Find("img")
			title := tr.Find("a", "class", "img_bdr").Find("img")
			subTitle := tr.Find("span", "class", "goods_sname")
			infos := tr.Find("div", "class", "goods_info")
			price := tr.Find("div", "class", "goods_price").Find("em", "class", "yes_b")
			// a := tr.Find("div", "class", "goods_rating")
			// log.Print(a.HTML())

			if checkNil := tr.Find("div", "class", "goods_rating"); len(checkNil.HTML()) < 100 {
				sellCount := "정보 없음"
				reviewCount := "정보 없음"
				starCount := "정보 없음"

				data[j].Cnt = cnt.Text()
				data[j].ImageURL = image.Attrs()["src"]
				data[j].Title = title.Attrs()["alt"]
				data[j].SubTitle = strings.Trim(subTitle.Text(), "\n ")
				data[j].Infos = strings.Trim(infos.FullText(), "\n ")
				data[j].Price = price.Text() + "원"
				data[j].SellCount = sellCount
				data[j].ReviewCount = reviewCount
				data[j].Star = starCount
				j++

			} else if checkNil := tr.Find("div", "class", "goods_rating"); len(checkNil.HTML()) > 140 && len(checkNil.HTML()) < 1000 {
				sellCount := tr.Find("span", "class", "gd_reviewCount")
				reviewCount := "정보 없음"
				starCount := "정보 없음"

				data[j].Cnt = cnt.Text()
				data[j].ImageURL = image.Attrs()["src"]
				data[j].Title = title.Attrs()["alt"]
				data[j].SubTitle = strings.Trim(subTitle.Text(), "\n ")
				data[j].Infos = strings.Trim(infos.FullText(), "\n ")
				data[j].Price = price.Text() + "원"
				data[j].SellCount = strings.Trim(sellCount.Text(), "\n ")
				data[j].ReviewCount = reviewCount
				data[j].Star = starCount
				j++
			} else if checkNil := tr.Find("div", "class", "goods_rating"); len(checkNil.HTML()) >= 1000 {
				sellCount := tr.Find("span", "class", "gd_reviewCount")
				reviewCount := tr.Find("em", "class", "txC_blue")
				starCount := tr.Find("span", "class", "gd_rating").Find("em", "class", "yes_b")

				data[j].Cnt = cnt.Text()
				data[j].ImageURL = image.Attrs()["src"]
				data[j].Title = title.Attrs()["alt"]
				data[j].SubTitle = strings.Trim(subTitle.Text(), "\n ")
				data[j].Infos = strings.Trim(infos.FullText(), "\n ")
				data[j].Price = price.Text() + "원"
				data[j].SellCount = strings.Trim(sellCount.Text(), "\n ")
				data[j].ReviewCount = "회원리뷰 " + reviewCount.Text() + "건"
				data[j].Star = "평균 " + starCount.Text() + "점"
				j++
			}
			log.Print("<-- ", j, " -->")

		}
	}

	re, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	log.Println(string(re))
	return c.String(200, string(re))
}

//# utf-8 to euc-kr ENCODER
func encodeToEUCKR(s string) (string, error) {
	var buf bytes.Buffer
	wr := transform.NewWriter(&buf, korean.EUCKR.NewEncoder())
	_, err := wr.Write([]byte(s))
	if err != nil {
		return "", err
	}
	defer wr.Close()
	return buf.String(), nil
}

// //#  euc-kr to utf-8 DECODER
// func encodeToUTF8(s string) (string, error) {
// 	var buf bytes.Buffer
// 	wr := transform.NewWriter(&buf, korean.EUCKR.NewDecoder())
// 	_, err := wr.Write([]byte(s))
// 	if err != nil {
// 		return "", err
// 	}
// 	defer wr.Close()
// 	return buf.String(), nil
// }
