package physical

/*
 * 路径的代价
 */
type Cost struct {
	Startup int /* 启动代价 */
	Run     int /* 运行代价 */
	Total   int /* 总代价 */
}

func CalculateCost(startup int, run int) *Cost {
	return &Cost{
		Startup: startup,
		Run:     run,
		Total:   startup + run,
	}
}
