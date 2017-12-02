package AliyunSMS

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

func randStr(a int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz-"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < a; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func getTime() string {
	return time.Now().Add(-8 * time.Hour).Format("2006-01-02T15:04:05Z")

}

func alReplace(value string) string {
	a := strings.Replace(value, "+", "%20", -1)
	b := strings.Replace(a, "*", "%2A", -1)
	return strings.Replace(b, "%7E", "~", -1)
}

func specialUrlEncode(value string) string {
	value = url.QueryEscape(value)
	a := strings.Replace(value, "+", "%20", -1)
	b := strings.Replace(a, "*", "%2A", -1)
	return strings.Replace(b, "%7E", "~", -1)
}
func sign(key, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))

}
