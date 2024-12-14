package service

import (
	"github.com/mmdali-dev/deymod/db"
	"github.com/mmdali-dev/deymod/db/model"
)

type locationService service

var (
	Location locationService = locationService{"LocatioService", "CRUD and auth for Locations", db.GetDB()}
)

func (s *locationService) CreateLocation(location *model.Location) (err error) {
	err = s.DB.Create(&location).Error
	return err
}

func (s *locationService) DeleteLocation(location *model.Location) (err error) {
	err = s.DB.Where(location).Delete(location).Error
	return err
}

func (s *locationService) UpdateLocation(location *model.Location) (err error) {
	err = s.DB.Save(location).Error
	return err
}

func (s *locationService) FindLocation(location *model.Location) (err error) {
	err = s.DB.Where(location).First(location).Error
	return err
}

func (s *locationService) FindLocations(location *model.Location, locations *[]model.Location) (err error) {
	err = s.DB.Where(location).Find(locations).Error
	return err
}

func (s *locationService) AllLocations(locations *[]model.Location) (err error) {
	err = s.DB.Find(&locations).Error
	return err
}

// public information for Location service
func (s *locationService) CreatePublicLocation(publoc *model.PublicLocation) (err error) {
	err = s.DB.Create(&publoc).Error
	return
}
func (s *locationService) UpdatePublicLocation(publoc *model.PublicLocation) (err error) {
	err = s.DB.Save(publoc).Error
	return err
}
func (s *locationService) FindPublicLocation(publoc *model.PublicLocation) (err error) {
	err = s.DB.Where(publoc).First(publoc).Error
	return err
}

func (s *locationService) FindPublicLocations(publoc *model.PublicLocation, publocs *[]model.PublicLocation) (err error) {
	err = s.DB.Where(publoc).Find(publocs).Error
	return err
}
func (s *locationService) AllPublicLocations(publocs *[]model.PublicLocation) (err error) {
	err = s.DB.Find(&publocs).Error
	return err
}
func (s *locationService) DeletePublicLocation(publoc *model.PublicLocation) (err error) {
	err = s.DB.Where(publoc).Delete(publoc).Error
	return err
}
