package mpos

import "encoding/json"

type baseResponse struct {
	Version string          `json:"version"`
	Runtime float64         `json:"runtime"`
	Data    json.RawMessage `json:"data"`
}

type response map[string]baseResponse

// BlocksFoundResponse is the json response returned
// by the getblocksfound API action for php-mpos.
type BlocksFoundResponse []struct {
	ID            int     `json:"id"`
	Height        int     `json:"height"`
	Blockhash     string  `json:"blockhash"`
	Confirmations int     `json:"confirmations"`
	Amount        float64 `json:"amount"`
	Difficulty    float64 `json:"difficulty"`
	Time          int     `json:"time"`
	Accounted     int     `json:"accounted"`
	AccountID     int     `json:"account_id"`
	WorkerName    string  `json:"worker_name"`
	Shares        int     `json:"shares"`
	ShareID       int64   `json:"share_id"`
	Finder        string  `json:"finder"`
	IsAnonymous   int     `json:"is_anonymous"`
	Estshares     int     `json:"estshares"`
}

// BlockStatsResponse is the json response returned
// by the getblockstats API action for php-mpos.
type BlockStatsResponse struct {
	Total                    int     `json:"Total"`
	TotalValid               string  `json:"TotalValid"`
	TotalOrphan              string  `json:"TotalOrphan"`
	TotalDifficulty          float64 `json:"TotalDifficulty"`
	TotalShares              string  `json:"TotalShares"`
	TotalEstimatedShares     int64   `json:"TotalEstimatedShares"`
	TotalAmount              float64 `json:"TotalAmount"`
	OneHourTotal             string  `json:"1HourTotal"`
	OneHourValid             string  `json:"1HourValid"`
	OneHourOrphan            string  `json:"1HourOrphan"`
	OneHourDifficulty        float64 `json:"1HourDifficulty"`
	OneHourShares            string  `json:"1HourShares"`
	OneHourEstimatedShares   int     `json:"1HourEstimatedShares"`
	OneHourAmount            float64 `json:"1HourAmount"`
	Two4HourTotal            string  `json:"24HourTotal"`
	Two4HourValid            string  `json:"24HourValid"`
	Two4HourOrphan           string  `json:"24HourOrphan"`
	Two4HourDifficulty       float64 `json:"24HourDifficulty"`
	Two4HourShares           string  `json:"24HourShares"`
	Two4HourEstimatedShares  int64   `json:"24HourEstimatedShares"`
	Two4HourAmount           float64 `json:"24HourAmount"`
	SevenDaysTotal           string  `json:"7DaysTotal"`
	SevenDaysValid           string  `json:"7DaysValid"`
	SevenDaysOrphan          string  `json:"7DaysOrphan"`
	SevenDaysDifficulty      float64 `json:"7DaysDifficulty"`
	SevenDaysShares          string  `json:"7DaysShares"`
	SevenDaysEstimatedShares int64   `json:"7DaysEstimatedShares"`
	SevenDaysAmount          float64 `json:"7DaysAmount"`
	FourWeeksTotal           string  `json:"4WeeksTotal"`
	FourWeeksValid           string  `json:"4WeeksValid"`
	FourWeeksOrphan          string  `json:"4WeeksOrphan"`
	FourWeeksDifficulty      float64 `json:"4WeeksDifficulty"`
	FourWeeksShares          string  `json:"4WeeksShares"`
	FourWeeksEstimatedShares int64   `json:"4WeeksEstimatedShares"`
	FourWeeksAmount          float64 `json:"4WeeksAmount"`
	One2MonthTotal           string  `json:"12MonthTotal"`
	One2MonthValid           string  `json:"12MonthValid"`
	One2MonthOrphan          string  `json:"12MonthOrphan"`
	One2MonthDifficulty      float64 `json:"12MonthDifficulty"`
	One2MonthShares          string  `json:"12MonthShares"`
	One2MonthEstimatedShares int64   `json:"12MonthEstimatedShares"`
	One2MonthAmount          float64 `json:"12MonthAmount"`
}

// DashboardDataResponse is the json response returned
// by the getdashboarddata API action for php-mpos.
type DashboardDataResponse struct {
	Raw struct {
		Personal struct {
			Hashrate float64 `json:"hashrate"`
		} `json:"personal"`
		Pool struct {
			Hashrate float64 `json:"hashrate"`
		} `json:"pool"`
		Network struct {
			Hashrate        float64 `json:"hashrate"`
			Esttimeperblock int     `json:"esttimeperblock"`
		} `json:"network"`
	} `json:"raw"`
	Personal struct {
		Hashrate        float64 `json:"hashrate"`
		Sharerate       int     `json:"sharerate"`
		Sharedifficulty int     `json:"sharedifficulty"`
		Shares          struct {
			Valid          float64 `json:"valid"`
			Invalid        float64 `json:"invalid"`
			InvalidPercent float64 `json:"invalid_percent"`
			Unpaid         float64 `json:"unpaid"`
		} `json:"shares"`
		Estimates struct {
			Block    float64 `json:"block"`
			Fee      float64 `json:"fee"`
			Donation float64 `json:"donation"`
			Payout   float64 `json:"payout"`
		} `json:"estimates"`
	} `json:"personal"`
	Balance struct {
		Confirmed   float64 `json:"confirmed"`
		Unconfirmed float64 `json:"unconfirmed"`
	} `json:"balance"`
	BalanceForAutoExchange struct {
		Confirmed   int `json:"confirmed"`
		Unconfirmed int `json:"unconfirmed"`
	} `json:"balance_for_auto_exchange"`
	BalanceOnExchange    int `json:"balance_on_exchange"`
	RecentCredits24Hours struct {
		Amount float64 `json:"amount"`
	} `json:"recent_credits_24hours"`
	RecentCredits []struct {
		Date   string  `json:"date"`
		Amount float64 `json:"amount"`
	} `json:"recent_credits"`
	Pool struct {
		Info struct {
			Name     string `json:"name"`
			Currency string `json:"currency"`
		} `json:"info"`
		Workers  int     `json:"workers"`
		Hashrate float64 `json:"hashrate"`
		Shares   struct {
			Valid          float64 `json:"valid"`
			Invalid        float64 `json:"invalid"`
			InvalidPercent float64 `json:"invalid_percent"`
			Estimated      float64 `json:"estimated"`
			Progress       float64 `json:"progress"`
		} `json:"shares"`
		Price      string `json:"price"`
		Difficulty int    `json:"difficulty"`
		TargetBits int    `json:"target_bits"`
	} `json:"pool"`
	System struct {
		Load []float64 `json:"load"`
	} `json:"system"`
	Network struct {
		Hashrate        float64     `json:"hashrate"`
		Difficulty      json.Number `json:"difficulty"`
		Block           json.Number `json:"block"`
		Esttimeperblock int         `json:"esttimeperblock"`
	} `json:"network"`
}

