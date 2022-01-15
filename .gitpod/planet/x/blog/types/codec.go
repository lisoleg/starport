package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePost{}, "blog/CreatePost", nil)
	cdc.RegisterConcrete(&MsgUpdatePost{}, "blog/UpdatePost", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "blog/DeletePost", nil)
	cdc.RegisterConcrete(&MsgCreateSentPost{}, "blog/CreateSentPost", nil)
	cdc.RegisterConcrete(&MsgUpdateSentPost{}, "blog/UpdateSentPost", nil)
	cdc.RegisterConcrete(&MsgDeleteSentPost{}, "blog/DeleteSentPost", nil)
	cdc.RegisterConcrete(&MsgCreateTimedoutPost{}, "blog/CreateTimedoutPost", nil)
	cdc.RegisterConcrete(&MsgUpdateTimedoutPost{}, "blog/UpdateTimedoutPost", nil)
	cdc.RegisterConcrete(&MsgDeleteTimedoutPost{}, "blog/DeleteTimedoutPost", nil)
	cdc.RegisterConcrete(&MsgSendIbcPost{}, "blog/SendIbcPost", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgUpdatePost{},
		&MsgDeletePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentPost{},
		&MsgUpdateSentPost{},
		&MsgDeleteSentPost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTimedoutPost{},
		&MsgUpdateTimedoutPost{},
		&MsgDeleteTimedoutPost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendIbcPost{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
