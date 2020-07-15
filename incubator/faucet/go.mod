module github.com/okwme/modules/incubator/faucet

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/gorilla/mux v1.7.4
	github.com/kyokomi/emoji v2.2.4+incompatible
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/tendermint v0.33.3
)

replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.5-0.20200715105500-2f18421bd970
