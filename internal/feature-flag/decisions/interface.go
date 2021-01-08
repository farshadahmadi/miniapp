package decisions

import (
	"context"

	"github.com/farshadahmadi/miniapp/internal/feature-flag/features"
)

// DecisionService is the interface that every service for desicion should implement
type DecisionService interface {
	GetFlag(ctx context.Context, params features.GetFlagParams) (treatment string, err error)
}
