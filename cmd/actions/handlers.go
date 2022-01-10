package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	actionstypes "github.com/forbole/bdjuno/v2/cmd/actions/types"
)

func accountBalancesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload actionstypes.AccountBalancesPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload: failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	result, err := getAccountBalances(actionPayload.Input.Address.Address)
	if err != nil {
		graphQlError(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

func totalSupplyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := getTotalSupply()
	if err != nil {
		graphQlError(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}
