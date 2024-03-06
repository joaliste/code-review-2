package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/service"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("Find a vehicle", func(t *testing.T) {
		// Given
		sv := service.NewVehicleDefaultMock()
		sv.FindByColorAndYearFunc = func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
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
		hd := handler.NewHandlerVehicle(sv)

		hdFunc := hd.FindByColorAndYear()

		expectedBodyOutput := `{"data":{"1":{"Id":1,"Brand":"A","Model":"B","Registration":"C","Color":"D","FabricationYear":1,"Capacity":1,"MaxSpeed":1,"FuelType":"E","Transmission":"F","Weight":1,"Height":1,"Length":1,"Width":1}},"message":"vehicles found"}`
		expectedStatusCode := http.StatusOK
		expectedHeaderOutput := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// When
		req := httptest.NewRequest(http.MethodGet, "/vehicles/color/D/year/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("color", "D")
		chiCtx.URLParams.Add("year", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)
		// Then
		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBodyOutput, res.Body.String())
		require.Equal(t, expectedHeaderOutput, res.Header())
		require.Equal(t, 1, sv.Spy.FindByColorAndYear)
	})

	t.Run("Invalid year", func(t *testing.T) {
		// Given
		sv := service.NewVehicleDefaultMock()
		hd := handler.NewHandlerVehicle(sv)

		hdFunc := hd.FindByColorAndYear()

		expectedBodyOutput := `{"message":"invalid year", "status":"Bad Request"}`
		expectedStatusCode := http.StatusBadRequest
		expectedHeaderOutput := http.Header{
			"Content-Type": []string{"application/json"},
		}
		// When
		req := httptest.NewRequest(http.MethodGet, "/vehicles/color/D/year/F", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("color", "D")
		chiCtx.URLParams.Add("year", "F")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)
		// Then
		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBodyOutput, res.Body.String())
		require.Equal(t, expectedHeaderOutput, res.Header())
		require.Equal(t, 0, sv.Spy.FindByColorAndYear)
	})

	t.Run("Unknown error", func(t *testing.T) {
		// Given
		sv := service.NewVehicleDefaultMock()
		sv.FindByColorAndYearFunc = func(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
			return nil, errors.New("unknown error")
		}

		hd := handler.NewHandlerVehicle(sv)

		hdFunc := hd.FindByColorAndYear()

		expectedBodyOutput := `{"message":"internal error", "status":"Internal Server Error"}`
		expectedStatusCode := http.StatusInternalServerError
		expectedHeaderOutput := http.Header{
			"Content-Type": []string{"application/json"},
		}
		// When
		req := httptest.NewRequest(http.MethodGet, "/vehicles/color/D/year/1", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("color", "D")
		chiCtx.URLParams.Add("year", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()
		hdFunc(res, req)
		// Then
		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBodyOutput, res.Body.String())
		require.Equal(t, expectedHeaderOutput, res.Header())
		require.Equal(t, 1, sv.Spy.FindByColorAndYear)
	})
}
