package app

import (
	"encoding/json"
	"io/ioutil"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/bandprotocol/d3n/chain/x/zoracle"
)

// GenesisState defines a type alias for the Gaia genesis application state.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		genaccounts.ModuleName: genaccounts.ModuleCdc.MustMarshalJSON(
			genaccounts.GenesisState{},
		),
		auth.ModuleName: auth.ModuleCdc.MustMarshalJSON(
			auth.GenesisState{
				Params: auth.Params{
					MaxMemoCharacters:      512,
					TxSigLimit:             7,
					TxSizeCostPerByte:      10,
					SigVerifyCostED25519:   590,
					SigVerifyCostSecp256k1: 1000,
				},
			},
		),
		bank.ModuleName: bank.ModuleCdc.MustMarshalJSON(bank.GenesisState{
			SendEnabled: true,
		}),
		staking.ModuleName: staking.ModuleCdc.MustMarshalJSON(staking.GenesisState{
			Params: staking.Params{
				UnbondingTime: time.Hour * 24 * 7 * 3, // 3 weeks
				BondDenom:     "uband",
				MaxEntries:    7,
				MaxValidators: 100,
			},
		}),
		mint.ModuleName: mint.ModuleCdc.MustMarshalJSON(mint.GenesisState{
			Minter: mint.Minter{
				AnnualProvisions: sdk.NewDecWithPrec(0, 0),
				Inflation:        sdk.NewDecWithPrec(135, 3), // 13.5%
			},
			Params: mint.Params{
				BlocksPerYear:       10519200,                  // 3 second  block times
				GoalBonded:          sdk.NewDecWithPrec(67, 2), // 67%
				InflationMax:        sdk.NewDecWithPrec(20, 2), // 20%
				InflationMin:        sdk.NewDecWithPrec(7, 2),  // 7%
				InflationRateChange: sdk.NewDecWithPrec(13, 2), // 13%
				MintDenom:           "uband",
			},
		}),
		distr.ModuleName: distr.ModuleCdc.MustMarshalJSON(distr.GenesisState{
			FeePool:                         distr.InitialFeePool(),
			CommunityTax:                    sdk.NewDecWithPrec(2, 2), // 2%
			BaseProposerReward:              sdk.NewDecWithPrec(1, 2), // 1%
			BonusProposerReward:             sdk.NewDecWithPrec(4, 2), // 4%
			WithdrawAddrEnabled:             true,
			DelegatorWithdrawInfos:          []distr.DelegatorWithdrawInfo{},
			PreviousProposer:                nil,
			OutstandingRewards:              []distr.ValidatorOutstandingRewardsRecord{},
			ValidatorAccumulatedCommissions: []distr.ValidatorAccumulatedCommissionRecord{},
			ValidatorHistoricalRewards:      []distr.ValidatorHistoricalRewardsRecord{},
			ValidatorCurrentRewards:         []distr.ValidatorCurrentRewardsRecord{},
			DelegatorStartingInfos:          []distr.DelegatorStartingInfoRecord{},
			ValidatorSlashEvents:            []distr.ValidatorSlashEventRecord{},
		}),
		gov.ModuleName: gov.ModuleCdc.MustMarshalJSON(gov.GenesisState{
			StartingProposalID: 1,
			DepositParams: gov.DepositParams{
				MinDeposit:       sdk.NewCoins(sdk.NewCoin("uband", sdk.TokensFromConsensusPower(1000))),
				MaxDepositPeriod: 86400 * 2 * time.Second, // 2 days
			},
			VotingParams: gov.VotingParams{
				VotingPeriod: 86400 * 2 * time.Second, // 2 days
			},
			TallyParams: gov.TallyParams{
				Quorum:    sdk.NewDecWithPrec(4, 1),   //  40%
				Threshold: sdk.NewDecWithPrec(5, 1),   // 50%
				Veto:      sdk.NewDecWithPrec(334, 3), // 33.4%
			},
		}),
		crisis.ModuleName: crisis.ModuleCdc.MustMarshalJSON(crisis.GenesisState{
			ConstantFee: sdk.NewCoin("uband", sdk.TokensFromConsensusPower(10000)),
		}),
		slashing.ModuleName: slashing.ModuleCdc.MustMarshalJSON(slashing.GenesisState{
			Params: slashing.Params{
				MaxEvidenceAge:          60 * 30240 * time.Second, // 3 weeks
				SignedBlocksWindow:      int64(30000),
				MinSignedPerWindow:      sdk.NewDecWithPrec(5, 2), // 5%
				DowntimeJailDuration:    60 * 10 * time.Second,    // 10 minutes
				SlashFractionDoubleSign: sdk.NewDecWithPrec(5, 2), // 5%
				SlashFractionDowntime:   sdk.NewDecWithPrec(1, 4), // 0.01%
			},
			SigningInfos: make(map[string]slashing.ValidatorSigningInfo),
			MissedBlocks: make(map[string][]slashing.MissedBlock),
		}),
		supply.ModuleName: supply.ModuleCdc.MustMarshalJSON(supply.GenesisState{
			Supply: sdk.NewCoins(),
		}),
		zoracle.ModuleName: zoracle.ModuleCdc.MustMarshalJSON(zoracle.DefaultGenesisState()),
		genutil.ModuleName: genutil.ModuleCdc.MustMarshalJSON(genutil.GenesisState{
			GenTxs: []json.RawMessage{},
		}),
	}
}

func GetDefaultDataSourcesAndOracleScripts(owner sdk.AccAddress) json.RawMessage {
	state := zoracle.DefaultGenesisState()
	dataSources := []struct {
		name string
		path string
	}{
		{"Coingecko script", "./datasources/coingecko_price.sh"},
		{"Crypto compare script", "./datasources/crypto_compare_price.sh"},
		{"Binance price", "./datasources/binance_price.sh"},
	}

	// TODO: Find a better way to specific path to data sources
	state.DataSources = make([]zoracle.DataSource, len(dataSources))
	for i, dataSource := range dataSources {
		script, err := ioutil.ReadFile(dataSource.path)
		if err != nil {
			panic(err)
		}
		state.DataSources[i] = zoracle.NewDataSource(
			owner,
			dataSource.name,
			sdk.Coins{},
			script,
		)
	}

	// TODO: Find a better way to specific path to oracle scripts
	oracleScripts := []struct {
		name string
		path string
	}{
		{"Crypto price script", "./owasm/res/crypto_price.wasm"},
	}
	state.OracleScripts = make([]zoracle.OracleScript, len(oracleScripts))
	for i, oracleScript := range oracleScripts {
		code, err := ioutil.ReadFile(oracleScript.path)
		if err != nil {
			panic(err)
		}
		state.OracleScripts[i] = zoracle.NewOracleScript(
			owner,
			oracleScript.name,
			code,
		)
	}
	return zoracle.ModuleCdc.MustMarshalJSON(state)
}
