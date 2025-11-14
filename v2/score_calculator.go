package main

import (
	"fmt"
	"math"
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

// normalizeScore 标准化分数，使用对数缩放减少极端值影响
func (sc *ScoreCalculator) normalizeScore(score float64) float64 {
	// 检查是否为无穷大或NaN
	if math.IsInf(score, 0) || math.IsNaN(score) {
		return 10000.0 // 为无穷大值设置一个合理的上限
	}
	if score <= 0 {
		return 0.1 // 避免log(0)
	}

	// 限制最大值，避免极端情况
	maxScore := 10000000.0 // 1000万分上限
	if score > maxScore {
		score = maxScore
	}

	// 使用对数缩放，让分数更稳定
	// log10(score/1000 + 1) * 1000 可以将大数值压缩到合理范围
	return math.Log10(score/1000.0+1.0) * 1000.0
}

// CalculateTotal 计算综合得分
func (sc *ScoreCalculator) CalculateTotal(results []BenchmarkResult) float64 {
	categoryScores := make(map[string][]float64)
	// 按类别分组得分
	for _, result := range results {
		normalizedScore := sc.normalizeScore(result.Score)
		categoryScores[result.Category] = append(categoryScores[result.Category], normalizedScore)
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
			normalizedScore := sc.normalizeScore(result.Score)
			scores = append(scores, normalizedScore)
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
	report.WriteString(fmt.Sprintf("综合得分: %.2f\n", totalScore))
	report.WriteString("\n")
	// 分类得分
	report.WriteString("分类得分:\n")
	categories := []string{"计算密集型", "内存性能", "并发性能", "加密性能", "浮点性能", "压缩性能", "算法性能"}
	for _, category := range categories {
		score := sc.GetCategoryScore(results, category)
		weight := sc.categoryWeights[category] * 100
		report.WriteString(fmt.Sprintf("  %-12s: %8.2f (权重: %.0f%%)\n", category, score, weight))
	}
	report.WriteString("\n")
	// 详细结果
	report.WriteString("详细测试结果:\n")
	for _, result := range results {
		normalizedScore := sc.normalizeScore(result.Score)

		// 避免除零错误
		var ratio float64
		if result.SingleRate > 0 {
			ratio = result.MultiRate / result.SingleRate
		} else {
			ratio = 0.0
		}

		report.WriteString(fmt.Sprintf("  %-6s | %-32s | 得分: %8.2f | 单核: %8.2f | 多核: %8.2f | 多核/单核: %.2f\n",
			result.Category, result.Name, normalizedScore, result.SingleRate, result.MultiRate, ratio))
	}
	report.WriteString("\n")

	return report.String()
}

// CompareWithBaseline 与基准对比
func (sc *ScoreCalculator) CompareWithBaseline(currentScore, baselineScore float64) string {
	// 避免除零错误
	if baselineScore <= 0 {
		return "基准分数无效，无法对比"
	}

	improvement := (currentScore - baselineScore) / baselineScore * 100
	if improvement > 0 {
		return fmt.Sprintf("性能提升: %.2f%%", improvement)
	} else {
		return fmt.Sprintf("性能下降: %.2f%%", -improvement)
	}
}
