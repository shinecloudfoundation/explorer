package lcd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shinecloudnet/explorer/backend/utils"
)

const (
	UrlAccount                                   = "%s/auth/accounts/%s"
	UrlBankTokenStats                            = "%s/supply/total"
	UrlValidator                                 = "%s/staking/validators/%s"
	UrlValidators                                = "%s/staking/validators?status=%s&page=%d&limit=%d"
	UrlDelegationByVal                           = "%s/staking/validators/%s/delegations"
	UrlDelegationsByDelegator                    = "%s/staking/delegators/%s/delegations"
	UrlUnbondingDelegationByDelegator            = "%s/staking/delegators/%s/unbonding_delegations"
	UrlDelegationsByValidator                    = "%s/staking/validators/%s/delegations"
	UrlUnbondingDelegationByValidator            = "%s/staking/validators/%s/unbonding_delegations"
	UrlRedelegationsByValidator                  = "%s/staking/validators/%s/redelegations"
	UrlSignInfo                                  = "%s/slashing/validators/%s/signing-info"
	UrlNodeInfo                                  = "%s/node_info"
	UrlNodeVersion                               = "%s/version"
	UrlGenesis                                   = "%s/genesis"
	UrlWithdrawAddress                           = "%s/distribution/delegators/%s/withdraw_address"
	UrlBlockLatest                               = "%s/blocks/latest"
	UrlBlock                                     = "%s/blocks/%d"
	UrlValidatorSet                              = "%s/validatorsets/%d"
	UrlValidatorSetLatest                        = "%s/validatorsets/latest"
	UrlStakePool                                 = "%s/staking/pool"
	UrlBlocksResult                              = "%s/blocks_results/%d"
	UrlGovParam                                  = "%s/params?module=%s"
	UrlDistributionRewardsByValidatorAcc         = "%s/distribution/delegators/%s/rewards"
	UrlValidatorsSigningInfoByConsensuPublicKey  = "%s/slashing/validators/%s/signing-info"
	UrlDistributionWithdrawAddressByValidatorAcc = "%s/distribution/delegators/%s/withdraw_address"
	UrlDistributionInfoOfValidatorAcc            = "%s/distribution/validators/%s"
	UrlTokenStatsSupply                          = "https://rpc.irisnet.org/token-stats/supply"
	UrlTokenStatsCirculation                     = "https://rpc.irisnet.org/token-stats/circulation"
	UrlLookupIconsByKeySuffix                    = "https://keybase.io/_/api/1.0/user/lookup.json?fields=pictures&key_suffix=%s"
	UrlAssetTokens                               = "%s/asset/list"
	UrlAssetGateways                             = "%s/asset/gateways"
	UrlHtlcInfo                                  = "%s/htlc/htlcs/%s"
	UrlProposalVoters                            = "%s/gov/proposals/%v/votes"
	CommunityTaxAddr                             = "iaa18rtw90hxz4jsgydcusakz6q245jh59kfma3e5h"

	NativeTokenDenom      = "uscds"
	NativeTokenPrecision  = 1000000
	UrlBalance            = "%s/bank/balances/%s"
	MintedTokenLockedAddr = "scloud1m3h30wlvsf8llruxtpukdvsy0km2kum80quuau"
	UrlTotalSupply        = "%s/supply/total"
	UrlTotalSupplyDenom   = "%s/supply/total/%s"
)

