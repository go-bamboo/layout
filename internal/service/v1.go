package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-bamboo/layout/internal/biz"
	"github.com/go-bamboo/pkg/log"
)

type V1Service struct {
	uc *biz.GreeterUsecase
}

func NewV1Service(uc *biz.GreeterUsecase) *V1Service {
	return &V1Service{
		uc: uc,
	}
}

func (s *V1Service) Hello(ctx *gin.Context) {
	ctx.Abort()
	ctx.String(http.StatusOK, "hello")
}

func (s *V1Service) ChainBotWebHook(ctx *gin.Context) {
	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Debug(string(bytes))
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
	return
}
