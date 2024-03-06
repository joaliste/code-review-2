package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepositoryReadVehicleMap_FindByColorAndYear(t *testing.T) {
	// Given
	db := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Color:           "A",
			FabricationYear: 2008,
		},
	}, 2: {
		Id: 2,
		VehicleAttributes: internal.VehicleAttributes{
			Color:           "A",
			FabricationYear: 2007,
		},
	}, 3: {
		Id: 3,
		VehicleAttributes: internal.VehicleAttributes{
			Color:           "B",
			FabricationYear: 2008,
		},
	}}
	rp := repository.NewRepositoryReadVehicleMap(db)

	t.Run("Find by A color and 2008 as year", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Color:           "A",
				FabricationYear: 2008,
			},
		}}
		// When
		result, err := rp.FindByColorAndYear("A", 2008)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Find by B color and 2008 as year", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{3: {
			Id: 3,
			VehicleAttributes: internal.VehicleAttributes{
				Color:           "B",
				FabricationYear: 2008,
			},
		}}
		// When
		result, err := rp.FindByColorAndYear("B", 2008)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("No vehicles found", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{}
		// When
		result, err := rp.FindByColorAndYear("C", 2008)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
