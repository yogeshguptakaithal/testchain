import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateComment } from "./types/testchain/testchain/tx";
import { MsgUpdateComment } from "./types/testchain/testchain/tx";
import { MsgDeleteComment } from "./types/testchain/testchain/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/yogeshguptakaithal.testchain.testchain.MsgCreateComment", MsgCreateComment],
    ["/yogeshguptakaithal.testchain.testchain.MsgUpdateComment", MsgUpdateComment],
    ["/yogeshguptakaithal.testchain.testchain.MsgDeleteComment", MsgDeleteComment],
    
];

export { msgTypes }