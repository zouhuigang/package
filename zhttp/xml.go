package zhttp

import (
	"encoding/xml"
)

//生成XML
func generateRequestXml(params interface{}) ([]byte, error) {

	data, err := xml.MarshalIndent(&params, "", " ")
	if err != nil {
		return nil, err
	}

	//string(data)=>strings.NewReader(requestXml)
	return data, nil
}

//发送xml数据,返回的也是xml .<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[mch_id参数格式错误]]></return_msg></xml>
func PostXmlWithUnmarshal(url string, req interface{}, resp interface{}) error {
	var data []byte
	var err error

	switch req.(type) {
	case []byte:
		data = req.([]byte)
	default:
		data, err = generateRequestXml(req)
		if err != nil {
			return err
		}
	}

	//发起请求
	data, err = POSTContentType(url, "text/xml", data)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}

	// fmt.Println("====xml====", string(data))

	err = xml.Unmarshal(data, resp)
	if err != nil {
		return err
	}
	return nil
}
