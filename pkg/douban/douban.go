package douban

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// 豆瓣搜索结果
type DouBanSearchResp struct {
	Count int `json:"count"`
	Start int `json:"start"`
	Total int `json:"total"`
	Books []struct {
		Rating struct {
			Max       int    `json:"max"`
			NumRaters int    `json:"numRaters"`
			Average   string `json:"average"`
			Min       int    `json:"min"`
		} `json:"rating"`
		Subtitle string   `json:"subtitle"`
		Author   []string `json:"author"`
		Pubdate  string   `json:"pubdate"`
		Tags     []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
			Title string `json:"title"`
		} `json:"tags"`
		OriginTitle string        `json:"origin_title"`
		Image       string        `json:"image"`
		Binding     string        `json:"binding"`
		Translator  []interface{} `json:"translator"`
		Catalog     string        `json:"catalog"`
		EbookURL    string        `json:"ebook_url,omitempty"`
		Pages       string        `json:"pages"`
		Images      struct {
			Small  string `json:"small"`
			Large  string `json:"large"`
			Medium string `json:"medium"`
		} `json:"images"`
		Alt         string `json:"alt"`
		ID          string `json:"id"`
		Publisher   string `json:"publisher"`
		Isbn10      string `json:"isbn10"`
		Isbn13      string `json:"isbn13"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		AltTitle    string `json:"alt_title"`
		AuthorIntro string `json:"author_intro"`
		Summary     string `json:"summary"`
		EbookPrice  string `json:"ebook_price,omitempty"`
		Series      struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"series,omitempty"`
		Price string `json:"price"`
	} `json:"books"`
}

func DouBanSearch(keyword string, start int, count int) (*DouBanSearchResp, error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://api.douban.com/v2/book/search", nil)
	q := req.URL.Query()
	q.Add("q", keyword)
	q.Add("apikey", "0ac44ae016490db2204ce0a042db2916")
	q.Add("start", strconv.Itoa(start))
	q.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("User-Agent", "MicroMessenger/")
	req.Header.Add("Referer", "https://servicewechat.com/wx2f9b06c1de1ccfca/91/page-frame.html")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result DouBanSearchResp
	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}
