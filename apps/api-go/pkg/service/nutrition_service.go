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
	// Create a prompt based on the input

	systemPrompt := `You are a nutritionist. You are given a picture (base64 encoded) of a meal/food and optionally a description of the food item.
		If a description is provided, use it to help you understand the food item better.
		You need to return the nutritional information for that food item.
    `

	description := ""
	if input.Description != nil {
		description = *input.Description
	}

	prompt := fmt.Sprintf(`Analyze the following food picture:
			Image: %s
			Description: %s`, input.ImageBase64, description)

	// Generate structured recipe data using the same schema
	recipe, _, err := genkit.GenerateData[ScanOutput](ctx, s.genkit,
		ai.WithSystem(systemPrompt),
		ai.WithPrompt(prompt),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate recipe: %w", err)
	}

	return recipe, nil
}
