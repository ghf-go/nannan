package gnet

import (
	"encoding/json"
	"fmt"
	"github.com/ghf-go/nannan/glog"
	"io/ioutil"
	"net/http"
	"net/url"
)

type WxConf struct {
	Appid  string
	Secret string
}
type WxAccessTokenH5 struct {
	WxConf
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
}
type WxUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

func (wx WxConf) H5RedirectURL(redirectUrl string) string {
	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=123#wechat_redirect", wx.Appid, url.QueryEscape(redirectUrl))
}
func (wx WxConf) H5GetAccessTokenByCode(code string) *WxAccessTokenH5 {
	errFormt := "wx->H5GetAccessTokenByCode %s"
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", wx.Appid, wx.Secret, code)
	r, e := http.Get(url)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	if r.StatusCode != 200 {
		glog.Error(errFormt, e.Error())
		return nil
	}
	defer r.Body.Close()
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	ret := &WxAccessTokenH5{WxConf: wx}
	e = json.Unmarshal(b, ret)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	return ret
}

func (wt *WxAccessTokenH5) H5RefreshToken() bool {
	errFormt := "wx->H5RefreshToken %s"
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s", wt.Appid, wt.RefreshToken)
	r, e := http.Get(url)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return false
	}
	if r.StatusCode != 200 {
		glog.Error(errFormt, e.Error())
		return false
	}
	defer r.Body.Close()
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return false
	}
	e = json.Unmarshal(b, wt)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return false
	}
	return true
}

func (wx WxConf) GetServerToken(code string) *WxAccessTokenH5 {
	errFormt := "wx->GetServerToken %s"
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wx.Appid, wx.Secret)
	r, e := http.Get(url)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	if r.StatusCode != 200 {
		glog.Error(errFormt, e.Error())
		return nil
	}
	defer r.Body.Close()
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	ret := &WxAccessTokenH5{WxConf: wx}
	e = json.Unmarshal(b, ret)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	return ret
}

func (wt *WxAccessTokenH5) H5GetUserInfo() *WxUserInfo {
	errFormt := "wx->H5RefreshToken %s"
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", wt.AccessToken, wt.Openid)
	r, e := http.Get(url)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	if r.StatusCode != 200 {
		glog.Error(errFormt, e.Error())
		return nil
	}
	defer r.Body.Close()
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	u := &WxUserInfo{}
	e = json.Unmarshal(b, u)
	if e != nil {
		glog.Error(errFormt, e.Error())
		return nil
	}
	return u
}
