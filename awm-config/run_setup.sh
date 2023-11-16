#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.


set -e # Stop on first error

# Needed for submodules
git config --global --add safe.directory '*'
if [[ $# -eq 1 ]] && [[ "$1" == "--local-tests-only" ]]; then
    dir_prefix=.
else
    dir_prefix=/code
fi

source $dir_prefix/scripts/utils.sh

# Signifies container is ready
rm -f $dir_prefix/NETWORK_READY
rm -f $dir_prefix/NETWORK_RUNNING
# Set up the network if this is the first time the container is starting
if [ ! -e $dir_prefix/NETWORK_RUNNING ]; then
    rm -f $dir_prefix/vars.sh

    cd subnet-evm

    # Source $AVALANCHEGO_VERSION from versions.sh
    source ./scripts/versions.sh

    # Build the subnet-evm
    export VM_BUILD_PATH=$"/tmp/subnet-evm-runner"
    ./scripts/build.sh $VM_BUILD_PATH/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
    cd ..

    echo "Avalanche cli version: $(avalanche --version --skip-update-check)"

    # Start the local Avalanche network
    avalanche network clean --skip-update-check

    # Configure the subnet genesis files for the three subnets to be created.
    # Avoid using sed -i due to docker + macos m1 issues
    rm -f ./subnetGenesis_*.json
    sed "s/\"<EVM_CHAIN_ID>\"/68430/g" ./docker/genesisTemplate.json > subnetGenesis_A.json
    sed "s/\"<EVM_CHAIN_ID>\"/68431/g" ./docker/genesisTemplate.json > subnetGenesis_B.json
    sed "s/\"<EVM_CHAIN_ID>\"/68432/g" ./docker/genesisTemplate.json > subnetGenesis_C.json

    # Deploy three test subnets to the local network.
    echo "Creating new subnet A..."
    avalanche subnet create subneta --force --custom --genesis ./subnetGenesis_A.json --config ./docker/defaultNodeConfig.json --vm $VM_BUILD_PATH/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy --log-level info --skip-update-check
    avalanche subnet configure subneta --config ./docker/defaultNodeConfig.json --chain-config ./docker/defaultChainConfig.json --skip-update-check
    avalanche subnet deploy subneta --local --avalanchego-version $AVALANCHEGO_VERSION --config ./docker/defaultNodeConfig.json --log-level info --skip-update-check

    echo "Creating new subnet B..."
    avalanche subnet create subnetb --force --custom --genesis ./subnetGenesis_B.json --config ./docker/defaultNodeConfig.json --vm $VM_BUILD_PATH/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy --log-level info --skip-update-check
    avalanche subnet configure subnetb --config ./docker/defaultNodeConfig.json --chain-config ./docker/defaultChainConfig.json --skip-update-check
    avalanche subnet deploy subnetb --local --avalanchego-version $AVALANCHEGO_VERSION --config ./docker/defaultNodeConfig.json --log-level info --skip-update-check

    # echo "Creating new subnet C..."
    # avalanche subnet create subnetc --force --custom --genesis ./subnetGenesis_C.json --config ./docker/defaultNodeConfig.json --vm $VM_BUILD_PATH/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy --log-level info --skip-update-check
    # avalanche subnet configure subnetc --config ./docker/defaultNodeConfig.json --chain-config ./docker/defaultChainConfig.json --skip-update-check
    # avalanche subnet deploy subnetc --local --avalanchego-version $AVALANCHEGO_VERSION --config ./docker/defaultNodeConfig.json --log-level info --skip-update-check

    # Find the proper Avalanche CLI log directory
    function getJsonVal () {
        python3 -c "import json,sys;sys.stdout.write(json.dumps(json.load(sys.stdin)$1).strip('\"'))";
    }

    subnet_a_chain_id=$(cat $HOME/.avalanche-cli/subnets/subneta/sidecar.json |  getJsonVal "['Networks']['Local Network']['BlockchainID']")
    subnet_a_subnet_id=$(cat $HOME/.avalanche-cli/subnets/subneta/sidecar.json |  getJsonVal "['Networks']['Local Network']['SubnetID']")
    subnet_b_chain_id=$(cat $HOME/.avalanche-cli/subnets/subnetb/sidecar.json |  getJsonVal "['Networks']['Local Network']['BlockchainID']")
    subnet_b_subnet_id=$(cat $HOME/.avalanche-cli/subnets/subnetb/sidecar.json |  getJsonVal "['Networks']['Local Network']['SubnetID']")
    # subnet_c_chain_id=$(cat $HOME/.avalanche-cli/subnets/subnetc/sidecar.json |  getJsonVal "['Networks']['Local Network']['BlockchainID']")
    # subnet_c_subnet_id=$(cat $HOME/.avalanche-cli/subnets/subnetc/sidecar.json |  getJsonVal "['Networks']['Local Network']['SubnetID']")
    c_chain_subnet_id=11111111111111111111111111111111LpoYY # hardcoded primary subnet ID
    c_chain_chain_id=$(curl -X POST --data '{"jsonrpc": "2.0","method": "platform.getBlockchains","params": {},"id": 1}' -H 'content-type:application/json;' 127.0.0.1:9650/ext/bc/P | getJsonVal "['result']['blockchains']" | python3 docker/getBlockChainId.py C-Chain)

    subnet_c_chain_id=2gLyawqthdiyrJktJmdnDAb1XVc6xwJXU6iJKu3Uwj21F2mXAK
    subnet_c_subnet_id=K4GygGTpKkNzzjiLfZVsmQduGqSFztJx4nk52CvA1afcFAhsH

    echo "Subnet A chain ID: $subnet_a_chain_id"
    echo "Subnet B chain ID: $subnet_b_chain_id"
    echo "Subnet C chain ID: $subnet_c_chain_id"
    echo "C-Chain chain ID: $c_chain_chain_id"

    user_private_key=0x56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027
    user_address_bytes=8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC
    user_address=0x$user_address_bytes # Address corresponding to the user_private_key

    export PATH="$PATH:$HOME/.foundry/bin"

    subnet_a_url="http://127.0.0.1:9650/ext/bc/$subnet_a_chain_id/rpc"
    subnet_b_url="http://127.0.0.1:9650/ext/bc/$subnet_b_chain_id/rpc"
    # subnet_c_url="http://127.0.0.1:9650/ext/bc/$subnet_c_chain_id/rpc"
    subnet_c_url="https://mevm.bbd.sh/v1"
    c_chain_url="http://127.0.0.1:9650/ext/bc/C/rpc"

    # Deploy TeleporterMessenger contract to each chain.
    cd contracts
    forge build
    cd ..
    go run utils/contract-deployment/contractDeploymentTools.go constructKeylessTx contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json
    teleporter_deploy_address=$(cat UniversalTeleporterDeployerAddress.txt)
    teleporter_deploy_tx=$(cat UniversalTeleporterDeployerTransaction.txt)
    teleporter_contract_address=$(cat UniversalTeleporterMessengerContractAddress.txt)
    echo $teleporter_deploy_address $teleporter_contract_address
    echo "Finished reading universal deploy address and transaction"

    cast send --private-key $user_private_key --value 50ether $teleporter_deploy_address --rpc-url $subnet_a_url
    cast send --private-key $user_private_key --value 50ether $teleporter_deploy_address --rpc-url $subnet_b_url
    # cast send --private-key $user_private_key --value 50ether $teleporter_deploy_address --rpc-url $subnet_c_url
    cast send --private-key $user_private_key --value 50ether $teleporter_deploy_address --rpc-url $c_chain_url
    echo "Sent ether to teleporter deployer on each subnet."
    # Verify that the transaction status was successful for the deployments
    status=$(cast publish --rpc-url $subnet_a_url $teleporter_deploy_tx |  getJsonVal "['status']")
    if [[ $status != "0x1" ]]; then
        echo "Error deploying Teleporter Messenger on subnet A."
        exit 1
    fi
    echo "Deployed TeleporterMessenger to Subnet A."
    status=$(cast publish --rpc-url $subnet_b_url $teleporter_deploy_tx |  getJsonVal "['status']")
    if [[ $status != "0x1" ]]; then
        echo "Error deploying Teleporter Messenger on subnet B."
        exit 1
    fi
    echo "Deployed TeleporterMessenger to Subnet B."
    # status=$(cast publish --rpc-url $subnet_c_url $teleporter_deploy_tx |  getJsonVal "['status']")
    # if [[ $status != "0x1" ]]; then
    #     echo "Error deploying Teleporter Messenger on subnet C."
    #     exit 1
    # fi
    # echo "Deployed TeleporterMessenger to Subnet C."
    status=$(cast publish --rpc-url $c_chain_url $teleporter_deploy_tx |  getJsonVal "['status']")
    if [[ $status != "0x1" ]]; then
        echo "Error deploying Teleporter Messenger on C-chain."
        exit 1
    fi
    echo "Deployed TeleporterMessenger to C-chain."
    # Deploy TeleporterRegistry to each chain.
    cd contracts
    registry_deploy_result_a=$(forge create --private-key $user_private_key \
        --rpc-url $subnet_a_url src/Teleporter/upgrades/TeleporterRegistry.sol:TeleporterRegistry --constructor-args "[(1,$teleporter_contract_address)]")
    registry_address_a=$(parseContractAddress "$registry_deploy_result_a")
    echo "TeleporterRegistry contract deployed to subnet A at $registry_address_a."

    registry_deploy_result_b=$(forge create --private-key $user_private_key \
        --rpc-url $subnet_b_url src/Teleporter/upgrades/TeleporterRegistry.sol:TeleporterRegistry --constructor-args "[(1,$teleporter_contract_address)]")
    registry_address_b=$(parseContractAddress "$registry_deploy_result_b")
    echo "TeleporterRegistry contract deployed to subnet B at $registry_address_b."
    # registry_deploy_result_c=$(forge create --private-key $user_private_key \
    #     --rpc-url $subnet_c_url src/Teleporter/upgrades/TeleporterRegistry.sol:TeleporterRegistry --constructor-args "[(1,$teleporter_contract_address)]")
    # registry_address_c=$(parseContractAddress "$registry_deploy_result_c")
    # echo "TeleporterRegistry contract deployed to subnet C at $registry_address_c."
    cd ..

    # Send tokens to cover gas costs for the relayers.
    relayer_private_key=C2CE4E001B7585F543982A01FBC537CFF261A672FA8BD1FAFC08A207098FE2DE
    relayer_address=0xA100fF48a37cab9f87c8b5Da933DA46ea1a5fb80
    cast send --private-key $user_private_key --value 500ether $relayer_address --rpc-url $subnet_a_url
    cast send --private-key $user_private_key --value 500ether $relayer_address --rpc-url $subnet_b_url
    # cast send --private-key $user_private_key --value 500ether $relayer_address --rpc-url $subnet_c_url
    cast send --private-key $user_private_key --value 500ether $relayer_address --rpc-url $c_chain_url
    echo "Sent ether to relayer account on each subnet."

    subnet_a_chain_id_hex=$(getBlockChainIDHex $subnet_a_chain_id)
    subnet_b_chain_id_hex=$(getBlockChainIDHex $subnet_b_chain_id)
    subnet_c_chain_id_hex=$(getBlockChainIDHex $subnet_c_chain_id)
    subnet_a_subnet_id_hex=$(getBlockChainIDHex $subnet_a_subnet_id)
    subnet_b_subnet_id_hex=$(getBlockChainIDHex $subnet_b_subnet_id)
    subnet_c_subnet_id_hex=$(getBlockChainIDHex $subnet_c_subnet_id)
    c_chain_chain_id_hex=$(getBlockChainIDHex $c_chain_chain_id)
    c_chain_subnet_id_hex=$(getBlockChainIDHex $c_chain_subnet_id)
    warp_messenger_precompile_addr=0x0200000000000000000000000000000000000005
    # Write all vars to file so that they can be imported from another container
    set > $dir_prefix/vars.sh

    # Indicate for future runs that we started up successfully
    touch $dir_prefix/NETWORK_RUNNING
else
    echo "Resuming network from previous run"
    source $dir_prefix/vars.sh || true
    avalanche network start --skip-update-check
    restart="restart"
fi

# Signal other containers this container is ready
touch $dir_prefix/NETWORK_READY

function cleanup()
{
    echo "Stopping the network..."
    avalanche network stop --skip-update-check
    echo "Gracefully shut down the network."
}

trap cleanup EXIT

# Stream the subnet logs
network_runner_dir=$(python3 docker/findCliLogDirectory.py $restart)
subnet_a_log_file=$HOME/.avalanche-cli/runs/$network_runner_dir/node1/logs/$subnet_a_chain_id.log
subnet_b_log_file=$HOME/.avalanche-cli/runs/$network_runner_dir/node1/logs/$subnet_b_chain_id.log
# subnet_c_log_file=$HOME/.avalanche-cli/runs/$network_runner_dir/node1/logs/$subnet_c_chain_id.log

echo "Streaming subnet A log file at $subnet_a_log_file"
echo "Streaming subnet B log file at $subnet_b_log_file"
# echo "Streaming subnet C log file at $subnet_c_log_file"

tail -f $subnet_a_log_file &
tail -f $subnet_b_log_file &
# tail -f $subnet_c_log_file &

wait $!
# x/warp/contract.go
# x/warp/contract_warp_handler.go