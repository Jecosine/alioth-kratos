package middleware

import (
	"fmt"
	"github.com/Jecosine/alioth-kratos/api/proto"
	"github.com/Jecosine/alioth-kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/errors"
	jwtv4 "github.com/golang-jwt/jwt/v4"
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

func GenerateTokenWithTime(ts time.Time, secret string) string {
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, utils.AuthClaim{
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: &jwtv4.NumericDate{Time: ts},
		},
		User:    testUser,
		Expired: ts.Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}
func SimpleGetClaims(secret string, tokenString string) (*jwtv4.Token, error) {
	token, err := jwtv4.ParseWithClaims(tokenString, &utils.AuthClaim{}, func(token *jwtv4.Token) (interface{}, error) {
		// validate algorithm
		if _, ok := token.Method.(*jwtv4.SigningMethodHMAC); !ok {
			// DEBUG: log if alg not paired
			fmt.Printf("Invalid alg: %s", token.Header["alg"])
			return nil, errors.Unauthorized("Unauthorized", "Unsupported Authentication Method")
		}
		return []byte("alioth"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func TestGenerateToken(t *testing.T) {
	t.Log(GenerateToken("alioth", testUser))
	past := time.Unix(time.Now().Unix()-1000, 0)
	future := time.Unix(time.Now().Unix()+100000, 0)
	t.Run("Expired token", func(t *testing.T) {
		tokenString := GenerateTokenWithTime(past, "alioth")
		token, err := SimpleGetClaims("alioth", tokenString)
		if assert.Error(t, err) {
			if assert.Equal(t, err.(*jwtv4.ValidationError).Errors, jwtv4.ValidationErrorExpired) {
				t.Log("expired token successfully detected")
			} else {
				t.Fatalf("unexpect error occurs: %v", err)
			}
		} else {
			t.Logf("token: %v", token)
			t.Fatal("expired token passes validation?")
		}

	})
	t.Run("Valid token", func(t *testing.T) {
		tokenString := GenerateTokenWithTime(future, "alioth")
		t.Log(tokenString)
		_, err := SimpleGetClaims("alioth", tokenString)
		if assert.NoError(t, err) {
			t.Log("valid token pass")
		} else {
			t.Fatalf("valid token but not pass: %v", err)
		}
	})
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
