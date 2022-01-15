import { Writer, Reader } from "protobufjs/minimal";
import { Post } from "../blog/post";
import { SentPost } from "../blog/sent_post";
import { TimedoutPost } from "../blog/timedout_post";
export declare const protobufPackage = "cosmonaut.planet.blog";
/** GenesisState defines the blog module's genesis state. */
export interface GenesisState {
    portId: string;
    postList: Post[];
    postCount: number;
    sentPostList: SentPost[];
    sentPostCount: number;
    timedoutPostList: TimedoutPost[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    timedoutPostCount: number;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
