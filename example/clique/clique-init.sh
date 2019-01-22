#!/bin/bash
set -u
set -e

GETH=geth
echo "use ${GETH}"

echo "[*] Cleaning up temporary data directories"
rm -rf clique
mkdir -p clique/logs

echo "[*] Configuring node 1 (static)"
mkdir -p clique/dd1/{keystore,geth}
cp static-nodes.json clique/dd1/
cp keys/key1 clique/dd1/keystore
cp keys/nodekey1 clique/dd1/geth/nodekey
${GETH} --datadir clique/dd1 init genesis.json

echo "[*] Configuring node 2 (static)"
mkdir -p clique/dd2/{keystore,geth}
cp static-nodes.json clique/dd2/
cp keys/key2 clique/dd2/keystore
cp keys/nodekey2 clique/dd2/geth/nodekey
${GETH} --datadir clique/dd2 init genesis.json

echo "[*] Configuring node 3 (static)"
mkdir -p clique/dd3/{keystore,geth}
cp static-nodes.json clique/dd3/
cp keys/key3 clique/dd3/keystore
cp keys/nodekey3 clique/dd3/geth/nodekey
${GETH} --datadir clique/dd3 init genesis.json

echo "[*] Configuring node 4 (static)"
mkdir -p clique/dd4/{keystore,geth}
cp static-nodes.json clique/dd4/
cp keys/key4 clique/dd4/keystore
cp keys/nodekey4 clique/dd4/geth/nodekey
${GETH} --datadir clique/dd4 init genesis.json

echo "[*] Configuring node 5"
mkdir -p clique/dd5/{keystore,geth}
cp static-nodes.json clique/dd5/
cp keys/key5 clique/dd5/keystore
cp keys/nodekey5 clique/dd5/geth/nodekey
${GETH} --datadir clique/dd5 init genesis.json
