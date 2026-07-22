package webconfiguration

type Repository interface {
	GetAll() ([]FeatureFlag, error)
	GetByName(name string) (*FeatureFlag, error)
	Update(id uint, enabled bool) error
	FindOne(model interface{}, query string, args ...any) error
}