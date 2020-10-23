package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func getProtocol(route netlink.Route) string {
	id := route.Protocol
	proto, ok := ProtocolTable[id]
	if !ok {
		return "none"
	}
	return proto
}

func getScope(route netlink.Route) string {
	id := int(route.Scope)
	scope, ok := ScopeTable[id]
	if !ok {
		return "none"
	}
	return scope
}

func getDestPrefix(route netlink.Route) string {
	dst := route.Dst
	if dst == nil {
		return "0.0.0.0/0"
	}
	return fmt.Sprintf("%s", route.Dst)
}

func getVia(route netlink.Route) string {
	gw := route.Gw
	if gw == nil {
		return "none"
	}
	return fmt.Sprintf("%s", gw)
}

func getPriority(route netlink.Route) string {
	return fmt.Sprintf("%d", route.Priority)
}

func getRoutingTable(route netlink.Route) string {
	return fmt.Sprintf("%d", route.Table)
}

func getSource(route netlink.Route) string {
	src := route.Src
	if src == nil {
		return "none"
	}
	return fmt.Sprintf("%s", route.Src)
}
