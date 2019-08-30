package zsend

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	S_URL = "http://sms-api.luosimao.com/v1/send.json" //api地址
)

//真实发送短信
func Sms(strmobile string, content string, s_COM string, s_API_KEY string) error {
	content = fmt.Sprintf("%s【%s】", content, s_COM)
	post_data := make(map[string]string) //post_data := map[string]string{}
	post_data["mobile"] = strmobile
	post_data["message"] = content
	return HttpPostH(S_URL, post_data, s_API_KEY)
}

type LuoSiMaoResponse struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

func HttpPostH(queryurl string, postdata map[string]string, s_API_KEY string) error {
	//postdata数据
	//form
	postform := url.Values{} //map[string][]string
	for key, val := range postdata {
		postform.Set(key, val)
	}
	req_new := ioutil.NopCloser(strings.NewReader(postform.Encode())) //把form数据编下码

	//模拟请求
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", queryurl, req_new)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded") //application/x-www-form-urlencoded application/json;charset=utf-8  application/x-www-form-urlencoded; param=value
	//realm := "api:key-9a59f8ecb5201cc2728bb31719bc35e4"
	//reqest.Header.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	reqest.SetBasicAuth("api", s_API_KEY)

	response, err := client.Do(reqest)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	res := LuoSiMaoResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	if res.Error == 0 {
		return nil
	}

	return errors.New(res.Msg)

}
