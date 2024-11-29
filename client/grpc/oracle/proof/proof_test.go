package proof

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto/tmhash"

	oracletypes "github.com/bandprotocol/chain/v3/x/oracle/types"
)

func hexToBytes(hexstr string) []byte {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return b
}

func leafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func innerHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func TestEncodeRelay(t *testing.T) {
	block := BlockRelayProof{
		MultiStoreProof: MultiStoreProof{
			OracleIAVLStateHash: hexToBytes(
				"5248463E932D16F7D092E268C0DED87B23D3B0E71856F1C6AE2AA91F6C713320",
			),
			MintStoreMerkleHash: hexToBytes(
				"7FD5F5C7C2920C187618542901CDC5717BE8204F24BE856E80902A1BB04737E4",
			),
			ParamsToRestakeStoresMerkleHash: hexToBytes(
				"F981716562A49DE06E3DCAFBFB6388C294BAA4FA9D45777E25740A92F81CF65E",
			),
			RollingseedToTransferStoresMerkleHash: hexToBytes(
				"E8E27CBB44BB654F64EEEF4667868AD48667CEB28E3DB5C4DF7A4B4B87F0C04B",
			),
			TssToUpgradeStoresMerkleHash: hexToBytes(
				"FC96CFFD30E5B8979EA66F9D0DA1CBAB16F69669E8B2A1FB2E1BEB457C9726E8",
			),
			AuthToIcahostStoresMerkleHash: hexToBytes(
				"C9C8849ED125CC7681329C4D27B83B1FC8ACF7A865C9D1D1DF575CCA56F48DBE",
			),
		},
		BlockHeaderMerkleParts: BlockHeaderMerkleParts{
			VersionAndChainIdHash: hexToBytes(
				"3F02642D9E70D5C1C493A4F732BFE9C9B95A4A42651703B816EDCFC8FADA5312",
			),
			Height:         25000,
			TimeSecond:     1629849931,
			TimeNanoSecond: 290650376,
			LastBlockIdAndOther: hexToBytes(
				"9B4825C99C3E739E1DC171FFB0E2BF34E99EEE41B34E407E40CF594834427B09",
			),
			NextValidatorHashAndConsensusHash: hexToBytes(
				"BF23413F237906B07202B3355E7311651ACE6BD2A34FD6FC3BD98EFE4FB78755",
			),
			LastResultsHash: hexToBytes(
				"9FB9C7533CAF1D218DA3AF6D277F6B101C42E3C3B75D784242DA663604DD53C2",
			),
			EvidenceAndProposerHash: hexToBytes(
				"7D11A74E40884411901BD7A70631734990B1FDBF5DE9E4C92C63B7650A6A6659",
			),
		},
		CommonEncodedVotePart: CommonEncodedVotePart{
			SignedDataPrefix: hexToBytes("080211A86100000000000022480A20"),
			SignedDataSuffix: hexToBytes(
				"1224080112206BF91EFBA26A4CD86EBBD0E54DCFC9BD2C790859CFA96215661A47E4921A6301",
			),
		},
		Signatures: []TMSignature{
			{
				R:                hexToBytes("84B8585B71240FEE0E674952B79ED25D793F1B31B42DD37B80F75B98510B5754"),
				S:                hexToBytes("1EC44DD7C5389474DF8E5C25CC6ED8B573CCA2E009AA824EE825BDC693935927"),
				V:                28,
				EncodedTimestamp: hexToBytes("08CD9296890610EAE9963D"),
			},
			{
				R:                hexToBytes("394365193F819CF539381366D31B6C5849AAA31AE8BA6F95C62C5C80656BFB5C"),
				S:                hexToBytes("6A07E4A3C0ABCEAE5F854D492DF699438FB84762F152F739DDEAC48DDCFCB5CC"),
				V:                28,
				EncodedTimestamp: hexToBytes("08CD9296890610EA928633"),
			},
			{
				R:                hexToBytes("5D7B4BE7B21B00D08AD7DBE48CF2761CECCB599E64AAB10B2901A0DD58F00325"),
				S:                hexToBytes("7160EF689A533C1E983707507FC8466DAEA1D0DC7A889E3A27D1BB1D09CEC030"),
				V:                28,
				EncodedTimestamp: hexToBytes("08CD929689061086FAB239"),
			},
			{
				R:                hexToBytes("5654A44FB89330C34CF2D862F940763194A145B72ED3BB0ADD5759E1E68FD145"),
				S:                hexToBytes("2AC795D02A9C574CF12343FDFC67FDCED8A24F88EC8138C7F8230F6EB442B726"),
				V:                28,
				EncodedTimestamp: hexToBytes("08CD9296890610F0E1F733"),
			},
		},
	}
	result, err := block.encodeToEthData()
	require.Nil(t, err)
	expect := hexToBytes(
		"5248463e932d16f7d092e268c0ded87b23d3b0e71856f1c6ae2aa91f6c7133207fd5f5c7c2920c187618542901cdc5717be8204f24be856e80902a1bb04737e4f981716562a49de06e3dcafbfb6388c294baa4fa9d45777e25740a92f81cf65ee8e27cbb44bb654f64eeef4667868ad48667ceb28e3db5c4df7a4b4b87f0c04bfc96cffd30e5b8979ea66f9d0da1cbab16f69669e8b2a1fb2e1beb457c9726e8c9c8849ed125cc7681329c4d27b83b1fc8acf7a865c9d1d1df575cca56f48dbe3f02642d9e70d5c1c493a4f732bfe9c9b95a4a42651703b816edcfc8fada531200000000000000000000000000000000000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000006125894b000000000000000000000000000000000000000000000000000000001152f9089b4825c99c3e739e1dc171ffb0e2bf34e99eee41b34e407e40cf594834427b09bf23413f237906b07202b3355e7311651ace6bd2a34fd6fc3bd98efe4fb787559fb9c7533caf1d218da3af6d277f6b101c42e3c3b75d784242da663604dd53c27d11a74e40884411901bd7a70631734990b1fdbf5de9e4c92c63b7650a6a6659000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000f080211a86100000000000022480a20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000261224080112206bf91efba26a4cd86ebbd0e54dcfc9bd2c790859cfa96215661a47e4921a63010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002c084b8585b71240fee0e674952b79ed25d793f1b31b42dd37b80f75b98510b57541ec44dd7c5389474df8e5c25cc6ed8b573cca2e009aa824ee825bdc693935927000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000b08cd9296890610eae9963d000000000000000000000000000000000000000000394365193f819cf539381366d31b6c5849aaa31ae8ba6f95c62c5c80656bfb5c6a07e4a3c0abceae5f854d492df699438fb84762f152f739ddeac48ddcfcb5cc000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000b08cd9296890610ea9286330000000000000000000000000000000000000000005d7b4be7b21b00d08ad7dbe48cf2761ceccb599e64aab10b2901a0dd58f003257160ef689a533c1e983707507fc8466daea1d0dc7a889e3a27d1bb1d09cec030000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000b08cd929689061086fab2390000000000000000000000000000000000000000005654a44fb89330c34cf2d862f940763194a145b72ed3bb0add5759e1e68fd1452ac795d02a9c574cf12343fdfc67fdced8a24f88ec8138c7f8230f6eb442b726000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000b08cd9296890610f0e1f733000000000000000000000000000000000000000000",
	)
	require.Equal(t, expect, result)
}