type AccountVo struct {
	Address   string   `json:"address"`
	Coins     []string `json:"coins"`
	PublicKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"public_key"`
	AccountNumber string `json:"account_number"`
	Sequence      string `json:"sequence"`
}

type StringResponse struct {
	Height string `json:"height"`
	Result string `json:"result"`
}

type ValidatorVoResponse struct {
	Height string         `json:"height"`
	Result ValidatorVoNew `json:"result"`
}

type ValidatorsVoResponse struct {
	Height string           `json:"height"`
	Result []ValidatorVoNew `json:"result"`
}

type ValidatorVo struct {
	OperatorAddress string      `json:"operator_address"`
	ConsensusPubkey string      `json:"consensus_pubkey"`
	Jailed          bool        `json:"jailed"`
	Status          int         `json:"status"`
	Tokens          string      `json:"tokens"`
	DelegatorShares string      `json:"delegator_shares"`
	Description     Description `json:"description"`
	BondHeight      string      `json:"bond_height"`
	UnbondingHeight string      `json:"unbonding_height"`
	UnbondingTime   time.Time   `json:"unbonding_time"`
	Commission      Commission  `json:"commission"`
	Uptime          float32     `json:"uptime"`
	SelfBond        string      `json:"self_bond"`
	DelegatorNum    int         `json:"delegator_num"`
	ProposerAddr    string      `json:"proposer_addr"`
	VotingRate      float32     `json:"voting_rate"`
	Icons           string      `json:"icons"`
}

type ValidatorVoNew struct {
	OperatorAddress string        `json:"operator_address"`
	ConsensusPubkey string        `json:"consensus_pubkey"`
	Jailed          bool          `json:"jailed"`
	Status          int           `json:"status"`
	Tokens          string        `json:"tokens"`
	DelegatorShares string        `json:"delegator_shares"`
	Description     Description   `json:"description"`
	BondHeight      string        `json:"bond_height"`
	UnbondingHeight string        `json:"unbonding_height"`
	UnbondingTime   time.Time     `json:"unbonding_time"`
	Commission      CommissionNew `json:"commission"`
	Uptime          float32       `json:"uptime"`
	SelfBond        string        `json:"self_bond"`
	DelegatorNum    int           `json:"delegator_num"`
	ProposerAddr    string        `json:"proposer_addr"`
	VotingRate      float32       `json:"voting_rate"`
	Icons           string        `json:"icons"`
}

func (v ValidatorVoNew) toValidatorVo() ValidatorVo {
	return ValidatorVo{
		OperatorAddress: v.OperatorAddress,
		ConsensusPubkey: v.ConsensusPubkey,
		Jailed:          v.Jailed,
		Status:          v.Status,
		Tokens:          v.Tokens,
		DelegatorShares: v.DelegatorShares,
		Description:     v.Description,
		BondHeight:      v.BondHeight,
		UnbondingHeight: v.UnbondingHeight,
		UnbondingTime:   v.UnbondingTime,
		Commission: Commission{
			Rate:          v.Commission.CommissionRates.Rate,
			MaxRate:       v.Commission.CommissionRates.MaxRate,
			MaxChangeRate: v.Commission.CommissionRates.MaxChangeRate,
			UpdateTime:    v.Commission.UpdateTime,
		},
		Uptime:       v.Uptime,
		SelfBond:     v.SelfBond,
		DelegatorNum: v.DelegatorNum,
		ProposerAddr: v.ProposerAddr,
		VotingRate:   v.VotingRate,
		Icons:        v.Icons,
	}
}

func (v ValidatorVo) String() string {
	return fmt.Sprintf(`
		OperatorAddress :%v
		ConsensusPubkey :%v
		Jailed          :%v
		Status          :%v
		Tokens          :%v
		DelegatorShares :%v
		Description     :%v
		BondHeight      :%v
		UnbondingHeight :%v
		UnbondingTime   :%v
		Commission      :%v
		Uptime          :%v
		SelfBond        :%v
		DelegatorNum    :%v
		ProposerAddr    :%v
		VotingRate      :%v
		Icons           :%v
		`, v.OperatorAddress, v.ConsensusPubkey, v.Jailed, v.Status, v.Tokens, v.DelegatorShares, v.Description, v.BondHeight, v.UnbondingHeight, v.UnbondingTime,
		v.Commission, v.Uptime, v.SelfBond, v.DelegatorNum, v.ProposerAddr, v.VotingRate, v.Icons)
}

type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
}
type Commission struct {
	Rate          string    `json:"rate"`
	MaxRate       string    `json:"max_rate"`
	MaxChangeRate string    `json:"max_change_rate"`
	UpdateTime    time.Time `json:"update_time"`
}

type CommissionRates struct {
	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
}

type CommissionNew struct {
	CommissionRates CommissionRates `json:"commission_rates"`
	UpdateTime      time.Time       `json:"update_time"`
}

type NodeInfoVo struct {
	ProtocolVersion struct {
		P2P   string `json:"p2p"`
		Block string `json:"block"`
		App   string `json:"app"`
	} `json:"protocol_version"`
	ID         string `json:"id"`
	ListenAddr string `json:"listen_addr"`
	Network    string `json:"network"`
	Version    string `json:"version"`
	Channels   string `json:"channels"`
	Moniker    string `json:"moniker"`
	Other      struct {
		TxIndex    string `json:"tx_index"`
		RPCAddress string `json:"rpc_address"`
	} `json:"other"`
}

type NodeInfoVoReponse struct {
	NodeInfoVo NodeInfoVo `json:"node_info"`
}

type GenesisVo struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		Genesis struct {
			GenesisTime     time.Time `json:"genesis_time"`
			ChainID         string    `json:"chain_id"`
			ConsensusParams struct {
				BlockSize struct {
					MaxBytes string `json:"max_bytes"`
					MaxGas   string `json:"max_gas"`
				} `json:"block_size"`
				Evidence struct {
					MaxAge string `json:"max_age"`
				} `json:"evidence"`
				Validator struct {
					PubKeyTypes []string `json:"pub_key_types"`
				} `json:"validator"`
			} `json:"consensus_params"`
			AppHash  string `json:"app_hash"`
			AppState struct {
				Accounts []struct {
					Address        string   `json:"address"`
					Coins          []string `json:"coins"`
					SequenceNumber string   `json:"sequence_number"`
					AccountNumber  string   `json:"account_number"`
				} `json:"accounts"`
				Auth struct {
					CollectedFee interface{} `json:"collected_fee"`
					Data         struct {
						NativeFeeDenom string `json:"native_fee_denom"`
					} `json:"data"`
					Params struct {
						GasPriceThreshold string `json:"gas_price_threshold"`
						TxSize            string `json:"tx_size"`
					} `json:"params"`
				} `json:"auth"`
				Stake struct {
					Pool struct {
						BondedTokens string `json:"bonded_tokens"`
					} `json:"pool"`
					Params struct {
						UnbondingTime string `json:"unbonding_time"`
						MaxValidators int    `json:"max_validators"`
					} `json:"params"`
					LastTotalPower       string      `json:"last_total_power"`
					LastValidatorPowers  interface{} `json:"last_validator_powers"`
					Validators           interface{} `json:"validators"`
					Bonds                interface{} `json:"bonds"`
					UnbondingDelegations interface{} `json:"unbonding_delegations"`
					Redelegations        interface{} `json:"redelegations"`
					Exported             bool        `json:"exported"`
				} `json:"stake"`
				Mint struct {
					Minter struct {
						LastUpdate        time.Time `json:"last_update"`
						MintDenom         string    `json:"mint_denom"`
						InflationBasement string    `json:"inflation_basement"`
					} `json:"minter"`
					Params struct {
						Inflation string `json:"inflation"`
					} `json:"params"`
				} `json:"mint"`
				Distr struct {
					Params struct {
						CommunityTax        string `json:"community_tax"`
						BaseProposerReward  string `json:"base_proposer_reward"`
						BonusProposerReward string `json:"bonus_proposer_reward"`
					} `json:"params"`
					FeePool struct {
						ValAccum struct {
							UpdateHeight string `json:"update_height"`
							Accum        string `json:"accum"`
						} `json:"val_accum"`
						ValPool       interface{} `json:"val_pool"`
						CommunityPool interface{} `json:"community_pool"`
					} `json:"fee_pool"`
					ValidatorDistInfos     interface{} `json:"validator_dist_infos"`
					DelegatorDistInfos     interface{} `json:"delegator_dist_infos"`
					DelegatorWithdrawInfos interface{} `json:"delegator_withdraw_infos"`
					PreviousProposer       string      `json:"previous_proposer"`
				} `json:"distr"`
				Gov struct {
					Params struct {
						CriticalDepositPeriod string `json:"critical_deposit_period"`
						CriticalMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"critical_min_deposit"`
						CriticalVotingPeriod   string `json:"critical_voting_period"`
						CriticalMaxNum         string `json:"critical_max_num"`
						CriticalThreshold      string `json:"critical_threshold"`
						CriticalVeto           string `json:"critical_veto"`
						CriticalParticipation  string `json:"critical_participation"`
						CriticalPenalty        string `json:"critical_penalty"`
						ImportantDepositPeriod string `json:"important_deposit_period"`
						ImportantMinDeposit    []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"important_min_deposit"`
						ImportantVotingPeriod  string `json:"important_voting_period"`
						ImportantMaxNum        string `json:"important_max_num"`
						ImportantThreshold     string `json:"important_threshold"`
						ImportantVeto          string `json:"important_veto"`
						ImportantParticipation string `json:"important_participation"`
						ImportantPenalty       string `json:"important_penalty"`
						NormalDepositPeriod    string `json:"normal_deposit_period"`
						NormalMinDeposit       []struct {
							Denom  string `json:"denom"`
							Amount string `json:"amount"`
						} `json:"normal_min_deposit"`
						NormalVotingPeriod  string `json:"normal_voting_period"`
						NormalMaxNum        string `json:"normal_max_num"`
						NormalThreshold     string `json:"normal_threshold"`
						NormalVeto          string `json:"normal_veto"`
						NormalParticipation string `json:"normal_participation"`
						NormalPenalty       string `json:"normal_penalty"`
						SystemHaltPeriod    string `json:"system_halt_period"`
					} `json:"params"`
				} `json:"gov"`
				Upgrade struct {
					GenesisVersion struct {
						UpgradeInfo struct {
							ProposalID string `json:"ProposalID"`
							Protocol   struct {
								Version   string `json:"version"`
								Software  string `json:"software"`
								Height    string `json:"height"`
								Threshold string `json:"threshold"`
							} `json:"Protocol"`
						} `json:"UpgradeInfo"`
						Success bool `json:"Success"`
					} `json:"GenesisVersion"`
				} `json:"upgrade"`
				Slashing struct {
					Params struct {
						MaxEvidenceAge          string `json:"max_evidence_age"`
						SignedBlocksWindow      string `json:"signed_blocks_window"`
						MinSignedPerWindow      string `json:"min_signed_per_window"`
						DoubleSignJailDuration  string `json:"double_sign_jail_duration"`
						DowntimeJailDuration    string `json:"downtime_jail_duration"`
						CensorshipJailDuration  string `json:"censorship_jail_duration"`
						SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
						SlashFractionDowntime   string `json:"slash_fraction_downtime"`
						SlashFractionCensorship string `json:"slash_fraction_censorship"`
					} `json:"params"`
					SigningInfos struct {
					} `json:"signing_infos"`
					MissedBlocks struct {
					} `json:"missed_blocks"`
					SlashingPeriods interface{} `json:"slashing_periods"`
				} `json:"slashing"`
				Service struct {
					Params struct {
						MaxRequestTimeout    string `json:"max_request_timeout"`
						MinDepositMultiple   string `json:"min_deposit_multiple"`
						ServiceFeeTax        string `json:"service_fee_tax"`
						SlashFraction        string `json:"slash_fraction"`
						ComplaintRetrospect  string `json:"complaint_retrospect"`
						ArbitrationTimeLimit string `json:"arbitration_time_limit"`
						TxSizeLimit          string `json:"tx_size_limit"`
					} `json:"params"`
				} `json:"service"`
				Guardian struct {
					Profilers []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"profilers"`
					Trustees []struct {
						Description string `json:"description"`
						Type        string `json:"type"`
						Address     string `json:"address"`
						AddedBy     string `json:"added_by"`
					} `json:"trustees"`
				} `json:"guardian"`
				Gentxs []struct {
					Type  string `json:"type"`
					Value struct {
						Msg []struct {
							Type  string `json:"type"`
							Value struct {
								Description struct {
									Moniker  string `json:"moniker"`
									Identity string `json:"identity"`
									Website  string `json:"website"`
									Details  string `json:"details"`
								} `json:"Description"`
								Commission struct {
									Rate          string `json:"rate"`
									MaxRate       string `json:"max_rate"`
									MaxChangeRate string `json:"max_change_rate"`
								} `json:"Commission"`
								DelegatorAddress string `json:"delegator_address"`
								ValidatorAddress string `json:"validator_address"`
								Pubkey           struct {
									Type  string `json:"type"`
									Value string `json:"value"`
								} `json:"pubkey"`
								Delegation struct {
									Denom  string `json:"denom"`
									Amount string `json:"amount"`
								} `json:"delegation"`
							} `json:"value"`
						} `json:"msg"`
						Fee struct {
							Amount interface{} `json:"amount"`
							Gas    string      `json:"gas"`
						} `json:"fee"`
						Signatures []struct {
							PubKey struct {
								Type  string `json:"type"`
								Value string `json:"value"`
							} `json:"pub_key"`
							Signature     string `json:"signature"`
							AccountNumber string `json:"account_number"`
							Sequence      string `json:"sequence"`
						} `json:"signatures"`
						Memo string `json:"memo"`
					} `json:"value"`
				} `json:"gentxs"`
			} `json:"app_state"`
		} `json:"genesis"`
	} `json:"result"`
}

