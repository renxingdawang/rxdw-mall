package paseto

import (
	"github.com/hertz-contrib/paseto"
	"strconv"
)

type TokenGenerator struct {
	paseto.GenTokenFunc
	paseto.ParseFunc
}

func NewTokenGenerator(asymmetricKey string, implicit []byte) (*TokenGenerator, error) {
	signFunc, err := paseto.NewV4SignFunc(asymmetricKey, implicit) //这里使用的是对称密钥
	if err != nil {
		return nil, err
	}
	verifyFunc, err := paseto.NewV4PublicParseFunc(asymmetricKey, implicit)
	if err != nil {
		return nil, err
	}
	return &TokenGenerator{signFunc, verifyFunc}, nil
}

func (g *TokenGenerator) CreateToken(claims *paseto.StandardClaims) (token string, err error) {
	return g.GenTokenFunc(claims, nil, nil)
}
func (g *TokenGenerator) ParseToken(token string) (int32, error) {
	parsedToken, err := g.ParseFunc(token)
	if err != nil {
		return 0, err
	}
	userIDstr, err := parsedToken.GetString("ID")
	if err != nil {
		return 0, err
	}
	userID, err := strconv.ParseInt(userIDstr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(userID), nil
}
