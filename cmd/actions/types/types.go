package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type GraphQLError struct {
	Message string `json:"message"`
}

// ========================= Account Balances =========================
type AccountBalancesPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            Account_balancesArgs   `json:"input"`
}

type Account_balancesArgs struct {
	Address Address
}

type Address struct {
	Address string
}

type Coins struct {
	Coins []sdk.Coin
}

// ========================= Total Supply =========================
type TotalSupplyPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            total_supplyArgs       `json:"input"`
}

type total_supplyArgs struct{}
