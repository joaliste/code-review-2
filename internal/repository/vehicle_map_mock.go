package repository

import "app/internal"

func NewVehicleMapMock() *VehicleMapMock {
	return &VehicleMapMock{}
}

type VehicleMapMock struct {
	FindByColorAndYearFunc      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FindByBrandAndYearRangeFunc func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)

	Spy struct {
		FindByColorAndYear      int
		FindByBrandAndYearRange int
	}
}

func (v2 *VehicleMapMock) FindAll() (v map[int]internal.Vehicle, err error) {
	//TODO implement me
	panic("implement me")
}

func (v2 *VehicleMapMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindByColorAndYear++
	return v2.FindByColorAndYearFunc(color, fabricationYear)
}

func (v2 *VehicleMapMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindByBrandAndYearRange++
	return v2.FindByBrandAndYearRangeFunc(brand, startYear, endYear)
}

func (v2 *VehicleMapMock) FindByBrand(brand string) (v map[int]internal.Vehicle, err error) {
	//TODO implement me
	panic("implement me")
}

func (v2 *VehicleMapMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	//TODO implement me
	panic("implement me")
}
