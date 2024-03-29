package _step

// 测试一下基于 json 的 rpc
// - 客户端实现、服务端运行方式 有变更
//
// 一下也不好看出什么效果，因为都是底层，所以需要抓包：
// 使用 WireShark 监听交互的端口，先运行 step2 的 客户端和服务端；再运行 step3 的 客户端和服务端
//
// 非 json
// 0000   02 00 00 00 45 00 00 9d 89 1f 40 00 80 06 00 00   ....E.....@.....
// 0010   7f 00 00 01 7f 00 00 01 c8 0f 04 d2 e7 9b 01 e2   ................
// 0020   84 a4 aa 68 50 18 27 f9 e7 6b 00 00 2f ff 81 03   ...hP.'..k../...
// 0030   01 01 07 52 65 71 75 65 73 74 01 ff 82 00 01 02   ...Request......
// 0040   01 0d 53 65 72 76 69 63 65 4d 65 74 68 6f 64 01   ..ServiceMethod.
// 0050   0c 00 01 03 53 65 71 01 06 00 00 00 17 ff 82 01   ....Seq.........
// 0060   12 48 65 6c 6c 6f 53 65 72 76 69 63 65 2e 48 65   .HelloService.He
// 0070   6c 6c 6f 00 1c ff 83 03 01 01 04 41 72 67 73 01   llo........Args.
// 0080   ff 84 00 01 01 01 05 56 61 6c 75 65 01 0c 00 00   .......Value....
// 0090   00 0f ff 84 01 0a 49 27 6d 20 63 6c 69 65 6e 74   ......I'm client
// 00a0   00
// .
// json
// 0000   02 00 00 00 45 00 00 71 fd 24 40 00 80 06 00 00   ....E..q.$@.....
// 0010   7f 00 00 01 7f 00 00 01 cb f7 04 d2 bf bf 9f b2   ................
// 0020   c6 42 14 e1 50 18 27 f9 ad d4 00 00 7b 22 6d 65   .B..P.'.....{"me
// 0030   74 68 6f 64 22 3a 22 48 65 6c 6c 6f 53 65 72 76   thod":"HelloServ
// 0040   69 63 65 2e 48 65 6c 6c 6f 22 2c 22 70 61 72 61   ice.Hello","para
// 0050   6d 73 22 3a 5b 7b 22 76 61 6c 75 65 22 3a 22 49   ms":[{"value":"I
// 0060   27 6d 20 63 6c 69 65 6e 74 22 7d 5d 2c 22 69 64   'm client"}],"id
// 0070   22 3a 30 7d 0a                                    ":0}.
