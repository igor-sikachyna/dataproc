#!/bin/bash

rm -r ~/.dataprocd || true
DATAPROCD_BIN=$(which dataprocd)
# configure dataprocd
$DATAPROCD_BIN config set client chain-id demo
$DATAPROCD_BIN config set client keyring-backend test
$DATAPROCD_BIN keys add alice
$DATAPROCD_BIN keys add bob
$DATAPROCD_BIN init test --chain-id demo --default-denom mini
# update genesis
$DATAPROCD_BIN genesis add-genesis-account alice 10000000mini --keyring-backend test
$DATAPROCD_BIN genesis add-genesis-account bob 1000mini --keyring-backend test
# create default validator
$DATAPROCD_BIN genesis gentx alice 1000000mini --chain-id demo
$DATAPROCD_BIN genesis collect-gentxs