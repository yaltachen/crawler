package fetcher

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// Fetch 获取网页
func Fetch(url string) ([]byte, error) {
	var (
		client  *http.Client
		req     *http.Request
		resp    *http.Response
		reader  *bufio.Reader
		encode  encoding.Encoding
		content []byte
		err     error
	)
	// log.Printf("fetching %s\r\n", url)
	client = &http.Client{}
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		log.Printf("new request failed, url: %s, err: %v\r\n", url, err)
		return nil, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	if resp, err = client.Do(req); err != nil {
		log.Printf("request failed, url: %s, err: %v\r\n", url, err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("request failed response code: %d\r\n", resp.StatusCode)
		return nil, err
	}

	defer resp.Body.Close()

	reader = bufio.NewReader(resp.Body)

	if encode, err = determineEncoding(reader); err != nil {
		log.Printf("can not get encode using utf-8 as default, url:%s, err: %v\r\n", url, err)
		err = nil
	}

	if content, err = ioutil.ReadAll(transform.NewReader(reader, encode.NewDecoder())); err != nil {
		log.Printf("read body failed, url: %s, err: %v\r\n", url, err)
		return nil, err
	}

	return content, nil
}

// determineEncoding 自动判断网页编码
func determineEncoding(r *bufio.Reader) (e encoding.Encoding, err error) {
	var (
		content []byte
	)
	if content, err = r.Peek(1024); err != nil {
		return unicode.UTF8, err
	}
	e, _, _ = charset.DetermineEncoding(content, "")
	return e, nil
}
