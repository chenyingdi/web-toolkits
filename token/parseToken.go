package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
)

func ParseToken(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(g.Cfg().GetString("app.AppSecret")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, InvalidToken
	}

	return token.Claims, nil
}
