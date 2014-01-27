package ds

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math"
	"strconv"
)

func strIsOperator(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/" || str == "^"
}

func strIsParen(str string) bool {
	return str == "(" || str == ")"
}

func isOperator(toke interface{}) bool {
	tstr := toke.(token.Token).String()
	return strIsOperator(tstr)
}

func IsParen(toke interface{}) bool {
	tstr := toke.(token.Token).String()
	return strIsParen(tstr)
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
	case token.ADD:
		return 1
	case token.SUB:
		return 1
	//"*", "/"
	case token.MUL:
		return 3
	case token.QUO:
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
	parseBTree BTree
	vnames     Set
	vval       map[string]float64
}

func (e *Equation) init() {
	e.parseBTree = MakeBTree()
	e.vnames = MakeEmptySet()
	e.vval = make(map[string]float64)
}

func MakeBasicEquation_Value(val float64) (e Equation) {
	e.init()
	e.parseBTree = MakeBTreeWithVal(val)
	return
}

func MakeBasicEquation_Variable(v string) (e Equation) {
	e.init()
	e.parseBTree = MakeBTreeWithVal(v)
	e.vnames = MakeEmptySet()
	e.vnames.AddElem(v)
	return
}

func (e Equation) Eval() float64 {
	return e.evalHelper(&e.parseBTree)

}

func (e Equation) evalHelper(pt *BTree) float64 {
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
	er.parseBTree = MakeBTreeWithSubtrees(&e1.parseBTree, &e2.parseBTree, "+")
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e1 Equation) Sub(e2 Equation) (er Equation) {
	er.init()
	er.parseBTree = MakeBTreeWithSubtrees(&e1.parseBTree, &e2.parseBTree, "-")
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e1 Equation) Mul(e2 Equation) (er Equation) {
	er.init()
	er.parseBTree = MakeBTreeWithSubtrees(&e1.parseBTree, &e2.parseBTree, "*")
	er.vnames = e1.vnames.Union(e2.vnames)
	return
}

func (e Equation) Print() {
	e.parseBTree.inOrderPrint()
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
			//fmt.Printf("%s %q is Variable or value\n", tok, lit)
			t := MakeBTreeWithVal(lit)
			outputQueue.Push(&(t))
		} else if IsNumLiteral(tok) {
			//fmt.Printf("%s %q is Number\n", tok, lit)
			v, err := strconv.ParseFloat(lit, 64)
			if err != nil {
				panic("oh no")
			}
			t := MakeBTreeWithVal(v)
			outputQueue.Push(&(t))
			//If the token is an operator, o1, then:
		} else if isOperator(tok) {
			//fmt.Println(tok, " is operator")
			//while there is an operator token, o2, at the top of the stack,
			// and either o1 is left-associative and its precedence is equal to that of o2,
			// or o1 has precedence less than that of o2,
			for !outputStack.Empty() &&
				isOperator(outputStack.Top()) &&
				((IsLeftAssoc(tok) &&
					GetPrecedence(tok) == GetPrecedence(outputStack.Top())) ||
					GetPrecedence(tok) < GetPrecedence(outputStack.Top())) {
				//pop o2 off the stack,
				//Make its children the first two things in the oq
				//And put it onto the output queue;
				popAndEnq(&outputQueue, outputStack.Top().(token.Token).String())
				outputStack.Pop()
			}
			//push o1 onto the stack.
			outputStack.Push(tok)
			//If the token is a left parenthesis, then push it onto the stack.
		} else if tok.String() == "(" {
			//fmt.Printf("%s %q is lparen\n", tok, lit)
			outputStack.Push(tok)
			//If the token is a right parenthesis:
		} else if tok.String() == ")" {
			//fmt.Printf("%s %q is rparen\n", tok, lit)
			//Until the token at the top of the stack is a left parenthesis
			for !outputStack.Empty() && !IsLParen(outputStack.Top()) {
				//pop operators off the stack onto the output queue.
				//pop o2 off the stack,
				//Make its children the first two things in the oq
				//And put it onto the output queue;
				popAndEnq(&outputQueue, outputStack.Top().(token.Token).String())
				outputStack.Pop()
			}
			//Pop the left parenthesis from the stack, but not onto the output queue.
			if !outputStack.Empty() && IsLParen(outputStack.Top()) {
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
			popAndEnq(&outputQueue, outputStack.Top().(token.Token).String())
			outputStack.Pop()
		}

	}
	//	outputQueue.Print()

	eq.init()
	//We can reuse the outputStack because we know its empty after the previous loop
	//Now we need to unroll the queue and make it into a full stack
	for outputQueue.Size() != 0 {
		//Pop variables and values onto the stack
		curBTree := outputQueue.Front().(*BTree)

		switch v := curBTree.val.(type) {
		case string:
			if !strIsOperator(v) {
				//Not a variable
				eq.vnames.AddElem(v)
				outputStack.Push(curBTree)
			} else {
				curBTree.right = outputStack.Top().(*BTree)
				outputStack.Pop()
				curBTree.left = outputStack.Top().(*BTree)
				outputStack.Pop()
				outputStack.Push(curBTree)
			}
		case float64:
			outputStack.Push(curBTree)
		}
		//When you hit an operator, pop two values off the stack
		//The first value is the right hand side
		if outputQueue.Size() != 1 {
			outputQueue.Pop()
		} else {
			eq.parseBTree = *curBTree
			outputQueue.Pop()
		}
	}
	return
}

func popAndEnq(outputQueue *Queue, val string) {
	toAdd := MakeBTreeWithVal(val)
	outputQueue.Push(&toAdd)
}
