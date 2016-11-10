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

        doc.Find(".productCategoryContainerWraper").Each(func(i int, s *goquery.Selection) {
                fmt.Println(i)
       
                title := s.Find("a").Text()
                link, _ := s.Find("a").Attr("href")
                fmt.Println(title)
                fmt.Println(link)
        })

		r.HTML(http.StatusOK, "index", nil)	
	})

	m.RunOnAddr(":8094")
}
