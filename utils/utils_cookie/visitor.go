package utils_cookie

import (
	"ben_gin_study/config"
	"ben_gin_study/utils/utils_password"
	"ben_gin_study/utils/utils_rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
// 访客系统相关
exists, cookie := utils_cookie.GetVisitorCookie(ctx)
if !exists {
	cookie = utils_cookie.MakeCookieValue()
	utils_cookie.SetVisitorCookie(ctx, cookie)
}
fmt.Println(cookie)
**/

// 获取访客 cookie
func GetVisitorCookie(ctx *gin.Context) (exists bool, cookie string) {
	var err error
	cookie, err = ctx.Cookie(config.VISITOR["CookieName"])
	if err != nil {
		exists = false
	} else {
		exists = true
	}
	return
}

// 设置访客 cookie
func SetVisitorCookie(ctx *gin.Context, value string) {
	// maxAge	int	生存期（秒）
	// path	string	有效域
	// domain	string	有效域名
	// secure	bool	是否安全传输 是则只走https
	// httpOnly	bool	是否仅网络使用 是则js无法获取
	ctx.SetCookie(config.VISITOR["CookieName"], value, config.GetVisitor_CookieMaxAge(), "/", config.VISITOR["DOMAIN"], false, true)
}

// 生成 cookie
func MakeCookieValue() string {
	randInt := time.Now().UnixNano()
	randStr := utils_rand.RandString(32, 3)
	return utils_password.MD5Encode(fmt.Sprintf("%v_%v", randInt, randStr))
}
