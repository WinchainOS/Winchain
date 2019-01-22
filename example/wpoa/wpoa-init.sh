#!/bin/bash
set -u
set -e

GETH=geth
echo "use ${GETH}"

echo "[*] Cleaning up temporary data directories"
rm -rf winchain
mkdir -p winchain/logs

echo "[*] Configuring node 1 (static)"
mkdir -p winchain/dd1/{keystore,geth}
cp static-nodes.json winchain/dd1/
cp keys/key1 winchain/dd1/keystore
cp keys/nodekey1 winchain/dd1/geth/nodekey
${GETH} --datadir winchain/dd1 init genesis.json

echo "[*] Configuring node 2 (static)"
mkdir -p winchain/dd2/{keystore,geth}
cp static-nodes.json winchain/dd2/
cp keys/key2 winchain/dd2/keystore
cp keys/nodekey2 winchain/dd2/geth/nodekey
${GETH} --datadir winchain/dd2 init genesis.json

echo "[*] Configuring node 3 (static)"
mkdir -p winchain/dd3/{keystore,geth}
cp static-nodes.json winchain/dd3/
cp keys/key3 winchain/dd3/keystore
cp keys/nodekey3 winchain/dd3/geth/nodekey
${GETH} --datadir winchain/dd3 init genesis.json

echo "[*] Configuring node 4 (static)"
mkdir -p winchain/dd4/{keystore,geth}
cp static-nodes.json winchain/dd4/
cp keys/key4 winchain/dd4/keystore
cp keys/nodekey4 winchain/dd4/geth/nodekey
${GETH} --datadir winchain/dd4 init genesis.json

echo "[*] Configuring node 5"
mkdir -p winchain/dd5/{keystore,geth}
cp static-nodes.json winchain/dd5/
cp keys/key5 winchain/dd5/keystore
cp keys/nodekey5 winchain/dd5/geth/nodekey
${GETH} --datadir winchain/dd5 init genesis.json
