/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package api

import (
	"github.com/nats-io/go-nats"
	"log"
)

type (
	NatsMsg struct {
		Channel   string `json:"channel"`
		Publisher string `json:"publisher"`
		Protocol  string `json:"protocol"`
		Payload   []byte `json:"payload"`
	}
)

var (
	NatsConn *nats.Conn
)

func NatsInit(host string, port string) error {
	/** Connect to NATS broker */
	var err error
	NatsConn, err = nats.Connect("nats://" + host + ":" + port)
	if err != nil {
		log.Fatalf("NATS: Can't connect: %v\n", err)
	}

	return err
}
