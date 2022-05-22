package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
)

type proxyUsecase struct {
	proxyRepo interfaces.ProxyRepository
}

func NewProxyUsecase(proxyRepo interfaces.ProxyRepository) interfaces.ProxyUsecase {
	return &proxyUsecase{
		proxyRepo: proxyRepo,
	}
}

func (u *proxyUsecase) UpdateProxyBytesHandler(ctx *gin.Context, kb int64, error_count int64) error {
	_, err := u.proxyRepo.UpdateProxyBytes(entity.TenantActivity{
		ID:                2,
		BYTES_TRANSFERRED: kb,
		IS_ERROR:          error_count,
	})

	if err != nil {
		return err
	}
	return nil
}
