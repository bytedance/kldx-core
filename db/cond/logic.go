package cond

import (
	"code.byted.org/apaas/goapi_core/db/op"
)

type Logic struct {
	Type        string
	Expressions []*Expression
	Logics      []*Logic
}

func newLogic(logicType string, expressions []*Expression, logics []*Logic) *Logic {
	return &Logic{
		Type:        logicType,
		Expressions: expressions,
		Logics:      logics,
	}
}

func NewAndLogic(expressions []*Expression, logics []*Logic) *Logic {
	return newLogic(op.And, expressions, logics)
}

func NewOrLogic(expressions []*Expression, logics []*Logic) *Logic {
	return newLogic(op.Or, expressions, logics)
}
