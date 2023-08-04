package protos

import (
	"context"
	"github.com/golang-jwt/jwt"
	"time"
	api "token/services/protos/kitex_gen/api"
)

var jwtSecret = []byte("go is the best language in the world!")

type UserJwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	id int64
}

// TokenServiceImpl implements the last service interface defined in the IDL.
type TokenServiceImpl struct{}

// UserAuthentication implements the TokenServiceImpl interface.
func (s *TokenServiceImpl) UserAuthentication(ctx context.Context, req *api.DouyinUserAuthenticationRequest) (resp *api.DouyinUserAuthenticationResponse, err error) {
	resp = new(api.DouyinUserAuthenticationResponse)
	// 鉴权
	token := req.UserToken

	if token == "" {
		resp.UserId = -1
		err = nil
		return
	}

	claims, err := jwt.ParseWithClaims(token, &UserJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		resp.UserId = -1
		return
	}

	if claims == nil {
		resp.UserId = -1
		return
	}

	tokenClaim, ok := claims.Claims.(*UserJwtClaims)

	// 格式正确并有效
	if ok && claims.Valid {
		resp.UserId = tokenClaim.id
	}

	return
}

// UserGetToken implements the TokenServiceImpl interface.
func (s *TokenServiceImpl) UserGetToken(ctx context.Context, req *api.DouyinUserGetTokenRequest) (resp *api.DouyinUserGetTokenResponse, err error) {
	resp = new(api.DouyinUserGetTokenResponse)

	id := req.UserId

	expireTime := time.Now().Add(time.Hour).Unix()

	claims := UserJwtClaims{
		&jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "tiktok",
		},
		id,
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ad salt
	signedToken, err := unsignedToken.SignedString(jwtSecret)

	resp.UserToken = signedToken

	return
}
