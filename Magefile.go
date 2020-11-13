// +build mage

package main

import (
	"log"

	"github.com/cardil/homedir/internal"
	"github.com/joho/godotenv"

	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"

	// mage:import
	_ "github.com/wavesoftware/go-magetasks/container"
)

// Default target is set to binary
var Default = magetasks.Binary

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	config.Binaries = append(config.Binaries, config.Binary{
		Name: internal.BinaryName,
	})
	config.VersionVariablePath = "github.com/cardil/homedir/internal.Version"
}
