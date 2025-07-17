package workflow_engine

var ConditionOperatorToFunctionMap = map[string]func(float64, float64) bool{
	"greater_than":          func(a, b float64) bool { return a > b },
	"less_than":             func(a, b float64) bool { return a < b },
	"equals":                func(a, b float64) bool { return a == b },
	"greater_than_or_equal": func(a, b float64) bool { return a >= b },
	"less_than_or_equal":    func(a, b float64) bool { return a <= b },
}

var ConditionOperatorToStringMap = map[string]string{
	"greater_than":          ">",
	"less_than":             "<",
	"equals":                "==",
	"greater_than_or_equal": ">=",
	"less_than_or_equal":    "<=",
}
