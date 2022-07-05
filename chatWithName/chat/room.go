package main

import (
	"net"
)

type room struct {
	name    string
	members map[net.Addr]*chat.client
}

func (r *room) broadcast(sender *chat.client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
