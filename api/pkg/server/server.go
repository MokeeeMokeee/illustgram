package server

import (
	"shalust/api/pkg/server/handler"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine, port string) {

	v1 := r.Group("/api")
	{
		v1.GET("/user/:user_id", handler.GetUserIllustratio)
		v1.GET("/user/:user_id/Illustratio", handler.GetUserIllustratio)
		v1.GET("/user/:user_id/commic", handler.GetUserCommic)
		v1.GET("/user/:user_id/graffiti", handler.GetUserGraffiti)
		v1.GET("/user/:user_id/rough", handler.GetUserRough)

		v1.GET("/getIllustratio", handler.GetIllustratio)
		v1.GET("/getGraffiti", handler.GetGraffiti)
		v1.GET("/getCommic", handler.GetCommic)
		v1.GET("/getRough", handler.GetRough)

		v1.GET("/getLikeIllustratio", handler.GetUserLikedIllustratio)
		v1.GET("/getLikeGraffiti", handler.GetUserLikedGraffiti)
		v1.GET("/getLikeCommic", handler.GetUserLikedCommic)
		v1.GET("/getLikeRough", handler.GetUserLikedRough)

		v1.GET("/getbookmarkIllustratio", handler.GetUserBookmarkedIllustratio)
		v1.GET("/getbookmarkGraffiti", handler.GetUserBookmarkedGraffiti)
		v1.GET("/getbookmarkCommic", handler.GetUserBookmarkedCommic)
		v1.GET("/getbookmarkRough", handler.GetUserbookmarkedRough)

		v1.POST("/createUser", handler.CreateUser)

		v1.POST("/postContent", handler.PostContent)

		v1.POST("/postManagement/illustratio", handler.PostIllustratioManagement)
		v1.POST("/postManagement/commic", handler.PostCommicManagement)
		v1.POST("/postManagement/graffiti", handler.PostGraffitiManagement)
		v1.POST("/postManagement/rough", handler.PostRoughManagement)

		v1.POST("/bookmarkContent", handler.BookmarkContent)
		v1.POST("/likeContent", handler.LikeContent)

		v1.PUT("/cahngeDiscloseContent", handler.CahngeDiscloseContent)
		v1.POST("/deleteContnt", handler.DeleteContent)

	}

	r.Run(port)
}
