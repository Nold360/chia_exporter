package main

type NetworkInfo struct {
	NetworkName   string `json:"network_name"`
	NetworkPrefix string `json:"network_prefix"`
	Success       bool
}

//TODO figure out what type some of these fields should be, set to interface{} for now
type BlockchainState struct {
	BlockchainState struct {
		Difficulty                  int
		GenesisChallengeInitialized bool `json:"genesis_challenge_initialized"`
		MempoolSize                 int  `json:"mempool_size"`
		Peak                        struct {
			ChallengeBlockInfoHash string `json:"challenge_block_info_hash"`
			ChallengeVDFOutput     struct {
				Data string
			} `json:"challenge_vdf_output"`
			Deficit                            int
			FarmerPuzzleHash                   string `json:"farmer_puzzle_hash"`
			Fees                               float64
			FinishedChallengeSlotHashes        interface{} `json:"finished_challenge_slot_hashes"`
			FinishedInfusedChallengeSlotHashes interface{} `json:"finished_infused_challenge_slot_hashes"`
			HeaderHash                         string      `json:"header_hash"`
			Height                             int
			InfusedChallengeVDFOutput          struct {
				Date string
			} `json:"infused_challenge_vdf_output"`
			Overflow                   bool
			PoolPuzzleHash             string      `json:"pool_puzzle_hash"`
			PrevHash                   string      `json:"prev_hash"`
			PrevTransactionBlockHash   string      `json:"prev_transaction_block_hash"`
			PrevTransactionBlockHeight int         `json:"prev_transaction_block_height"`
			RequiredIters              int         `json:"required_iters"`
			RewardClaimsIncorporated   interface{} `json:"reward_claims_incorporated"`
			RewardInfusionNewChallenge string      `json:"reward_infusion_new_challenge"`
			SignagePointIndex          int         `json:"signage_point_index"`
			SubEpochSummaryIncluded    interface{} `json:"sub_epoch_summary_included"`
			SubSlotIters               int         `json:"sub_slot_iters"`
			Timestamp                  interface{}
			TotalIters                 int `json:"total_iters"`
			Weight                     int
		}
		Space        float64
		SubSlotIters int `json:"sub_slot_iters"`
		Sync         struct {
			SyncMode           bool `json:"sync_mode"`
			SyncProgressHeight int  `json:"sync_progress_height"`
			SyncTipHeight      int  `json:"sync_tip_height"`
			Synced             bool
		}
	} `json:"blockchain_state"`
	Success bool
}

// Chia node types from server/outbound_message.py
const (
	NodeTypeNone = iota
	NodeTypeFullNode
	NodeTypeHarvester
	NodeTypeFarmer
	NodeTypeTimelord
	NodeTypeIntroducer
	NodeTypeWallet
	NumNodeTypes = 6
)

type NodeType int

type Connections struct {
	Connections []struct {
		BytesRead       int     `json:"bytes_read"`
		BytesWritten    int     `json:"bytes_written"`
		CreationTime    float64 `json:"creation_time"`
		LastMessageTime float64 `json:"last_message_time"`
		LocalPort       int     `json:"local_port"`
		NodeId          string  `json:"node_id"`
		PeakHash        string  `json:"peak_hash"`
		PeakHeight      int     `json:"peak_height"`
		PeakWeight      int     `json:"peak_weight"`
		PeerHost        string  `json:"peer_host"`
		PeerPort        int     `json:"peer_port"`
		PeerServerPort  int     `json:"peer_server_port"`
		Type            NodeType
	}
	Success bool
}

type Wallet struct {
	ID        int
	Name      string
	Type      int
	Data      string
	StringID  string
	PublicKey string
}

type Wallets struct {
	Wallets []Wallet
	Success bool
}

type WalletBalance struct {
	WalletBalance struct {
		ConfirmedBalance   int64 `json:"confirmed_wallet_balance"`
		MaxSendAmount      int64 `json:"max_send_amount"`
		PendingChange      int64 `json:"pending_change"`
		SpendableBalance   int64 `json:"spendable_balance"`
		UnconfirmedBalance int64 `json:"unconfirmed_wallet_balance"`
		WalletID           int   `json:"wallet_id"`
	} `json:"wallet_balance"`
	Success bool
}

type WalletSyncStatus struct {
	GenesisInitialized bool `json:"genesis_initialized"`
	Synced             bool
	Syncing            bool
	Succes             bool
}

type WalletHeightInfo struct {
	Height  int64
	Success bool
}

type WalletPublicKeys struct {
	PublicKeyFingerprints []int `json:"public_key_fingerprints"`
	Success               bool
}

type FarmedAmount struct {
	FarmedAmount     int64 `json:"farmed_amount"`
	RewardAmount     int64 `json:"farmer_reward_amount"`
	FeeAmount        int64 `json:"fee_amount"`
	LastHeightFarmed int64 `json:"last_height_farmed"`
	PoolRewardAmount int64 `json:"pool_reward_amount"`
	Success          bool
}

