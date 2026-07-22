package webconfiguration

import (
	"cryptox/internal/modules/market"
	"log"
)

type WSMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Service struct {
	repo Repository
	hub  *market.Hub
}

func NewService(repo Repository, hub *market.Hub) *Service {
	return &Service{repo: repo, hub: hub}
}

func (s *Service) GetFeatures() ([]FeatureFlag, error) {
	return s.repo.GetAll()
}

func (s *Service) IsEnabled(name string) bool {

	feature, err := s.repo.GetByName(name)

	if err != nil {
		return false
	}

	return feature.Enabled
}

func (s *Service) Update(id uint, enabled bool) error {

	if err := s.repo.Update(id, enabled); err != nil {
		return err
	}

	features, err := s.GetFeaturesMap()
	if err != nil {
		return err
	}

	log.Println(
		"broadcasting feature:",
		features,
	)

	if s.hub != nil {
		s.hub.BroadcastGlobal(
			WSMessage{
				Type: "feature_update",
				Data: features,
			},
		)
	}
	return nil
}

func (s *Service) GetFeaturesMap() (map[string]bool, error) {

	features, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	result := make(map[string]bool)

	for _, feature := range features {

		result[feature.FeatureName] = feature.Enabled
	}

	return result, nil
}

func (s *Service) Find(email string) any {
	var user User
	if err := s.repo.FindOne(&user, "email = ?", email); err != nil {
		return err
	}
	return user.Role
}

// func (s *Service) UpdateFeature(id uint,enabled bool) error {

// 	err:=s.repo.UpdateFeature(
// 		id,
// 		enabled,
// 	)

// 	if err!=nil{
// 		return err
// 	}

// 	features,_:=
// 	s.GetFeaturesMap()

// 	s.hub.BroadcastGlobal(
// 		fiber.Map{
// 			"type":"feature_update",
// 			"data":features,
// 		},
// 	)

// 	return nil
// }
