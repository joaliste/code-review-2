package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceVehicleDefault_FindByColorAndYear(t *testing.T) {
	// Given
	rp := repository.NewVehicleMapMock()
	rp.FindByColorAndYearFunc = func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
		return map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "A",
				Model:           "B",
				Registration:    "C",
				Color:           "D",
				FabricationYear: 1,
				Capacity:        1,
				MaxSpeed:        1,
				FuelType:        "E",
				Transmission:    "F",
				Weight:          1,
				Dimensions: internal.Dimensions{
					Height: 1,
					Length: 1,
					Width:  1,
				},
			},
		}}, nil
	}
	sv := service.NewServiceVehicleDefault(rp)

	expectedResult := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "A",
			Model:           "B",
			Registration:    "C",
			Color:           "D",
			FabricationYear: 1,
			Capacity:        1,
			MaxSpeed:        1,
			FuelType:        "E",
			Transmission:    "F",
			Weight:          1,
			Dimensions: internal.Dimensions{
				Height: 1,
				Length: 1,
				Width:  1,
			},
		},
	}}
	// When
	result, err := sv.FindByColorAndYear("D", 1)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, 1, rp.Spy.FindByColorAndYear)

}

func TestServiceVehicleDefault_FindByBrandAndYearRange(t *testing.T) {
	// Given
	rp := repository.NewVehicleMapMock()
	rp.FindByBrandAndYearRangeFunc = func(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
		return map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "A",
				Model:           "B",
				Registration:    "C",
				Color:           "D",
				FabricationYear: 1,
				Capacity:        1,
				MaxSpeed:        1,
				FuelType:        "E",
				Transmission:    "F",
				Weight:          1,
				Dimensions: internal.Dimensions{
					Height: 1,
					Length: 1,
					Width:  1,
				},
			},
		}}, nil
	}

	sv := service.NewServiceVehicleDefault(rp)

	expectedResult := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "A",
			Model:           "B",
			Registration:    "C",
			Color:           "D",
			FabricationYear: 1,
			Capacity:        1,
			MaxSpeed:        1,
			FuelType:        "E",
			Transmission:    "F",
			Weight:          1,
			Dimensions: internal.Dimensions{
				Height: 1,
				Length: 1,
				Width:  1,
			},
		},
	}}
	// When
	result, err := sv.FindByBrandAndYearRange("A", 0, 2)
	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, 1, rp.Spy.FindByBrandAndYearRange)
}
