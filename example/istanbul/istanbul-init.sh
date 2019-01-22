#!/bin/bash
set -u
set -e

GETH=geth

echo "use ${GETH}"

echo "[*] Cleaning up temporary data directories"
rm -rf qdata
mkdir -p qdata/logs

echo "[*] Configuring node 1"
mkdir -p qdata/dd1/{keystore,geth}
cp permissioned-nodes.json qdata/dd1/static-nodes.json
cp keys/key1 qdata/dd1/keystore
cp keys/nodekey1 qdata/dd1/geth/nodekey
${GETH} --datadir qdata/dd1 init istanbul-genesis.json

echo "[*] Configuring node 2"
mkdir -p qdata/dd2/{keystore,geth}
cp permissioned-nodes.json qdata/dd2/static-nodes.json
cp keys/key2 qdata/dd2/keystore
cp keys/nodekey2 qdata/dd2/geth/nodekey
${GETH} --datadir qdata/dd2 init istanbul-genesis.json

echo "[*] Configuring node 3"
mkdir -p qdata/dd3/{keystore,geth}
cp permissioned-nodes.json qdata/dd3/static-nodes.json
cp keys/key3 qdata/dd3/keystore
cp keys/nodekey3 qdata/dd3/geth/nodekey
${GETH} --datadir qdata/dd3 init istanbul-genesis.json

echo "[*] Configuring node 4 as voter"
mkdir -p qdata/dd4/{keystore,geth}
cp permissioned-nodes.json qdata/dd4/static-nodes.json
cp keys/key4 qdata/dd4/keystore
cp keys/nodekey4 qdata/dd4/geth/nodekey
${GETH} --datadir qdata/dd4 init istanbul-genesis.json

echo "[*] Configuring node 5 as voter"
mkdir -p qdata/dd5/{keystore,geth}
cp permissioned-nodes.json qdata/dd5/static-nodes.json
cp keys/key5 qdata/dd5/keystore
cp keys/nodekey5 qdata/dd5/geth/nodekey
${GETH} --datadir qdata/dd5 init istanbul-genesis.json

echo "[*] Configuring node 6"
mkdir -p qdata/dd6/{keystore,geth}
cp permissioned-nodes.json qdata/dd6/static-nodes.json
cp keys/key6 qdata/dd6/keystore
cp keys/nodekey6 qdata/dd6/geth/nodekey
${GETH} --datadir qdata/dd6 init istanbul-genesis.json

echo "[*] Configuring node 7"
mkdir -p qdata/dd7/{keystore,geth}
cp permissioned-nodes.json qdata/dd7/static-nodes.json
cp keys/key7 qdata/dd7/keystore
cp keys/nodekey7 qdata/dd7/geth/nodekey
${GETH} --datadir qdata/dd7 init istanbul-genesis.json
