package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	bxapiReq "github.com/flipped-aurora/gin-vue-admin/server/model/bxapi/request"
)

type BxJWT struct {
	SigningKey []byte
}

var (
	BxTokenExpired     = errors.New("BxToken is expired")
	BxTokenNotValidYet = errors.New("BxToken not active yet")
	BxTokenMalformed   = errors.New("That's not even a bxToken")
	BxTokenInvalid     = errors.New("Couldn't handle this bxToken:")
)

func NewBxJWT() *BxJWT {
	return &BxJWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

func (j *BxJWT) CreateClaims(baseClaims bxapiReq.BaseClaims) bxapiReq.CustomClaims {
	expiresDuration, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := bxapiReq.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: 60 * 60 * 24 * 7, // 7天缓冲时间
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"BxAPI"},                           // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),           // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresDuration)), // 过期时间
			Issuer:    global.GVA_CONFIG.JWT.Issuer,                        // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *BxJWT) CreateToken(claims bxapiReq.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *BxJWT) CreateTokenByOldToken(oldToken string, claims bxapiReq.CustomClaims) (string, error) {
	v, err, _ := global.GVA_Concurrency_Control.Do("BxJWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *BxJWT) ParseToken(tokenString string) (*bxapiReq.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &bxapiReq.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, BxTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, BxTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, BxTokenNotValidYet
			} else {
				return nil, BxTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*bxapiReq.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, BxTokenInvalid

	} else {
		return nil, BxTokenInvalid
	}
}
