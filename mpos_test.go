package mpos

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	if action := req.URL.Query().Get("action"); action != "" {
		if val, ok := mockResults[action]; ok {
			r := ioutil.NopCloser(bytes.NewReader(val))
			return func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			}(req)
		}
	}
	return mockFunc(req)
}

func init() {
	// determine testing mode
	if os.Getenv("MPOS_API") == "" {
		client = &mockClient{}
	} else {
		client = &http.Client{}
	}
	pool = NewMiningPoolHub(os.Getenv("MPOS_API"), "BASE_URL")

	// load mock results
	b, err := os.ReadFile("responses.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&mockResults); err != nil {
		panic(err)
	}
}

func TestAccount(t *testing.T) {
	if _, err := pool.Account(); err != nil {
		t.Error(err)
	}
}

func TestBlockCount(t *testing.T) {
	if _, err := pool.BlockCount(); err != nil {
		t.Error(err)
	}
}

func TestBlocksFound(t *testing.T) {
	if _, err := pool.BlocksFound(); err != nil {
		t.Error(err)
	}
}

func TestBlockStats(t *testing.T) {
	if _, err := pool.BlockStats(); err != nil {
		t.Error(err)
	}
}

func TestCurrentWorkers(t *testing.T) {
	if _, err := pool.CurrentWorkers(); err != nil {
		t.Error(err)
	}
}

func TestDashboardData(t *testing.T) {
	if _, err := pool.DashboardData(); err != nil {
		t.Error(err)
	}
}

func TestDifficulty(t *testing.T) {
	if _, err := pool.Difficulty(); err != nil {
		t.Error(err)
	}
}

func TestEstimatedTime(t *testing.T) {
	if _, err := pool.EstimatedTime(); err != nil {
		t.Error(err)
	}
}

func TestPoolHashrate(t *testing.T) {
	if _, err := pool.PoolHashrate(); err != nil {
		t.Error(err)
	}
}

func TestPoolInfo(t *testing.T) {
	if _, err := pool.PoolInfo(); err != nil {
		t.Error(err)
	}
}

func TestPoolStatus(t *testing.T) {
	if _, err := pool.PoolStatus(); err != nil {
		t.Error(err)
	}
}

func TestTimeSinceLastBlock(t *testing.T) {
	if _, err := pool.TimeSinceLastBlock(); err != nil {
		t.Error(err)
	}
}

func TestTopContributors(t *testing.T) {
	if _, err := pool.TopContributors(); err != nil {
		t.Error(err)
	}
}

// TODO: handle live-testing non-admin user
func TestUserBalance(t *testing.T) {
	if _, err := pool.User(1000).Balance(); err != nil {
		t.Error(err)
	}
}

func TestUserHashrate(t *testing.T) {
	if _, err := pool.User(1000).Hashrate(); err != nil {
		t.Error(err)
	}
}

func TestUserShareRate(t *testing.T) {
	if _, err := pool.User(1000).Sharerate(); err != nil {
		t.Error(err)
	}
}

func TestUserStatus(t *testing.T) {
	if _, err := pool.User(1000).Status(); err != nil {
		t.Error(err)
	}
}

func TestUserTransactions(t *testing.T) {
	if _, err := pool.User(1000).Transactions(); err != nil {
		t.Error(err)
	}

}

func TestUserWorkers(t *testing.T) {
	if _, err := pool.User(1000).Workers(); err != nil {
		t.Error(err)
	}
}

func TestPublic(t *testing.T) {
	if _, err := pool.Public(); err != nil {
		t.Error(err)
	}
}

func ExampleMiningPoolHub() {
	pool := NewMiningPoolHub("API_KEY", "BASE_URL")

	info, err := pool.PoolInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Pool is mining %v, with a %v%% fee\n", info.Coinname, info.Fees)

	// Output: Pool is mining Ethereum, with a 0.9% fee
}
