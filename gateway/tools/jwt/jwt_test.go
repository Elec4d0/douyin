package jwt

import (
	"fmt"
	"testing"
)

var Token string

func TestGenerateToken(t *testing.T) {
	fmt.Println("开始测试——————生成密钥—————————————————")
	var userId int64 = 1000004
	fmt.Println(userId)
	userName := "我是0"
	Token, _ = GenerateToken(userId, userName)
	fmt.Println(Token)
	fmt.Println("结束测试——————生成密钥—————————————")
}

/*
func TestParseToken(t *testing.T) {
	fmt.Println(Token)
	fmt.Println("开始测试——————解析密钥—————————————————")
	for i := 0; i < 1000000; i++ {
		ParseToken(Token)
	}
	//fmt.Println(userId)
	fmt.Println("结束测试—————解析密钥—————————————")
}
*/

func TestGetCaims(t *testing.T) {
	fmt.Println("开始测试——————解析密钥—————————————————")
	uid := ParseToken(Token)
	//fmt.Println(cam.RegisteredClaims.Subject)
	fmt.Println(uid)
	testT := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJSZWdpc3RlcmVkQ2xhaW1zIjp7ImlzcyI6IkFwaUdhdGVXYXkiLCJzdWIiOiLljZflroHljIXlj7jku6TlrpjmlrkiLCJhdWQiOlsiQW5kcm9pZF9BUFAiLCJJT1NfQVBQIl0sImV4cCI6MTY5MjUzMjM3NiwibmJmIjoxNjkyNTI4Nzc2LCJpYXQiOjE2OTI1Mjg3NzYsImp0aSI6ImtNdUk2MEpNWGIifX0.XtfFudLdUN6Q0pVLt1k8JAcVGZjG570AtiE50aPwB6drFllUCFzvB-jy6zWGNNTxobKdKqXf3q8SrgklBRLy6w"
	auid := ParseToken(testT)
	fmt.Println(auid)
	fmt.Println("结束测试—————解析密钥—————————————")
}
