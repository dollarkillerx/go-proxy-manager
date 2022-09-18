package utils

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/enum"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/request"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var JWT *jwt.JWT

func InitJWT() {
	JWT = jwt.NewJwt(uuid.New().String())
}

// GetAuthModel GetAuthModel
func GetAuthModel(ctx *gin.Context) request.AuthJWT {
	get, exists := ctx.Get(enum.AuthModel.String())
	if !exists {
		panic("what fuck JWTToken is not exists")
	}

	model, ok := get.(request.AuthJWT)
	if !ok {
		panic("what fuck JWTToken is not exists 2")
	}

	return model
}
