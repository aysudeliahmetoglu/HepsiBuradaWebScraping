package main

import(
	"fmt"

	"github.com/gocolly/colly"

)

type item struct{
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string  `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.hepsiburada.com"),

	)
	c.OnHTML("div[moria-ProductCard-joawUM cqPuXj sz7amsy3dwb]", func(h *colly.HTMLElement){
		fmt.Println(h.Text)

	})
    c.Visit("https://www.hepsiburada.com/iphone-14-pro-max-c-80857057")
    
}