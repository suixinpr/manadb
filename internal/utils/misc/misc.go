package misc

/* 向上对齐 */
func AlignUp(x int, align int) int {
	return (x + align - 1) & ^(align - 1)
}

/* 向下对齐 */
func AlignDown(x int, align int) int {
	return (x & (^(align - 1)))
}
