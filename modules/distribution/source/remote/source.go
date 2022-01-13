package remote

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/forbole/juno/v2/node/remote"

	distrsource "github.com/forbole/bdjuno/v2/modules/distribution/source"
)

var (
	_ distrsource.Source = &Source{}
)

// Source implements distrsource.Source querying the data from a remote node
type Source struct {
	*remote.Source
	distrClient distrtypes.QueryClient
}

// NewSource returns a new Source instace
func NewSource(source *remote.Source, distrClient distrtypes.QueryClient) *Source {
	return &Source{
		Source:      source,
		distrClient: distrClient,
	}
}

// ValidatorCommission implements distrsource.Source
func (s Source) ValidatorCommission(valOperAddr string, height int64) (sdk.DecCoins, error) {
	timeNow := time.Now()
	res, err := s.distrClient.ValidatorCommission(
		s.Ctx,
		&distrtypes.QueryValidatorCommissionRequest{ValidatorAddress: valOperAddr},
		remote.GetHeightRequestHeader(height),
	)
	if err != nil {
		return nil, err
	}

	fmt.Println("Time(Seconds) spent for distribution/ ValidatorCommission: ", time.Since(timeNow).Seconds())

	return res.Commission.Commission, nil
}

// DelegatorTotalRewards implements distrsource.Source
func (s Source) DelegatorTotalRewards(delegator string, height int64) ([]distrtypes.DelegationDelegatorReward, error) {
	timeNow := time.Now()

	res, err := s.distrClient.DelegationTotalRewards(
		s.Ctx,
		&distrtypes.QueryDelegationTotalRewardsRequest{DelegatorAddress: delegator},
		remote.GetHeightRequestHeader(height),
	)
	if err != nil {
		return nil, fmt.Errorf("error while getting delegation total rewards for for delegator %s at height %v: %s", delegator, height, err)
	}

	fmt.Println("Time(Seconds) spent for distribution/ DelegatorTotalRewards: ", time.Since(timeNow).Seconds())

	return res.Rewards, nil
}

// DelegatorWithdrawAddress implements distrsource.Source
func (s Source) DelegatorWithdrawAddress(delegator string, height int64) (string, error) {
	timeNow := time.Now()

	res, err := s.distrClient.DelegatorWithdrawAddress(
		s.Ctx,
		&distrtypes.QueryDelegatorWithdrawAddressRequest{DelegatorAddress: delegator},
		remote.GetHeightRequestHeader(height),
	)
	if err != nil {
		return "", err
	}
	fmt.Println("Time(Seconds) spent for distribution/ DelegatorWithdrawAddress: ", time.Since(timeNow).Seconds())

	return res.WithdrawAddress, nil
}

// CommunityPool implements distrsource.Source
func (s Source) CommunityPool(height int64) (sdk.DecCoins, error) {
	timeNow := time.Now()

	res, err := s.distrClient.CommunityPool(
		s.Ctx,
		&distrtypes.QueryCommunityPoolRequest{},
		remote.GetHeightRequestHeader(height),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("Time(Seconds) spent for distribution/ CommunityPool: ", time.Since(timeNow).Seconds())

	return res.Pool, nil
}

// Params implements distrsource.Source
func (s Source) Params(height int64) (distrtypes.Params, error) {
	timeNow := time.Now()

	res, err := s.distrClient.Params(
		s.Ctx,
		&distrtypes.QueryParamsRequest{},
		remote.GetHeightRequestHeader(height),
	)
	if err != nil {
		return distrtypes.Params{}, err
	}
	fmt.Println("Time(Seconds) spent for distribution/ Params: ", time.Since(timeNow).Seconds())

	return res.Params, nil
}
