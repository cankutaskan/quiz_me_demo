package get

type PerformanceResponse struct {
	Performance          float64 `json:"performance"`
	ComparisonMessage    string  `json:"comparison_message"`
	ComparisonPercentage float64 `json:"comparison_percentage"`
}
