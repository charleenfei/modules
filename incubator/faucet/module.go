package faucet

import (
	"context"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/charleenfei/modules/incubator/faucet/client/cli"
	"github.com/charleenfei/modules/incubator/faucet/keeper"
	"github.com/charleenfei/modules/incubator/faucet/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	abci "github.com/tendermint/tendermint/abci/types"
)

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec doesn't support amino
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// DefaultGenesis returns default genesis state as raw bytes for the faucet
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}

// ValidateGenesis performs genesis state validation for the faucet module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, _ client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}

// RegisterRESTRoutes doesn't support legacy REST routes.
// We don't want to support the legacy rest server here
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the distribution module.
// also implements app modeul basic
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// RegisterInterfaces implements app module basic
func (b AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

type AppModule struct {
	AppModuleBasic
	msgServer   keeper.MsgServer
	queryServer keeper.QueryServer
}

// NewAppModule creates a new AppModule Object
func NewAppModule(m keeper.MsgServer, q keeper.QueryServer) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		msgServer:      m,
		queryServer:    q,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.msgServer))
}

func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns a nil Querier.
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.msgServer)
	types.RegisterQueryServer(cfg.QueryServer(), am.queryServer)
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return nil
}