type BlockVo struct {
	BlockMeta struct {
		BlockID struct {
			Hash  string `json:"hash"`
			Parts struct {
				Total string `json:"total"`
				Hash  string `json:"hash"`
			} `json:"parts"`
		} `json:"block_id"`
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID     string    `json:"chain_id"`
			Height      string    `json:"height"`
			Time        time.Time `json:"time"`
			NumTxs      string    `json:"num_txs"`
			TotalTxs    string    `json:"total_txs"`
			LastBlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"last_block_id"`
			LastCommitHash     string `json:"last_commit_hash"`
			DataHash           string `json:"data_hash"`
			ValidatorsHash     string `json:"validators_hash"`
			NextValidatorsHash string `json:"next_validators_hash"`
			ConsensusHash      string `json:"consensus_hash"`
			AppHash            string `json:"app_hash"`
			LastResultsHash    string `json:"last_results_hash"`
			EvidenceHash       string `json:"evidence_hash"`
			ProposerAddress    string `json:"proposer_address"`
		} `json:"header"`
	} `json:"block_meta"`
	Block struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID     string    `json:"chain_id"`
			Height      string    `json:"height"`
			Time        time.Time `json:"time"`
			NumTxs      string    `json:"num_txs"`
			TotalTxs    string    `json:"total_txs"`
			LastBlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"last_block_id"`
			LastCommitHash     string `json:"last_commit_hash"`
			DataHash           string `json:"data_hash"`
			ValidatorsHash     string `json:"validators_hash"`
			NextValidatorsHash string `json:"next_validators_hash"`
			ConsensusHash      string `json:"consensus_hash"`
			AppHash            string `json:"app_hash"`
			LastResultsHash    string `json:"last_results_hash"`
			EvidenceHash       string `json:"evidence_hash"`
			ProposerAddress    string `json:"proposer_address"`
		} `json:"header"`
		Data struct {
			Txs interface{} `json:"txs"`
		} `json:"data"`
		Evidence struct {
			Evidence interface{} `json:"evidence"`
		} `json:"evidence"`
		LastCommit struct {
			BlockID struct {
				Hash  string `json:"hash"`
				Parts struct {
					Total string `json:"total"`
					Hash  string `json:"hash"`
				} `json:"parts"`
			} `json:"block_id"`
			Precommits []struct {
				Type    int    `json:"type"`
				Height  string `json:"height"`
				Round   string `json:"round"`
				BlockID struct {
					Hash  string `json:"hash"`
					Parts struct {
						Total string `json:"total"`
						Hash  string `json:"hash"`
					} `json:"parts"`
				} `json:"block_id"`
				Timestamp        time.Time `json:"timestamp"`
				ValidatorAddress string    `json:"validator_address"`
				ValidatorIndex   string    `json:"validator_index"`
				Signature        string    `json:"signature"`
			} `json:"precommits"`
		} `json:"last_commit"`
	} `json:"block"`
}

