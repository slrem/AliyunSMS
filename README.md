# AliyunSMS - 阿里云短信服务

阿里云平台的短信服务

## Installation



    $ go get github.com/slrem/AliyunSMS


## Examples

```Go
func main() {
  //发送短信 bizId 发送流水号 用来查询发送状态
	bizId,err:=SendSms("testId", "testSecret", "阿里云短信测试专用", "15300000001", "SMS_71390007", "{\"customer\":\"test\"}")

  //查询发送状态
	r, _ := QuerySendDetails("testId", "testSecret", "15300000001", bizId, "20171129", "10", "1")
	a, _ := json.Marshal(r)
	log.Println(string(a))
}

```
