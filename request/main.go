package request

import (
	"ace/model"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	Url    string
	Header map[string]string
	Params map[string]string
	Body   map[string]any
}

var Client HttpClient

func (c *HttpClient) Get(uri string) ([]byte, error) {
	params := url.Values{}
	u, err := url.Parse(c.Url + uri)
	if err != nil {
		return nil, err
	}
	for s, a := range c.Params {
		params.Set(s, a)
	}
	u.RawQuery = params.Encode()
	urlPath := u.String()

	request, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		return nil, err
	}

	for s, a := range c.Header {
		request.Header.Add(s, a)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	c.Reset()
	return body, nil
}

func (c *HttpClient) Post(uri string) ([]byte, error) {
	da := make([]string, 0)
	for s, s2 := range c.Body {
		da = append(da, fmt.Sprintf("%v=%v", s, s2))
	}
	payload := strings.Join(da, "&")

	u, err := url.Parse(c.Url + uri)
	if err != nil {
		return nil, err
	}

	c.Header["Content-Type"] = "application/x-www-form-urlencoded"

	request, err := http.NewRequest("POST", u.String(), strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	for s, a := range c.Header {
		request.Header.Set(s, a)
	}

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	c.Reset()
	return body, nil
}

func (c *HttpClient) Reset() {
	c.Header = map[string]string{}
	c.Body = map[string]any{}
	c.Params = map[string]string{}
}

func Setup(r model.Request) {
	header := make(map[string]string)
	Client = HttpClient{
		Url:    r.Url,
		Header: header,
	}
}
