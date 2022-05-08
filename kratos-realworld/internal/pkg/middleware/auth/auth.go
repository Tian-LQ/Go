package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt/v4"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func GenerateToken(secret string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenStr
}

func JWTAuth(secret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				tokenStr := tr.RequestHeader().Get("Authorization")
				spew.Dump(tokenStr)
				auths := strings.SplitN(tokenStr, " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
					return nil, errors.New("jwt token missing")
				}
				token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(secret), nil
				})
				if err != nil {
					return nil, err
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					spew.Dump(claims["username"])
				} else {
					return nil, errors.New("token invalid")
				}
			}
			return handler(ctx, req)
		}
	}
}
