package core

import (
	dbImpl "github.com/bytedance/kldx-core/db/impl"
	rs "github.com/bytedance/kldx-core/resources"
)

var (
	DB        = dbImpl.NewDB()
	Resources = rs.NewResources()
)
