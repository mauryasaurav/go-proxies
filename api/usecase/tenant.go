package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauryasaurav/go-proxy/domain/dto"
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
)

type tenantUsecase struct {
	tenantRepo interfaces.TenantRepository
}

func NewTenantUsecase(tenantRepo interfaces.TenantRepository) interfaces.TenantUsecase {
	return &tenantUsecase{
		tenantRepo: tenantRepo,
	}
}

func (u *tenantUsecase) CreateTenantHandler(ctx *gin.Context, req *dto.Tenant) error {
	uuid, _ := uuid.NewUUID()

	_, err := u.tenantRepo.CreateTenant(entity.TenantSchema{
		ID:   int64(uuid.ID()),
		NAME: req.NAME,
	})

	if err != nil {
		return err
	}
	return nil
}

func (u *tenantUsecase) GetTenantsHandler(ctx *gin.Context) ([]*entity.Tenant, error) {
	tenant, err := u.tenantRepo.GetTenants()

	if err != nil {
		return nil, err
	}
	return tenant, nil
}

func (u *tenantUsecase) GetTenantHandler(ctx *gin.Context, tenantId int64) (*entity.TenantSchema, error) {

	tenant, err := u.tenantRepo.GetTenant(tenantId)

	if err != nil {
		return nil, err
	}
	return tenant, nil
}
