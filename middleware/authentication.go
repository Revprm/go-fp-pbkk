package middleware

import (
	"net/http"
	"strings"

	"github.com/Revprm/go-fp-pbkk/constants"
	"github.com/Revprm/go-fp-pbkk/dto"
	"github.com/Revprm/go-fp-pbkk/service"
	"github.com/Revprm/go-fp-pbkk/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_DENIED_ACCESS, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId, userRole, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set(constants.CTX_KEY_TOKEN, authHeader)
		ctx.Set(constants.CTX_KEY_USER_ID, userId)
		ctx.Set(constants.CTX_KEY_ROLE_NAME, userRole)
		ctx.Next()
	}
}
