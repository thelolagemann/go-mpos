package mpos

import (
	"os"
	"testing"
)

var (
	pool = NewMiningPoolHub(os.Getenv("MPOS_API"), "https://ethereum.miningpoolhub.com/index.php")
)

// TODO: mock tests

func TestBlockCount(t *testing.T) {
	_, err := pool.BlockCount()
	if err != nil {
		t.Error(err)
	}
}

func TestBlocksFound(t *testing.T) {
	_, err := pool.BlocksFound()
	if err != nil {
		t.Error(err)
	}
}

func TestBlockStats(t *testing.T) {
	_, err := pool.BlockStats()
	if err != nil {
		t.Error(err)
	}
}

func TestCurrentWorkers(t *testing.T) {
	_, err := pool.CurrentWorkers()
	if err != nil {
		t.Error(err)
	}
}

func TestDashboardData(t *testing.T) {
	_, err := pool.DashboardData()
	if err != nil {
		t.Error(err)
	}
}

func TestDifficulty(t *testing.T) {
	_, err := pool.Difficulty()
	if err != nil {
		t.Error(err)
	}
}

func TestEstimatedTime(t *testing.T) {
	_, err := pool.EstimatedTime()
	if err != nil {
		t.Error(err)
	}
}

func TestPoolHashrate(t *testing.T) {
	_, err := pool.PoolHashrate()
	if err != nil {
		t.Error(err)
	}
}

func TestPoolInfo(t *testing.T) {
	_, err := pool.PoolInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestPoolStatus(t *testing.T) {
	_, err := pool.PoolStatus()
	if err != nil {
		t.Error(err)
	}
}

func TestTimeSinceLastBlock(t *testing.T) {
	_, err := pool.TimeSinceLastBlock()
	if err != nil {
		t.Error(err)
	}
}

func TestTopContributors(t *testing.T) {
	_, err := pool.TopContributors()
	if err != nil {
		t.Error(err)
	}
}

func TestUserBalance(t *testing.T) {
	_, err := pool.UserBalance("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserBalance("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestUserHashrate(t *testing.T) {
	_, err := pool.UserHashrate("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserHashrate("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestUserShareRate(t *testing.T) {
	_, err := pool.UserShareRate("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserShareRate("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestUserStatus(t *testing.T) {
	_, err := pool.UserStatus("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserStatus("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestUserTransactions(t *testing.T) {
	_, err := pool.UserTransactions("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserTransactions("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}

	_, err = pool.UserTransactions("")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestUserWorkers(t *testing.T) {
	_, err := pool.UserWorkers("")
	if err != nil {
		t.Error(err)
	}

	_, err = pool.UserWorkers("invalid-perms")
	if err == nil {
		t.Error("expecting error, got none")
	}
}

func TestPublic(t *testing.T) {
	_, err := pool.Public()
	if err != nil {
		t.Error(err)
	}
}
