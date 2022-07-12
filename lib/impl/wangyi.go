package impl

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/codecodify/news/lib"
	"log"
	"strings"
)

type Wangyi struct{}

func (c *Wangyi) Get(index int) (lib.Response, error) {
	url := "https://www.163.com/dy/media/T1603594732083.html"
	var resp lib.Response
	body, err := lib.Fetch(url)
	if err != nil {
		log.Printf("fetch wangyi 163 url error:%s\n", err)
		return resp, err
	}
	//fmt.Printf("%s", body)

	buffer := bytes.NewBuffer(body)
	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		log.Printf("wangyi 163 go-goquery document error:%s\n", err)
		return resp, err
	}
	selection := doc.Find("a[class='title']")
	index = lib.GetCorrectionIndex(index, selection.Length())
	var newUrl string
	selection.Each(func(i int, selection *goquery.Selection) {
		if i == index {
			newUrl = selection.AttrOr("href", "")
		}
	})

	if len(newUrl) == 0 {
		return resp, errors.New("fetch news href error")
	}

	//fmt.Println(newUrl)
	if body, err = lib.Fetch(newUrl); err != nil {
		log.Printf("fetch wangyi 163 news url error:%s\n", err)
		return resp, err
	}

	if doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(body)); err != nil {
		log.Printf("fetch wangyi 163 news url error:%s\n", err)
		return resp, err
	}

	var lists []string
	doc.Find("div[class='post_body']").Each(func(i int, selection *goquery.Selection) {
		html, _ := selection.Html()
		lists = append(lists, strings.Split(html, "<br/>")...)
	})

	for _, title := range lists {
		title = strings.TrimSpace(strings.ReplaceAll(title, "\u200b", ""))
		if strings.Contains(title, "<") || strings.Contains(title, ">") || len(title) == 0 {
			continue
		}
		if strings.Contains(title, "、") {
			str := strings.Join(strings.Split(title, "、")[1:], "、")
			resp.Data.News = append(resp.Data.News, str)
		}
		resp.AllData = append(resp.AllData, title)
	}

	if len(resp.AllData) > 0 {
		resp.Data.Title = resp.AllData[0]
		resp.Data.Date = resp.AllData[1]
		resp.Data.Weiyu = resp.AllData[len(resp.AllData)-1]
	}

	return resp, nil
}
