package cond

type Expression struct {
	LeftValue  string      `json:"leftValue"`
	Operator   string      `json:"operator"`
	RightValue interface{} `json:"rightValue"`
}

func NewExpression(leftValue string, operator string, rightValue interface{}) *Expression {
	return &Expression{Operator: operator, LeftValue: leftValue, RightValue: rightValue}
}
