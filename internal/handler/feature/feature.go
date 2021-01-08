package feature

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/farshadahmadi/miniapp/internal/feature-flag/decisions"
	"github.com/farshadahmadi/miniapp/internal/feature-flag/features"
	"github.com/farshadahmadi/miniapp/internal/generated/swagger/server/miniapp/model"
	"github.com/farshadahmadi/miniapp/internal/generated/swagger/server/miniapp/operation/feature"
)

type GetFeatureFlag struct {
	decisionService decisions.DecisionService
}

func NewGetFeatureFlag(decisionService decisions.DecisionService) *GetFeatureFlag {
	return &GetFeatureFlag{
		decisionService: decisionService,
	}
}

func (gff *GetFeatureFlag) GetFeatureFlag(params feature.FeatureFlagParams) middleware.Responder {
	attributes := map[string]interface{}{}
	if len(params.Body.Attributes) != 0 {
		attributes = params.Body.Attributes
	}

	gfp := features.GetFlagParams{
		Feature:    params.Feature,
		TrafficID:  string(params.Body.TrafficID),
		Attributes: attributes,
	}

	flagValue, err := gff.decisionService.GetFlag(params.HTTPRequest.Context(), gfp)
	if err != nil {
		return feature.NewFeatureFlagInternalServerError()
	}

	payload := &model.FeatureFlagResponse{
		Flag:  params.Feature,
		Value: flagValue,
	}
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return feature.NewFeatureFlagInternalServerError()
	}
	return feature.NewFeatureFlagOK().WithPayload(payload)
}
