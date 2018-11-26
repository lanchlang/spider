package main

import (
	"github.com/gocolly/colly"
	"time"
	"strings"
	"github.com/gocolly/colly/extensions"
	"encoding/json"
)

type Posts struct{
	Tumbelog struct{
		Title string `json:"title"`
		Description string `json:"description"`
		Name string `json:"name"`
	} `json:"tumbelog"`
	PostsStart int `json:"posts-start"`
	PostsTotal int `json:"posts-total"`
	PostsType bool `json:"posts-type"`
	Posts []struct{
		Id interface{} `json:"id"`
		Url string `json:"url"`
		Date string `json:"date"`
		Timestamp int64 `json:"unix-timestamp"`
		Type string `json:"type"`
		Width int `json:"width"`
		Height int `json:"height"`
		PhotoUrl1280 string `json:"photo-url-1280"`
		PhotoUrl500 string `json:"photo-url-500"`
		PhotoUrl400 string `json:"photo-url-400"`
		PhotoUrl250 string `json:"photo-url-250"`
		PhotoUrl100 string `json:"photo-url-100"`
		RebloggedRootUrl string `json:"reblogged-root-url"`
		Tags []string `json:"tags"`
		Photos  []struct{
			Width int `json:"width"`
			Height int `json:"height"`
			PhotoUrl1280 string `json:"photo-url-1280"`
			PhotoUrl500 string `json:"photo-url-500"`
			PhotoUrl400 string `json:"photo-url-400"`
			PhotoUrl250 string `json:"photo-url-250"`
			PhotoUrl100 string `json:"photo-url-100"`
			Caption string `json:"caption"`
		} `json:"photos"`
		VideoCaption string `json:"video-caption"`
		VideoSource string `json:"video-source"`
		VideoPlayer string `json:"video-player"`
		VideoPlayer500 string `json:"video-player-500"`
		VideoPlayer250 string `json:"video-player-250"`
	} `json:"posts"`
}

/**
* tumblr
*/

//运行Meitulu
func main() {
	//groutineChan:=make(chan int,20)
	// Instantiate default collector
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*tumblr.*",
		Parallelism: 4,
		Delay:      10*time.Second,
	})
	c.SetRequestTimeout( time.Duration(60)*time.Second)
	extensions.Referrer(c)
	extensions.RandomUserAgent(c)

	c.OnResponse(func(response *colly.Response) {
		body:=string(response.Body)
		startIndex :=strings.Index(body,"{")
		if startIndex >0{
			body=strings.TrimSpace(body)
			endIndex:=strings.LastIndex(body,";")
			body=body[startIndex:endIndex]
			println(body)
			posts:=Posts{}
			err:=json.Unmarshal([]byte(body),&posts)
			if err!=nil{
				println(err.Error())
			}else{
				println(posts.PostsType)
			}
		}

	})
	c.OnError(func(response *colly.Response, e error) {
		println(e.Error())
	})
	c.Visit("http://novinhasmusa.tumblr.com/api/read/json")
	c.Wait()
}