type PoolState struct {
	PoolState []struct {
		CurrentDificulty      int64        `json:"current_difficulty"`
		CurrentPoints         int64        `json:"current_points"`
		PointsAcknowledged24h [][2]float64 `json:"points_acknowledged_24h"`
		PointsFound24h        [][2]float64 `json:"points_found_24h"`
		PoolConfig            struct {
			LauncherId string `json:"launcher_id"`
			PoolURL    string `json:"pool_url"`
		} `json:"pool_config"`
	} `json:"pool_state"`
	Success bool
}

type PlotData struct {
	FileSize      int64   `json:"file_size"`
	Filename      string  `json:"filename"`
	PlotSeed      string  `json:"plot-seed"`
	PlotID        string  `json:"plot_id"`
	PublicKey     string  `json:"plot_public_key"`
	PoolContract  string  `json:"pool_contract_puzzle_hash"`
	PoolPublicKey string  `json:"pool_public_key"`
	Size          int64   `json:"size"`
	TimeModified  float64 `json:"time_modified"`
}

type PlotFiles struct {
	FailedToOpen []string   `json:"failed_to_open_filenames"`
	NotFound     []string   `json:"not_found_filenames"`
	Plots        []PlotData `json:"plots"`
	Success      bool
}

type ForkPort struct {
	Name            string  `json:"name"`
  Symbol          string  `json:symbol`
	FullNodePort    int			`json:"full_node_port"`
	WalletPort      int			`json:"wallet_port"`
	HarvesterPort   int			`json:"harvester_port"`
	FarmerPort      int			`json:"farmer_port"`
}

