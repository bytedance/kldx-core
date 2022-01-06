package core

import (
	dbImpl "code.byted.org/apaas/goapi_core/db/impl"
	rs "code.byted.org/apaas/goapi_core/resources"
)

var (
	DB        = dbImpl.NewDB()
	Resources = rs.NewResources()
)
