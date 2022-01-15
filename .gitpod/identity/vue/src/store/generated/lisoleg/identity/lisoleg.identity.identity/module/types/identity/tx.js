/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "lisoleg.identity.identity";
const baseMsgCreateIdentity = {
    creator: "",
    id: "",
    pubKey: "",
    certificate: "",
    credentials: "",
    owner: "",
    data: "",
};
export const MsgCreateIdentity = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateIdentity };
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
    fromJSON(object) {
        const message = { ...baseMsgCreateIdentity };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = String(object.pubKey);
        }
        else {
            message.pubKey = "";
        }
        if (object.certificate !== undefined && object.certificate !== null) {
            message.certificate = String(object.certificate);
        }
        else {
            message.certificate = "";
        }
        if (object.credentials !== undefined && object.credentials !== null) {
            message.credentials = String(object.credentials);
        }
        else {
            message.credentials = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = String(object.data);
        }
        else {
            message.data = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = { ...baseMsgCreateIdentity };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = object.pubKey;
        }
        else {
            message.pubKey = "";
        }
        if (object.certificate !== undefined && object.certificate !== null) {
            message.certificate = object.certificate;
        }
        else {
            message.certificate = "";
        }
        if (object.credentials !== undefined && object.credentials !== null) {
            message.credentials = object.credentials;
        }
        else {
            message.credentials = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = "";
        }
        return message;
    },
};
const baseMsgCreateIdentityResponse = {};
export const MsgCreateIdentityResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateIdentityResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgCreateIdentityResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateIdentityResponse,
        };
        return message;
    },
};
const baseMsgUpdateIdentity = {
    creator: "",
    id: "",
    pubKey: "",
    certificate: "",
    credentials: "",
    owner: "",
    data: "",
};
export const MsgUpdateIdentity = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgUpdateIdentity };
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
    fromJSON(object) {
        const message = { ...baseMsgUpdateIdentity };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = String(object.id);
        }
        else {
            message.id = "";
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = String(object.pubKey);
        }
        else {
            message.pubKey = "";
        }
        if (object.certificate !== undefined && object.certificate !== null) {
            message.certificate = String(object.certificate);
        }
        else {
            message.certificate = "";
        }
        if (object.credentials !== undefined && object.credentials !== null) {
            message.credentials = String(object.credentials);
        }
        else {
            message.credentials = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = String(object.owner);
        }
        else {
            message.owner = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = String(object.data);
        }
        else {
            message.data = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
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
    fromPartial(object) {
        const message = { ...baseMsgUpdateIdentity };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = "";
        }
        if (object.pubKey !== undefined && object.pubKey !== null) {
            message.pubKey = object.pubKey;
        }
        else {
            message.pubKey = "";
        }
        if (object.certificate !== undefined && object.certificate !== null) {
            message.certificate = object.certificate;
        }
        else {
            message.certificate = "";
        }
        if (object.credentials !== undefined && object.credentials !== null) {
            message.credentials = object.credentials;
        }
        else {
            message.credentials = "";
        }
        if (object.owner !== undefined && object.owner !== null) {
            message.owner = object.owner;
        }
        else {
            message.owner = "";
        }
        if (object.data !== undefined && object.data !== null) {
            message.data = object.data;
        }
        else {
            message.data = "";
        }
        return message;
    },
};
const baseMsgUpdateIdentityResponse = {};
export const MsgUpdateIdentityResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgUpdateIdentityResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgUpdateIdentityResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgUpdateIdentityResponse,
        };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateIdentity(request) {
        const data = MsgCreateIdentity.encode(request).finish();
        const promise = this.rpc.request("lisoleg.identity.identity.Msg", "CreateIdentity", data);
        return promise.then((data) => MsgCreateIdentityResponse.decode(new Reader(data)));
    }
    UpdateIdentity(request) {
        const data = MsgUpdateIdentity.encode(request).finish();
        const promise = this.rpc.request("lisoleg.identity.identity.Msg", "UpdateIdentity", data);
        return promise.then((data) => MsgUpdateIdentityResponse.decode(new Reader(data)));
    }
}
