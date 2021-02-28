package mpos

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var (
	pool *MiningPoolHub

	mockFunc    func(req *http.Request) (*http.Response, error)
	mockResults map[string]json.RawMessage
)

type mockClient struct{}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return mockFunc(req)
}

func init() {
	// determine testing mode
	if os.Getenv("MPOS_API") == "" {
		pool = &MiningPoolHub{apiKey: "", BaseURL: "https://example.mpos.com", client: &mockClient{}}
	} else {
		pool = NewMiningPoolHub(os.Getenv("MPOS_API"), "https://ethereum.miningpoolhub.com/index.php")
	}

	// load mock results
	b, err := os.ReadFile("responses.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&mockResults); err != nil {
		panic(err)
	}
}

func mockFuncBody(body []byte) {
	r := ioutil.NopCloser(bytes.NewReader(body))
	mockFunc = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
}

func TestBlockCount(t *testing.T) {
	mockFuncBody(mockResults["getblockcount"])

	if _, err := pool.BlockCount(); err != nil {
		t.Error(err)
	}
}

func TestBlocksFound(t *testing.T) {
	mockFuncBody(mockResults["getblocksfound"])

	if _, err := pool.BlocksFound(); err != nil {
		t.Error(err)
	}
}

func TestBlockStats(t *testing.T) {
	mockFuncBody(mockResults["getblockstats"])
	if _, err := pool.BlockStats(); err != nil {
		t.Error(err)
	}
}

func TestCurrentWorkers(t *testing.T) {
	mockFuncBody(mockResults["getcurrentworkers"])
	if _, err := pool.CurrentWorkers(); err != nil {
		t.Error(err)
	}
}

func TestDashboardData(t *testing.T) {
	mockFuncBody(mockResults["getdashboarddata"])
	if _, err := pool.DashboardData(); err != nil {
		t.Error(err)
	}
}

func TestDifficulty(t *testing.T) {
	mockFuncBody(mockResults["getdifficulty"])
	if _, err := pool.Difficulty(); err != nil {
		t.Error(err)
	}
}

func TestEstimatedTime(t *testing.T) {
	mockFuncBody(mockResults["getestimatedtime"])
	if _, err := pool.EstimatedTime(); err != nil {
		t.Error(err)
	}
}

func TestPoolHashrate(t *testing.T) {
	mockFuncBody(mockResults["getpoolhashrate"])
	if _, err := pool.PoolHashrate(); err != nil {
		t.Error(err)
	}
}

func TestPoolInfo(t *testing.T) {
	mockFuncBody(mockResults["getpoolinfo"])
	if _, err := pool.PoolInfo(); err != nil {
		t.Error(err)
	}
}

func TestPoolStatus(t *testing.T) {
	mockFuncBody(mockResults["getpoolstatus"])
	if _, err := pool.PoolStatus(); err != nil {
		t.Error(err)
	}
}

func TestTimeSinceLastBlock(t *testing.T) {
	mockFuncBody(mockResults["gettimesincelastblock"])
	if _, err := pool.TimeSinceLastBlock(); err != nil {
		t.Error(err)
	}
}

func TestTopContributors(t *testing.T) {
	mockFuncBody(mockResults["gettopcontributors"])
	if _, err := pool.TopContributors(); err != nil {
		t.Error(err)
	}
}

func TestUserBalance(t *testing.T) {
	mockFuncBody(mockResults["getuserbalance"])
	if _, err := pool.UserBalance(""); err != nil {
		t.Error(err)
	}
}

func TestUserHashrate(t *testing.T) {
	mockFuncBody(mockResults["getuserhashrate"])
	if _, err := pool.UserHashrate(""); err != nil {
		t.Error(err)
	}
}

func TestUserShareRate(t *testing.T) {
	mockFuncBody(mockResults["getusersharerate"])
	if _, err := pool.UserShareRate(""); err != nil {
		t.Error(err)
	}
}

func TestUserStatus(t *testing.T) {
	mockFuncBody(mockResults["getuserstatus"])
	if _, err := pool.UserStatus(""); err != nil {
		t.Error(err)
	}
}

func TestUserTransactions(t *testing.T) {
	mockFuncBody(mockResults["getusertransactions"])
	if _, err := pool.UserTransactions("1310"); err != nil {
		t.Error(err)
	}

	if _, err := pool.UserTransactions(""); err == nil {
		t.Error("expecting error, got none")
	}

}

func TestUserWorkers(t *testing.T) {
	mockFuncBody(mockResults["getuserworkers"])
	if _, err := pool.UserWorkers(""); err != nil {
		t.Error(err)
	}
}

func TestPublic(t *testing.T) {
	mockFuncBody(mockResults["getpublic"])
	if _, err := pool.Public(); err != nil {
		t.Error(err)
	}
}
