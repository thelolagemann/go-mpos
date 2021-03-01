// Package mpos is an unofficial PHP-MPOS API client for golang.
package mpos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var (
	client httpClient
)

// MiningPoolHub is an API client for php-mpos.
type MiningPoolHub struct {
	// BaseURL is the base URL that all requests are sent to.
	// Useful for instances that have multiple pools under
	// different subdomains whilst utilizing the same API key.
	BaseURL string

	apiKey string
	client httpClient
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewMiningPoolHub returns a MiningPoolHub client embedded with
// the provided api key and baseurl. A httpClient is embedded that
// is used to send all requests.
func NewMiningPoolHub(apiKey string, baseURL string) *MiningPoolHub {
	return &MiningPoolHub{
		apiKey:  apiKey,
		BaseURL: baseURL,
		client:  client,
	}
}

func (m *MiningPoolHub) get(action string, params *url.Values, bind interface{}) (*http.Response, error) {
	// build url
	base, err := url.Parse(m.BaseURL)
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

	// send get request
	req, err := http.NewRequest("GET", base.String(), nil)
	if err != nil {
		return nil, err
	}
	res, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%v returned non-ok http code, %v", action, res.StatusCode)
	}

	// if bind
	if bind != nil {
		var decode response

		if err := json.NewDecoder(res.Body).Decode(&decode); err != nil {

			return nil, fmt.Errorf("error decoding json: %v", err)
		}

		if val, ok := decode[action]; ok {
			if err := json.Unmarshal(val.Data, &bind); err != nil {
				return nil, fmt.Errorf("error unmarshalling json: %v, body: %v", err, string(val.Data))
			}
		}

	}

	return res, err
}

// BlockCount get current block height in blockchain.
func (m *MiningPoolHub) BlockCount() (blockCount json.Number, err error) {
	_, err = m.get("getblockcount", nil, &blockCount)
	return blockCount, err
}

// BlocksFound get last N blocks found as configured in
// admin panel.
func (m *MiningPoolHub) BlocksFound() (bRes *BlocksFoundResponse, err error) {
	_, err = m.get("getblocksfound", nil, &bRes)
	return bRes, err
}

// BlockStats get pool block stats.
func (m *MiningPoolHub) BlockStats() (bRes *BlockStatsResponse, err error) {
	_, err = m.get("getblockstats", nil, &bRes)
	return bRes, err
}

// CurrentWorkers get amount of current active workers.
func (m *MiningPoolHub) CurrentWorkers() (cRes uint64, err error) {
	_, err = m.get("getcurrentworkers", nil, &cRes)
	return cRes, err
}

// DashboardData fetch all dashboard related information.
func (m *MiningPoolHub) DashboardData() (dRes *DashboardDataResponse, err error) {
	_, err = m.get("getdashboarddata", nil, &dRes)
	return dRes, err
}

// Difficulty get current difficulty in blockchain.
func (m *MiningPoolHub) Difficulty() (dRes json.Number, err error) {
	_, err = m.get("getdifficulty", nil, &dRes)
	return dRes, err
}

// EstimatedTime get estimated time to next block based on
// pool hashrate (seconds).
func (m *MiningPoolHub) EstimatedTime() (eRes float64, err error) {
	_, err = m.get("getestimatedtime", nil, &eRes)
	return eRes, err
}

// NavbarData - disabled on miningpoolhub.com
func (m *MiningPoolHub) NavbarData() {
	// TBI
}

// PoolHashrate get current pool hashrate.
func (m *MiningPoolHub) PoolHashrate() (pRes float64, err error) {
	_, err = m.get("getpoolhashrate", nil, &pRes)
	return pRes, err
}

// PoolInfo get the information on pool settings.
func (m *MiningPoolHub) PoolInfo() (pRes *PoolInfoResponse, err error) {
	_, err = m.get("getpoolinfo", nil, &pRes)
	return pRes, err
}

// PoolShareRate - always returns 0 on miningpoolhub.com
func (m *MiningPoolHub) PoolShareRate() {
	// TBI
}

// PoolStatus fetch overall pool status
func (m *MiningPoolHub) PoolStatus() (pRes *PoolStatusResponse, err error) {
	_, err = m.get("getpoolstatus", nil, &pRes)
	return pRes, err
}

// TimeSinceLastBlock get time since last block found (seconds)
func (m *MiningPoolHub) TimeSinceLastBlock() (tRes uint, err error) {
	_, err = m.get("gettimesincelastblock", nil, &tRes)
	return tRes, err
}

// TopContributors fetch top contributors data
func (m *MiningPoolHub) TopContributors() (tRes *TopContributorsResponse, err error) {
	_, err = m.get("gettopcontributors", nil, &tRes)
	return tRes, err
}

// User returns a MiningPoolHubUser struct.
func (m *MiningPoolHub) User(id int32) *MiningPoolHubUser { return &MiningPoolHubUser{m, id} }

// Account returns a MiningPoolHubUser embedded with
// your own account ID. It does this by first sending
// a request to the getuserstatus endpoint, and
// accessing your ID from the result. AFAIK this is the
// only reliable way to acquire your own ID.
func (m *MiningPoolHub) Account() (*MiningPoolHubUser, error) {
	var status UserStatusResponse
	_, err := m.get("getuserstatus", nil, &status)
	if err != nil {
		return nil, err
	}
	return &MiningPoolHubUser{
		id: status.Shares.ID,
	}, nil
}

// Public fetch public pool statistics, no authentication required
func (m *MiningPoolHub) Public() (pRes *PublicResponse, err error) {
	res, err := m.get("public", nil, nil)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(res.Body).Decode(&pRes)
	return pRes, err
}
