package mpos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MiningPoolHub ...
type MiningPoolHub struct {
	apiKey  string
	baseURL string
}

// NewMiningPoolHub ...
func NewMiningPoolHub(apiKey string, baseURL string) *MiningPoolHub {
	// TODO validate apiKey
	return &MiningPoolHub{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

// get processes a request to be sent to an php-mpos api.
func (m *MiningPoolHub) get(action string, params *url.Values, bind interface{}) (*http.Response, error) {
	// build url
	base, err := url.Parse(m.baseURL)
	if err != nil {
		return nil, err
	}
	if params == nil {
		params = &url.Values{}
	}
	params.Add("action", action)
	params.Add("page", "api")

	// add api key
	if m.apiKey != "" {
		params.Add("api_key", m.apiKey)
	}
	base.RawQuery = params.Encode()

	res, err := http.Get(base.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v returned non-ok http code, %v", action, res.StatusCode)
	}
	if bind != nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(body, &bind); err != nil {
			return nil, fmt.Errorf("error unmarshaling JSON %v, body: %v", err, string(body))
		}
	}

	return res, err
}

// BlockCount get current block height in blockchain.
func (m *MiningPoolHub) BlockCount() (bRes BlockCountResponse, err error) {
	_, err = m.get("getblockcount", nil, &bRes)
	return bRes, err
}

// BlocksFound get last N blocks found as configured in
// admin panel.
func (m *MiningPoolHub) BlocksFound() (bRes BlocksFoundResponse, err error) {
	_, err = m.get("getblocksfound", nil, &bRes)
	return bRes, err
}

// BlockStats get pool block stats.
func (m *MiningPoolHub) BlockStats() (bRes BlockStatsResponse, err error) {
	_, err = m.get("getblockstats", nil, &bRes)
	return bRes, err
}

// CurrentWorkers get amount of current active workers.
func (m *MiningPoolHub) CurrentWorkers() (cRes CurrentWorkersResponse, err error) {
	_, err = m.get("getcurrentworkers", nil, &cRes)
	return cRes, err
}

// DashboardData fetch all dashboard related information.
func (m *MiningPoolHub) DashboardData() (dRes DashboardDataResponse, err error) {
	_, err = m.get("getdashboarddata", nil, &dRes)
	return dRes, err
}

// Difficulty get current difficulty in blockchain.
func (m *MiningPoolHub) Difficulty() (dRes DifficultyResponse, err error) {
	_, err = m.get("getdifficulty", nil, &dRes)
	return dRes, err
}

// EstimatedTime get estimated time to next block based on
// pool hashrate (seconds).
func (m *MiningPoolHub) EstimatedTime() (eRes EstimatedTimeResponse, err error) {
	_, err = m.get("getestimatedtime", nil, &eRes)
	return eRes, err
}

// NavbarData - disabled on miningpoolhub.com
func (m *MiningPoolHub) NavbarData() {
	// TBI
}

// PoolHashrate get current pool hashrate.
func (m *MiningPoolHub) PoolHashrate() (pRes PoolHashrateResponse, err error) {
	_, err = m.get("getpoolhashrate", nil, &pRes)
	return pRes, err
}

// PoolInfo get the information on pool settings.
func (m *MiningPoolHub) PoolInfo() (pRes PoolInfoResponse, err error) {
	_, err = m.get("getpoolinfo", nil, &pRes)
	return pRes, err
}

// PoolShareRate - always returns 0 on miningpoolhub.com
func (m *MiningPoolHub) PoolShareRate() {
	// TBI
}

// PoolStatus fetch overall pool status
func (m *MiningPoolHub) PoolStatus() (pRes PoolStatusResponse, err error) {
	_, err = m.get("getpoolstatus", nil, &pRes)
	return pRes, err
}

// TimeSinceLastBlock get time since last block found (seconds)
func (m *MiningPoolHub) TimeSinceLastBlock() (tRes TimeSinceLastBlockResponse, err error) {
	_, err = m.get("gettimesincelastblock", nil, &tRes)
	return tRes, err
}

// TopContributors fetch top contributors data
func (m *MiningPoolHub) TopContributors() (tRes TopContributorsResponse, err error) {
	_, err = m.get("gettopcontributors", nil, &tRes)
	return tRes, err
}

// UserBalance fetch a users balance. If not admin, will
// fetch current users balance.
func (m *MiningPoolHub) UserBalance(id string) (uRes UserBalanceResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getuserbalance", params, &uRes)
	return uRes, err
}

// UserHashrate fetch a users hashrate. If not admin, will
// fetch current users hashrate.
func (m *MiningPoolHub) UserHashrate(id string) (uRes UserHashrateResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getuserhashrate", params, &uRes)
	return uRes, err
}

// UserShareRate fetch a users hashrate. If not admin, will
// fetch current users hashrate.
func (m *MiningPoolHub) UserShareRate(id string) (uRes UserShareRateResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getusersharerate", params, &uRes)
	return uRes, err
}

// UserStatus fetch a users overall status, both id and
// username work for id. If not admin, will fetch current
// users status.
func (m *MiningPoolHub) UserStatus(id string) (uRes UserStatusResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getuserstatus", params, &uRes)
	return uRes, err
}

// UserTransactions get a users transactions
func (m *MiningPoolHub) UserTransactions(id string) (uRes UserTransactionsResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getusertransactions", params, &uRes)
	return uRes, err
}

// UserWorkers fetch a users worker status, both
// id and username work for id. If not admin, will
// fetch current users workers.
func (m *MiningPoolHub) UserWorkers(id string) (uRes UserTransactionsResponse, err error) {
	params := &url.Values{}
	if id != "" {
		params.Add("id", id)
	}
	_, err = m.get("getuserworkers", params, &uRes)
	return uRes, err
}

// Public fetch public pool statistics, no authentication required
func (m *MiningPoolHub) Public() (pRes PublicResponse, err error) {
	_, err = m.get("public", nil, &pRes)
	return pRes, err
}
