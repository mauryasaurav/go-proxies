package repozitory

import (
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
	"gorm.io/gorm"
)

const dbTableName string = "tenant_activities"

type tenantActivityRepository struct {
	db *gorm.DB
}

func NewTenantActivityRepository(db *gorm.DB) interfaces.TenantActivityRepository {
	return &tenantActivityRepository{db: db}
}

func (ta *tenantActivityRepository) GetTenantActivity(tenantActivityID int64) (*entity.TenantActivityResponse, error) {
	var tenant entity.TenantActivity
	err := ta.db.Table(dbTableName).
		Where("id = ? ", tenantActivityID).
		First(&tenant).
		Error

	if err != nil {
		return nil, err
	}

	return &entity.TenantActivityResponse{
		IS_ERROR:          tenant.IS_ERROR,
		BYTES_TRANSFERRED: tenant.BYTES_TRANSFERRED,
	}, nil
}
