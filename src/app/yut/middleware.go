package main

import (
	"app/yut/manager/userManager"
	"app/yut/proto/netproto"
	"context"
	l4g "github.com/alecthomas/log4go"
	"github.com/go-chi/chi"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func ApiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_ = r.ParseForm()

		var d string = ""
		if r.Method == "POST" {
			d = r.PostForm.Encode()
		}

		l4g.Debug("request ===> url:[%s] method:[%s] data:[%s]", r.URL.Path, r.Method, d)

		if r.Method != "POST" {
			resp := &netproto.NetNotLoginResponse{}
			resp.Msg = "Not Support Method:" + r.Method
			resp.ResponseError(w)
			return
		}

		if !userManager.GetUsrSessionMgr().UserHasLogin(r) {
			if r.URL.Path != "/api/user/login" &&
				r.URL.Path != "/api/user/logout" &&
				r.URL.Path != "/api/user/upload-image" {
				resp := &netproto.NetNotLoginResponse{}
				resp.Msg = "Not Allow Request, Please Login First"
				resp.ResponseError(w)
				return
			}
		} else {
			// 已经登录 权限合法性判断
			if r.URL.Path != "/api/user/logout" {
				if !userManager.GetUsrSessionMgr().CheckAuth(r) {
					resp := &netproto.NetNoAuthResponse{}
					resp.Msg = r.URL.Path +  ": No Auth, Please Contact To Admin Manager"
					resp.ResponseError(w)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func PubBlogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		blog_id := chi.URLParam(r, "blog_id")

		l4g.Debug("PubBlogMiddleware:", blog_id)
		// redis 缓存获取,缓存存在,直接返回

		id, err := strconv.Atoi(blog_id)
		if err != nil {
			resp := &netproto.NetInvalidParamResponse{}
			resp.Msg = "invalid blog_id type." + err.Error()
			resp.ResponseError(w)
			return
		}
		l4g.Debug("PubBlogMiddleware:", blog_id)

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isValidReferer(referer string) bool {
	reg := regexp.MustCompile(`^http://(.*?):(\d{4})(.*?)$`)
	params := reg.FindStringSubmatch(referer)

	if len(params) >= 3 && params[1] == "localhost" && params[2] == "8086" {
		return true
	}

	return false
}

func StaticMiddleware(prefix string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p

			//if !isValidReferer(r2.Referer()) {
			//	resp := &netproto.NetResponse{}
			//	resp.Msg = r.URL.Path +  " 404 not found"
			//	resp.ResponseError(w)
			//	return
			//}

			//w.Header().Set("Content-Type", "application/json; charset=utf-8")
			//w.Header().Set("Access-Control-Allow-Origin", "*")
			//w.WriteHeader(200)

			next.ServeHTTP(w, r2)
		} else {
			resp := &netproto.NetResponse{}
			resp.Msg = r.URL.Path +  " 404 not found"
			resp.ResponseError(w)
			return
		}
	})
}