type ValidatorSetVoResponse struct {
	Height string         `json:"height"`
	Result ValidatorSetVo `json:"result"`
}

type ValidatorSetVo struct {
	BlockHeight string `json:"block_height"`
	Validators  []struct {
		Address          string `json:"address"`
		PubKey           string `json:"pub_key"`
		ProposerPriority string `json:"proposer_priority"`
		VotingPower      string `json:"voting_power"`
	} `json:"validators"`
}

type StakePoolVoResponse struct {
	Height string      `json:"height"`
	Result StakePoolVo `json:"result"`
}

type StakePoolVo struct {
	NotBondedTokens string `json:"not_bonded_tokens"`
	BondedTokens    string `json:"bonded_tokens"`
	TotalSupply     string `json:"total_supply"`
	BondedRatio     string `json:"bonded_ratio"`
}

type DelegationVoResponse struct {
	Height string         `json:"height"`
	Result []DelegationVo `json:"result"`
}

type DelegationVo struct {
	DelegatorAddr string `json:"delegator_address"`
	ValidatorAddr string `json:"validator_address"`
	Shares        string `json:"shares"`
	Balance       string `json:"balance"`
}

type ValidatorDelegations []DelegationVo

func (sort ValidatorDelegations) Len() int {
	return len(sort)
}
func (sort ValidatorDelegations) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}
func (sort ValidatorDelegations) Less(i, j int) bool {
	return sort[i].Balance > sort[j].Balance
}

