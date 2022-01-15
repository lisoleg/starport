import { Reader, Writer } from "protobufjs/minimal";
import { SellOrderBook } from "../dex/sell_order_book";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { DenomTrace } from "../dex/denom_trace";
export declare const protobufPackage = "cosmonaut.interchange.dex";
export interface QueryGetSellOrderBookRequest {
    index: string;
}
export interface QueryGetSellOrderBookResponse {
    sellOrderBook: SellOrderBook | undefined;
}
export interface QueryAllSellOrderBookRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSellOrderBookResponse {
    sellOrderBook: SellOrderBook[];
    pagination: PageResponse | undefined;
}
export interface QueryGetDenomTraceRequest {
    index: string;
}
export interface QueryGetDenomTraceResponse {
    denomTrace: DenomTrace | undefined;
}
export interface QueryAllDenomTraceRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllDenomTraceResponse {
    denomTrace: DenomTrace[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetSellOrderBookRequest: {
    encode(message: QueryGetSellOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSellOrderBookRequest;
    fromJSON(object: any): QueryGetSellOrderBookRequest;
    toJSON(message: QueryGetSellOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSellOrderBookRequest>): QueryGetSellOrderBookRequest;
};
export declare const QueryGetSellOrderBookResponse: {
    encode(message: QueryGetSellOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSellOrderBookResponse;
    fromJSON(object: any): QueryGetSellOrderBookResponse;
    toJSON(message: QueryGetSellOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSellOrderBookResponse>): QueryGetSellOrderBookResponse;
};
export declare const QueryAllSellOrderBookRequest: {
    encode(message: QueryAllSellOrderBookRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSellOrderBookRequest;
    fromJSON(object: any): QueryAllSellOrderBookRequest;
    toJSON(message: QueryAllSellOrderBookRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSellOrderBookRequest>): QueryAllSellOrderBookRequest;
};
export declare const QueryAllSellOrderBookResponse: {
    encode(message: QueryAllSellOrderBookResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSellOrderBookResponse;
    fromJSON(object: any): QueryAllSellOrderBookResponse;
    toJSON(message: QueryAllSellOrderBookResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSellOrderBookResponse>): QueryAllSellOrderBookResponse;
};
export declare const QueryGetDenomTraceRequest: {
    encode(message: QueryGetDenomTraceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDenomTraceRequest;
    fromJSON(object: any): QueryGetDenomTraceRequest;
    toJSON(message: QueryGetDenomTraceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetDenomTraceRequest>): QueryGetDenomTraceRequest;
};
export declare const QueryGetDenomTraceResponse: {
    encode(message: QueryGetDenomTraceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDenomTraceResponse;
    fromJSON(object: any): QueryGetDenomTraceResponse;
    toJSON(message: QueryGetDenomTraceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetDenomTraceResponse>): QueryGetDenomTraceResponse;
};
export declare const QueryAllDenomTraceRequest: {
    encode(message: QueryAllDenomTraceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDenomTraceRequest;
    fromJSON(object: any): QueryAllDenomTraceRequest;
    toJSON(message: QueryAllDenomTraceRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllDenomTraceRequest>): QueryAllDenomTraceRequest;
};
export declare const QueryAllDenomTraceResponse: {
    encode(message: QueryAllDenomTraceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDenomTraceResponse;
    fromJSON(object: any): QueryAllDenomTraceResponse;
    toJSON(message: QueryAllDenomTraceResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllDenomTraceResponse>): QueryAllDenomTraceResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a sellOrderBook by index. */
    SellOrderBook(request: QueryGetSellOrderBookRequest): Promise<QueryGetSellOrderBookResponse>;
    /** Queries a list of sellOrderBook items. */
    SellOrderBookAll(request: QueryAllSellOrderBookRequest): Promise<QueryAllSellOrderBookResponse>;
    /** Queries a denomTrace by index. */
    DenomTrace(request: QueryGetDenomTraceRequest): Promise<QueryGetDenomTraceResponse>;
    /** Queries a list of denomTrace items. */
    DenomTraceAll(request: QueryAllDenomTraceRequest): Promise<QueryAllDenomTraceResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    SellOrderBook(request: QueryGetSellOrderBookRequest): Promise<QueryGetSellOrderBookResponse>;
    SellOrderBookAll(request: QueryAllSellOrderBookRequest): Promise<QueryAllSellOrderBookResponse>;
    DenomTrace(request: QueryGetDenomTraceRequest): Promise<QueryGetDenomTraceResponse>;
    DenomTraceAll(request: QueryAllDenomTraceRequest): Promise<QueryAllDenomTraceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
