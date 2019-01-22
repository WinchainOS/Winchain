#!/bin/bash
set -u
set -e

GETH=geth

echo "use ${GETH}"

NETWORK_ID=$(cat istanbul-genesis.json | grep chainId | awk -F " " '{print $2}' | awk -F "," '{print $1}')

if [ $NETWORK_ID -eq 1 ]
then
    echo "  Quorum should not be run with a chainId of 1 (Ethereum mainnet)"
    echo "  please set the chainId in the istanbul-genesis.json to another value "
    echo "  1337 is the recommend ChainId for Geth private clients."
fi

mkdir -p qdata/logs

echo "[*] Starting Ethereum nodes with ChainID and NetworkId of $NETWORK_ID"
set -v
ARGS="--nodiscover --istanbul.blockperiod 5 --networkid $NETWORK_ID --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,istanbul --verbosity 3"
nohup ${GETH} --datadir qdata/dd1 $ARGS --rpcport 22000 --port 21000 --unlock 0 --password passwords.txt 2>>qdata/logs/1.log &
nohup ${GETH} --datadir qdata/dd2 $ARGS --rpcport 22001 --port 21001 --unlock 0 --password passwords.txt 2>>qdata/logs/2.log &
nohup ${GETH} --datadir qdata/dd3 $ARGS --rpcport 22002 --port 21002 --unlock 0 --password passwords.txt 2>>qdata/logs/3.log &
nohup ${GETH} --datadir qdata/dd4 $ARGS --rpcport 22003 --port 21003 --unlock 0 --password passwords.txt 2>>qdata/logs/4.log &
nohup ${GETH} --datadir qdata/dd5 $ARGS --rpcport 22004 --port 21004 --unlock 0 --password passwords.txt 2>>qdata/logs/5.log &
nohup ${GETH} --datadir qdata/dd6 $ARGS --rpcport 22005 --port 21005 --unlock 0 --password passwords.txt 2>>qdata/logs/6.log &
nohup ${GETH} --datadir qdata/dd7 $ARGS --rpcport 22006 --port 21006 --unlock 0 --password passwords.txt 2>>qdata/logs/7.log &
set +v

echo
echo "All nodes configured. See 'qdata/logs' for logs, and run e.g. 'geth attach qdata/dd1/geth.ipc' to attach to the first Geth node."

exit 0
