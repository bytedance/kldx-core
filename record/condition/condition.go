package cond

import "code.byted.org/apaas/goapi_core/record/operator"

type A []interface{}
type M map[string]interface{}

func And(exps ...interface{}) M {
	if len(exps) == 0 {
		return nil
	}
	return M{op.And: exps}
}

func Or(exps ...interface{}) M {
	if len(exps) == 0 {
		return nil
	}
	return M{op.Or: exps}
}

func Eq(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Eq, rightValue)
}

func Neq(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Neq, rightValue)
}

func Gt(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Gt, rightValue)
}

func Gte(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Gte, rightValue)
}

func Lt(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Lt, rightValue)
}

func Lte(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Lte, rightValue)
}

func Contain(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Contain, rightValue)
}

func NotContain(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.NotContain, rightValue)
}

func In(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.In, rightValue)
}

func NotIn(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.NotIn, rightValue)
}

func Empty(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.Empty, rightValue)
}

func NotEmpty(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.NotEmpty, rightValue)
}

func HasAnyOf(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.HasAnyOf, rightValue)
}

func HasNoneOf(leftValue string, rightValue interface{}) *Expression {
	return NewExpression(leftValue, op.HasNoneOf, rightValue)
}
