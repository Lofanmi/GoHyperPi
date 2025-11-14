package main

import (
	"fmt"
	"os"
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
