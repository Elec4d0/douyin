package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
)

const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBAPMUShKtwM+cClgVWyoc6ZxkSGy+To30QJ0nUgaSv3t/zh+NHcMq
2qHoJhJ2rZoq1t3qCDKsf5+w47QDiMq5JX8CAwEAAQJBAOCX6Yzyn8jzKxeRu+bg
SfTnL4fSGnDMsnrB3ucV5fhrI/8j8FyeBmMRklmXx0EZl2Hl7G84Y44W0HXuUY9X
daECIQD79hC8T12ddLu7Zyy+wmlFvjhSaU2yV1iEapJbeWCgMQIhAPb5xspShhqI
xnjrg/efzrAI4VVw32JXQ2EdRI0zAOSvAiBjZ7Eymh1VAbkPNqVwnULrQSD3YpRE
yDEkDOexLzHwAQIhALiylNKronRngx3c62UdEuIc0f8mmTgfEFmpHKIH2YwrAiB5
+4m6XfYX8Qlt/aPU/2jPc6YHOPHAs7ZIDQWNOjpPzg==
-----END RSA PRIVATE KEY-----
`

const publicKey = `
-----BEGIN RSA PUBLIC KEY-----
MEgCQQDzFEoSrcDPnApYFVsqHOmcZEhsvk6N9ECdJ1IGkr97f84fjR3DKtqh6CYS
dq2aKtbd6ggyrH+fsOO0A4jKuSV/AgMBAAE=
-----END RSA PUBLIC KEY-----
`

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}

type Claims struct {
	UserID           int64
	RegisteredClaims *jwt.RegisteredClaims
}

func (c Claims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.ExpiresAt, nil
}

func (c Claims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.IssuedAt, nil
}

func (c Claims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.RegisteredClaims.NotBefore, nil
}

func (c Claims) GetIssuer() (string, error) {
	return c.RegisteredClaims.Issuer, nil
}

func (c Claims) GetSubject() (string, error) {
	return c.RegisteredClaims.Subject, nil
}

func (c Claims) GetAudience() (jwt.ClaimStrings, error) {
	return c.RegisteredClaims.Audience, nil
}

func parsePrivateKey(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse key error")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}

func GenerateToken(userId int64, userName string) (string, error) {
	//登录与注册传参，生成Claims对象
	claim := Claims{
		UserID: userId,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "ApiGateWay",                                  // 签发者
			Subject:   userName,                                      // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},    //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now()),                //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                //签发时间
			ID:        randStr(10),                                   // jwt ID, 类似于盐值
		},
	}

	//解析RSA私钥对象
	rsaPriKey, err := parsePrivateKey([]byte(privateKey))
	if err != nil {
		fmt.Println("RSA私钥解析错误")
		return "", err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claim).SignedString(rsaPriKey)
	if err != nil {
		fmt.Println("Claims对象 RSA私钥签名错误")
		return "", err
	}

	return token, err
}

func parsePublicKey(pubKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		fmt.Println("Base64解析公钥错误")
		return nil, errors.New("block nil")
	}
	pubRet, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("X509解析公钥为RSA公钥对象错误")
		return nil, errors.New("x509.ParsePKCS1PublicKey error")
	}

	return pubRet, nil
}

func GetCaims(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		pub, err := parsePublicKey([]byte(publicKey))
		if err != nil {
			fmt.Println("公钥解析对象时错误", err)
			return nil, err
		}
		return pub, nil
	})

	if err != nil {
		fmt.Println("公钥解析token 错误")
		fmt.Println(err)
		return nil, err
	}

	if !token.Valid {
		fmt.Println("Token非法")
		return nil, errors.New("Token invalid")
	}

	//断言，假定解析完成的Claims 是自定义的Claims类型
	Claims, ok := token.Claims.(*Claims)
	if !ok {
		fmt.Println("Claims 接口断言类型错误， 非法的Claims 类型")
		return nil, errors.New("invalid claim type")
	}

	return Claims, nil
}

func ParseToken(tokenString string) (userId int64) {
	claims, err := GetCaims(tokenString)
	if err != nil {
		return -1
	}
	return claims.UserID
}
