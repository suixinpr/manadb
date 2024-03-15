package types

type JoinType int

const (
	CrossJoin JoinType = iota + 1
	NaturalJoin
	InnerJoin
	LeftJoin
	RightJoin
	FullJoin
)
