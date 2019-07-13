package distancematrix

import (
	"context"
	"errors"
	"googlemaps.github.io/maps"
	"strings"
)

type DistanceMatrix struct {
	Client *maps.Client
}

type DistanceConfig struct {
	Url    string
	ApiKey string
}

type DistanceService interface {
	CalculateDistance(ctx context.Context, origin []string, dest []string) (int, error)
}

func Init(url string, apiKey string) (*DistanceMatrix, error) {
	client, err := maps.NewClient(maps.WithAPIKey(apiKey), maps.WithBaseURL(url))
	if err != nil {
		return nil, err
	}
	return &DistanceMatrix{Client: client}, nil
}

func (d *DistanceMatrix) CalculateDistance(ctx context.Context, origin []string, dest []string) (int, error) {

	row, err := d.Client.DistanceMatrix(ctx, &maps.DistanceMatrixRequest{
		Origins:      []string{strings.Join(origin, ",")},
		Destinations: []string{strings.Join(dest, ",")},
	})

	if err != nil {
		return 0, err
	}

	if len(row.Rows) > 0 && len(row.Rows[0].Elements) > 0 {

		if row.Rows[0].Elements[0].Status == "OK" {
			return row.Rows[0].Elements[0].Distance.Meters, nil
		}

		return 0, errors.New("distance cannot be calculated - " + row.Rows[0].Elements[0].Status)

	} else {
		return 0, errors.New("DISTANCE_MATRIX_SERVER_ERROR")
	}
}
