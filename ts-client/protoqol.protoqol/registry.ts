import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdatePost } from "./types/protoqol/protoqol/tx";
import { MsgCreatePost } from "./types/protoqol/protoqol/tx";
import { MsgDeletePost } from "./types/protoqol/protoqol/tx";
import { MsgWriteComment } from "./types/protoqol/protoqol/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/protoqol.protoqol.MsgUpdatePost", MsgUpdatePost],
    ["/protoqol.protoqol.MsgCreatePost", MsgCreatePost],
    ["/protoqol.protoqol.MsgDeletePost", MsgDeletePost],
    ["/protoqol.protoqol.MsgWriteComment", MsgWriteComment],
    
];

export { msgTypes }