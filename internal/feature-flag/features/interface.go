package features

// FeatureService is an interface that allows application to get flags based on the given params
type FeatureService interface {
	Get(params GetFlagParams) (string, error)
}
