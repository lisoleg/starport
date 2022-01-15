// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendBuyOrder } from "./types/dex/tx";
import { MsgSendCreatePair } from "./types/dex/tx";
import { MsgCancelSellOrder } from "./types/dex/tx";
const types = [
    ["/cosmonaut.interchange.dex.MsgSendBuyOrder", MsgSendBuyOrder],
    ["/cosmonaut.interchange.dex.MsgSendCreatePair", MsgSendCreatePair],
    ["/cosmonaut.interchange.dex.MsgCancelSellOrder", MsgCancelSellOrder],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendBuyOrder: (data) => ({ typeUrl: "/cosmonaut.interchange.dex.MsgSendBuyOrder", value: data }),
        msgSendCreatePair: (data) => ({ typeUrl: "/cosmonaut.interchange.dex.MsgSendCreatePair", value: data }),
        msgCancelSellOrder: (data) => ({ typeUrl: "/cosmonaut.interchange.dex.MsgCancelSellOrder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
