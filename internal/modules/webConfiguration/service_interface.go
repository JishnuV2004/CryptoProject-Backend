package webconfiguration

type FeatureService interface {
	IsEnabled(feature string) bool
	GetFeatures() ([]FeatureFlag, error)
	Update(id uint, enabled bool) error
	Find(email string) any
}