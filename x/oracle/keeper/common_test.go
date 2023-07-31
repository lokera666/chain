package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/oracle/types"
)

var (
	BasicName          = "BASIC_NAME"
	BasicDesc          = "BASIC_DESCRIPTION"
	BasicSchema        = "BASIC_SCHEMA"
	BasicSourceCodeURL = "BASIC_SOURCE_CODE_URL"
	BasicFilename      = "BASIC_FILENAME"
	BasicCalldata      = []byte("BASIC_CALLDATA")
	BasicClientID      = "BASIC_CLIENT_ID"
	BasicReport        = []byte("BASIC_REPORT")
	BasicResult        = []byte("BASIC_RESULT")
	BasicSigningResult = types.SigningResult{
		SigningID:      1,
		ErrorCodespace: "",
		ErrorCode:      0,
	}
	CoinsZero         = sdk.NewCoins()
	Coins10uband      = sdk.NewCoins(sdk.NewInt64Coin("uband", 10))
	Coins20uband      = sdk.NewCoins(sdk.NewInt64Coin("uband", 20))
	Coins1000000uband = sdk.NewCoins(sdk.NewInt64Coin("uband", 1000000))
)
