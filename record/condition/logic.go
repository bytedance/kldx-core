package cond

import "github.com/kldx/core/record/operator"

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
