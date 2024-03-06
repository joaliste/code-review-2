package repository

import "app/internal"

func NewVehicleMapMock() *VehicleMapMock {
	return &VehicleMapMock{}
}

type VehicleMapMock struct {
	FindByColorAndYearFunc      func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error)
	FindByBrandAndYearRangeFunc func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error)
	FindByBrandFunc             func(brand string) (v map[int]internal.Vehicle, err error)
	FindAllFunc                 func() (v map[int]internal.Vehicle, err error)
	FindByWeightRangeFunc       func(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error)

	Spy struct {
		FindByColorAndYear      int
		FindByBrandAndYearRange int
		FindByBrand             int
		FindAll                 int
		FindByWeightRange       int
	}
}

func (v2 *VehicleMapMock) FindAll() (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindAll++
	return v2.FindAllFunc()
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
	v2.Spy.FindByBrand++
	return v2.FindByBrandFunc(brand)
}

func (v2 *VehicleMapMock) FindByWeightRange(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
	v2.Spy.FindByWeightRange++
	return v2.FindByWeightRangeFunc(fromWeight, toWeight)
}
