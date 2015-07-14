package glasso

import (
	"testing"

	"github.com/bmizerany/assert"
	. "github.com/timkaye11/glasso/util"
)

// make sure everything is constructed OK
func TestMakeDF(t *testing.T) {
	// make dataframe
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}
	labels := []string{"a", "", "c"}

	df, err := DF(data, labels)
	assert.Equal(t, err, nil)
	assert.Equal(t, df.rows, 3)
	assert.Equal(t, df.cols, 3)
	assert.Equal(t, df.colToIdx["a"], 0)
}

// test transform dataframe function
func TestTransformDF(t *testing.T) {
	// a, b, c |
	//---------|
	// 1, 2, 3 |
	// 4, 5, 6 |
	// 7, 8, 9 |
	//---------/
	// sum : 12, 15, 18

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	labels := []string{"a", "b", "c"}
	df, _ := DF(data, labels)

	// silyl transformation function
	add1 := func(x float64) float64 {
		return x + 1
	}

	// add 1 to every number in cols "a" & "c"
	df.Transform(add1, "a", "c")
	newA, _ := df.GetCol("a")
	newB, _ := df.GetCol("b")
	newC, _ := df.GetCol("c")

	assert.Equal(t, Sum(newA), 15.0)
	assert.Equal(t, Sum(newB), 15.0) // shouldn't change
	assert.Equal(t, Sum(newC), 21.0)
}

// test apply function
func TestApplyDF(t *testing.T) {
	// make data
	data := []float64{
		1, 2, 3,
		2, 3, 1,
		3, 1, 2,
	}
	labels := []string{"a", "b", "c"}
	df, _ := DF(data, labels)

	colProds := df.Apply(Mult, true, 0, 1, 2)
	rowProds := df.Apply(Mult, false, 0, 1, 2)

	// all the products should equal 6
	for i := 0; i < 3; i++ {
		assert.T(t, colProds[i] == rowProds[i])
	}

}