// PoolInfoResponse is the json response returned
// by the getpoolinfo API action for php-mpos.
type PoolInfoResponse struct {
	Currency             string  `json:"currency"`
	Coinname             string  `json:"coinname"`
	Cointarget           string  `json:"cointarget"`
	Coindiffchangetarget int     `json:"coindiffchangetarget"`
	Stratumport          string  `json:"stratumport"`
	PayoutSystem         string  `json:"payout_system"`
	Confirmations        int     `json:"confirmations"`
	MinApThreshold       float64 `json:"min_ap_threshold"`
	MaxApThreshold       int     `json:"max_ap_threshold"`
	RewardType           string  `json:"reward_type"`
	Reward               int     `json:"reward"`
	Txfee                float64 `json:"txfee"`
	TxfeeManual          float64 `json:"txfee_manual"`
	TxfeeAuto            float64 `json:"txfee_auto"`
	Fees                 float64 `json:"fees"`
}

// PoolStatusResponse is the json response returned
// by the getpoolstatus API action for php-mpos.
type PoolStatusResponse struct {
	PoolName            string  `json:"pool_name"`
	Hashrate            float64 `json:"hashrate"`
	Efficiency          float64 `json:"efficiency"`
	Workers             int     `json:"workers"`
	Currentnetworkblock string  `json:"currentnetworkblock"`
	Nextnetworkblock    int     `json:"nextnetworkblock"`
	Lastblock           int     `json:"lastblock"`
	Networkdiff         string  `json:"networkdiff"`
	Esttime             float64 `json:"esttime"`
	Estshares           float64 `json:"estshares"`
	Timesincelast       int     `json:"timesincelast"`
	Nethashrate         float64 `json:"nethashrate"`
}

// TopContributorsResponse is the json response returned
// by the gettopcontributors API action for php-mpos.
type TopContributorsResponse struct {
	Hashes struct {
		Account  string  `json:"account"`
		Hashrate float64 `json:"hashrate"`
	} `json:"hashes"`
}

// UserBalanceResponse is the json response returned
// by the getuserbalance API action for php-mpos.
type UserBalanceResponse struct {
	Confirmed   float64 `json:"confirmed"`
	Unconfirmed float64 `json:"unconfirmed"`
}

// UserStatusResponse is the json response returned
// by the getuserstatus API action for php-mpos.
type UserStatusResponse struct {
	Username string `json:"username"`
	Shares   struct {
		Valid         int     `json:"valid"`
		Valid1        int     `json:"valid1"`
		Valid2        int     `json:"valid2"`
		Valid3        int     `json:"valid3"`
		Valid4        int     `json:"valid4"`
		Valid5        int     `json:"valid5"`
		Valid6        int     `json:"valid6"`
		Invalid       int     `json:"invalid"`
		ID            int32   `json:"id"`
		DonatePercent float64 `json:"donate_percent"`
		IsAnonymous   int     `json:"is_anonymous"`
		Username      string  `json:"username"`
	} `json:"shares"`
	Hashrate  float64 `json:"hashrate"`
	Sharerate int     `json:"sharerate"`
}

// UserTransactionsResponse is the json response returned
// by the getuserstatus API action for php-mpos.
type UserTransactionsResponse struct {
	Transactions []struct {
		ID            int64       `json:"id"`
		Username      string      `json:"username"`
		Type          string      `json:"type"`
		Amount        float64     `json:"amount"`
		CoinAddress   interface{} `json:"coin_address"`
		Timestamp     string      `json:"timestamp"`
		Txid          interface{} `json:"txid"`
		Height        int         `json:"height"`
		Blockhash     string      `json:"blockhash"`
		Confirmations int         `json:"confirmations"`
	} `json:"transactions"`
}

// UserWorkersResponse is the json response returned
// by the getuserworkers API action for php-mpos.
type UserWorkersResponse []struct {
	ID         int     `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Monitor    int     `json:"monitor"`
	Hashrate   float64 `json:"hashrate"`
	Difficulty int     `json:"difficulty"`
}

// PublicResponse is the json response returned
// by the public API action for php-mpos.
type PublicResponse struct {
	PoolName        string  `json:"pool_name"`
	Hashrate        float64 `json:"hashrate"`
	Workers         int     `json:"workers"`
	SharesThisRound float64 `json:"shares_this_round"`
	LastBlock       int     `json:"last_block"`
	NetworkHashrate float64 `json:"network_hashrate"`
}
