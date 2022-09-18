package middleware

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/enum"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/errs"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/request"
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/response"
	"github.com/dollarkillerx/go-proxy-manager/internal/utils"
	"github.com/dollarkillerx/jwt"
	"github.com/gin-gonic/gin"

	"log"
)

func UAAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			response.Return(ctx, errs.PleaseSignIn)
			return
		}
		log.Println(tokenStr)

		token, err := jwt.TokenFormatString(tokenStr)
		if err != nil {
			log.Println(err)
			response.Return(ctx, errs.PleaseSignIn)
			return
		}

		err = utils.JWT.VerificationSignature(token)
		if err != nil {
			log.Println(err)
			response.Return(ctx, errs.PleaseSignIn)
			return
		}

		var tk request.AuthJWT
		err = token.Payload.Unmarshal(&tk)
		if err != nil {
			log.Println(err)
			response.Return(ctx, errs.PleaseSignIn)
			return
		}

		ctx.Set(enum.AuthModel.String(), tk)
	}
}
