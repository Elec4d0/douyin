package jwt

import (
	"fmt"
	"testing"
)

var Token string

func TestGenerateToken(t *testing.T) {
	fmt.Println("开始测试——————生成密钥—————————————————")
	var userId int64 = 1000004
	userName := "我是0"
	Token, _ = GenerateToken(userId, userName)
	fmt.Println("结束测试——————生成密钥—————————————")
}

func TestParseToken(t *testing.T) {
	fmt.Println("开始测试——————解析密钥—————————————————")
	userId := ParseToken(Token)
	fmt.Println(userId)
	fmt.Println("结束测试—————解析密钥—————————————")
}

func TestgetCaims(t *testing.T) {
	fmt.Println("开始测试——————解析密钥—————————————————")
	cam := ParseToken(Token)
	fmt.Println(cam)
	fmt.Println("结束测试—————解析密钥—————————————")
}
