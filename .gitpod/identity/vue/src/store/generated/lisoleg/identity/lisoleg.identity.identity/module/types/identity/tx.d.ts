import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "lisoleg.identity.identity";
export interface MsgCreateIdentity {
    creator: string;
    id: string;
    pubKey: string;
    certificate: string;
    credentials: string;
    owner: string;
    data: string;
}
export interface MsgCreateIdentityResponse {
}
export interface MsgUpdateIdentity {
    creator: string;
    id: string;
    pubKey: string;
    certificate: string;
    credentials: string;
    owner: string;
    data: string;
}
export interface MsgUpdateIdentityResponse {
}
export declare const MsgCreateIdentity: {
    encode(message: MsgCreateIdentity, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentity;
    fromJSON(object: any): MsgCreateIdentity;
    toJSON(message: MsgCreateIdentity): unknown;
    fromPartial(object: DeepPartial<MsgCreateIdentity>): MsgCreateIdentity;
};
export declare const MsgCreateIdentityResponse: {
    encode(_: MsgCreateIdentityResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateIdentityResponse;
    fromJSON(_: any): MsgCreateIdentityResponse;
    toJSON(_: MsgCreateIdentityResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateIdentityResponse>): MsgCreateIdentityResponse;
};
export declare const MsgUpdateIdentity: {
    encode(message: MsgUpdateIdentity, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateIdentity;
    fromJSON(object: any): MsgUpdateIdentity;
    toJSON(message: MsgUpdateIdentity): unknown;
    fromPartial(object: DeepPartial<MsgUpdateIdentity>): MsgUpdateIdentity;
};
export declare const MsgUpdateIdentityResponse: {
    encode(_: MsgUpdateIdentityResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateIdentityResponse;
    fromJSON(_: any): MsgUpdateIdentityResponse;
    toJSON(_: MsgUpdateIdentityResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateIdentityResponse>): MsgUpdateIdentityResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateIdentity(request: MsgCreateIdentity): Promise<MsgCreateIdentityResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    UpdateIdentity(request: MsgUpdateIdentity): Promise<MsgUpdateIdentityResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateIdentity(request: MsgCreateIdentity): Promise<MsgCreateIdentityResponse>;
    UpdateIdentity(request: MsgUpdateIdentity): Promise<MsgUpdateIdentityResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
