package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

// writeReportToFile 将报告写入文件
func writeReportToFile(report, filename string) error {
	return os.WriteFile(filename, []byte(report), 0644)
}

// formatDuration 格式化持续时间
func formatDuration(d float64) string {
	if d < 1.0 {
		return fmt.Sprintf("%.2fms", d*1000)
	} else if d < 60.0 {
		return fmt.Sprintf("%.2fs", d)
	} else {
		minutes := int(d / 60)
		seconds := d - float64(minutes*60)
		return fmt.Sprintf("%dm %.2fs", minutes, seconds)
	}
}

// timeToScore 耗时转为分数
func timeToScore(duration time.Duration) float64 {
	// 根据实际测试结果调整系数
	const k = 40000000.0 // 4000万系数

	if duration.Nanoseconds() == 0 {
		return math.Inf(1) // 如果时间为0，返回无穷大
	}

	// 将纳秒转换为毫秒来计算
	msElapsed := float64(duration.Nanoseconds()) / 1000000.0
	return k / msElapsed
}
