package request

import (
	"ace/model"
	"ace/pkg"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HttpClient struct {
	ContentGeneration string
	Analytics         string
	Header            map[string]string
	Params            map[string]string
	Body              map[string]any
}

var Client HttpClient

func (c *HttpClient) Get(uri string) ([]byte, error) {
	params := url.Values{}
	u, err := url.Parse(c.ContentGeneration + uri)
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

func (c *HttpClient) Post(str string, t bool) ([]byte, error) {
	var request *http.Request
	if t {
		f := fmt.Sprintf("%v/%v", pkg.Upload.Path, str)
		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", str)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
		err = writer.Close()
		if err != nil {
			return nil, err
		}
		u, err := url.Parse(fmt.Sprintf("%v%v", c.Analytics, "/upload"))
		if err != nil {
			return nil, err
		}
		request, err = http.NewRequest("POST", u.String(), body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		values := url.Values{}
		for s, s2 := range c.Body {
			values.Add(s, fmt.Sprintf("%v", s2))
		}
		payload := values.Encode()
		u, err := url.Parse(c.ContentGeneration + str)
		if err != nil {
			return nil, err
		}

		if _, ok := c.Header["Content-Type"]; !ok {
			c.Header["Content-Type"] = "application/x-www-form-urlencoded"
		}

		request, err = http.NewRequest("POST", u.String(), strings.NewReader(payload))
		if err != nil {
			return nil, err
		}

		for s, a := range c.Header {
			request.Header.Set(s, a)
		}
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
		ContentGeneration: r.ContentGeneration,
		Analytics:         r.Analytics,
		Header:            header,
	}
}
