import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdatePool } from "./types/protoqol/dex/tx";
import { MsgAddLiquidity } from "./types/protoqol/dex/tx";
import { MsgCreatePool } from "./types/protoqol/dex/tx";
import { MsgDeletePool } from "./types/protoqol/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/protoqol.dex.MsgUpdatePool", MsgUpdatePool],
    ["/protoqol.dex.MsgAddLiquidity", MsgAddLiquidity],
    ["/protoqol.dex.MsgCreatePool", MsgCreatePool],
    ["/protoqol.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }