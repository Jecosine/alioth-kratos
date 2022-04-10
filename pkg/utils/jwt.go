package utils

import (
	"context"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/go-kratos/kratos/v2/errors"
	jwtv4 "github.com/golang-jwt/jwt/v4"
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

func AddClaimsToContext(ctx context.Context, claims *AuthClaim) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func GetClaimsFromContext(ctx context.Context) (*AuthClaim, error) {
	value := ctx.Value(claimsKey)
	if value != nil {
		attempt, ok := value.(*AuthClaim)
		if ok {
			return attempt, nil
		}
	}
	return nil, errors.Unauthorized("Unauthorized", "Cannot get claims from context")
}
