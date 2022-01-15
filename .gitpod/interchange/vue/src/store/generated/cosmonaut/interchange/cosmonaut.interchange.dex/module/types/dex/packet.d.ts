import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "cosmonaut.interchange.dex";
export interface DexPacketData {
    noData: NoData | undefined;
    /** this line is used by starport scaffolding # ibc/packet/proto/field */
    buyOrderPacket: BuyOrderPacketData | undefined;
    /** this line is used by starport scaffolding # ibc/packet/proto/field/number */
    createPairPacket: CreatePairPacketData | undefined;
}
export interface NoData {
}
/** CreatePairPacketData defines a struct for the packet payload */
export interface CreatePairPacketData {
    sourceDenom: string;
    targetDenom: string;
}
/** CreatePairPacketAck defines a struct for the packet acknowledgment */
export interface CreatePairPacketAck {
}
/** BuyOrderPacketData defines a struct for the packet payload */
export interface BuyOrderPacketData {
    amountDenom: string;
    amount: number;
    priceDenom: string;
    price: number;
}
/** BuyOrderPacketAck defines a struct for the packet acknowledgment */
export interface BuyOrderPacketAck {
    remainingAmount: number;
    purchase: number;
}
export declare const DexPacketData: {
    encode(message: DexPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): DexPacketData;
    fromJSON(object: any): DexPacketData;
    toJSON(message: DexPacketData): unknown;
    fromPartial(object: DeepPartial<DexPacketData>): DexPacketData;
};
export declare const NoData: {
    encode(_: NoData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NoData;
    fromJSON(_: any): NoData;
    toJSON(_: NoData): unknown;
    fromPartial(_: DeepPartial<NoData>): NoData;
};
export declare const CreatePairPacketData: {
    encode(message: CreatePairPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CreatePairPacketData;
    fromJSON(object: any): CreatePairPacketData;
    toJSON(message: CreatePairPacketData): unknown;
    fromPartial(object: DeepPartial<CreatePairPacketData>): CreatePairPacketData;
};
export declare const CreatePairPacketAck: {
    encode(_: CreatePairPacketAck, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): CreatePairPacketAck;
    fromJSON(_: any): CreatePairPacketAck;
    toJSON(_: CreatePairPacketAck): unknown;
    fromPartial(_: DeepPartial<CreatePairPacketAck>): CreatePairPacketAck;
};
export declare const BuyOrderPacketData: {
    encode(message: BuyOrderPacketData, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketData;
    fromJSON(object: any): BuyOrderPacketData;
    toJSON(message: BuyOrderPacketData): unknown;
    fromPartial(object: DeepPartial<BuyOrderPacketData>): BuyOrderPacketData;
};
export declare const BuyOrderPacketAck: {
    encode(message: BuyOrderPacketAck, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): BuyOrderPacketAck;
    fromJSON(object: any): BuyOrderPacketAck;
    toJSON(message: BuyOrderPacketAck): unknown;
    fromPartial(object: DeepPartial<BuyOrderPacketAck>): BuyOrderPacketAck;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
