package auth

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const AnonymousUserID = "anonymous-user"

type InternalRequestMiddleWare struct {
	expectedSecret string
}

func NewInternalRequestMiddleWare(expectedSecret string) InternalRequestMiddleWare {
	return InternalRequestMiddleWare{
		expectedSecret: expectedSecret,
	}

}

func (m InternalRequestMiddleWare) CheckInternalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := m.validate(c,true); err != nil {
			log.Println(err)
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}

func (m InternalRequestMiddleWare) CheckInternalAuthNotRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := m.validate(c,false);err != nil {
			log.Println(err)
		}
		c.Next()
	}
}

func GetRequestUserID(ctx *gin.Context,allowAnonymous bool) (string,error) {
	userID := ctx.GetHeader("X-User-ID")
	if userID != "" {
		return userID,nil
	}
	if allowAnonymous {
		return AnonymousUserID,nil
	}
	return "", errors.New("failed to get user id")
}

func (m InternalRequestMiddleWare) validate(ctx *gin.Context,requireUserID bool) error {
	if os.Getenv("AUTH_ENABLE") != "true" {
		return nil
	}
	internalSecret := ctx.GetHeader("X-Internal-Secret")

	if m.expectedSecret == "" {
		return errors.New("internal service secret is not configured")
	}
	if internalSecret == "" {
		return errors.New("missing internal secret")
	}
	if internalSecret != m.expectedSecret {
		return errors.New("invalid internal secret")
	}
	if requireUserID && ctx.GetHeader("X-User-ID") == "" {
		return errors.New("missing user id")
	}
	return nil
}