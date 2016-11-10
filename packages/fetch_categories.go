package fetch_categories

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	//"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	//"github.com/go-sql-driver/mysql"

	"net/http"
	//"fmt"
	//"log"
)

type Category struct {
	gorm.Model
	Category string
	Urls []string
	Titles []string
}


func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		db, err := gorm.Open("mysql", "user=root password=852456 dbname=interior sslmode=disable")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&Category{})

		/* fetch data
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
		*/
		r.HTML(http.StatusOK, "index", nil)	
	})

	m.RunOnAddr(":8095")
}
