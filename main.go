package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"net/http"
	"fmt"
	"log"
)

type Category struct {
	gorm.Model
	CategoryId int
	CategoryName string
	Url string
	Title string
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		db, err := gorm.Open("mysql", "root:852456@/interior?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println(err)
			panic("failed to connect database")
		}
		defer db.Close()

		db.CreateTable(&Category{})
		
		

		//Fetch data
		doc, err := goquery.NewDocument("http://www.ikea.com/ru/ru/catalog/allproducts/")
		if err != nil {
			log.Fatal(err)
		}

		axilary := doc.Find(".rightContent")

		axilary.Find(".header").Each(func(i int, s *goquery.Selection) {
			header := s.Text()
			category := Category {
				CategoryId: i,
				CategoryName: header,
				Url: "null",
				Title: "null",
			}
			db.Create(&category)
			db.Save(&category)
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

	m.RunOnAddr(":8098")
}
