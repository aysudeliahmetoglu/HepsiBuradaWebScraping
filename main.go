package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	
        
)

func main() {
	
	// url := "https://www.hepsiburada.com/iphone-14-pro-max-512-gb-p-HBCV00002W46NM"
    
	// res,_ := http.Get(url)

	// if res.StatusCode != 200 {
	// 	fmt.Println("Hata:",res.StatusCode)
	// 	return
	// }
	client := &http.Client{
		
	}
	// resp, err := http.Get("https://www.hepsiburada.com/iphone-14-pro-max-512-gb-p-HBCV00002W46NM")
	
	
	req, err := http.NewRequest("GET", "https://www.hepsiburada.com/iphone-14-pro-max-512-gb-p-HBCV00002W46NM", nil)
	

	req.Header.Add("If-None-Match", `W/" Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0"`)
	
	resp, err := client.Do(req)
	fmt.Println(resp)
	
	if err != nil {
		fmt.Println("error")
	}
	
	
	doc,_ := goquery.NewDocumentFromReader(resp.Body)


    title := doc.Find(".productDetailRight .product-information").Text()	
	price := doc.Find(".currentPriceBeforePoint").Text()
    fmt.Println(title)
	fmt.Println(price)
}
