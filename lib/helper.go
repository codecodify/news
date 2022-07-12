package lib

import (
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	defer func() {
		if req != nil && req.Body != nil {
			_ = req.Body.Close()
		}
	}()
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	return ioutil.ReadAll(resp.Body)
}

func GetCorrectionIndex(index, max int) int {
	if index < 0 {
		return 0
	}
	if index > max {
		return max
	}
	return index
}
