package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

type FixUrl struct {
	HostInfo        string //host，ip+端口格式	123.123.123.123:8080
	FixedHostInfo   string //host，由 scheme://host 组成，不一定带端口，默认情况下会省略	https://123.123.123.123
	IP              string //IP	123.123.123.123
	Port            string //端口	8080
	Path            string //请求 URI ，不携带参数
	Dir             string
	Method          string //请求方式，默认为 GET
	ContentType     string
	Data            string  //请求 POST Data
	PostDataType    string  //请求参数类型
	u               url.URL //url 结构体
	mustIP          string
	ParseWithScheme bool //是否解析的时候自带 scheme
}

type HttpResponse struct {
	RawBody    string //原始响应内容
	StatusCode int    //返回状态吗
	//Utf8Html     string       //UTF-8 编码后的 HTML
	//Title        string       //UTF-8 编码后的标题
	//Size         int          //UTF-8 编码后的 HTML 长度
	//HeaderString HeaderString //响应头
	Cookie   string         //响应 Cookie
	Response *http.Response //原始 HTTP 响应 Response
}

type Dict map[string][]string

func (h Dict) Store(key, value string) {
	textproto.MIMEHeader(h).Add(key, value)
}

type RequestConfig struct {
	URI                      string //请求 URI	/getrecords.php
	Method                   string //请求方式，主要分为 GET、POST 等	GET
	Data                     string //请求携带数据
	Following                int    //跳转次数	3
	FollowRedirect           bool   //是否跳转	true
	DenyFollwRedirectOutHost bool   //是否支持跳转到非 ip:port 的另外的网站
	Header                   Dict   //Header 头
	Timeout                  int    //请求超时时间	15
	VerifyTls                bool   //是否验证 Tls 协议	true
	TrackFunction            func(string, string)
	BasicAuth                string  //认证信息 user:pass 格式	admin:admin
	Proxy                    url.URL //代理
	InvalidResponse          bool    //是否会返回错误的格式
}

// get请求配置
func NewGetRequestConfig(uri string) *RequestConfig {
	cfg := new(RequestConfig)
	cfg.Method = "GET"
	cfg.Header = make(map[string][]string)
	cfg.Header.Store("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0")
	cfg.URI = uri
	return cfg
}

// post请求配置
func NewPostRequestConfig(uri string) *RequestConfig {
	cfg := new(RequestConfig)
	cfg.Method = "POST"
	cfg.Header = make(map[string][]string)
	cfg.Header.Store("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:102.0) Gecko/20100101 Firefox/102.0")
	cfg.URI = uri
	return cfg
}

// 发送http请求
func DoHttpRequest(hostinfo *FixUrl, req *RequestConfig) (*HttpResponse, error) {
	goby_resp := new(HttpResponse)
	client := http.Client{}
	fmt.Println(hostinfo.IP + req.URI)
	http_request, err := http.NewRequest(req.Method, hostinfo.IP+req.URI, strings.NewReader(req.Data))
	for s, i := range req.Header {
		http_request.Header.Add(s, i[0])
	}
	http_resp, err := client.Do(http_request)
	if http_resp != nil {
		defer http_resp.Body.Close()
	}
	body, err := ioutil.ReadAll(http_resp.Body)
	goby_resp.RawBody = string(body)
	goby_resp.StatusCode = http_resp.StatusCode
	goby_resp.Cookie = http_resp.Header.Get("Set-Cookie")
	goby_resp.Response = http_resp
	return goby_resp, err
}