type DistributionRewardsResponse struct {
	Height string              `json:"height"`
	Result DistributionRewards `json:"result"`
}

type DistributionRewards struct {
	Total   utils.CoinsAsStr         `json:"total"`
	Rewards []RewardsFromDelegations `json:"rewards"`
}

type RewardsFromDelegations struct {
	ValidatorAddress string           `json:"validator_address"`
	Reward           utils.CoinsAsStr `json:"reward"`
}

type ValidatorDistributionInfoResponse struct {
	Height string                    `json:"height"`
	Result ValidatorDistributionInfo `json:"result"`
}

type ValidatorDistributionInfo struct {
	Validator       string           `json:"validator"`
	SelfBondRewards utils.CoinsAsStr `json:"self_bond_rewards"`
	ValCommission   utils.CoinsAsStr `json:"val_commission"`
}

type ValidatorSigningInfo struct {
	StartHeight       string `json:"start_height"`
	IndexOffset       string `json:"index_offset"`
	JailedUntil       string `json:"jailed_until"`
	MissedBlocksCount string `json:"missed_blocks_counter"`
}

type ReDelegationsResponse struct {
	Height string          `json:"height"`
	Result []ReDelegations `json:"result"`
}

type ReDelegations struct {
	DelegatorAddr    string `json:"delegator_addr"`
	ValidatorSrcAddr string `json:"validator_src_addr"`
	ValidatorDstAddr string `json:"validator_dst_addr"`
	CreationHeight   string `json:"creation_height"`
	MinTime          int64  `json:"min_time"`
	InitialBalance   string `json:"initial_balance"`
	Balance          string `json:"balance"`
	SharesSrc        string `json:"shares_src"`
	SharesDst        string `json:"shares_dst"`
}

