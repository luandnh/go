package auth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"go-project/common/log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/token"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
)

var cacheObj libcache.Cache
var strategy union.Union
var tokenStrategy auth.Strategy

type AuthClaim struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

type AuthMdw struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	alg        string
	method     jwt.SigningMethod
	expiredIn  time.Duration
}

var authMdw *AuthMdw

func NewAuthMdw() {
	publicKeyData, err := os.ReadFile("middleware/auth/key/auth-public.pem")
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyData)
	if err != nil {
		panic(err)
	}
	privateKeyData, err := os.ReadFile("middleware/auth/key/auth-private.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		panic(err)
	}
	cacheObj = libcache.FIFO.New(0)
	cacheObj.SetTTL(time.Minute * 10)
	tokenStrategy = token.New(validateTokenAuth, cacheObj)
	strategy = union.New(tokenStrategy)
	method := jwt.GetSigningMethod("RS256")
	authMdw = &AuthMdw{
		privateKey: privateKey,
		publicKey:  publicKey,
		alg:        "RS256",
		method:     method,
		expiredIn:  time.Minute * 60,
	}
}

func validateTokenAuth(ctx context.Context, r *http.Request, tokenStr string) (auth.Info, time.Time, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return authMdw.publicKey, nil
	})
	if err != nil {
		return nil, time.Time{}, err
	}
	tokenSign, err := token.SigningString()
	if err != nil {
		return nil, time.Time{}, err
	}
	err = authMdw.method.Verify(tokenSign, token.Signature, authMdw.publicKey)
	if err != nil && token.Valid {
		return nil, time.Time{}, fmt.Errorf("error while verifying key: %v", err)
	} else if err == nil && !token.Valid {
		return nil, time.Time{}, fmt.Errorf("invalid key passed validation")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Info(claims)
		return auth.NewUserInfo(claims["id"].(string), claims["id"].(string), nil, nil), time.Now(), nil
	}
	return nil, time.Time{}, errors.New("invalid token")
}

func GenerateJWT(data map[string]interface{}) (string, error) {
	t := jwt.New(authMdw.method)
	tcl := jwt.MapClaims{
		"aud": "luandnh",
		"exp": time.Now().Add(authMdw.expiredIn).Unix(),
	}
	for k, v := range data {
		tcl[k] = v
	}
	t.Claims = tcl
	return t.SignedString(authMdw.privateKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, user, err := strategy.AuthenticateRequest(c.Request)
		if err != nil {
			log.Error(err)
			log.Error("invalid credentials")
			c.JSON(
				http.StatusUnauthorized,
				map[string]interface{}{
					"error": http.StatusText(http.StatusUnauthorized),
				},
			)
			c.Abort()
			return
		}
		c.Set("user", user)
	}
}
