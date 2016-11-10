package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/PuerkitoBio/goquery"
	
	"net/http"
	"fmt"
	"log"
	//"io/ioutil"
)



func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {

		doc, err := goquery.NewDocument("http://www.ikea.com/ru/ru/catalog/allproducts/")
		if err != nil {
				log.Fatal(err)
		}

		axilary := doc.Find(".rightContent")

		axilary.Find(".header").Each(func(i int, s *goquery.Selection) {
			//header := s.Text()
		})

		axilary.Find(".textContainer").Each(func(i int, s *goquery.Selection) {
			s.Find("a").Each(func(j int, t *goquery.Selection) {
				//title := t.Text()
			})

			s.Find("a").Each(func(j int, t *goquery.Selection) {
				//url, _ := s.Attr("href")
			})
		})	

		r.HTML(http.StatusOK, "index", nil)	
	})

	m.RunOnAddr(":8095")
}
