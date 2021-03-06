package router

import (
	"net/http"

	"github.com/FlorentinDUBOIS/achievements/api"
	"github.com/FlorentinDUBOIS/achievements/api/accomplished"
	"github.com/FlorentinDUBOIS/achievements/api/achievements"
	"github.com/FlorentinDUBOIS/achievements/api/ranking"
	"github.com/FlorentinDUBOIS/achievements/api/users"
)

// APIv0 group all handlers of the api at version 0
var APIv0 = Handlers{
	Route{
		Method:  http.MethodGet,
		Path:    "/ranking/",
		Handler: ranking.Get,
	},
	Route{
		Method:  http.MethodGet,
		Path:    "/achievements/",
		Handler: achievements.Get,
	},
	CRUD{
		Path: "/achievements",
		Create: Route{
			Path:    "/",
			Handler: achievements.Create,
		},
		Read: Route{
			Path:    "/:id",
			Handler: achievements.GetOne,
		},
		Update: Route{
			Path:    "/:id",
			Handler: achievements.Update,
		},
		Delete: Route{
			Path:    "/:id",
			Handler: achievements.Delete,
		},
	},
	Route{
		Method:  http.MethodGet,
		Path:    "/users/",
		Handler: users.Get,
	},
	CRUD{
		Path: "/users",
		Create: Route{
			Path:    "/",
			Handler: users.Create,
		},
		Read: Route{
			Path:    "/:id",
			Handler: users.GetOne,
		},
		Update: Route{
			Path:    "/:id",
			Handler: users.Update,
		},
		Delete: Route{
			Path:    "/:id",
			Handler: users.Delete,
		},
	},
	Route{
		Method:  http.MethodGet,
		Path:    "/users/:uid/achievements/",
		Handler: accomplished.Get,
	},
	CRUD{
		Path: "/users/:uid/achievements",
		Create: Route{
			Path:    "/:aid",
			Handler: accomplished.SetAccomplished,
		},
		Read: Route{
			Path:    "/:aid",
			Handler: api.NotImplemented,
		},
		Update: Route{
			Path:    "/:aid",
			Handler: api.NotImplemented,
		},
		Delete: Route{
			Path:    "/:aid",
			Handler: accomplished.Delete,
		},
	},
}
