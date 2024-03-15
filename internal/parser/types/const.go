package types

type ConstType int

const (
	/* ConstBoolean represents nil in evaluation. */
	ConstNull ConstType = iota

	/* ConstBoolean represents type bool in evaluation. */
	ConstBoolean

	/* ConstInt represents type int64 in evaluation. */
	ConstInt

	/* ConstInt represents type uint64 in evaluation. */
	ConstUint

	/* ConstFloat represents type float64 in evaluation. */
	ConstFloat

	/* ConstString represents type string in evaluation. */
	ConstString
)
