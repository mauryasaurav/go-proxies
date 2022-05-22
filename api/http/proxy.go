package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mauryasaurav/go-proxy/domain/entity"
	"github.com/mauryasaurav/go-proxy/domain/interfaces"
	"github.com/mauryasaurav/go-proxy/utils/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type proxyHandler struct {
	proxyUsecase interfaces.ProxyUsecase
}

func rewriteBody(resp *http.Response) error {
	b, err := ioutil.ReadAll(resp.Body) //Read html

	if err != nil {
		return err
	}
	err = resp.Body.Close()

	if err != nil {
		return err
	}

	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1) // replace html
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	bytes := int64(resp.ContentLength)
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	error_count := int64(0)
	if bytes == 2 {
		error_count = error_count + 1
	}

	tenantActivity := &entity.TenantActivity{
		ID:                2,
		BYTES_TRANSFERRED: bytes,
		IS_ERROR:          error_count,
	}

	dbConn, err := gorm.Open(postgres.Open(constants.DB_URL), &gorm.Config{})

	dbConn.Model(&tenantActivity).Where("id = 2").UpdateColumn("bytes_transferred", gorm.Expr("bytes_transferred + ?", z))
	dbConn.Model(&tenantActivity).Where("id = 2").UpdateColumn("is_error", gorm.Expr("is_error + ?", error_count))

	if err != nil {
		return err
	}
	return err
}

func NewProxyHandler(route *gin.RouterGroup, u interfaces.ProxyUsecase) {
	handler := proxyHandler{proxyUsecase: u}
	route.POST("/*proxyPath", handler.proxy)
	route.GET("/*proxyPath", handler.proxy)
}

func (h *proxyHandler) proxy(ctx *gin.Context) {
	var remote *url.URL
	var err error
	a := ctx.Param("proxyPath")
	i, err := strconv.Atoi(a[1:len(a)])
	v := (reflect.TypeOf(a).Kind())

	// for integer case
	if i != 0 {

		remote, err = url.Parse("https://httpstat.us/i")
		if err != nil {
			log.Println(err)
		}

	}
	if v == reflect.String && i == 0 {

		remote, err = url.Parse("https://jsonplaceholder.typicode.com/a")\
		if err != nil {
			log.Println(err)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ModifyResponse = rewriteBody

	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = ctx.Param("proxyPath")
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
