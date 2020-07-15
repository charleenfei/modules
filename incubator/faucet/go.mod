module github.com/okwme/modules/incubator/faucet

go 1.13

require (
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/gorilla/mux v1.7.4
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/tendermint v0.33.3
	github.com/tmdvs/Go-Emoji-Utils v1.1.0
)

replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.5-0.20200715162801-4fd244eef297

// replace github.com/cosmos/cosmos-sdk v0.38.4 => /Users/billy/GitHub.com/okwme/cosmos-sdk
