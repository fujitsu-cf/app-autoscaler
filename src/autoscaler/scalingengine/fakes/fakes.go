package fakes

//go:generate counterfeiter -o ./fake_cf_client.go ../../cf CfClient
//go:generate counterfeiter -o ./fake_policy_db.go ../../db PolicyDB
//go:generate counterfeiter -o ./fake_history_db.go ../../db HistoryDB