package mpos

import (
	"net/http"
	"net/url"
)

// MiningPoolHubUser is a convenience struct that wraps
// around MiningPoolHub, embedding an id alongside it to
// be used with API actions pertaining to user interaction
// and management.
type MiningPoolHubUser struct {
	*MiningPoolHub
	id int32
}

func (m *MiningPoolHubUser) getUser(action string, bind interface{}) (*http.Response, error) {
	params := &url.Values{}
	params.Add("id", string(m.id))
	return m.get(action, params, bind)
}

// Balance fetch the users balance.
func (m *MiningPoolHubUser) Balance() (uRes *UserBalanceResponse, err error) {
	_, err = m.getUser("getuserbalance", &uRes)
	return uRes, err
}

// Hashrate fetch the users hash rate.
func (m *MiningPoolHubUser) Hashrate() (uRes float64, err error) {
	_, err = m.getUser("getuserhashrate", &uRes)
	return uRes, err
}

// Sharerate fetch the users share rate.
func (m *MiningPoolHubUser) Sharerate() (uRes float64, err error) {
	_, err = m.getUser("getusersharerate", &uRes)
	return uRes, err
}

// Status fetch the users overall status.
func (m *MiningPoolHubUser) Status() (uRes UserStatusResponse, err error) {
	_, err = m.getUser("getuserstatus", &uRes)
	return uRes, err
}

// Transactions get the users transactions.
func (m *MiningPoolHubUser) Transactions() (uRes UserTransactionsResponse, err error) {
	_, err = m.getUser("getusertransactions", &uRes)
	return uRes, err
}

// Workers fetch the users worker status.
func (m *MiningPoolHubUser) Workers() (uRes UserWorkersResponse, err error) {
	_, err = m.getUser("getuserworkers", &uRes)
	return uRes, err
}
