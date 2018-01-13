package router

import (
	"github.com/FlorentinDUBOIS/achievements/api"
	"github.com/FlorentinDUBOIS/achievements/api/achievements"
)

// APIv0 group all handlers of the api at version 0
var APIv0 = Handlers{
	CRUD{
		Path: "/users",
		Create: Route{
			Path:    "/",
			Handler: api.NotImplemented,
		},
		Read: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
		Update: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
		Delete: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
	},
	CRUD{
		Path: "/achievements",
		Create: Route{
			Path:    "/",
			Handler: achievements.Create,
		},
		Read: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
		Update: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
		Delete: Route{
			Path:    "/:id",
			Handler: api.NotImplemented,
		},
	},
}
