package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/entity"
)

type ProxyUsecase interface {
	UpdateProxyBytesHandler(ctx *gin.Context, kb int64, error_count int64) error
}

type ProxyRepository interface {
	UpdateProxyBytes(tenantActivity entity.TenantActivity) (*entity.TenantActivity, error)
}
