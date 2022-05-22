package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
)

type tenantActivityUsecase struct {
	tenantActivityRepo interfaces.TenantActivityRepository
}

func NewTenantActivityUsecase(tenantActivityRepo interfaces.TenantActivityRepository) interfaces.TenantActivityUsecase {
	return &tenantActivityUsecase{
		tenantActivityRepo: tenantActivityRepo,
	}
}

func (u *tenantActivityUsecase) GetTenantActivityHandler(ctx *gin.Context, tenantActivityId int64) (*entity.TenantActivityResponse, error) {

	tenant, err := u.tenantActivityRepo.GetTenantActivity(tenantActivityId)

	if err != nil {
		return nil, err
	}
	return tenant, nil
}
