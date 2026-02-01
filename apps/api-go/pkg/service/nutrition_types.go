package service

import "log/slog"

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

type ScanOutput struct {
	FoodName    string     `json:"food_name"`
	Confidence  float64    `json:"confidence"`
	Macros      *MacroData `json:"macros"`
	ServingSize string     `json:"serving_size"`
}

// LogValue implements slog.LogValuer for structured logging
func (s *ScanOutput) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("food_name", s.FoodName),
		slog.Float64("confidence", s.Confidence),
		slog.String("serving_size", s.ServingSize),
		slog.Any("macros", s.Macros),
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
