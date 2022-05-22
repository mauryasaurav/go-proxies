package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
)

type tenantActivityHandler struct {
	tenantActivityUsecase interfaces.TenantActivityUsecase
}

func NewTenantActivityHandler(route *gin.RouterGroup, u interfaces.TenantActivityUsecase) {
	handler := tenantActivityHandler{tenantActivityUsecase: u}
	route.GET("/:tenantActivityId", handler.GetTenantActivity)
}

func (ta *tenantActivityHandler) GetTenantActivity(ctx *gin.Context) {
	s := ctx.Param("tenantActivityId")

	tenantActivityId, err1 := strconv.Atoi(s)
	if err1 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}

	tenantActivity, err := ta.tenantActivityUsecase.GetTenantActivityHandler(ctx, int64(tenantActivityId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tenantActivity)
}
