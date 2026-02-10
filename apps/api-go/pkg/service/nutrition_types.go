package service

import (
	"log/slog"
	"math"
)

type flowName string

const (
	FoodScanFlow flowName = "foodScanFlow"
)

type ScanInput struct {
	ImageBase64 string  `json:"image_base64"`
	Description *string `json:"description,omitempty"`
}

// LogValue implements slog.LogValuer for structured logging
func (s *ScanInput) LogValue() slog.Value {
	attrs := []slog.Attr{
		slog.Int("image_size_bytes", len(s.ImageBase64)),
	}
	if s.Description != nil {
		attrs = append(attrs, slog.String("description", *s.Description))
	}
	return slog.GroupValue(attrs...)
}

// Ingredient represents a single food component with its nutritional data
// Note: Macro fields are inlined to avoid Genkit schema bug with repeated types
type Ingredient struct {
	Name        string  `json:"name"`
	WeightGrams int     `json:"weight_grams"`
	Calories    int     `json:"calories"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Fat         float64 `json:"fat"`
	Fiber       float64 `json:"fiber"`
}

// LogValue implements slog.LogValuer for structured logging
func (i *Ingredient) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", i.Name),
		slog.Int("weight_grams", i.WeightGrams),
		slog.Int("calories", i.Calories),
		slog.Float64("protein", i.Protein),
		slog.Float64("carbs", i.Carbs),
		slog.Float64("fat", i.Fat),
		slog.Float64("fiber", i.Fiber),
	)
}

// ScanOutput represents the AI scan result with per-ingredient data.
// Total macros and weight are computed from ingredients, not returned by AI.
type ScanOutput struct {
	IsFood         bool         `json:"is_food"`         // Whether image contains food
	DetectedObject string       `json:"detected_object"` // What was detected (for logging)
	FoodName       string       `json:"food_name"`
	Confidence     float64      `json:"confidence"`
	Ingredients    []Ingredient `json:"ingredients"`
}

// TotalMacros computes total macros by summing all ingredient macros
func (s *ScanOutput) TotalMacros() MacroData {
	var totals MacroData
	for _, ing := range s.Ingredients {
		totals.Calories += ing.Calories
		totals.Protein += ing.Protein
		totals.Carbs += ing.Carbs
		totals.Fat += ing.Fat
		totals.Fiber += ing.Fiber
	}
	// Round floats to 1 decimal place for clean output
	totals.Protein = math.Round(totals.Protein*10) / 10
	totals.Carbs = math.Round(totals.Carbs*10) / 10
	totals.Fat = math.Round(totals.Fat*10) / 10
	totals.Fiber = math.Round(totals.Fiber*10) / 10
	return totals
}

// TotalWeight computes total weight by summing all ingredient weights
func (s *ScanOutput) TotalWeight() int {
	total := 0
	for _, ing := range s.Ingredients {
		total += ing.WeightGrams
	}
	return total
}

// LogValue implements slog.LogValuer for structured logging
func (s *ScanOutput) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Bool("is_food", s.IsFood),
		slog.String("detected_object", s.DetectedObject),
		slog.String("food_name", s.FoodName),
		slog.Float64("confidence", s.Confidence),
		slog.Int("total_weight", s.TotalWeight()),
		slog.Any("total_macros", s.TotalMacros()),
		slog.Any("ingredients", s.Ingredients),
	)
}

type MacroData struct {
	Calories int     `json:"calories"`
	Protein  float64 `json:"protein"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Fiber    float64 `json:"fiber"`
}

// LogValue implements slog.LogValuer for structured logging
func (m *MacroData) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("calories", m.Calories),
		slog.Float64("protein", m.Protein),
		slog.Float64("carbs", m.Carbs),
		slog.Float64("fat", m.Fat),
		slog.Float64("fiber", m.Fiber),
	)
}
