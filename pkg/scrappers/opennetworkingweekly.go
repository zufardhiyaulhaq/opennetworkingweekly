package scrappers

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/zufardhiyaulhaq/opennetworkingweekly/models"
)

type OpenNetworkingWeekly struct{}

func (s *OpenNetworkingWeekly) GetOpenNetworkingWeekly(currentContent models.OpenNetworkingContents) models.OpenNetworkingContents {
	content := getOpenNetworkingNews(currentContent)

	return content
}

func getOpenNetworkingNews(currentContent models.OpenNetworkingContents) models.OpenNetworkingContents {
	res, err := http.Get("https://opennetworking.org/category/news-and-events/latest-news/")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".blog-item").Each(func(i int, data *goquery.Selection) {
		url, _ := data.Find("a").Attr("href")
		title, _ := data.Find("a").Attr("title")
		doAppend := true

		for _, v := range currentContent.Content {
			if v.Url == url {
				doAppend = false
			}
		}

		if doAppend {
			content := models.Content{
				Title: title, Url: url, Kind: "Open Networking News", IsDelivered: false,
			}
			currentContent.Content = append(currentContent.Content, content)
		}
	})

	return currentContent
}
