package webconfiguration

import "gorm.io/gorm"

type ConfigRepo struct {
	db *gorm.DB
}

func NewConfigRepo(db *gorm.DB) *ConfigRepo {
	return &ConfigRepo{
		db: db,
	}
}

func (r *ConfigRepo) GetAll() ([]FeatureFlag, error) {

	var features []FeatureFlag

	err := r.db.Find(&features).Error

	return features, err
}

func (r *ConfigRepo) GetByName(name string) (*FeatureFlag, error) {

	var feature FeatureFlag

	err := r.db.
		Where("feature_name=?", name).
		First(&feature).Error

	return &feature, err
}

func (r *ConfigRepo) Update(id uint, enabled bool) error {

	return r.db.
		Model(&FeatureFlag{}).
		Where("id = ?", id).
		Update("enabled", enabled).
		Error
}

func (r *ConfigRepo) FindOne(model interface{}, query string, args ...any) error {
	return r.db.Where(query, args...).First(model).Error
}