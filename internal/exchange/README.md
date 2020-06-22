# Implementing a new exchange

Follow this instructions to implement a new worker for an exchange.

Always refer to [\mercadobitcoin](mercadobitcoin/worker.go) as a base implementation.

## Interface

A worker must implement 3 methods of the [interface](exchange.go) in order to run from :

- Init
- SyncData
- GetName

### Init

Init method takes care of database connection and initialization and must return a `gorm.DB` pointer.

### SyncData

This method will be called by our main app and must retrieve new trade data from the API.

### Get Name

This method tells the app which worker is running.
