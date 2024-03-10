#!/bin/bash

# This script is used to deploy the UniV2 contract to the target network
# This avoids trying to predict what address it might be deployed to
set -e

evm_endpoint=$1

echo "Deploying UniswapV2 contracts to $evm_endpoint"

cd loadtest/contracts/evm || exit 1

bigNumber=100000000000000000000000000000000 # 10^32


# deploy the uniswapV2 factory contract
feeCollector=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 # first anvil address, just need a random address for fees
wallet=0xF87A299e6bC7bEba58dbBe5a5Aa21d49bCD16D52

echo "Deploying factory contract..."
factoryAddress=$(forge create -r "$evm_endpoint" --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e src/univ2/UniswapV2Factory.sol:UniswapV2Factory --json --constructor-args $feeCollector | jq -r '.deployedTo')

echo "Deploying router contract..."
routerAddress=$(forge create -r "$evm_endpoint" --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e src/univ2/UniswapV2Router.sol:UniswapV2Router --json --constructor-args $factoryAddress $feeCollector | jq -r '.deployedTo')

# create ERC20s
echo "Deploying token1 contract..."
token1Address=$(forge create -r "$evm_endpoint" --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e src/ERC20Token.sol:ERC20Token --json --constructor-args "Token1" "T1" | jq -r '.deployedTo')

echo "Deploying token2 contract..."
token2Address=$(forge create -r "$evm_endpoint" --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e src/ERC20Token.sol:ERC20Token --json --constructor-args "Token2" "T2" | jq -r '.deployedTo')

echo "Minting tokens..."
cast send -r "$evm_endpoint" $token1Address --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "mint(address,uint256)" $wallet $bigNumber --legacy --json 1> /dev/null
cast send -r "$evm_endpoint" $token2Address --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "mint(address,uint256)" $wallet $bigNumber --legacy --json 1> /dev/null

echo "Creating a pool..."
cast send -r "$evm_endpoint" $factoryAddress --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "createPair(address,address)" $token1Address $token2Address --legacy --json

# get the pair address
pairAddress=$(cast call -r "$evm_endpoint" $factoryAddress "getAllPairsIndex(uint256)" 0)

echo "Approving router to spend tokens..."
cast send -r "$evm_endpoint" $token1Address --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "approve(address,uint256)" $routerAddress $bigNumber --legacy --json
cast send -r "$evm_endpoint" $token2Address --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "approve(address,uint256)" $routerAddress $bigNumber --legacy --json

echo "Adding liquidity to the pool..."
cast send -r "$evm_endpoint" $routerAddress --private-key 57acb95d82739866a5c29e40b0aa2590742ae50425b7dd5b5d279a986370189e "addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)" $token1Address $token2Address $bigNumber $bigNumber 0 0 $wallet 1000000000000000000 --legacy --json

# print addresses out for use in other scripts
echo "UniswapV2Factory Address: \"$factoryAddress\""
echo "UniswapV2Router Address: \"$routerAddress\""
echo "Token1 Address: \"$token1Address\""
echo "Token2 Address: \"$token2Address\""
echo "Pair Address: \"$pairAddress\""
