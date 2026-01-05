#!/bin/bash

function create_binding {
    contract_dir=$1
    contract=$2
    binding_dir=$3
    echo $contract
    mkdir -p $binding_dir/${contract}
    contract_json="../out/${contract}.sol/${contract}.json"
    solc_abi=$(cat ${contract_json} | jq -r '.abi')
    solc_bin=$(cat ${contract_json} | jq -r '.bytecode.object')

    mkdir -p data
    echo ${solc_abi} > data/tmp.abi
    echo ${solc_bin} > data/tmp.bin

    rm -f $binding_dir/${contract}/binding.go
    abigen --bin=data/tmp.bin --abi=data/tmp.abi --pkg=contract${contract} --out=$binding_dir/${contract}/binding.go
}

forge clean
forge build

create_binding ./programs/ "Membership" ../operator/bindings
create_binding ./programs/ "Misbehaviour" ../operator/bindings
create_binding ./programs/ "UpdateClient" ../operator/bindings
create_binding ./light-clients/ "SP1ICS07Tendermint" ../operator/bindings
create_binding ./ "ICS20Transfer" ../operator/bindings
create_binding ./ "ICS26Router" ../operator/bindings

# ./compile.sh ./ ICS26Router ./bindings 