//create database and user
// use sync-iris
// db.createUser(
//     {
//         user:"iris",
//         pwd:"irispassword",
//         roles:[{role:"root",db:"admin"}]
//     }
// )

// db.createCollection("ex_power_change");
// db.createCollection("ex_uptime_change");
db.createCollection("ex_tx_num_stat");
db.createCollection("ex_config");
db.createCollection("ex_val_black_list");
db.createCollection("ex_validator");
db.createCollection("ex_gov_params");
db.createCollection("ex_asset_tokens");
db.createCollection("ex_asset_gateways");

// create index
// db.ex_power_change.createIndex({"height": 1, "address": 1}, {"unique": true});
db.ex_tx_num_stat.createIndex({"date": -1}, {"unique": true});
db.ex_config.createIndex({"env": 1, "chain_id": 1}, {"unique": true, "background": true});
db.ex_val_black_list.createIndex({"operator_addr": 1}, {"unique": true});
db.ex_validator.createIndex({"operator_address": 1}, {"unique": true});
db.ex_validator.createIndex({"proposer_addr": 1}, {"unique": true,"background": true});
db.ex_validator.createIndex({"jailed": -1, "status": -1, "voting_power": -1}, {"background": true});
db.ex_gov_params.createIndex({"key": 1}, {"unique": true});
db.ex_asset_tokens.createIndex({"token_id": 1}, {"unique": true, "background": true});
db.ex_asset_tokens.createIndex({"source": 1}, {"background": true});
db.ex_asset_gateways.createIndex({"moniker": 1}, {"unique": true, "background": true});

// init data
db.ex_config.insert({
    "network_name": "shinecloudnet",
    "env": "mainnet",
    "host": "http://3.113.34.107:8080",
    "chain_id": "shinecloudnet",
    "show_faucet": 0
});
