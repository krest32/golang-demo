package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
	通过body发送Post请求
*/
func Post(url,contentType,requestBody string) (response string)  {

	resp, err := http.Post(url, contentType, strings.NewReader(requestBody))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return ""
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return ""
	}
	return string(b)
}

