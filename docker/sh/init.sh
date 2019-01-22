#!/bin/bash
set -u

echo
echo "========================[init]=========================="
GETH=geth
echo "use ${GETH}"

echo "[*] Cleaning up temporary data directories"
rm -rf /data/winchain
mkdir -p /data/winchain/logs

echo "[*] Configuring node (static)"
mkdir -p /data/winchain/keystore
mkdir -p /data/winchain/geth
cp /example/static-nodes.json /data/winchain/

# 生成nodekey
if [ ! -f "/data/winchain/geth/nodekey" ]; then
    cp /example/nodekgstey /data/winchain/geth/
fi

${GETH} --datadir /data/winchain init /example/genesis.json