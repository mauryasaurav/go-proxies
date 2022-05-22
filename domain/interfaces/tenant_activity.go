package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/entity"
)

type TenantActivityUsecase interface {
	GetTenantActivityHandler(ctx *gin.Context, tenantActivityId int64) (*entity.TenantActivityResponse, error)
}

type TenantActivityRepository interface {
	GetTenantActivity(tenantActivityId int64) (*entity.TenantActivityResponse, error)
}