type UnbondingDelegationsResponse struct {
	Height string                     `json:"height"`
	Result UnbondingDelegationNewList `json:"result"`
}

type UnbondingDelegation struct {
	DelegatorAddr  string `json:"delegator_addr"`
	ValidatorAddr  string `json:"validator_addr"`
	InitialBalance string `json:"initial_balance"`
	Balance        string `json:"balance"`
	CreationHeight int64  `json:"creation_height,string"`
	MinTime        string `json:"min_time"`
}

type UnbondingDelegationNewList []UnbondingDelegationNew

func (list UnbondingDelegationNewList) toUnbondingDelegationList() (result []UnbondingDelegation) {
	for _, ubd := range list {
		for _, entry := range ubd.Entries {
			height, _ := strconv.Atoi(entry.CreationHeight)
			result = append(result, UnbondingDelegation{
				DelegatorAddr:  ubd.DelegatorAddr,
				ValidatorAddr:  ubd.ValidatorAddr,
				InitialBalance: entry.InitialBalance,
				Balance:        entry.Balance,
				CreationHeight: int64(height),
				MinTime:        entry.CompletionTime,
			})
		}
	}
	return
}

type UnbondingDelegationNew struct {
	DelegatorAddr string    `json:"delegator_address"`
	ValidatorAddr string    `json:"validator_address"`
	Entries       []Entries `json:"entries"`
}

