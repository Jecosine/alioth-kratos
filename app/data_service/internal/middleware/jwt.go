package middleware

import (
	"context"
	"fmt"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/go-kratos/kratos/v2/errors"
	m "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var (
	claimsKey struct{}
	//expiredKey struct{}
)

type AuthClaim struct {
	jwtv4.RegisteredClaims
	User    *proto.UserProto
	Expired int64
}

// GenerateToken by received user claims
func GenerateToken(secret string, user *proto.UserProto) string {
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, AuthClaim{
		User:    user,
		Expired: time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// NewJwtAuth is middleware of jwt
func NewJwtAuth(secret string) m.Middleware {
	return func(handler m.Handler) m.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				authString := tr.RequestHeader().Get("Authentication")
				// token start with `Token`
				claims, err := GetClaimsFromAuthString(authString, secret)
				if err != nil {
					return nil, err
				}
				ctx = context.WithValue(ctx, claimsKey, claims)
			}
			return handler(ctx, req)
		}
	}
}

func GetClaimsFromAuthString(authString, secret string) (*AuthClaim, error) {
	splits := strings.SplitN(authString, " ", 2)
	if len(splits) != 2 || !strings.EqualFold(splits[0], "Token") {
		return nil, errors.Unauthorized("Unauthorized", "Invalid Authentication Token")
	}
	token, err := jwtv4.ParseWithClaims(splits[1], &AuthClaim{}, func(token *jwtv4.Token) (interface{}, error) {
		// validate algorithm
		if _, ok := token.Method.(*jwtv4.SigningMethodHMAC); !ok {
			// DEBUG: log if alg not paired
			fmt.Printf("Invalid alg: %s", token.Header["alg"])
			return nil, errors.Unauthorized("Unauthorized", "Unsupported Authentication Method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		if time.Now().Unix() > claims.Expired {
			// expired
			return nil, errors.Unauthorized("Unauthorized", "Token Expired")
		}
		// TODO: user authentication
		if true {
			return claims, nil
		} else {
			return nil, errors.Unauthorized("Unauthorized", "Invalid Authentication Claims")
		}
	} else {
		return nil, errors.Unauthorized("Unauthorized", "Invalid Authentication Token")
	}

}
