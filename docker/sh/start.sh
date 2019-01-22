#!/bin/bash
set -u

echo
echo "========================[start]=========================="
GETH=geth
echo "Use ${GETH}"

NETWORK_ID=2018
mkdir -p /data/winchain/logs

echo "[*] Starting Winchain nodes with ChainID and NetworkId of $NETWORK_ID"
echo
echo "Node configured. See '/data/winchain/logs' for logs, and run e.g. 'geth attach /data/winchain/geth.ipc' to attach to the first Winchain node."
geth --datadir /data/winchain --txpool.nolocals --nodiscover --networkid $NETWORK_ID --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi db,eth,debug,miner,net,shh,txpool,personal,web3,wpoa --rpccorsdomain='*' --verbosity 5 --rpcport 9545 --port 30303