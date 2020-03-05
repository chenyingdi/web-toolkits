package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"time"
)

func GeneToken(args g.Map) (string, error) {
	claim := make(jwt.MapClaims)

	claim["exp"] = time.Now().Add(time.Hour * time.Duration(24*3)).Unix()
	claim["iat"] = time.Now().Unix()

	if args != nil {
		for k, v := range args {
			claim[k] = v
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = claim

	return token.SignedString([]byte(g.Cfg().GetString("app.AppSecret")))

}
