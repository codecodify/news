package api

import (
	"encoding/json"
	"github.com/codecodify/news/lib"
	"github.com/codecodify/news/lib/impl"
	"log"
	"net/http"
	"strconv"
	"time"
)

var crawlers map[string]lib.ICrawler

func init() {
	crawlers = map[string]lib.ICrawler{
		// 知乎
		"zhihu": &impl.Zhihu{},
		// 网易163
		"163": &impl.Wangyi{},
	}
}

func Crawler(w http.ResponseWriter, r *http.Request) {
	var resp lib.Response
	origin := r.URL.Query().Get("origin")
	index, err := strconv.Atoi(r.URL.Query().Get("index"))
	source, ok := crawlers[origin]
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	result := func() {
		if err != nil {
			log.Printf("error: %s", err)
		}
		resp.Time = time.Now().Format("2006-01-02 15:04:05")
		resp.Suc = true
		b, _ := json.Marshal(resp)
		_, _ = w.Write(b)
	}

	if !ok {
		result()
		return
	}
	resp, err = source.Get(index)
	result()

}
