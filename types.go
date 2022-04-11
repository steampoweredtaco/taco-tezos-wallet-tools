package main

import "fmt"

type ExportedKukaiWallet struct {
	Provider         string
	Version          int
	WalletType       int
	EncryptedSeed    string
	EncryptedEntropy string
	Iv               string
}

func (w *ExportedKukaiWallet) String() string {
	return fmt.Sprintf("Provider: %s\nVersion: %d\nWallet Type: %d\nEncrypted Seed: %s\nEncryptedEntropy: %s\nIV: %s",
		w.Provider, w.Version, w.WalletType, w.EncryptedSeed, w.EncryptedEntropy, w.Iv)
}
