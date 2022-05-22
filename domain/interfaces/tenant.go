package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/dto"
	"github.com/mauryasaurav/go-proxy/domain/entity"
)

type TenantUsecase interface {
	CreateTenantHandler(ctx *gin.Context, request *dto.Tenant) error
	GetTenantHandler(ctx *gin.Context, id int64) (*entity.TenantSchema, error)
	GetTenantsHandler(ctx *gin.Context) ([]*entity.Tenant, error)
}

type TenantRepository interface {
	CreateTenant(tenant entity.TenantSchema) (*entity.TenantSchema, error)
	GetTenant(tenantId int64) (*entity.TenantSchema, error)
	GetTenants() ([]*entity.Tenant, error)
}
