package Router

import (
	"calbmp-back/security"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 解决同源策略问题
	r.Use(security.Cors())

	v1 := r.Group("/api/v1")
	{
		GetDataRouter(v1)
		GetUserRouter(v1)
		GetInputRouter(v1)
		GetResultRouter(v1)
		GetBmpRouter(v1)
		GetHistoryRouter(v1)
		GetVfsmRouter(v1)
	}

	return r
}
