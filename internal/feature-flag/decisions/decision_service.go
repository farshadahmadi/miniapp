package decisions

import (
	"context"

	"github.com/farshadahmadi/miniapp/internal/feature-flag/features"
)

// Decisions is a service for making toggline decisions in a sinle place
type Decisions struct {
	featureService features.FeatureService
}

// NewDecisions instanciates Decisions service
func NewDecisions(featureService features.FeatureService) *Decisions {
	return &Decisions{
		featureService: featureService,
	}
}

// GetFlag is a proxy for client to request feature flags from split.io
func (d *Decisions) GetFlag(ctx context.Context, params features.GetFlagParams) (treatment string, err error) {
	treatment, err = d.featureService.Get(params)

	return
}
