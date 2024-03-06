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

func TestRepositoryReadVehicleMap_FindByBrandAndYearRange(t *testing.T) {
	// Given
	db := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "A",
			FabricationYear: 2008,
		},
	}, 2: {
		Id: 2,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "A",
			FabricationYear: 2005,
		},
	}, 3: {
		Id: 3,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           "B",
			FabricationYear: 2010,
		},
	}}
	rp := repository.NewRepositoryReadVehicleMap(db)

	t.Run("Find by brand A and with the year between 2000 and 2010", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "A",
				FabricationYear: 2008,
			},
		}, 2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "A",
				FabricationYear: 2005,
			},
		}}
		// When
		result, err := rp.FindByBrandAndYearRange("A", 2000, 2010)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Find by brand A and with the year between 2000 and 2006", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "A",
				FabricationYear: 2005,
			},
		}}
		// When
		result, err := rp.FindByBrandAndYearRange("A", 2000, 2006)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Find by brand B and with the year between 2000 and 2010", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{3: {
			Id: 3,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           "B",
				FabricationYear: 2010,
			},
		}}
		// When
		result, err := rp.FindByBrandAndYearRange("B", 2000, 2010)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Brand not found", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{}
		// When
		result, err := rp.FindByBrandAndYearRange("C", 2000, 2010)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Year range not found", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{}
		// When
		result, err := rp.FindByBrandAndYearRange("A", 2000, 2002)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}

func TestRepositoryReadVehicleMap_FindByBrand(t *testing.T) {
	// Given
	db := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Brand: "A",
		},
	}, 2: {
		Id: 2,
		VehicleAttributes: internal.VehicleAttributes{
			Brand: "A",
		},
	}, 3: {
		Id: 3,
		VehicleAttributes: internal.VehicleAttributes{
			Brand: "B",
		},
	}}
	rp := repository.NewRepositoryReadVehicleMap(db)

	t.Run("Find by brand A", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Brand: "A",
			},
		}, 2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Brand: "A",
			},
		}}
		// When
		result, err := rp.FindByBrand("A")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Find by brand B", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{3: {
			Id: 3,
			VehicleAttributes: internal.VehicleAttributes{
				Brand: "B",
			},
		}}
		// When
		result, err := rp.FindByBrand("B")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Brand not found", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{}
		// When
		result, err := rp.FindByBrand("C")
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}

func TestRepositoryReadVehicleMap_FindByWeightRange(t *testing.T) {
	// Given
	db := map[int]internal.Vehicle{1: {
		Id: 1,
		VehicleAttributes: internal.VehicleAttributes{
			Weight: 1,
		},
	}, 2: {
		Id: 2,
		VehicleAttributes: internal.VehicleAttributes{
			Weight: 10,
		},
	}, 3: {
		Id: 3,
		VehicleAttributes: internal.VehicleAttributes{
			Weight: 4,
		},
	}}
	rp := repository.NewRepositoryReadVehicleMap(db)

	t.Run("Find by 0 to 10 weight range", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{1: {
			Id: 1,
			VehicleAttributes: internal.VehicleAttributes{
				Weight: 1,
			},
		}, 2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Weight: 10,
			},
		}, 3: {
			Id: 3,
			VehicleAttributes: internal.VehicleAttributes{
				Weight: 4,
			},
		}}
		// When
		result, err := rp.FindByWeightRange(0, 10)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Find by 5 to 10 weight range", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{2: {
			Id: 2,
			VehicleAttributes: internal.VehicleAttributes{
				Weight: 10,
			},
		}}
		// When
		result, err := rp.FindByWeightRange(5, 10)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("Vehicles not found", func(t *testing.T) {
		// Given
		expectedResult := map[int]internal.Vehicle{}
		// When
		result, err := rp.FindByWeightRange(11, 30)
		// Then
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
