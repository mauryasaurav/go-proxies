package repozitory

import (
	"errors"

	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
	"gorm.io/gorm"
)

const tableName string = "tenants"

type tenantRepository struct {
	db *gorm.DB
}

func ToTenantDTO(tenant entity.Tenant) *entity.Tenant {
	return &entity.Tenant{
		ID:   tenant.ID,
		NAME: tenant.NAME,
	}
}

func ToTenantSliceDTO(adminUserSlice []entity.Tenant) []*entity.Tenant {
	adminUsers := make([]*entity.Tenant, 0, len(adminUserSlice))
	for _, adminUser := range adminUserSlice {
		adminUsers = append(adminUsers, ToTenantDTO(adminUser))
	}
	return adminUsers
}

func NewTenantRepository(db *gorm.DB) interfaces.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) CreateTenant(tenant entity.TenantSchema) (*entity.TenantSchema, error) {
	err := r.db.Table(tableName).
		Create(&tenant).
		Error
	if err != nil {
		return nil, err
	}
	return &tenant, err
}

func (r *tenantRepository) GetTenant(tenantID int64) (*entity.TenantSchema, error) {
	var tenant entity.Tenant
	err := r.db.Table(tableName).
		Where("id = ? ", tenantID).
		First(&tenant).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &entity.TenantSchema{
		ID:   tenant.ID,
		NAME: tenant.NAME,
	}, nil
}

func (r *tenantRepository) GetTenants() ([]*entity.Tenant, error) {
	var tenant []entity.Tenant

	err := r.db.Find(&tenant).Error
	if err != nil {
		return nil, err
	}

	return ToTenantSliceDTO(tenant), nil
}
