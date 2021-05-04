module github.com/MinterTeam/minter-hub-connector

go 1.13

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

// replace github.com/MinterTeam/mhub/chain => /Users/daniillashin/Desktop/mhub/chain

require (
	github.com/MinterTeam/mhub/chain v0.0.0-20210417174508-bac3972b7846
	github.com/MinterTeam/minter-go-sdk/v2 v2.2.0-alpha1.0.20210312102425-6b1675c84520
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/cosmos/go-bip39 v1.0.0
	github.com/ethereum/go-ethereum v1.9.22
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.9
	google.golang.org/grpc v1.35.0
)
