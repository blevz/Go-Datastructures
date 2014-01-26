package ds

import (
	"fmt"
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
		ParseAndReturnEquation("x+(x*1)")
		fmt.Println("")
		ParseAndReturnEquation("(x+x)*1")
		fmt.Println("")
		ParseAndReturnEquation("z+(x+3.5)")
		fmt.Println("")
		ParseAndReturnEquation("x^(2-35.2)")
		fmt.Println("")
		ParseAndReturnEquation("1+(2*3)")

	})

}
