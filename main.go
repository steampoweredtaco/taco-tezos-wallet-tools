package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/tyler-smith/go-bip39"
	"os"
)

type CFGFlags struct {
	Input string `short:"k" description:"Input file to decode"`
}

var CFG CFGFlags

func main() {
	p := flags.NewParser(nil, flags.Default)
	_, err := p.AddGroup("Kukai Wallet Options", "Given the file created when a Kukai wallet was created and a known password will print the random seed with these options.", &CFG)
	if err != nil {
		panic(err)
	}
	_, err = p.ParseArgs(os.Args)
	if err != nil {
		return
	}
	f, err := os.Open(CFG.Input)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(f)

	wallet := &ExportedKukaiWallet{}
	err = decoder.Decode(wallet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Processing:\n%s\n", wallet)
	password, err := getPassword()
	if err != nil {
		panic(err)
	}
	biv, err := hex.DecodeString(wallet.Iv)
	if err != nil {
		panic(err)
	}

	seed, err := decrypt_v2orV3(wallet.EncryptedSeed, password, biv)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted seed:", hex.EncodeToString(seed))
	// Overflows are handled exactly how desired

	biv[13] += 1
	entropy, err := decrypt_v2orV3(wallet.EncryptedEntropy, password, biv)
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println("Seed Phrase: ", mnemonic)
}
