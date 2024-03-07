package service

import "app/internal"

func NewVehicleDefaultMock() *VehicleDefaultMock {
	return &VehicleDefaultMock{}
}

type VehicleDefaultMock struct {
	FindByColorAndYearFunc      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FindByBrandAndYearRangeFunc func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)

	Spy struct {
		FindByColorAndYear      int
		FindByBrandAndYearRange int
	}
}

func (v2 *VehicleDefaultMock) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindByColorAndYear++
	return v2.FindByColorAndYearFunc(color, fabricationYear)
}

func (v2 *VehicleDefaultMock) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindByBrandAndYearRange++
	return v2.FindByBrandAndYearRangeFunc(brand, startYear, endYear)
}

func (v2 *VehicleDefaultMock) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	//TODO implement me
	panic("implement me")
}

func (v2 *VehicleDefaultMock) AverageCapacityByBrand(brand string) (a int, err error) {
	//TODO implement me
	panic("implement me")
}

func (v2 *VehicleDefaultMock) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	//TODO implement me
	panic("implement me")
}
