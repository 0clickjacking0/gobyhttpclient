package httpclient

import "net/url"

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

type HeaderString struct {
}

type HttpResponse struct {
	RawBody      string       //原始响应内容
	Utf8Html     string       //UTF-8 编码后的 HTML
	Title        string       //UTF-8 编码后的标题
	size         int          //UTF-8 编码后的 HTML 长度
	HeaderString HeaderString //响应头
	Cookie       string       //响应 Cookie
	//*http.Response              //原始 HTTP 响应 Response
}

type Dict struct {
	key   string
	value string
}

func (dict *Dict) Store(headname, headvalue string) *Dict {
	dict.key = headname
	dict.value = headvalue
	return dict
}

type RequestConfig struct {
	URI                      string //请求 URI	/getrecords.php
	Method                   string //请求方式，主要分为 GET、POST 等	GET
	Data                     int    //请求携带数据
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
	return cfg
}

// post请求配置
func NewPostRequestConfig(uri string) *RequestConfig {
	cfg := new(RequestConfig)
	return cfg
}

// 发送http请求
func DoHttpRequest(hostinfo *FixUrl, req *RequestConfig) (*HttpResponse, error) {
	var err error
	resp := new(HttpResponse)
	return resp, err
}
