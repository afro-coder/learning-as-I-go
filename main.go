package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"url>loc"`
}

type News struct {
	Titles    []string `xml:"url>name"`
	Keywords  []string `xml:"url>keywords"`
	Locations []string `url:"url>loc"`
}
type NewsMap struct {
	Keyword  string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-business-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
		fmt.Println(news_map)
		fmt.Println(n)
		for idx, _ := range n.Titles {
			fmt.Println("Getting the titles")
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
			break

		}
		fmt.Println(news_map)
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)

	}

}
