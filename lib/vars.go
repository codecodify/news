package lib

// Response 声明响应
type Response struct {
	Suc  bool   `json:"suc"`
	Time string `json:"time"`
	Data struct {
		Title string   `json:"title"`
		Date  string   `json:"date"`
		News  []string `json:"news"`
		Weiyu string   `json:"weiyu"`
	} `json:"data"`
	AllData []string `json:"all_data"`
}

// ICrawler 抓取功能抽象接口
type ICrawler interface {
	Get(int) (Response, error)
}

var Sources = map[string]ICrawler{}
