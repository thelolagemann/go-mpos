package mpos

import (
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

func (m *MiningPoolHubUser) getUser(action string, bind interface{}) error {
	params := &url.Values{}
	params.Add("id", string(m.id))
	_, err := m.get(action, params, bind)
	return err
}

// Balance fetch the users balance.
func (m *MiningPoolHubUser) Balance() (uRes *UserBalanceResponse, err error) {
	return uRes, m.getUser("getuserbalance", &uRes)
}

// Hashrate fetch the users hash rate.
func (m *MiningPoolHubUser) Hashrate() (uRes float64, err error) {
	return uRes, m.getUser("getuserhashrate", &uRes)
}

// Sharerate fetch the users share rate.
func (m *MiningPoolHubUser) Sharerate() (uRes float64, err error) {
	return uRes, m.getUser("getusersharerate", &uRes)
}

// Status fetch the users overall status.
func (m *MiningPoolHubUser) Status() (uRes UserStatusResponse, err error) {
	return uRes, m.getUser("getuserstatus", &uRes)
}

// Transactions get the users transactions.
func (m *MiningPoolHubUser) Transactions() (uRes UserTransactionsResponse, err error) {
	return uRes, m.getUser("getusertransactions", &uRes)
}

// Workers fetch the users worker status.
func (m *MiningPoolHubUser) Workers() (uRes UserWorkersResponse, err error) {
	return uRes, m.getUser("getuserworkers", &uRes)
}
