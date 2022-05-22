package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	handler "github.com/mauryasaurav/go-proxy/api/http"
	"github.com/mauryasaurav/go-proxy/api/repozitory"
	"github.com/mauryasaurav/go-proxy/api/usecase"
	"github.com/mauryasaurav/go-proxy/domain/dto"
	"github.com/mauryasaurav/go-proxy/utils/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	dbConn, err := gorm.Open(postgres.Open(constants.DB_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("dbConn not connected")
	}

	dbConn.AutoMigrate(&dto.TenantActivity{}, &dto.Tenant{})

	route := gin.Default()

	// tenant route
	tenantRepo := repozitory.NewTenantRepository(dbConn)
	tenantRoute := route.Group("/tenant")

	tenantUsecase := usecase.NewTenantUsecase(tenantRepo)
	handler.NewTenantHandler(tenantRoute, tenantUsecase)

	// tenant activity route
	tenantActivityRepo := repozitory.NewTenantActivityRepository(dbConn)
	tenantActivityRoute := route.Group("/tenant_activity")

	tenantActivityUsecase := usecase.NewTenantActivityUsecase(tenantActivityRepo)
	handler.NewTenantActivityHandler(tenantActivityRoute, tenantActivityUsecase)

	// proxy route
	proxyRepo := repozitory.NewProxyRepository(dbConn)
	proxyRoute := route.Group("/metrics")

	proxyUsecase := usecase.NewProxyUsecase(proxyRepo)
	handler.NewProxyHandler(proxyRoute, proxyUsecase)

	route.Run(":8000")
}
