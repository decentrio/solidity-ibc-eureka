echo '{
 "messages": [
  {
   "@type": "/ibc.core.client.v2.MsgRegisterCounterparty",
   "client_id": "08-wasm-0",
   "counterparty_merkle_prefix": [""],
   "counterparty_client_id": "cosmoshub-1",
   "signer": "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"
  }
 ],
 "metadata": "ipfs://CID",
 "deposit": "10000000stake",
 "title": "register counterparty",
 "summary": "register counterparty",
 "expedited": false
}' > register_counterparty_proposal.json
# gaiad tx ibc client add-counterparty "cosmoshub-1" "08-wasm-0" "" --from test1 --keyring-backend test --gas-prices 1stake -y

gaiad tx gov submit-proposal register_counterparty_proposal.json --from test1 --keyring-backend test --gas 200000000 --gas-prices 1stake -y

sleep 5
gaiad tx gov vote 3 yes --from test --keyring-backend test --gas-prices 1stake -y

sleep 30