package main

import (
	"fmt"
	"strings"
)

// ScoreCalculator 综合评分系统
type ScoreCalculator struct {
	categoryWeights map[string]float64
}

// NewScoreCalculator 创建评分计算器
func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{
		categoryWeights: map[string]float64{
			"计算密集型": 0.2,
			"内存性能":  0.15,
			"并发性能":  0.15,
			"加密性能":  0.15,
			"浮点性能":  0.15,
			"压缩性能":  0.1,
			"算法性能":  0.1,
		},
	}
}

// CalculateTotal 计算综合得分
func (sc *ScoreCalculator) CalculateTotal(results []BenchmarkResult) float64 {
	categoryScores := make(map[string][]float64)
	// 按类别分组得分
	for _, result := range results {
		categoryScores[result.Category] = append(categoryScores[result.Category], result.Score)
	}
	// 计算加权总分
	totalScore := 0.0
	for category, scores := range categoryScores {
		if len(scores) == 0 {
			continue
		}
		// 取该类别的平均分
		categoryAvg := 0.0
		for _, score := range scores {
			categoryAvg += score
		}
		categoryAvg /= float64(len(scores))
		// 应用权重
		if weight, exists := sc.categoryWeights[category]; exists {
			totalScore += categoryAvg * weight
		}
	}

	return totalScore
}

// GetCategoryScore 获取特定类别得分
func (sc *ScoreCalculator) GetCategoryScore(results []BenchmarkResult, category string) float64 {
	var scores []float64
	for _, result := range results {
		if result.Category == category {
			scores = append(scores, result.Score)
		}
	}
	if len(scores) == 0 {
		return 0
	}
	avg := 0.0
	for _, score := range scores {
		avg += score
	}
	return avg / float64(len(scores))
}

// GenerateReport 生成性能报告
func (sc *ScoreCalculator) GenerateReport(results []BenchmarkResult) string {
	var report strings.Builder
	// 基本信息
	report.WriteString("=== GoHyperPi v2 性能测试报告 ===\n\n")

	// 综合得分
	totalScore := sc.CalculateTotal(results)
	report.WriteString(fmt.Sprintf("综合得分: %.0f\n", totalScore))
	report.WriteString("\n")
	// 分类得分
	report.WriteString("分类得分:\n")
	categories := []string{"计算密集型", "内存性能", "并发性能", "加密性能", "浮点性能", "压缩性能", "算法性能"}
	for _, category := range categories {
		score := sc.GetCategoryScore(results, category)
		weight := sc.categoryWeights[category] * 100
		report.WriteString(fmt.Sprintf("  %-6s: %8.0f (权重: %.0f%%)\n", category, score, weight))
	}
	report.WriteString("\n")
	// 详细结果
	report.WriteString("详细测试结果:\n")
	for _, result := range results {
		report.WriteString(fmt.Sprintf("  %-6s | %-32s | 得分: %8.0f | 单核耗时: %s | 多核耗时: %s | 多核/单核: %.2f\n",
			result.Category, result.Name,
			result.Score,
			formatDuration(result.SingleDuration.Seconds()), formatDuration(result.MultiDuration.Seconds()),
			result.Ratio))
	}
	report.WriteString("\n")

	return report.String()
}
