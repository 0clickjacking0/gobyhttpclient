package httpclient

type FixUrl struct {
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
	VerifyTls      bool
	FollowRedirect bool
	Timeout        int
	Header         Dict
	Data           string
}

// get请求配置
func NewGetRequestConfig(uri string) *RequestConfig {
	var cfg *RequestConfig
	return cfg
}

// post请求配置
func NewPostRequestConfig(uri string) *RequestConfig {
	var cfg *RequestConfig
	return cfg
}

// 发送http请求
func DoHttpRequest(hostinfo *FixUrl, req *RequestConfig) (*HttpResponse, error) {
	var err error
	var resp *HttpResponse
	return resp, err
}