func TestEncodeVerify(t *testing.T) {
	data := OracleDataProof{
		Version: 217,
		Result: oracletypes.Result{
			ClientID:       "",
			OracleScriptID: 1,
			Calldata:       hexToBytes("000000010000000342544300000000000186a0"),
			AskCount:       1,
			MinCount:       1,
			RequestID:      1,
			AnsCount:       1,
			RequestTime:    1629803667,
			ResolveTime:    1629803671,
			ResolveStatus:  1,
			Result:         hexToBytes("000000010000000124ec078c"),
		},
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 217,
				SiblingHash:    hexToBytes("EB739BB22F48B7F3053A90BA2BA4FE07FAB262CADF8664489565C50FF505B8BD"),
			},
			{
				false,
				2,
				4,
				217,
				hexToBytes("1847107507D5E7B4CD9941EB6FFE1694264AF34C685C19DC478BEADDA265A578"),
			},
			{
				true,
				3,
				6,
				217,
				hexToBytes("E80AAE581EC004239854C4D90D8148E85F1F90D0704A73668FD2DA44DC0CEA53"),
			},
			{
				false,
				5,
				16,
				24998,
				hexToBytes("3174174561E430E12C1EAD3DC5B08BAE7D4315F417D1202E7277879C1433E658"),
			},
		},
	}

	result, err := data.encodeToEthData(25000)
	require.Nil(t, err)
	expect := hexToBytes(
		"00000000000000000000000000000000000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000d900000000000000000000000000000000000000000000000000000000000002800000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000006124d493000000000000000000000000000000000000000000000000000000006124d497000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013000000010000000342544300000000000186a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000010000000124ec078c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000d9eb739bb22f48b7f3053a90ba2ba4fe07fab262cadf8664489565c50ff505b8bd00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000d91847107507d5e7b4cd9941eb6ffe1694264af34c685c19dc478beadda265a57800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000d9e80aae581ec004239854c4d90d8148e85f1f90d0704a73668fd2da44dc0cea5300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000061a63174174561e430e12c1ead3dc5b08bae7d4315f417d1202e7277879c1433e658",
	)
	require.Equal(t, expect, result)
}

func TestEncodeVerifyCount(t *testing.T) {
	data := RequestsCountProof{
		Version: 215,
		Count:   1,
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("7DB3EEB5BDBEBCDE9EC49489CED0BD8A1ABAB7E653C720F18DF8149572699F1F"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  2,
				SubtreeSize:    4,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("24427D84B35482BBDD44DAF9C13A2C573C8B9DD8FBD4E1BC3B0F9201D707EC6B"),
			},
			{
				IsDataOnRight:  false,
				SubtreeHeight:  4,
				SubtreeSize:    10,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("E7A2CA3EC55627E7D6B06189FF7EF46C13ECC3B1900D127C7D437A7BA3EE3FFA"),
			},
			{
				IsDataOnRight:  false,
				SubtreeHeight:  5,
				SubtreeSize:    16,
				SubtreeVersion: 24998,
				SiblingHash:    hexToBytes("CE859DBB78E1B401CF0C91864E6A94E20A647B180804D36982FA2297796F0908"),
			},
		},
	}

	result, err := data.encodeToEthData(25000)
	require.Nil(t, err)
	expect := hexToBytes(
		"00000000000000000000000000000000000000000000000000000000000061a8000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000d70000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000061a67db3eeb5bdbebcde9ec49489ced0bd8a1abab7e653c720f18df8149572699f1f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000061a624427d84b35482bbdd44daf9c13a2c573c8b9dd8fbd4e1bc3b0f9201d707ec6b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000061a6e7a2ca3ec55627e7d6b06189ff7ef46c13ecc3b1900d127c7d437a7ba3ee3ffa00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000061a6ce859dbb78e1b401cf0c91864e6a94e20a647b180804d36982fa2297796f0908",
	)
	require.Equal(t, expect, result)
}
