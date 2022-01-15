/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "lisoleg.identity.identity";

export interface MsgCreateIdentity {
  creator: string;
  id: string;
  pubKey: string;
  certificate: string;
  credentials: string;
  owner: string;
  data: string;
}

export interface MsgCreateIdentityResponse {}

export interface MsgUpdateIdentity {
  creator: string;
  id: string;
  pubKey: string;
  certificate: string;
  credentials: string;
  owner: string;
  data: string;
}

export interface MsgUpdateIdentityResponse {}

const baseMsgCreateIdentity: object = {
  creator: "",
  id: "",
  pubKey: "",
  certificate: "",
  credentials: "",
  owner: "",
  data: "",
};

export const MsgCreateIdentity = {
  encode(message: MsgCreateIdentity, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.pubKey !== "") {
      writer.uint32(26).string(message.pubKey);
    }
    if (message.certificate !== "") {
      writer.uint32(34).string(message.certificate);
    }
    if (message.credentials !== "") {
      writer.uint32(42).string(message.credentials);
    }
    if (message.owner !== "") {
      writer.uint32(50).string(message.owner);
    }
    if (message.data !== "") {
      writer.uint32(58).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateIdentity } as MsgCreateIdentity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.pubKey = reader.string();
          break;
        case 4:
          message.certificate = reader.string();
          break;
        case 5:
          message.credentials = reader.string();
          break;
        case 6:
          message.owner = reader.string();
          break;
        case 7:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateIdentity {
    const message = { ...baseMsgCreateIdentity } as MsgCreateIdentity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey);
    } else {
      message.pubKey = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = String(object.certificate);
    } else {
      message.certificate = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      message.credentials = String(object.credentials);
    } else {
      message.credentials = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgCreateIdentity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    message.certificate !== undefined &&
      (obj.certificate = message.certificate);
    message.credentials !== undefined &&
      (obj.credentials = message.credentials);
    message.owner !== undefined && (obj.owner = message.owner);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateIdentity>): MsgCreateIdentity {
    const message = { ...baseMsgCreateIdentity } as MsgCreateIdentity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = object.certificate;
    } else {
      message.certificate = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      message.credentials = object.credentials;
    } else {
      message.credentials = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgCreateIdentityResponse: object = {};

export const MsgCreateIdentityResponse = {
  encode(
    _: MsgCreateIdentityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateIdentityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateIdentityResponse,
    } as MsgCreateIdentityResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateIdentityResponse {
    const message = {
      ...baseMsgCreateIdentityResponse,
    } as MsgCreateIdentityResponse;
    return message;
  },

  toJSON(_: MsgCreateIdentityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateIdentityResponse>
  ): MsgCreateIdentityResponse {
    const message = {
      ...baseMsgCreateIdentityResponse,
    } as MsgCreateIdentityResponse;
    return message;
  },
};

const baseMsgUpdateIdentity: object = {
  creator: "",
  id: "",
  pubKey: "",
  certificate: "",
  credentials: "",
  owner: "",
  data: "",
};

export const MsgUpdateIdentity = {
  encode(message: MsgUpdateIdentity, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.pubKey !== "") {
      writer.uint32(26).string(message.pubKey);
    }
    if (message.certificate !== "") {
      writer.uint32(34).string(message.certificate);
    }
    if (message.credentials !== "") {
      writer.uint32(42).string(message.credentials);
    }
    if (message.owner !== "") {
      writer.uint32(50).string(message.owner);
    }
    if (message.data !== "") {
      writer.uint32(58).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateIdentity {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateIdentity } as MsgUpdateIdentity;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.pubKey = reader.string();
          break;
        case 4:
          message.certificate = reader.string();
          break;
        case 5:
          message.credentials = reader.string();
          break;
        case 6:
          message.owner = reader.string();
          break;
        case 7:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateIdentity {
    const message = { ...baseMsgUpdateIdentity } as MsgUpdateIdentity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey);
    } else {
      message.pubKey = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = String(object.certificate);
    } else {
      message.certificate = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      message.credentials = String(object.credentials);
    } else {
      message.credentials = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateIdentity): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    message.certificate !== undefined &&
      (obj.certificate = message.certificate);
    message.credentials !== undefined &&
      (obj.credentials = message.credentials);
    message.owner !== undefined && (obj.owner = message.owner);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateIdentity>): MsgUpdateIdentity {
    const message = { ...baseMsgUpdateIdentity } as MsgUpdateIdentity;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = object.certificate;
    } else {
      message.certificate = "";
    }
    if (object.credentials !== undefined && object.credentials !== null) {
      message.credentials = object.credentials;
    } else {
      message.credentials = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgUpdateIdentityResponse: object = {};

export const MsgUpdateIdentityResponse = {
  encode(
    _: MsgUpdateIdentityResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateIdentityResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateIdentityResponse,
    } as MsgUpdateIdentityResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateIdentityResponse {
    const message = {
      ...baseMsgUpdateIdentityResponse,
    } as MsgUpdateIdentityResponse;
    return message;
  },

  toJSON(_: MsgUpdateIdentityResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateIdentityResponse>
  ): MsgUpdateIdentityResponse {
    const message = {
      ...baseMsgUpdateIdentityResponse,
    } as MsgUpdateIdentityResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateIdentity(
    request: MsgCreateIdentity
  ): Promise<MsgCreateIdentityResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdateIdentity(
    request: MsgUpdateIdentity
  ): Promise<MsgUpdateIdentityResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateIdentity(
    request: MsgCreateIdentity
  ): Promise<MsgCreateIdentityResponse> {
    const data = MsgCreateIdentity.encode(request).finish();
    const promise = this.rpc.request(
      "lisoleg.identity.identity.Msg",
      "CreateIdentity",
      data
    );
    return promise.then((data) =>
      MsgCreateIdentityResponse.decode(new Reader(data))
    );
  }

  UpdateIdentity(
    request: MsgUpdateIdentity
  ): Promise<MsgUpdateIdentityResponse> {
    const data = MsgUpdateIdentity.encode(request).finish();
    const promise = this.rpc.request(
      "lisoleg.identity.identity.Msg",
      "UpdateIdentity",
      data
    );
    return promise.then((data) =>
      MsgUpdateIdentityResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
