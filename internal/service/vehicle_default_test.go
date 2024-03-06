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
	// Then
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
	// Then
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, 1, rp.Spy.FindByBrandAndYearRange)
}

func TestServiceVehicleDefault_AverageMaxSpeedByBrand(t *testing.T) {
	t.Run("find two vehicles with an average of 5", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindByBrandFunc = func(brand string) (v map[int]internal.Vehicle, err error) {
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
			}, 2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "A",
					Model:           "B",
					Registration:    "C",
					Color:           "D",
					FabricationYear: 1,
					Capacity:        1,
					MaxSpeed:        5,
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

		expectedResult := 3.0
		// When
		result, err := sv.AverageMaxSpeedByBrand("A")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
		assert.Equal(t, 1, rp.Spy.FindByBrand)
	})

	t.Run("no vehicles found", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindByBrandFunc = func(brand string) (v map[int]internal.Vehicle, err error) {
			return map[int]internal.Vehicle{}, nil
		}

		sv := service.NewServiceVehicleDefault(rp)

		expectedError := internal.ErrServiceNoVehicles
		// When
		_, err := sv.AverageMaxSpeedByBrand("A")
		// Then
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, 1, rp.Spy.FindByBrand)
	})
}

func TestServiceVehicleDefault_AverageCapacityByBrand(t *testing.T) {
	t.Run("find two vehicles with an average capacity of 5", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindByBrandFunc = func(brand string) (v map[int]internal.Vehicle, err error) {
			return map[int]internal.Vehicle{1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "A",
					Model:           "B",
					Registration:    "C",
					Color:           "D",
					FabricationYear: 1,
					Capacity:        5,
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
			}, 2: {
				Id: 2,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "A",
					Model:           "B",
					Registration:    "C",
					Color:           "D",
					FabricationYear: 1,
					Capacity:        5,
					MaxSpeed:        5,
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

		expectedResult := 5
		// When
		result, err := sv.AverageCapacityByBrand("A")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
		assert.Equal(t, 1, rp.Spy.FindByBrand)
	})

	t.Run("no vehicles found", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindByBrandFunc = func(brand string) (v map[int]internal.Vehicle, err error) {
			return map[int]internal.Vehicle{}, nil
		}

		sv := service.NewServiceVehicleDefault(rp)

		expectedError := internal.ErrServiceNoVehicles
		// When
		_, err := sv.AverageCapacityByBrand("A")
		// Then
		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
		assert.Equal(t, 1, rp.Spy.FindByBrand)
	})
}

func TestServiceVehicleDefault_SearchByWeightRange(t *testing.T) {
	t.Run("Find without query", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindAllFunc = func() (v map[int]internal.Vehicle, err error) {
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
		result, err := sv.SearchByWeightRange(internal.SearchQuery{}, false)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
		assert.Equal(t, 1, rp.Spy.FindAll)
	})

	t.Run("Find with query", func(t *testing.T) {
		// Given
		rp := repository.NewVehicleMapMock()
		rp.FindByWeightRangeFunc = func(fromWeight float64, toWeight float64) (v map[int]internal.Vehicle, err error) {
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
		result, err := sv.SearchByWeightRange(internal.SearchQuery{FromWeight: 0, ToWeight: 2}, true)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
		assert.Equal(t, 0, rp.Spy.FindAll)
		assert.Equal(t, 1, rp.Spy.FindByWeightRange)
	})
}
