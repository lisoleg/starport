// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteSentPost } from "./types/blog/tx";
import { MsgUpdatePost } from "./types/blog/tx";
import { MsgCreateTimedoutPost } from "./types/blog/tx";
import { MsgDeletePost } from "./types/blog/tx";
import { MsgSendIbcPost } from "./types/blog/tx";
import { MsgDeleteTimedoutPost } from "./types/blog/tx";
import { MsgUpdateTimedoutPost } from "./types/blog/tx";
import { MsgCreateSentPost } from "./types/blog/tx";
import { MsgUpdateSentPost } from "./types/blog/tx";
import { MsgCreatePost } from "./types/blog/tx";


const types = [
  ["/cosmonaut.planet.blog.MsgDeleteSentPost", MsgDeleteSentPost],
  ["/cosmonaut.planet.blog.MsgUpdatePost", MsgUpdatePost],
  ["/cosmonaut.planet.blog.MsgCreateTimedoutPost", MsgCreateTimedoutPost],
  ["/cosmonaut.planet.blog.MsgDeletePost", MsgDeletePost],
  ["/cosmonaut.planet.blog.MsgSendIbcPost", MsgSendIbcPost],
  ["/cosmonaut.planet.blog.MsgDeleteTimedoutPost", MsgDeleteTimedoutPost],
  ["/cosmonaut.planet.blog.MsgUpdateTimedoutPost", MsgUpdateTimedoutPost],
  ["/cosmonaut.planet.blog.MsgCreateSentPost", MsgCreateSentPost],
  ["/cosmonaut.planet.blog.MsgUpdateSentPost", MsgUpdateSentPost],
  ["/cosmonaut.planet.blog.MsgCreatePost", MsgCreatePost],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgDeleteSentPost: (data: MsgDeleteSentPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgDeleteSentPost", value: data }),
    msgUpdatePost: (data: MsgUpdatePost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgUpdatePost", value: data }),
    msgCreateTimedoutPost: (data: MsgCreateTimedoutPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgCreateTimedoutPost", value: data }),
    msgDeletePost: (data: MsgDeletePost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgDeletePost", value: data }),
    msgSendIbcPost: (data: MsgSendIbcPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgSendIbcPost", value: data }),
    msgDeleteTimedoutPost: (data: MsgDeleteTimedoutPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgDeleteTimedoutPost", value: data }),
    msgUpdateTimedoutPost: (data: MsgUpdateTimedoutPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgUpdateTimedoutPost", value: data }),
    msgCreateSentPost: (data: MsgCreateSentPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgCreateSentPost", value: data }),
    msgUpdateSentPost: (data: MsgUpdateSentPost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgUpdateSentPost", value: data }),
    msgCreatePost: (data: MsgCreatePost): EncodeObject => ({ typeUrl: "/cosmonaut.planet.blog.MsgCreatePost", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
