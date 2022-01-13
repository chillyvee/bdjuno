package remote

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/forbole/juno/v2/node/remote"

	mintsource "github.com/forbole/bdjuno/v2/modules/mint/source"
)

var (
	_ mintsource.Source = &Source{}
)

// Source implements mintsource.Source using a remote node
type Source struct {
	*remote.Source
	querier minttypes.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, querier minttypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetInflation implements mintsource.Source
func (s Source) GetInflation(height int64) (sdk.Dec, error) {
	timeNow := time.Now()
	res, err := s.querier.Inflation(s.Ctx, &minttypes.QueryInflationRequest{}, remote.GetHeightRequestHeader(height))
	if err != nil {
		return sdk.Dec{}, err
	}
	fmt.Println("Time(Seconds) spent for mint/ GetInflation: ", time.Since(timeNow).Seconds())

	return res.Inflation, nil
}

// Params implements mintsource.Source
func (s Source) Params(height int64) (minttypes.Params, error) {
	timeNow := time.Now()

	res, err := s.querier.Params(s.Ctx, &minttypes.QueryParamsRequest{}, remote.GetHeightRequestHeader(height))
	if err != nil {
		return minttypes.Params{}, nil
	}
	fmt.Println("Time(Seconds) spent for mint/ Params: ", time.Since(timeNow).Seconds())

	return res.Params, nil
}
