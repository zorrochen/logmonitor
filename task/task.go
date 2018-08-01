package task

func Init() {
	// 监控数据分析
	go TaskMonitorStat()
	// 监控触发
	go TaskMonitorTrigger()
}
