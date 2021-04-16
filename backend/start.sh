#!/bin/bash

export DB_ADDR="127.0.0.1:27017"
export DB_DATABASE="shinecloudnet-explorer"
export DB_USER="shinecloudnet"
export DB_PASSWD="shinecloudnetpassword"
export ADDR_NODE_SERVER="http://localhost:1317"
export ADDR_HUB_RPC="http://3.112.87.138:26657"
export CHAIN_ID="shinecloudnet"

nohup ./build/scloudplorer rest-server > scloudplorer_backend.log 2>&1 &
