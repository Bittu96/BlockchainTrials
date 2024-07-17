# BlockchainTrials

A blockchain implementation.

## Install
Fist, use go get to install the latest version of the library:
```
go get -u github.com/Bittu96/BlockchainTrials
```

Next, include this package in your application:
```
import "github.com/Bittu96/BlockchainTrials"
```

## Example
```
bc := blockchain.New()
bc.AddNewBlock(transaction.GenerateRecord()) // add your own transaction record
bc.Validate()
```