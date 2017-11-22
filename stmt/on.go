package stmt

import (
	"github.com/ulule/loukoum/types"
)

type On struct {
	Left  Column
	Right Column
}

func NewOn(left, right Column) On {
	return On{
		Left:  left,
		Right: right,
	}
}

// Write expose statement as a SQL query.
func (on On) Write(ctx *types.Context) {
	ctx.Write("ON ")
	ctx.Write(on.Left.Name)
	ctx.Write(" = ")
	ctx.Write(on.Right.Name)
}

// IsEmpty return true if statement is undefined.
func (on On) IsEmpty() bool {
	return on.Left.IsEmpty() || on.Right.IsEmpty()
}