package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// NutritionService is a service for nutritional information
type NutritionService struct {
	genkit *genkit.Genkit
	flows  map[string]*core.Flow[*ScanInput, *ScanOutput, struct{}]
}

// NewNutritionService creates a new nutrition service
func NewNutritionService(genkit *genkit.Genkit) *NutritionService {
	svc := &NutritionService{
		genkit: genkit,
		flows:  map[string]*core.Flow[*ScanInput, *ScanOutput, struct{}]{},
	}
	svc.initializeFlows()
	return svc
}

// initializeFlows initializes all AI flows for genkit and stores them in the flows map
func (s *NutritionService) initializeFlows() {
	s.flows = map[string]*core.Flow[*ScanInput, *ScanOutput, struct{}]{
		"foodScanFlow": genkit.DefineFlow(s.genkit, "foodScanFlow", s.foodScanFlow),
	}
}

// ScanFood scans the food in the image and returns the nutritional information
func (s *NutritionService) ScanFood(ctx context.Context, input *ScanInput) (*ScanOutput, error) {
	slog.Info("received food scan request", "input", input)

	flow := s.flows["foodScanFlow"]
	response, err := flow.Run(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to run food scan flow: %w", err)
	}

	slog.Info("food scan response", "response", response)

	return response, nil
}

func (s *NutritionService) foodScanFlow(ctx context.Context, input *ScanInput) (*ScanOutput, error) {
	systemPrompt := `You are an expert nutritionist and food recognition AI.
Analyze the provided food image and estimate its nutritional content.

ANALYSIS STEPS:
1. Identify all visible food items, ingredients, and portion sizes
2. For each ingredient, estimate its weight in grams and calculate individual macros
3. Consider cooking methods (fried, grilled, steamed, etc.) as they affect calories
4. Estimate serving sizes relative to standard references (e.g., a fist ≈ 1 cup, palm ≈ 3oz protein)
5. Sum all ingredient macros to get the total meal macros

OUTPUT REQUIREMENTS:
- food_name: Overall meal/dish name
- confidence: How clearly the food is identifiable (0.0-1.0)
- serving_size: Total serving in grams or standard units (e.g., "1 plate, ~350g")
- macros: Total combined macros for the entire meal
- ingredients: Array of each component with:
  - name: Ingredient name (e.g., "Grilled Chicken Breast")
  - weight_grams: Estimated weight in grams
  - macros: Individual macros for this ingredient

GUIDELINES:
- Always break down complex meals into their visible components
- Use reasonable middle-ground estimates when portions are unclear
- Include fiber in macro calculations when applicable`

	// Build the user prompt text
	userPrompt := "Analyze this food image and provide nutritional information."
	if input.Description != nil && *input.Description != "" {
		userPrompt = fmt.Sprintf("Analyze this food image. Additional context: %s", *input.Description)
	}

	// Build the image data URL for multimodal input
	imageDataURL := "data:image/jpeg;base64," + input.ImageBase64

	// Generate structured output using proper multimodal input
	result, _, err := genkit.GenerateData[ScanOutput](ctx, s.genkit,
		ai.WithSystem(systemPrompt),
		ai.WithMessages(
			ai.NewUserMessage(
				ai.NewMediaPart("image/jpeg", imageDataURL),
				ai.NewTextPart(userPrompt),
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze food image: %w", err)
	}

	return result, nil
}
