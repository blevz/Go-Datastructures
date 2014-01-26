package ds

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math"
)

func isOperator(toke interface{}) bool {
	tstr := toke.(token.Token).String()
	return tstr == "+" || tstr == "-" || tstr == "*" || tstr == "/" || tstr == "^"
}

func IsParen(toke interface{}) bool {
	tstr := toke.(token.Token).String()
	return tstr == "(" || tstr == ")"
}

func IsLParen(toke interface{}) bool {
	tstr := toke.(token.Token).String()
	return tstr == "("
}

// For the Variables
func IsValue(toke interface{}) bool {
	t := toke.(token.Token)
	tstr := t.String()
	return t.IsLiteral() && tstr == "IDENT"
}

//None variable value eg 5, 2.5, or -2.5
func IsNumLiteral(toke interface{}) bool {
	t := toke.(token.Token)
	tstr := t.String()
	return t.IsLiteral() && tstr != "IDENT"
}

func IsLeftAssoc(toke interface{}) bool {
	op := toke.(token.Token)
	return op == token.ADD || op == token.SUB || op == token.MUL || op == token.QUO
}

/*
Variables
Binary Operators
Unary Operators
Functions
*/

func GetPrecedence(toke interface{}) int {
	op := toke.(token.Token)
	switch op {
	// "+", "-"
	case token.ADD, token.SUB:
		return 1

	//"*", "/"
	case token.MUL, token.QUO:
		return 3

	// "^"
	case token.XOR:
		return 5

	}
	return 0
}

type BToken struct {
	literal string
	ty      int
}

// An Equation consists of an AST with several types of elements
//

type Equation struct {
	parseTree Tree
	vnames    Set
	vval      map[string]float64
}

func (e *Equation) init() {
	e.parseTree = MakeTree()
	e.vnames = MakeEmptySet()
	e.vval = make(map[string]float64)
}

func MakeBasicEquation_Value(val float64) (e Equation) {
	e.init()
	e.parseTree = MakeTreeWithVal(val)
	return
}

func MakeBasicEquation_Variable(v string) (e Equation) {
	e.init()
	e.parseTree = MakeTreeWithVal(v)
	e.vnames = MakeEmptySet()
	e.vnames.AddElem(v)
	return
}

func (e Equation) Eval() float64 {
	return e.evalHelper(&e.parseTree)

}

func (e Equation) evalHelper(pt *Tree) float64 {
	switch t := pt.val.(type) {
	case int64:
		return float64(t)
	case float64:
		return t
	case string:
		switch t {

		case "^":
			return math.Pow(e.evalHelper(pt.left), e.evalHelper(pt.right))
		case "+":
			return e.evalHelper(pt.left) + e.evalHelper(pt.right)
		case "*":
			return e.evalHelper(pt.left) * e.evalHelper(pt.right)
		case "-":
			return e.evalHelper(pt.left) - e.evalHelper(pt.right)
		case "/":
			return e.evalHelper(pt.left) / e.evalHelper(pt.right)
		default:
			if e.vnames.Exist(t) {
				return e.vval[t]
			} else {
				fmt.Println("unrecognized variable")
			}
		}
	}
	return -1
}

func (e1 *Equation) SetVal(v string, x float64) {
	e1.vval[v] = x
}

//Methods for compounding equations

func (e1 Equation) Add(e2 Equation) (er Equation) {
	er.init()
	er.parseTree = MakeTreeWithSubtrees(&e1.parseTree, &e2.parseTree)
	er.parseTree.val = "+"
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e1 Equation) Sub(e2 Equation) (er Equation) {
	er.init()
	er.parseTree = MakeTreeWithSubtrees(&e1.parseTree, &e2.parseTree)
	er.parseTree.val = "-"
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e1 Equation) Mul(e2 Equation) (er Equation) {
	er.init()
	er.parseTree = MakeTreeWithSubtrees(&e1.parseTree, &e2.parseTree)
	er.parseTree.val = "*"
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e Equation) Print() {
	e.parseTree.inOrderPrint()
}

func (e Equation) GetVal() float64 {
	switch t := e.parseTree.val.(type) {
	case float64:
		return t
	case string:
		return float64(-1)
	}
	return 0
}

func ParseAndReturnEquation(str string) (eq Equation) {
	src := []byte(str)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	outputQueue := MakeQueue()
	outputStack := MakeStack()

	for {

		//While there are tokens to be read
		//Read a token
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		if tok.String() == ";" {
			break
		}

		//If the token is a number or a variable or represents a value
		//Then add it to the output queue.
		if IsValue(tok) {
			fmt.Printf("%s %q is Variable or value\n", tok, lit)
			outputQueue.Push(lit)
		}

		if IsNumLiteral(tok) {
			fmt.Printf("%s %q is Number\n", tok, lit)
			outputQueue.Push(lit)
		}

		//If the token is an operator, o1, then:
		if isOperator(tok) {
			fmt.Println(tok, " is operator")
			//while there is an operator token, o2, at the top of the stack,
			// and either o1 is left-associative and its precedence is equal to that of o2,
			// or o1 has precedence less than that of o2,
			for !outputStack.Empty() &&
				isOperator(outputStack.Top()) &&
				((IsLeftAssoc(tok) &&
					GetPrecedence(tok) == GetPrecedence(outputStack.Top())) ||
					GetPrecedence(tok) < GetPrecedence(outputStack.Top())) {
				//pop o2 off the stack, onto the output queue;
				outputQueue.Push(outputStack.Top())
				outputStack.Pop()
			}
			//push o1 onto the stack.
			outputStack.Push(tok)
		}

		//If the token is a left parenthesis, then push it onto the stack.
		if tok.String() == "(" {
			fmt.Printf("%s %q is lparen\n", tok, lit)
			outputStack.Push(tok)
			//If the token is a right parenthesis:
		} else if tok.String() == ")" {
			fmt.Printf("%s %q is rparen\n", tok, lit)
			//Until the token at the top of the stack is a left parenthesis
			for !outputStack.Empty() && !IsParen(outputStack.Top()) {
				//pop operators off the stack onto the output queue.
				outputQueue.Push(outputStack.Top())
				outputStack.Pop()
			}
			//Pop the left parenthesis from the stack, but not onto the output queue.
			if !outputStack.Empty() && outputStack.Top() != "(" {
				outputStack.Pop()
			}
			//If the token at the top of the stack is a function token, pop it onto the output queue.
			//If the stack runs out without finding a left parenthesis, then there are mismatched parentheses.
		}
	}

	//While there are still operator tokens in the stack:
	for !outputStack.Empty() {
		//If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses.
		if IsParen(outputStack.Top()) {
			panic("parenthesis problem")
		} else {
			//Pop the operator onto the output queue.
			outputQueue.Push(outputStack.Top())
			outputStack.Pop()
		}

	}

	outputQueue.Print()

	return
}
