package repository

import "app/internal"

func NewVehicleMapMock() *VehicleMapMock {
	return &VehicleMapMock{}
}

type VehicleMapMock struct {
	FindByColorAndYearFunc func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)

	Spy struct {
		FindByColorAndYear int
	}
}
