/* eslint-disable */
import { SellOrderBook } from "../dex/sell_order_book";
import { DenomTrace } from "../dex/denom_trace";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "cosmonaut.interchange.dex";
const baseGenesisState = { portId: "" };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.portId !== "") {
            writer.uint32(10).string(message.portId);
        }
        for (const v of message.sellOrderBookList) {
            SellOrderBook.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.denomTraceList) {
            DenomTrace.encode(v, writer.uint32(26).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.sellOrderBookList = [];
        message.denomTraceList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.portId = reader.string();
                    break;
                case 2:
                    message.sellOrderBookList.push(SellOrderBook.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.denomTraceList.push(DenomTrace.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.sellOrderBookList = [];
        message.denomTraceList = [];
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = String(object.portId);
        }
        else {
            message.portId = "";
        }
        if (object.sellOrderBookList !== undefined &&
            object.sellOrderBookList !== null) {
            for (const e of object.sellOrderBookList) {
                message.sellOrderBookList.push(SellOrderBook.fromJSON(e));
            }
        }
        if (object.denomTraceList !== undefined && object.denomTraceList !== null) {
            for (const e of object.denomTraceList) {
                message.denomTraceList.push(DenomTrace.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.portId !== undefined && (obj.portId = message.portId);
        if (message.sellOrderBookList) {
            obj.sellOrderBookList = message.sellOrderBookList.map((e) => e ? SellOrderBook.toJSON(e) : undefined);
        }
        else {
            obj.sellOrderBookList = [];
        }
        if (message.denomTraceList) {
            obj.denomTraceList = message.denomTraceList.map((e) => e ? DenomTrace.toJSON(e) : undefined);
        }
        else {
            obj.denomTraceList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.sellOrderBookList = [];
        message.denomTraceList = [];
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = object.portId;
        }
        else {
            message.portId = "";
        }
        if (object.sellOrderBookList !== undefined &&
            object.sellOrderBookList !== null) {
            for (const e of object.sellOrderBookList) {
                message.sellOrderBookList.push(SellOrderBook.fromPartial(e));
            }
        }
        if (object.denomTraceList !== undefined && object.denomTraceList !== null) {
            for (const e of object.denomTraceList) {
                message.denomTraceList.push(DenomTrace.fromPartial(e));
            }
        }
        return message;
    },
};