var forkPorts = []ForkPort {
	ForkPort {
		Name:"chia",
		Symbol: "XCH",
		FullNodePort: 8555,
		WalletPort: 9256,
		HarvesterPort: 8560,
		FarmerPort: 8559,
	},
	ForkPort {
		Name:"flax",
		Symbol: "XFX",
		FullNodePort: 6755,
		WalletPort: 6761,
		HarvesterPort: 6760,
		FarmerPort: 6759,
	},
	ForkPort {
		Name:"n-chain",
		Symbol: "NCH",
		FullNodePort: 38555,
		WalletPort: 39256,
		HarvesterPort: 38560,
		FarmerPort: 38559,
	},
	ForkPort {
		Name:"chives",
		Symbol: "XCC",
		FullNodePort: 9755,
		WalletPort: 9856,
		HarvesterPort: 9760,
		FarmerPort: 11759,
	},
	ForkPort {
		Name:"spare",
		Symbol: "SPARE",
		FullNodePort: 9555,
		WalletPort: 7256,
		HarvesterPort: 9560,
		FarmerPort: 9559,
	},
	ForkPort {
		Name:"silicoin",
		Symbol: "SIT",
		FullNodePort: 11555,
		WalletPort: 11256,
		HarvesterPort: 11560,
		FarmerPort: 11559,
	},
	ForkPort {
		Name:"flora",
		Symbol: "XFL",
		FullNodePort: 18755,
		WalletPort: 19456,
		HarvesterPort: 18760,
		FarmerPort: 18759,
	},
	ForkPort {
		Name:"hddcoin",
		Symbol: "HDD",
		FullNodePort: 28555,
		WalletPort: 29256,
		HarvesterPort: 28560,
		FarmerPort: 28559,
	},
	ForkPort {
		Name:"greendoge",
		Symbol: "GDOG",
		FullNodePort: 6655,
		WalletPort: 7356,
		HarvesterPort: 6660,
		FarmerPort: 6659,
	},
	ForkPort {
		Name:"dogechia",
		Symbol: "XDG",
		FullNodePort: 6769,
		WalletPort: 46761,
		HarvesterPort: 6770,
		FarmerPort: 46759,
	},
	ForkPort {
		Name:"apple",
		Symbol: "APPLE",
		FullNodePort: 26665,
		WalletPort: 26669,
		HarvesterPort: 26662,
		FarmerPort: 26670,
	},
	ForkPort {
		Name:"kale",
		Symbol: "XKA",
		FullNodePort: 6355,
		WalletPort: 7431,
		HarvesterPort: 6460,
		FarmerPort: 6459,
	},
	ForkPort {
		Name:"avocado",
		Symbol: "AVO",
		FullNodePort: 7544,
		WalletPort: 8753,
		HarvesterPort: 6860,
		FarmerPort: 6749,
	},
	ForkPort {
		Name:"maize",
		Symbol: "XMZ",
		FullNodePort: 8655,
		WalletPort: 8656,
		HarvesterPort: 8660,
		FarmerPort: 8659,
	},
	ForkPort {
		Name:"wheat",
		Symbol: "WHEAT",
		FullNodePort: 21555,
		WalletPort: 21256,
		HarvesterPort: 21560,
		FarmerPort: 21559,
	},
	ForkPort {
		Name:"socks",
		Symbol: "SOCK",
		FullNodePort: 58455,
		WalletPort: 59256,
		HarvesterPort: 58560,
		FarmerPort: 58559,
	},
	ForkPort {
		Name:"tad",
		Symbol: "TAD",
		FullNodePort: 4555,
		WalletPort: 4456,
		HarvesterPort: 4458,
		FarmerPort: 4457,
	},
	ForkPort {
		Name:"cryptodoge",
		Symbol: "XCD",
		FullNodePort: 16795,
		WalletPort: 16791,
		HarvesterPort: 16790,
		FarmerPort: 16799,
	},
	ForkPort {
		Name:"taco",
		Symbol: "XTX",
		FullNodePort: 18735,
		WalletPort: 19432,
		HarvesterPort: 18737,
		FarmerPort: 18736,
	},
	ForkPort {
		Name:"chiarose",
		Symbol: "XCR",
		FullNodePort: 8025,
		WalletPort: 9520,
		HarvesterPort: 8561,
		FarmerPort: 8459,
	},
	ForkPort {
		Name:"chaingreen",
		Symbol: "CGN",
		FullNodePort: 8855,
		WalletPort: 9556,
		HarvesterPort: 8860,
		FarmerPort: 8859,
	},
	ForkPort {
		Name:"melati",
		Symbol: "XMX",
		FullNodePort: 2555,
		WalletPort: 2256,
		HarvesterPort: 2560,
		FarmerPort: 2559,
	},
	ForkPort {
		Name:"cannabis",
		Symbol: "CANS",
		FullNodePort: 5540,
		WalletPort: 9656,
		HarvesterPort: 5569,
		FarmerPort: 5459,
	},
	ForkPort {
		Name:"covid",
		Symbol: "COV",
		FullNodePort: 18135,
		WalletPort: 19132,
		HarvesterPort: 18137,
		FarmerPort: 18136,
	},
	ForkPort {
		Name:"cactus",
		Symbol: "CAC",
		FullNodePort: 11555,
		WalletPort: 12256,
		HarvesterPort: 11560,
		FarmerPort: 11559,
	},
	ForkPort {
		Name:"btcgreen",
		Symbol: "XBTC",
		FullNodePort: 18942,
		WalletPort: 19544,
		HarvesterPort: 18585,
		FarmerPort: 18695,
	},
	ForkPort {
		Name:"lucky",
		Symbol: "SIX",
		FullNodePort: 16665,
		WalletPort: 16656,
		HarvesterPort: 16660,
		FarmerPort: 16659,
	},
	ForkPort {
		Name:"tranzact",
		Symbol: "TRZ",
		FullNodePort: 8673,
		WalletPort: 8675,
		HarvesterPort: 8653,
		FarmerPort: 8650,
	},
	ForkPort {
		Name:"sector",
		Symbol: "XSC",
		FullNodePort: 5555,
		WalletPort: 5256,
		HarvesterPort: 5560,
		FarmerPort: 5559,
	},
	ForkPort {
		Name:"goji",
		Symbol: "XGJ",
		FullNodePort: 7555,
		WalletPort: 8256,
		HarvesterPort: 7560,
		FarmerPort: 7559,
	},
	ForkPort {
		Name:"stai",
		Symbol: "STAI",
		FullNodePort: 8155,
		WalletPort: 1736,
		HarvesterPort: 1490,
		FarmerPort: 8553,
	},
	ForkPort {
		Name:"fork",
		Symbol: "XFK",
		FullNodePort: 16500,
		WalletPort: 17431,
		HarvesterPort: 16624,
		FarmerPort: 16525,
	},
	ForkPort {
		Name:"scam",
		Symbol: "SCM",
		FullNodePort: 9655,
		WalletPort: 9661,
		HarvesterPort: 9660,
		FarmerPort: 9773,
	},
	ForkPort {
		Name:"seno",
		Symbol: "XSE",
		FullNodePort: 18555,
		WalletPort: 19256,
		HarvesterPort: 18560,
		FarmerPort: 18559,
	},
	ForkPort {
		Name:"equality",
		Symbol: "XEQ",
		FullNodePort: 9547,
		WalletPort: 9761,
		HarvesterPort: 9760,
		FarmerPort: 9756,
	},
	ForkPort {
		Name:"skynet",
		Symbol: "XNT",
		FullNodePort: 9989,
		WalletPort: 9991,
		HarvesterPort: 9993,
		FarmerPort: 9992,
	},
	ForkPort {
		Name:"aedgecoin",
		Symbol: "AEC",
		FullNodePort: 9067,
		WalletPort: 9768,
		HarvesterPort: 9072,
		FarmerPort: 9071,
	},
	ForkPort {
		Name:"mint",
		Symbol: "XKM",
		FullNodePort: 29333,
		WalletPort: 29556,
		HarvesterPort: 29360,
		FarmerPort: 29339,
	},
}
