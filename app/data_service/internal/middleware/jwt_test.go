package middleware

import (
	"fmt"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testUser *proto.UserProto = &proto.UserProto{
		Id:          1,
		Nickname:    "banana",
		Password:    "114514",
		Avatar:      "https://baidu.com",
		CreatedTime: time.Now().Unix(),
	}
)

func TestGenerateToken(t *testing.T) {
	t.Log(GenerateToken("alioth", testUser))
}

func TestGetClaimsFromAuthString(t *testing.T) {
	authString := fmt.Sprintf("Token %s", GenerateToken("alioth", testUser))
	t.Logf("auth string: %s", authString)
	// error secret
	t.Run("Error secret", func(t *testing.T) {
		claims, err := GetClaimsFromAuthString(authString, "banana")
		if err != nil {
			t.Log(err)
			assert.Error(t, err)
		} else {
			t.Log(claims)
		}
	})
	// correct secret
	t.Run("Correct secret", func(t *testing.T) {
		claims, err := GetClaimsFromAuthString(authString, "alioth")
		if err != nil {
			t.Log(err)
			assert.Error(t, err)
		} else {
			t.Log(claims)
		}
	})
}
