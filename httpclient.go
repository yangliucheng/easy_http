package easy_http

import (
	"easy_http/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type RequestGen struct {
	Host    string
	Routers RouterArray
}

func NewRequestGen(host string, router RouterArray) *RequestGen {

	// 添加路由格式检查

	return &RequestGen{
		Host:    host,
		Routers: router,
	}
}

func createHttpClient(timeout time.Duration) *http.Client {

	client := &http.Client{
		Timeout: timeout,
	}

	return client
}

/**
 * the function declare SSH transmit
 */
func CreateHttpsClient() {

}

/**
 * @param handler
 * @param param : param of url ,such as /:name -> /ylc
 * @param body
 * @param fParam : param of table ,such as /getname?xxx
 */
func (requestGen *RequestGen) CreateRequest(handler string, param Mapstring, body io.Reader, header Mapstring, fParam string) (*http.Request, error) {
	// fParam is a form param of url
	router := requestGen.lookUrl(handler)
	path := router.Path
	//设置路由参数
	if len(param) > 0 {
		path = utils.ParaseUrlParam(router.Path, param)
	}

	endpoint := utils.StringJoin(requestGen.Host, path)
	// 设置表单参数
	if !strings.EqualFold(fParam, "") {
		endpoint = utils.StringJoin(endpoint, "?query=", fParam)
	}

	request, err := http.NewRequest(router.Method, endpoint, body)

	if err != nil {
		return request, err
	}

	// 设置请求头
	if header != nil {
		for key, value := range header {
			request.Header.Set(key, value)
		}
	}

	return request, nil
}

func (requestGen *RequestGen) lookUrl(handler string) *Router {

	router := new(Router)

	for _, r := range requestGen.Routers {

		if strings.EqualFold(r.Handler, handler) {
			*router = r
		}
	}

	return router
}

func (requestGen *RequestGen) DoHttpRequest(handler string, param Mapstring, body io.Reader, header Mapstring, fParam string) (*http.Response, error) {

	timeout := 100 * time.Second
	client := createHttpClient(timeout)

	request, err := requestGen.CreateRequest(handler, param, body, header, fParam)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return response, err
	}

	return response, nil
}
