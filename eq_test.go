package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEquations(t *testing.T) {

	Convey("Test Basic creation", t, func() {
		a := MakeBasicEquation_Value(5)
		So(a.Eval(), ShouldEqual, 5)
		b := MakeBasicEquation_Variable("x")
		//Because variables are always zero until set otherwise
		So(b.Eval(), ShouldEqual, 0)
		c := a.Add(b)

		So(c.Eval(), ShouldEqual, 5)
		c.SetVal("x", 5)
		So(c.Eval(), ShouldEqual, 10)
		c.SetVal("x", 5.5)
		So(c.Eval(), ShouldEqual, 10.5)
		c.SetVal("x", 1)

		d := MakeBasicEquation_Value(20)
		e := c.Mul(d)
		e.SetVal("x", 1)
		So(e.Eval(), ShouldEqual, 120)
		f := c.Mul(MakeBasicEquation_Variable("y"))
		So(f.Eval(), ShouldEqual, 0)
	})

	Convey("Test Equation Parsing", t, func() {
		a := ParseAndReturnEquation("x")
		b := ParseAndReturnEquation("a*x^2+b*x+c")
		b.SetVal("a", 1)
		b.SetVal("b", 1)
		b.SetVal("c", 1)
		b.SetVal("x", 1)
		a.SetVal("x", 1)
		So(a.Eval(), ShouldEqual, 1)
		So(b.Eval(), ShouldEqual, 3)
		a.SetVal("x", 2)
		b.SetVal("x", 2)
		So(a.Eval(), ShouldEqual, 2)
		So(b.Eval(), ShouldEqual, 4+2+1)
		a.SetVal("x", 3)
		b.SetVal("x", 3)
		So(a.Eval(), ShouldEqual, 3)
		So(b.Eval(), ShouldEqual, 9+3+1)
		a.SetVal("x", 4)
		b.SetVal("x", 4)
		So(a.Eval(), ShouldEqual, 4)
		So(b.Eval(), ShouldEqual, 16+4+1)

	})
	/*
		Convey("Test Equation Parsing", t, func() {
			e := ParseAndReturnEquation("z+(x*y)")
			e.Print()
			e.parseBTree.PrettyPrint()
			fmt.Println("(x+x)/2")
			ParseAndReturnEquation("(x+x)/2").parseBTree.PrettyPrint()
			fmt.Println("x+x*2")
			ParseAndReturnEquation("x+x*2").parseBTree.PrettyPrint()
			fmt.Println("")
			ParseAndReturnEquation("z+(x+3.5)")
			fmt.Println("")
			ParseAndReturnEquation("x^(2-35.2)")
			fmt.Println("")
			ParseAndReturnEquation("1+(2*3)")

			fmt.Println("1+2/3-4*5")
			e2 := ParseAndReturnEquation("1+2/3-4*5")
			e2.Print()
			e2.parseBTree.PrettyPrint()
		})*/

}
