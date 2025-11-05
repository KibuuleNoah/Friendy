package routes

import (
	"Friendy/controllers"
	"os"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)


func Routes(se *core.ServeEvent) error {
	// serves static files from the provided public dir (if exists)
	se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))


	apiCtrlV1 := controllers.ApiControllerV1{}

	// PROTECTED ENDPOINTS V1 (Auth Required)
	friendyV1 := se.Router.Group("/api/friendy/v1").Bind(apis.RequireAuth())
	friendyV1.GET("/all/",apiCtrlV1.AllFriends)

	// FRIEND CRUD ROUTES
	friendV1CRUD := friendyV1.Group("/friend")
	friendV1CRUD.GET("/",apiCtrlV1.ReadFriend)
	friendV1CRUD.POST("/",apiCtrlV1.CreateFriend)
	friendV1CRUD.PUT("/",apiCtrlV1.UpdateFriend)
	friendV1CRUD.DELETE("/",apiCtrlV1.DeleteFriend)


	return se.Next()
}

