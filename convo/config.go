// Copyright 2017 David Lazar. All rights reserved.
// Use of this source code is governed by the GNU AGPL
// license that can be found in the LICENSE file.

package convo

import (
	"golang.org/x/crypto/ed25519"

	"vuvuzela.io/alpenhorn/config"
	"vuvuzela.io/alpenhorn/errors"
	"vuvuzela.io/vuvuzela/mixnet"
)

// Use github.com/davidlazar/easyjson:
//go:generate easyjson .

func init() {
	config.RegisterService("Convo", &ConvoConfig{})
}

type ConvoConfig struct {
	Coordinator CoordinatorConfig
	MixServers  []mixnet.PublicServerConfig
}

//easyjson:readable
type CoordinatorConfig struct {
	Key     ed25519.PublicKey
	Address string
}

func (c *ConvoConfig) Validate() error {
	if len(c.Coordinator.Key) != ed25519.PublicKeySize {
		return errors.New("invalid coordinator key: %#v", c.Coordinator.Key)
	}
	if c.Coordinator.Address == "" {
		return errors.New("empty coordinator address")
	}

	if len(c.MixServers) == 0 {
		return errors.New("no mix servers defined for convo protocol")
	}

	for i, mix := range c.MixServers {
		if len(mix.Key) != ed25519.PublicKeySize {
			return errors.New("invalid key for mixer %d: %s", i, mix.Key)
		}
		if mix.Address == "" {
			return errors.New("empty address for mix server %d", i)
		}
	}

	return nil
}
