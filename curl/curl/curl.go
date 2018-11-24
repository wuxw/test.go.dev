package curl

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Brower struct {
	cookies []*http.Cookie
	client *http.Client
}

func NewBrowser() *Brower {
	br := & Brower{}
	br.client = &http.Client{}
	br.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			for _, v := range br.GetCookie() {
				req.AddCookie(v)
			}
		}
		return  nil
	}

	return br

}

func (self *Brower) setRequestCookie(req *http.Request)  {
	for _, v := range self.GetCookie() {
		req.AddCookie(v)
	}
}

func (self *Brower) encodeParams(params map[string]string) string   {
	paramsData := url.Values{}
	for k, v := range params {
		paramsData.Set(k, v)
	}
	return paramsData.Encode()
}

/*func (self *Brower) SetProxyUrl(proxyUrl string)  {
	proxy := func(_ *http.Request ) (*url.URL, error) {
		return url.Parse(proxyUrl)
	}
	transport := & http.Transport{Proxy:proxy}
	self.client.Transport = transport
}*/

/*func (self *Brower) AddCookie(cookies []*http.Cookie)  {
	self.cookies = append(self.cookies, cookies...)
}*/

func (self *Brower) GetCookie() ([]*http.Cookie) {
	return self.cookies
}

func (self *Brower) Get(requestUrl string)  ([]byte, int){
	requset , _ := http.NewRequest("GET", requestUrl, nil)
	self.setRequestCookie(requset)
	response , _ := self.client.Do(requset)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	return data, response.StatusCode
}

func (self *Brower) Post(requsetUrl string, params map[string]string, headers map[string]string)   ([]byte){
	postData := self.encodeParams(params)
	requset, _ := http.NewRequest("POST", requsetUrl, strings.NewReader(postData) )
	//这个固定的
	requset.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		requset.Header.Set(k, v)
	}
//	requset.Header.Set("Referer", "http://www.qqt2.ph/auth/signin")
	self.setRequestCookie(requset)

	response , _ := self.client.Do(requset)
	defer response.Body.Close()

	//保存响应的cookie
	rpsCks := response.Cookies()
	self.cookies = append(self.cookies, rpsCks...)

	data, _ := ioutil.ReadAll(response.Body)
	return data
}
