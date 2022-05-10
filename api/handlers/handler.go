package handlers

import (
	"fmt"
	"github.com/husanmusa/pro-book-service-api_gateway/config"
	"github.com/husanmusa/pro-book-service-api_gateway/pkg/response"
	"net/http/httputil"
	"net/url"

	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/gin-gonic/gin"
)

// Handler ...
type Handler struct {
	cfg config.Config
	log logger.LoggerI
}

// New ...
func New(cfg config.Config, log logger.LoggerI) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
	}
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, response.SuccessModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, response.ErrorModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Error:   err,
	})
}

// func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
//   return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DefaultOffset))
// }

// func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
//   return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DefaultLimit))
// }

func (h *Handler) makeProxy(c *gin.Context, proxyURL string, endpoint string) (err error) {
	fmt.Println(proxyURL)
	fmt.Println(endpoint)

	// parse the url
	url, err := url.Parse(proxyURL)
	if err != nil {
		h.log.Error("error in parse addr: %v", logger.Error(err))
		return err
	}

	// create the  reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	c.Request.URL.Host = url.Host
	c.Request.URL.Scheme = url.Scheme
	c.Request.URL.Path = endpoint
	c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
	c.Request.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(c.Writer, c.Request)

	return nil
}
