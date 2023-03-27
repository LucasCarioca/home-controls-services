package services

import (
	"errors"
	"fmt"

	"github.com/LucasCarioca/home-controls-services/pkg/datasource"
	"github.com/LucasCarioca/home-controls-services/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SwitchService switch crud service
type SwitchService struct {
	db *gorm.DB
}

// NewSwitchService create a new switch service
func NewSwitchService() *SwitchService {
	return &SwitchService{
		db: datasource.GetDataSource(),
	}
}

// GetAll get all switches
func (s *SwitchService) GetAll() []models.Switch {
	switches := make([]models.Switch, 0)
	s.db.Preload(clause.Associations).Find(&switches)
	return switches
}

// GetByID get switch by id
func (s *SwitchService) GetByID(id int) (*models.Switch, error) {
	fmt.Println(id)
	sw := models.Switch{}
	var c int64
	s.db.Preload(clause.Associations).Find(&sw, id).Count(&c)
	if c > 0 {
		return &sw, nil
	}
	return nil, errors.New("SWITCH_NOT_FOUND")
}

// DeleteByID delete switch by id
func (s *SwitchService) DeleteByID(id int) error {
	sw, err := s.GetByID(id)
	if err != nil {
		return err
	}
	s.db.Delete(sw)
	return nil
}

// UpdateDesiredStateByID update the desired state of a given switch
func (s *SwitchService) UpdateDesiredStateByID(id int, desiredState models.SwitchState) (*models.Switch, error) {
	sw, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	sw.DesiredState = desiredState
	s.db.Save(&sw)
	return sw, nil
}

// UpdateStateByID update the actual state of a given switch
func (s *SwitchService) UpdateStateByID(id int, state models.SwitchState) (*models.Switch, error) {
	sw, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	sw.State = state
	s.db.Save(&sw)
	return sw, nil
}

// Create creates a new switch
func (s *SwitchService) Create(name string, state models.SwitchState) models.Switch {
	sw := &models.Switch{
		Name:         name,
		State:        state,
		DesiredState: state,
	}
	s.db.Create(sw)
	return *sw
}
