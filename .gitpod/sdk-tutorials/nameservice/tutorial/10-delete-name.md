---
order: 11
---

# Delete Name

## MsgDeleteName


Now it is time to update the `Msg` for deleting names. Let's rename our `MsgDeleteWhois.go` to `MsgDeleteName.go`

```
mv x/nameservice/types/MsgDeleteWhois.go x/nameservice/types/MsgDeleteName.go
```

add it to the `./x/nameservice/types/MsgDeleteName.go` file. 

Replace `MsgDeleteWhois` by `MsgDeleteName`:
```
mv x/nameservice/types/MsgDeleteWhois.go x/nameservice/types/MsgDeleteName.go
```

<<< @/nameservice/nameservice/x/nameservice/types/MsgDeleteName.go

Replace `handlerMsgDeleteWhois` by `handlerMsgDeleteName`:
```
mv x/nameservice/handlerMsgDeleteWhois.go x/nameservice/handlerMsgDeleteName.go
```

Finally, define the `DeleteName` `handler` function which performs the state transitions triggered by the message. Keep in mind that at this point the message has had its `ValidateBasic` function run so there has been some input verification. However, `ValidateBasic` cannot query application state. Validation logic that is dependent on network state (e.g. account balances) should be performed in the `handler` function.

<<< @/nameservice/nameservice/x/nameservice/handlerMsgDeleteName.go

Afterwards, we'll follow the same steps as earlier and add the `MsgDeleteName` handler to the module router in `./x/nameservice/handler.go`:

```go
// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		case MsgDeleteName:
			return handleMsgDeleteName(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type()))
		}
	}
}
```

Let's rename our `handlerMsgDeleteWhois.go` to `handlerMsgDeleteName.go`

```
mv x/nameservice/handlerMsgDeleteWhois.go x/nameservice/handlerMsgDeleteName.go
```

<<< @/nameservice/nameservice/x/nameservice/handlerMsgDeleteName.go

### Now that you have your `Msgs` and `Handlers` defined it's time to learn about making the data from these transactions [available for querying.
