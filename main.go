package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/99designs/keyring"
	"github.com/alecthomas/kingpin"
)

var (
	version string

	// vault   = kingpin.Arg("vault", "Vault to use").Required().String()
	command = kingpin.Arg("command", "Command to run").Required().String()

	keyringDefaults = keyring.Config{
		ServiceName:              "goenvchain",
		KeychainName:             "goenvchain",
		KeychainTrustApplication: true,
	}
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	kingpin.Parse()

	keychain, err := keyring.Open(keyringDefaults)
	if err != nil {
		log.Fatal(err)
	}

	i := keyring.Item{
		Key:  "KEY1",
		Data: []byte("VAL1"),
	}
	i.Label = fmt.Sprintf("%s (%s)", "goenvchain", i.Key)

	keychain.Set(i)

	cmd := exec.Command(*command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	keys, _ := keychain.Keys()
	fmt.Println(keys)
	for _, k := range keys {
		v, _ := keychain.Get(k)
		fmt.Println(k, string(v.Data))
	}

	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
