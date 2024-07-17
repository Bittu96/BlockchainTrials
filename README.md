# BlockchainTrials

This is a blockchain implementation where you can play around with blockchain, blocks and validations! 

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

Please mail me if you need any help or help me improve this package :)