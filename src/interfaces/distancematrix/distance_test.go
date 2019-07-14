package distancematrix

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"googlemaps.github.io/maps"
	"context"
	"errors"
	"flag"
)

var integration = flag.Bool("integration", false, "run database integration tests")


type DistanceCalculatorMock struct {
	mock.Mock
}

func (m *DistanceCalculatorMock) DistanceMatrix(ctx context.Context, r *maps.DistanceMatrixRequest) (*maps.DistanceMatrixResponse, error) {
	args := m.Called(ctx, r)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*maps.DistanceMatrixResponse), args.Error(1)
}

func GetMockDistanceCalculator() *DistanceCalculatorMock {
	distanceCalculator := new(DistanceCalculatorMock)
	distanceCalculator.On("DistanceMatrix", nil, &maps.DistanceMatrixRequest{
		Origins:      []string{"abc"},
		Destinations: []string{"xyz"},
	}).Return(nil, errors.New("wrong origin and destination"))

	distanceCalculator.On("DistanceMatrix", nil, &maps.DistanceMatrixRequest{
		Origins:      []string{"40.43206,-80.38992"},
		Destinations: []string{"41.43206,-80.388"},
	}).Return(&maps.DistanceMatrixResponse{
		Rows: []maps.DistanceMatrixElementsRow{
			{
				Elements: []*maps.DistanceMatrixElement{
					{
						Status: "OK",
						Distance: maps.Distance{
							Meters: 131758,
						},
					},
				},
			},
		},
	}, nil)

	distanceCalculator.On("DistanceMatrix", nil, &maps.DistanceMatrixRequest{
		Origins:      []string{"0.43206,-80.38992"},
		Destinations: []string{"1.43206,-80.388"},
	}).Return(&maps.DistanceMatrixResponse{
		Rows: []maps.DistanceMatrixElementsRow{
			{
				Elements: []*maps.DistanceMatrixElement{
					{
						Status: "ZERO_RESULTS",
						Distance: maps.Distance{
							Meters: 0,
						},
					},
				},
			},
		},
	}, nil)

	return distanceCalculator
}

func TestDistanceMatrix_CalculateDistance(t *testing.T) {
	dm := &DistanceMatrix{GetMockDistanceCalculator()}

	_, err := dm.CalculateDistance(nil, []string{"abc"}, []string{"xyz"})

	if err == nil {
		t.Error("error should be thrown for wrong origin and destination")
	}

	_, err = dm.CalculateDistance(nil, []string{"0.43206", "-80.38992"}, []string{"1.43206", "-80.388"})

	if err == nil {
		t.Error("error should be thrown for no route found")
	}

	distance, err := dm.CalculateDistance(nil, []string{"40.43206", "-80.38992"}, []string{"41.43206", "-80.388"})

	if err != nil {
		t.Error(err)
	}

	if distance <= 0 {
		t.Error("distance must be greater than zero")
	}

}
