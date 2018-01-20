package zhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpPostJson(queryurl string, postdata map[string]string) string {
	data, err := json.Marshal(postdata)
	if err != nil {
		return err.Error()
	}
	body := bytes.NewBuffer([]byte(data))

	retstr, err := http.Post(queryurl, "application/json;charset=utf-8", body)
	if err != nil {
		return err.Error()
	}
	result, err := ioutil.ReadAll(retstr.Body)
	retstr.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

func HttpGet(queryurl string) string {
	u, _ := url.Parse(queryurl)
	retstr, err := http.Get(u.String())
	if err != nil {
		return err.Error()
	}
	result, err := ioutil.ReadAll(retstr.Body)
	retstr.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(result)
}

//发起http post请求
func HttpPost(data string, url string) (response string) {
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		fmt.Errorf("网络错误: %s", err)
		return "网络错误"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("数据解析错误：%s", err)
		return "数据解析错误"
	}
	stats := string(body)
	return stats
}
