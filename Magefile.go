// +build mage

package main

import (
	"github.com/cardil/knative-homedir/internal"
	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
)

// Default target is set to binary
var Default = magetasks.Binary

func init() {
	config.Binaries = append(config.Binaries, config.Binary{
		Name: internal.BinaryName,
	})
	config.VersionVariablePath = "github.com/cardil/knative-homedir/internal.Version"
}
