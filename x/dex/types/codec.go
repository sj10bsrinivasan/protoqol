package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePool{}, "dex/CreatePool", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "dex/UpdatePool", nil)
	cdc.RegisterConcrete(&MsgDeletePool{}, "dex/DeletePool", nil)
	cdc.RegisterConcrete(&MsgAddLiquidity{}, "dex/AddLiquidity", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePool{},
		&MsgUpdatePool{},
		&MsgDeletePool{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddLiquidity{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
