/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "lto.lto";

export interface MsgCreateAnchor {
  creator: string;
  hash: string;
  encoding: string;
}

export interface MsgCreateAnchorResponse {
  success: boolean;
}

function createBaseMsgCreateAnchor(): MsgCreateAnchor {
  return { creator: "", hash: "", encoding: "" };
}

export const MsgCreateAnchor = {
  encode(message: MsgCreateAnchor, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.hash !== "") {
      writer.uint32(18).string(message.hash);
    }
    if (message.encoding !== "") {
      writer.uint32(26).string(message.encoding);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAnchor {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAnchor();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.hash = reader.string();
          break;
        case 3:
          message.encoding = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAnchor {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      hash: isSet(object.hash) ? String(object.hash) : "",
      encoding: isSet(object.encoding) ? String(object.encoding) : "",
    };
  },

  toJSON(message: MsgCreateAnchor): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.hash !== undefined && (obj.hash = message.hash);
    message.encoding !== undefined && (obj.encoding = message.encoding);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAnchor>, I>>(object: I): MsgCreateAnchor {
    const message = createBaseMsgCreateAnchor();
    message.creator = object.creator ?? "";
    message.hash = object.hash ?? "";
    message.encoding = object.encoding ?? "";
    return message;
  },
};

function createBaseMsgCreateAnchorResponse(): MsgCreateAnchorResponse {
  return { success: false };
}

export const MsgCreateAnchorResponse = {
  encode(message: MsgCreateAnchorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAnchorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAnchorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAnchorResponse {
    return { success: isSet(object.success) ? Boolean(object.success) : false };
  },

  toJSON(message: MsgCreateAnchorResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAnchorResponse>, I>>(object: I): MsgCreateAnchorResponse {
    const message = createBaseMsgCreateAnchorResponse();
    message.success = object.success ?? false;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAnchor(request: MsgCreateAnchor): Promise<MsgCreateAnchorResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAnchor = this.CreateAnchor.bind(this);
  }
  CreateAnchor(request: MsgCreateAnchor): Promise<MsgCreateAnchorResponse> {
    const data = MsgCreateAnchor.encode(request).finish();
    const promise = this.rpc.request("lto.lto.Msg", "CreateAnchor", data);
    return promise.then((data) => MsgCreateAnchorResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
