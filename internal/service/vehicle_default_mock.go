package service

import "app/internal"

func NewVehicleDefaultMock() *VehicleDefaultMock {
	return &VehicleDefaultMock{}
}

type VehicleDefaultMock struct {
	FindByColorAndYearFunc      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FindByBrandAndYearRangeFunc func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)
	AverageMaxSpeedByBrandFunc  func(brand string) (a float64, err error)
	AverageCapacityByBrandFunc  func(brand string) (a int, err error)
	SearchByWeightRangeFunc     func(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error)

	Spy struct {
		FindByColorAndYear      int
		FindByBrandAndYearRange int
		AverageMaxSpeedByBrand  int
		AverageCapacityByBrand  int
		SearchByWeightRange     int
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
	v2.Spy.AverageMaxSpeedByBrand++
	return v2.AverageMaxSpeedByBrandFunc(brand)
}

func (v2 *VehicleDefaultMock) AverageCapacityByBrand(brand string) (a int, err error) {
	v2.Spy.AverageCapacityByBrand++
	return v2.AverageCapacityByBrandFunc(brand)
}

func (v2 *VehicleDefaultMock) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	v2.Spy.SearchByWeightRange++
	return v2.SearchByWeightRangeFunc(query, ok)
}
