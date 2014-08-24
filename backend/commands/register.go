// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	"code.google.com/p/log4go"
	"github.com/limetext/lime/backend"
)

type cmd struct {
	name string
	cmd  backend.Command
}

func register(cmds []cmd) {
	ch := backend.GetEditor().CommandHandler()
	for _, cmd := range cmds {
		if err := ch.Register(cmd.name, cmd.cmd); err != nil {
			log4go.Error("Failed to register command %s: %s", cmd.name, err)
		}
	}
}
