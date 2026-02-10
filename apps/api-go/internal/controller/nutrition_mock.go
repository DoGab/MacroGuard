package controller

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// NutritionMockController is a controller for mocking the nutrition handlers
type NutritionMockController struct{}

// NewNutritionMockController creates a new nutrition mock controller
func NewNutritionMockController() *NutritionMockController {
	return &NutritionMockController{}
}

// Register registers the nutrition mock controller with the API
func (c *NutritionMockController) Register(api huma.API) {
	huma.Register(api, huma.Operation{
		Path:        "/api/nutrition/scan",
		Method:      http.MethodPost,
		OperationID: "scan-food",
		Summary:     "Scan food image for nutritional information",
		Description: "Upload a base64-encoded food image and optionally provide a description. Returns detected food name and macro breakdown.",
		Tags:        []string{"nutrition"},
	}, c.ScanHandler)
}

// ScanHandler handles the scan request
func (c *NutritionMockController) ScanHandler(ctx context.Context, input *ScanInput) (*ScanOutput, error) {
	return &ScanOutput{
		Body: &ScanOutputBody{
			FoodName:   "Grilled Chicken Salad",
			Confidence: 0.95,
			Macros: &MacroData{
				Calories: 476,
				Protein:  47,
				Carbs:    11,
				Fat:      27,
				Fiber:    3,
			},
			ServingSize: "400g",
			Ingredients: []IngredientBody{
				{
					Name:        "Grilled Chicken Breast",
					WeightGrams: 150,
					Macros: &MacroData{
						Calories: 248,
						Protein:  38,
						Carbs:    0,
						Fat:      10,
						Fiber:    0,
					},
				},
				{
					Name:        "Mixed Greens",
					WeightGrams: 100,
					Macros: &MacroData{
						Calories: 20,
						Protein:  2,
						Carbs:    3,
						Fat:      0,
						Fiber:    2,
					},
				},
				{
					Name:        "Cherry Tomatoes",
					WeightGrams: 60,
					Macros: &MacroData{
						Calories: 18,
						Protein:  1,
						Carbs:    4,
						Fat:      0,
						Fiber:    1,
					},
				},
				{
					Name:        "Feta Cheese",
					WeightGrams: 40,
					Macros: &MacroData{
						Calories: 105,
						Protein:  6,
						Carbs:    2,
						Fat:      8,
						Fiber:    0,
					},
				},
				{
					Name:        "Olive Oil Dressing",
					WeightGrams: 20,
					Macros: &MacroData{
						Calories: 80,
						Protein:  0,
						Carbs:    1,
						Fat:      9,
						Fiber:    0,
					},
				},
				{
					Name:        "Cucumber",
					WeightGrams: 30,
					Macros: &MacroData{
						Calories: 5,
						Protein:  0,
						Carbs:    1,
						Fat:      0,
						Fiber:    0,
					},
				},
			},
		},
	}, nil
}
