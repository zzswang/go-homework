// 爬虫
package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var n sync.WaitGroup

func getCarInfo() {
	defer n.Done()
	doc, err := goquery.NewDocument("http://xianxing.dunkun.com/")
	if err != nil {
		log.Fatal(err)
	}

	todaynum := doc.Find("strong.pink").Eq(0).Text()
	tomorrownum := doc.Find("strong.pink").Eq(1).Text()
	fmt.Printf("今天限号: %s\n", todaynum)
	fmt.Printf("明天限号: %s\n", tomorrownum)
}

func getWeather() {
	defer n.Done()
	doc, err := goquery.NewDocument("http://www.bjmb.gov.cn/")
	if err != nil {
		log.Fatal(err)
	}

	todayTemp := doc.Find(".ri_div span").First().Text()
	todayWeather := doc.Find(".ri_div label").First().Text()
	fmt.Printf("当前天气：%s, %s\n", todayTemp, todayWeather)
}

func getPM2point5() {
	defer n.Done()
	doc, err := goquery.NewDocument("http://aqicn.org/city/beijing/us-embassy/cn/")
	if err != nil {
		log.Fatal(err)
	}

	pm := doc.Find("#aqiwgtvalue").First().Text()
	fmt.Printf("当前PM2.5：%s\n", pm)
}

func main() {
	n.Add(3)
	go getCarInfo()
	go getWeather()
	go getPM2point5()
	n.Wait()
}
