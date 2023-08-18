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

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBAL00QsML/ovZle3Lq3C7QBo9s00ivsLhG2xlamhHOZDrjTGJX4OA
H27qQbDREcYXpUt5JqOt+KzB4MA/vUKCbT0CAwEAAQJBAINbkS5RWXxGqCzcRj6S
AkM1qxJWmRI7rwpmrqWPLYxKiS1i/i3bwSA3H+NODWIk1p2BWtycWzx5s3cNLn4b
gIECIQD6WuNzXxZHRIxRJQDRyEeWLsrRv9nkZJXHde78DoIZuQIhAMF4ZOgQX2hV
+y9YZmca2tW7etwGPmVjFWQd6JFtjyGlAiBFR9GZo76uijGqYusPIrVswhYuZUEP
CybHw8MWzY0DQQIgc4DDDWCo9QtP+MYX7Lo1p6BUCwOXQMRUwv6wGBKGfxkCIQDn
EKF3Ee6bnLT5DMfrnGY20RNg1Yes+14KkEyYsx0++Q==
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN RSA PUBLIC KEY-----
MEgCQQC9NELDC/6L2ZXty6twu0AaPbNNIr7C4RtsZWpoRzmQ640xiV+DgB9u6kGw
0RHGF6VLeSajrfisweDAP71Cgm09AgMBAAE=
-----END RSA PUBLIC KEY-----`

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(strLen int) string {
	randBytes := make([]rune, strLen)
	for i := range randBytes {
		randBytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(randBytes)
}

type Claims struct {
	userId           int64
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
		userId: userId,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "ApiGateWay",                                    // 签发者
			Subject:   userName,                                        // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // jwt ID, 类似于盐值
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

func getCaims(tokenString string) (*Claims, error) {
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
	claims, err := getCaims(tokenString)
	if err != nil {
		return -1
	}
	return claims.userId
}
