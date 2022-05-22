package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/dto"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
)

type tenantHandler struct {
	tenantUsecase interfaces.TenantUsecase
}

func NewTenantHandler(route *gin.RouterGroup, u interfaces.TenantUsecase) {
	handler := tenantHandler{tenantUsecase: u}
	route.POST("", handler.CreateTenant)
	route.GET("/:tenantId", handler.GetTenant)
	route.GET("", handler.GetTenants)
}

func (t *tenantHandler) CreateTenant(ctx *gin.Context) {
	tenant := new(dto.Tenant)
	if err := ctx.Bind(tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := t.tenantUsecase.CreateTenantHandler(ctx, tenant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "tenant created successfully"})
}

func (t *tenantHandler) GetTenant(ctx *gin.Context) {
	s := ctx.Param("tenantId")

	tenantId, err1 := strconv.Atoi(s)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}

	tenant, err := t.tenantUsecase.GetTenantHandler(ctx, int64(tenantId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tenant)
}

func (t *tenantHandler) GetTenants(ctx *gin.Context) {
	req := new(dto.Tenant)
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenants, err := t.tenantUsecase.GetTenantsHandler(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tenants)
}
