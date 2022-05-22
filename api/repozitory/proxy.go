package repozitory

import (
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
	"gorm.io/gorm"
)

type proxyRepository struct {
	db *gorm.DB
}

func NewProxyRepository(db *gorm.DB) interfaces.ProxyRepository {
	return &proxyRepository{db: db}
}

func (r *proxyRepository) UpdateProxyBytes(tenantActivity entity.TenantActivity) (*entity.TenantActivity, error) {
	err := r.db.Table(tableName).
		Updates(&tenantActivity).
		Error
	if err != nil {
		return nil, err
	}
	return nil, err
}
