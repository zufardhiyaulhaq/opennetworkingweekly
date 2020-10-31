package main

import (
	"log"
	"strings"
	"time"

	communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	"github.com/zufardhiyaulhaq/opennetworkingweekly/handlers"
	"github.com/zufardhiyaulhaq/opennetworkingweekly/models"
	"github.com/zufardhiyaulhaq/opennetworkingweekly/pkg/scrappers"
	"gopkg.in/yaml.v2"
)

func main() {
	// start GitHub handler
	handler := handlers.Github{}
	handler.Start()

	// Init scrapper
	scrapper := scrappers.OpenNetworkingWeekly{}

	// initialize content
	var content models.OpenNetworkingContents

	// check current content file exist
	if !handler.FileExist("content.yaml") {
		newContent, err := yaml.Marshal(content)
		if err != nil {
			log.Fatal(err)
		}

		CreateFile(handler, "content.yaml", "init OpenNetworking Weekly content file", newContent)
	}

	// get current content
	contentTmpl := handler.GetFile("content.yaml")
	err := yaml.Unmarshal(contentTmpl, &content)
	if err != nil {
		log.Fatal(err)
	}
	currentContentLength := len(content.Content)

	// feed current content to scrapper
	openNetworkingWeeklycontent := scrapper.GetOpenNetworkingWeekly(content)
	openNetworkingWeeklycontentByte, err := yaml.Marshal(openNetworkingWeeklycontent)
	if err != nil {
		log.Fatal(err)
	}

	newContentLength := len(openNetworkingWeeklycontent.Content)

	if newContentLength != currentContentLength {
		handler.UpdateFile("content.yaml", "update OpenNetworking Weekly content", openNetworkingWeeklycontentByte)
		log.Printf("[opennetworkingweekly] Update content with new weekly content")
	} else {
		log.Printf("[opennetworkingweekly] no update about Open Networking")
		return
	}

	// communityv1alpha1 "github.com/cloudnative-id/community-operator/pkg/apis/community/v1alpha1"
	var weeklyName string
	var weeklyData []communityv1alpha1.ArticleSpec

	// populate weeklyData from opennetworkingWeeklycontent
	for index, value := range openNetworkingWeeklycontent.Content {
		if !value.IsDelivered {
			var data communityv1alpha1.ArticleSpec
			data.Title = value.Title
			data.Url = value.Url
			data.Type = value.Kind
			weeklyData = append(weeklyData, data)
			openNetworkingWeeklycontent.Content[index].IsDelivered = true
		}
	}

	// push the updated content.yaml
	openNetworkingWeeklycontentByte, err = yaml.Marshal(openNetworkingWeeklycontent)
	if err != nil {
		log.Fatal(err)
	}
	handler.UpdateFile("content.yaml", "update OpenNetworking Weekly content", openNetworkingWeeklycontentByte)
	log.Printf("[opennetworkingweekly] Update delivered status content")

	//Init builder
	builder := Builder{}

	// Build
	time := time.Now().Format("02-01-2006")
	weeklyName = "Open Networking Weekly " + time

	builder.build(weeklyName, weeklyData)
	weeklyCRD, err := yaml.Marshal(builder)

	commitMessage := "Weekly: Add " + weeklyName
	CreateFile(handler, strings.ToLower(strings.ReplaceAll(weeklyName, " ", "-"))+".yaml", commitMessage, weeklyCRD)
	log.Printf("[opennetworkingweekly] Create Weekly manifest")
}
