# Taco Export
## Description
An application written in go to export Kukai version 3 or HD Wallets encrypted data stores to seed phrase.

# How to build
```bash
>go build
```
# Usage
```bash
Kukai Wallet Options:
  -k:          Input file to decode

Help Options:
  -?           Show this help message
  -h, --help   Show this help message
```
# How to run
```bash
>./taco-export -k hd-wallet.tez
## enter password for wallet and then you'll get the unencrypted seed and seed phrase
```

# Example Run:
```powershell
PS C:\Users\beebl\GolandProjects\taco-exports> .\taco-exports.exe -h
Usage:
  C:\Users\beebl\GolandProjects\taco-exports\taco-exports.exe [OPTIONS]

Kukai Wallet Options:
  /k:         Input file to decode

Help Options:
  /?          Show this help message
  /h, /help   Show this help message

panic: Usage:
  C:\Users\beebl\GolandProjects\taco-exports\taco-exports.exe [OPTIONS]

Kukai Wallet Options:
  /k:         Input file to decode
Version: 3
Wallet Type: 4
Encrypted Seed: 8535aa09ae8da755b36a8bbf762f52a103f2122d8d0419a4b88f84ba86630e6d345353b943d888a2380cebf9e4d01488785890f1c40af235bf34a0230916443e==268df969c4a63261e3c1e1d759bbe06d
EncryptedEntropy: 34ec8e04bb9860f9f0b176fe8b5b157fac03f50ca6114a090995b1283c8bdb0f==d93c0583b7dbb37055765793509e653b
IV: 05e84d147c823ed68f23ada6e6c13dbf
Enter Password:
Decrypted seed: 47c8545034b796d51da60f43379e56935678c2347bc15193e9246acae5c87e40b5a721d26d9fa2eed9db9cd5ce7f2b67b98ca4757d7dab84f6b523a1441836ed
Seed Phrase:  stem fortune bonus ethics scout stomach unaware copy earth uncle unknown drip pelican bullet stick cluster ranch argue canoe chase cause tool syrup circle
PS C:\Users\beebl\GolandProjects\taco-exports>

```

![Demo](static/keyexportdemo.gif?raw=true "Demo")

## Warning
Obviously don't use this revealed private key, ever for anything except maybe demoing features.

You should build this yourself, but a copy has been provided in releases for window users. This is a safe program but always do your own research before running any strange executables from a random stranger off the interent.

## Dependencies
	github.com/jessevdk/go-flags v1.5.0
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
    golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
    golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1
