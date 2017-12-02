package AliyunSMS

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var aliyun = "http://dysmsapi.aliyuncs.com/?Signature="

type result struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}

// accessKeyId:阿里云生成的accessKeyId
// accessSecret:阿里云生成的accessSecret
// signName:在阿里云申请的短信签名
// phoneNumbers:发送的手机号码
// templateCode:在阿里云申请的短信模板
// templateParam:消息变量替换 JSON字符串 如{"code":"657893"}
// return bizId:发送流水号
func SendSms(accessKeyId, accessSecret, signName, phoneNumbers, templateCode, templateParam string) (bizId string, err error) {
	m := make(map[string]string)
	m["PhoneNumbers"] = phoneNumbers
	m["SignName"] = signName
	m["TemplateParam"] = templateParam
	m["TemplateCode"] = templateCode
	surl := getUrl("SendSms", accessKeyId, accessSecret, m)
	r, err := http.Get(surl)
	if err != nil {
		return
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var res result
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	if res.Code != "OK" {
		err = errors.New(res.Message)
		return
	}
	bizId = res.BizId
	return
}

type A struct {
	B []SmsSendDetailDTO `json:"SmsSendDetailDTO"`
}
type SmsSendDetailDTO struct {
	OutId        string
	SendDate     string
	SendStatus   int
	ReceiveDate  string
	ErrCode      string
	TemplateCode string
	Content      string
	PhoneNum     string
}
type SendStatusResp struct {
	TotalCount        int
	Code              string
	Message           string
	RequestId         string
	SmsSendDetailDTOs A `json:"SmsSendDetailDTOs"`
}

// accessKeyId:阿里云生成的accessKeyId
// accessSecret:阿里云生成的accessSecret
// phoneNumbers:发送的手机号码
// bizId:发送流水号,从调用发送接口返回值中获取
// sendDate:短信发送日期格式yyyyMMdd,支持最近30天记录查询
// pageSize:页大小Max=50
// currentPage:当前页码
func QuerySendDetails(accessKeyId, accessSecret, phoneNumber, bizId, sendDate, pageSize, currentPage string) (resp SendStatusResp, err error) {
	m := make(map[string]string)
	m["PhoneNumber"] = phoneNumber
	m["BizId"] = bizId
	m["SendDate"] = sendDate
	m["PageSize"] = pageSize
	m["CurrentPage"] = currentPage
	surl := getUrl("QuerySendDetails", accessKeyId, accessSecret, m)
	r, err := http.Get(surl)
	if err != nil {
		return
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	log.Println(string(b))
	err = json.Unmarshal(b, &resp)
	return
}

func getUrl(action, accessKeyId, accessSecret string, m map[string]string) (surl string) {
	data := url.Values{}
	data.Add("SignatureMethod", "HMAC-SHA1")
	data.Add("SignatureNonce", randStr(36))
	data.Add("AccessKeyId", accessKeyId)
	data.Add("SignatureVersion", "1.0")
	data.Add("Timestamp", getTime())
	data.Add("Format", "JSON")

	data.Add("Action", action)
	data.Add("Version", "2017-05-25")
	data.Add("RegionId", "cn-hangzhou")
	for k, v := range m {
		data.Add(k, v)
	}

	data.Del("Signature")
	sortedQueryString := alReplace(data.Encode())
	Signature := getSignature(accessSecret, sortedQueryString)
	surl = aliyun + Signature + "&" + sortedQueryString
	return
}

func getSignature(accessSecret, sortedQueryString string) (Signature string) {
	xxx := "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortedQueryString)
	sign := sign(accessSecret+"&", xxx)
	Signature = specialUrlEncode(sign)
	return
}
