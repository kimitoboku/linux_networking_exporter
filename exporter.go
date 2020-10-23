package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vishvananda/netlink"
)

const (
	namespace = "network"
)

var (
	routeLables = []string{
		"dest",
		"src",
		"gateway",
		"device",
		"metrics",
		"protocol",
		"scope",
		"table",
	}
	routeGaugeVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "route",
			Help:      "routing entries",
		},
		routeLables,
	)
)

type routeCollector struct{}

func init() {
	var c routeCollector
	prometheus.MustRegister(c)
}

func (c routeCollector) Collect(ch chan<- prometheus.Metric) {
	links, err := netlink.LinkList()
	if err != nil {
		return
	}

	for _, link := range links {
		routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
		if err != nil {
			return
		}

		for _, route := range routes {
			dev := link.Attrs().Name
			proto := getProtocol(route)
			scope := getScope(route)
			dst := getDestPrefix(route)
			via := getVia(route)
			metric := getPriority(route)
			table := getRoutingTable(route)
			src := getSource(route)
			routeGaugeVec.WithLabelValues(
				dst,
				src,
				via,
				dev,
				metric,
				proto,
				scope,
				table,
			).Set(1)
		}
	}

	routeGaugeVec.Collect(ch)

}

func (c routeCollector) Describe(ch chan<- *prometheus.Desc) {
	routeGaugeVec.Describe(ch)
}
