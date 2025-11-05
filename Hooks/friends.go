package hooks

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
)

func BindFriendsHooks(app core.App) {
    app.OnRecordUpdateRequest("Friends").BindFunc(func(e *core.RecordRequestEvent) error {
        // ignore for superusers
        if e.HasSuperuserAuth() {
            return e.Next()
        }

        // overwrite submitted status
        e.Record.Set("status", "pending")

				fmt.Println("**********",e.Record,"**********")
        // validate role
        // status := e.Record.GetString("status")
        // if (status != "pending" &&
        //     (e.Auth == nil || e.Auth.GetString("role") != "editor")) {
        //     return e.BadRequestError("Only editors can set a status different from pending", nil)
        // }

        return nil
    })
    // app.OnRecordDeleteRequest("Friends").BindFunc(func(e *core.RecordRequestEvent) error {
}

