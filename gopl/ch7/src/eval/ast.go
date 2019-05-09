package eval

//Var标识一种类型的变量  eg  x
type Var string
//literal标识一种类型的变量  eg 3.141
type literal float64

//简写一元表达式付 eg -x
type unary struct{
	op rune  // +，-其中一个
	x Expr
}

//二元运算符表达式
type binary struct{
	op rune // +,-,*,/
	x,y Expr
}

//函数，eg sinx(x)
type call struct{
	fn string // pow,sin.sqrt
	args []Expr
}


type Expr interface{
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}