type Entries struct {
	InitialBalance string `json:"initial_balance"`
	Balance        string `json:"balance"`
	CreationHeight string `json:"creation_height"`
	CompletionTime string `json:"completion_time"`
}

func (un UnbondingDelegation) String() string {
	return fmt.Sprintf(`
		DelegatorAddr  :%v
		ValidatorAddr  :%v
		InitialBalance :%v
		Balance        :%v
		CreationHeight :%v
		MinTime        :%v
		`, un.DelegatorAddr, un.ValidatorAddr, un.InitialBalance, un.Balance, un.CreationHeight, un.MinTime)

}

func (d DelegationVo) String() string {
	return fmt.Sprintf(`
		DelegatorAddr :%v
		ValidatorAddr :%v
		Shares        :%v
		Balance       :%v
		`, d.DelegatorAddr, d.ValidatorAddr, d.Shares, d.Balance)
}

type SignInfoVoResponse struct {
	Height string     `json:"height"`
	Result SignInfoVo `json:"result"`
}

type SignInfoVo struct {
	StartHeight         string    `json:"start_height"`
	IndexOffset         string    `json:"index_offset"`
	JailedUntil         time.Time `json:"jailed_until"`
	MissedBlocksCounter string    `json:"missed_blocks_counter"`
}

type BlockResultVo struct {
	Height  string `json:"height"`
	Results struct {
		DeliverTx []struct {
			Code      int         `json:"code"`
			Data      interface{} `json:"data"`
			Log       string      `json:"log"`
			Info      string      `json:"info"`
			GasWanted string      `json:"gas_wanted"`
			GasUsed   string      `json:"gas_used"`
			Events    []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"deliver_tx"`
		EndBlock struct {
			ValidatorUpdates []struct {
				PubKey struct {
					Type string `json:"type"`
					Data string `json:"data"`
				} `json:"pub_key"`
				Power string `json:"power"`
			} `json:"validator_updates"`
			ConsensusParamUpdates interface{} `json:"consensus_param_updates"`
			Events                []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"end_block"`
		BeginBlock struct {
			Events []struct {
				Type       string `json:"type"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"events"`
		} `json:"begin_block"`
	} `json:"results"`
}

type AssetTokens struct {
	BaseToken BaseToken `json:"base_token"`
}

type BaseToken struct {
	Id              string `json:"id"`
	Family          string `json:"family"`
	Source          string `json:"source"`
	Gateway         string `json:"gateway"`
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	Decimal         int    `json:"decimal"`
	CanonicalSymbol string `json:"canonical_symbol"`
	MinUnitAlias    string `json:"min_unit_alias"`
	InitialSupply   string `json:"initial_supply"`
	MaxSupply       string `json:"max_supply"`
	Mintable        bool   `json:"mintable"`
	Owner           string `json:"owner"`
}

type IssuedTokenListResponse struct {
	Height string        `json:"height"`
	Result []IssuedToken `json:"result"`
}

type IssuedTokenResponse struct {
	Height string      `json:"height"`
	Result IssuedToken `json:"result"`
}

type IssuedToken struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Decimals    int    `json:"decimals"`
	TotalSupply string `json:"total_supply"`
	Mintable    bool   `json:"mintable"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

type AssetGateways struct {
	Owner    string `json:"owner"`
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Details  string `json:"details"`
	Website  string `json:"website"`
}
