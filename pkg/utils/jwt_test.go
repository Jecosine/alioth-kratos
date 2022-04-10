package utils

import (
	"context"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/go-kratos/kratos/v2/errors"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetClaimsFromContext(t *testing.T) {
	var valueKey struct{}
	t.Run("Empty fetch context value", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		value := ctx.Value(valueKey)
		assert.Nil(t, value)
	})
	t.Run("Fetch existed context value", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		ctx = AddClaimsToContext(ctx, &AuthClaim{
			RegisteredClaims: jwtv4.RegisteredClaims{},
			User: &proto.UserProto{
				Id:          1,
				Nickname:    "banana",
				Password:    "1234",
				Avatar:      "https://baidu.com",
				CreatedTime: time.Now().Unix(),
			},
			Expired: time.Now().Unix(),
		})
		claims, err := GetClaimsFromContext(ctx)
		if err != nil {
			assert.True(t, errors.IsUnauthorized(err))
		}
		t.Logf("value: %v", claims.User.Nickname)
	})
}
