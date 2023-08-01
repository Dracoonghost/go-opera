package launcher

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/ethereum/go-ethereum/params"

	"github.com/Fantom-foundation/go-opera/opera"
	"github.com/Fantom-foundation/go-opera/opera/genesis"
	"github.com/Fantom-foundation/go-opera/opera/genesisstore"
)

var (
	Bootnodes = map[string][]string{
		"main": {
			"enode://969654fc30f10e45903ee2d107cfc165feaff6f9acbd2a9472120a3581f7c717ded76113aa116e35b8389dfcc9129ecfc85a409f6c46fc864ce077cb85316ac8@3.252.226.57:5050",
			"enode://15dde1f31ff775da40d06a5529c15cce5bcdfec6055f8e8782b0df3f2ad3eab30b31fbb91037c6b3dacb8e612e177bdfaf4d82e9080feda31e54412efb56a34a@34.251.109.79:5050",
			"enode://41943a0269fc573355ff6a5e5980195cea68eb7bc703cc9937a3cd0debc9a795396aa9560bce65c8f329a9343f88ee4084daabcca76c5a6a8054aa163973a735@3.10.205.205:5050",
			"enode://b781fa90f63ef497c02ef6da8f32c2e14cc47921403d21083d5ebc69b5a85f77d6f124f084e3abbf3984e07aa5810b250658c04d6d3f3bc76e3735b2512847db@3.73.43.156:5050",
			"enode://fe97ec5b6c2073e4074067f3e513e66cf0a157c449bbb179dc8876803dc9cd3a18b03e735b87393c843cdc3be9c12989c2dd8ddaf88eda2f73a7183f2d65e95e@54.242.2.27:5050",
			"enode://2591357319aaa566897854f225d3a60194a8d193410de28e9ea74df2bb2615af63766351d8950c0582c970e3688ad85a8e454c5a601b9b091f094869fe603350@18.234.231.118:5050",
			"enode://7c6dbe419f79852536f3fc850912eaf657f50f378f6f6bc238bfe16bcf830c06ef3c0358774313eb94d959475a8bfa221bb53a4b0dcf3025ac071396199a3462@13.212.202.99:5050",
			"enode://e6c658978b95bcb29354f5a4a808b3d17d62201bcbd6ac36fa937e134c18b89ca23ee6f9f3893beabf07ec2f3e07b19d02d705ebc67ba671bff9830fa002c635@3.26.147.192:5050",
			"enode://9554ca82fb2cc3a2e3f30b847304e26e0ec6e73993c19610443a9fcd7c2a38796089a0ebe24c2c0bde6fb4155433c04104e049e9c6bd73a758c01e8cb93a7d9d@13.125.201.204:5050",
			"enode://b3699ce7ee4a8c6efb07707256a2a08872576f438664c6437cb263989116747d8d7332fb9cfc35c8f3ba04a024ecc6893cda1782db6fed221d9de2ec11746cfd@18.170.86.32:5050",
		},
		"test": {
			"enode://52c84c99a4cca9524dc626261e932aa8b1c88a103523132a1d6c30fd7d1f3eab0cb105403971baa7255f9eb5eadde9761db23bb277c4c7d6c3eefaf133dcb35f@170.64.156.90:7946",
			"enode://fffb2ee36f5571d3fa640afa38fcf3cd9389a51668f318c42b839cb31432155b34a12e94923a2230d4ff2d786b68ea30a1b9990ad9cefb663ae776279dbb63e5@157.245.202.7:7946",
			"enode://fecedd0b4dd953cb64547e00bf2c8d749911b8ba4870cce20dcbff171897ff2f89561428018892c8b63340a29711a3b271f26f2176fe97261d21dc103010c4af@165.232.96.92:7946",
			"enode://46a8746ac22f8b2b60aaed4820eae914193af873271897476f0b420d50951c479a8fb42591d42d17ec3233f52617646df45e9b6632fe7e418a390d99cc5d53ee@167.71.206.77:7946",
		},
	}

	mainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0x0edf665432f7d59e36783a63e8f9f77b3f0901c139bda1ad37bf2b847df076a8"),
		NetworkID:   opera.MainNetworkID,
		NetworkName: "main",
	}

	testnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xc4a5fc96e575a16a9a0c7349d44dc4d0f602a54e0a8543360c2fee4c3937b49e"),
		NetworkID:   opera.TestNetworkID,
		NetworkName: "test",
	}

	AllowedOperaGenesis = []GenesisTemplate{
		{
			Name:   "Mainnet-8888 with pruned MPT",
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0): hash.HexToHash("0x01c45fe3343344c852ae54e4492d9fe1ea6f912830469949459b310928a73be3"),
				genesisstore.BlocksSection(0): hash.HexToHash("0xea083affb4f2a0a9018f7f1832f16d6276c6e09aba482f6b3c061d7e9cc00d42"),
				genesisstore.EvmSection(0):    hash.HexToHash("0xa7c16548e78eddc20dfad091604aad04913726ae2f5bb14b916bf1a9ea1e9f25"),
			},
		},
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
