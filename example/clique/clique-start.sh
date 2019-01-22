#!/bin/bash
set -u
set -e

GETH=geth
echo "use ${GETH}"

NETWORK_ID=$(cat genesis.json | grep chainId | awk -F " " '{print $2}' | awk -F "," '{print $1}')
mkdir -p clique/logs

echo "[*] Starting Ethereum nodes with ChainID and NetworkId of $NETWORK_ID"
set -v
ARGS="--txpool.nolocals --nodiscover --networkid $NETWORK_ID --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,wpoa --verbosity 3"
nohup  ${GETH} --datadir clique/dd1 $ARGS --rpcport 22000 --port 21000 --unlock 0 --password passwords.txt 2>>clique/logs/1.log &
nohup  ${GETH} --datadir clique/dd2 $ARGS --rpcport 22001 --port 21001 --unlock 0 --password passwords.txt 2>>clique/logs/2.log &
nohup  ${GETH} --datadir clique/dd3 $ARGS --rpcport 22002 --port 21002 --unlock 0 --password passwords.txt 2>>clique/logs/3.log &
nohup  ${GETH} --datadir clique/dd4 $ARGS --rpcport 22003 --port 21003 --unlock 0 --password passwords.txt 2>>clique/logs/4.log &
nohup  ${GETH} --datadir clique/dd5 $ARGS --rpcport 22004 --port 21004 --unlock 0 --password passwords.txt 2>>clique/logs/5.log &
set +v

echo
echo "All nodes configured. See 'clique/logs' for logs, and run e.g. 'geth attach clique/dd1/geth.ipc' to attach to the first Geth node."

exit 0
