package main

// Protocol
// from https://github.com/shemminger/iproute2/blob/e8e8f16ed155dfbe026ad3c22458d1277e17794e/lib/rt_names.c#L122-L145
const (
	RTPROT_UNSPEC     = 0
	RTPROT_REDIRECT   = 1
	RTPROT_KERNEL     = 2
	RTPROT_BOOT       = 3
	RTPROT_STATIC     = 4
	RTPROT_GATED      = 8
	RTPROT_RA         = 9
	RTPROT_MRT        = 10
	RTPROT_ZEBRA      = 11
	RTPROT_BIRD       = 12
	RTPROT_DNROUTED   = 13
	RTPROT_XORP       = 14
	RTPROT_NTK        = 15
	RTPROT_DHCP       = 16
	RTPROT_MROUTED    = 17
	RTPROT_KEEPALIVED = 18
	RTPROT_BABEL      = 42
	RTPROT_BGP        = 186
	RTPROT_ISIS       = 187
	RTPROT_OSPF       = 188
	RTPROT_RIP        = 189
	RTPROT_EIGRP      = 192
)

var ProtocolTable = map[int]string{
	RTPROT_UNSPEC:     "unspec",
	RTPROT_REDIRECT:   "redirect",
	RTPROT_KERNEL:     "kernel",
	RTPROT_BOOT:       "boot",
	RTPROT_STATIC:     "static",
	RTPROT_GATED:      "gated",
	RTPROT_RA:         "ra",
	RTPROT_MRT:        "mrt",
	RTPROT_ZEBRA:      "zebra",
	RTPROT_BIRD:       "bird",
	RTPROT_DNROUTED:   "dnrouted",
	RTPROT_XORP:       "xorp",
	RTPROT_NTK:        "ntk",
	RTPROT_DHCP:       "dhcp",
	RTPROT_MROUTED:    "mrouted",
	RTPROT_KEEPALIVED: "keepalived",
	RTPROT_BABEL:      "babel",
	RTPROT_BGP:        "bgp",
	RTPROT_ISIS:       "isis",
	RTPROT_OSPF:       "ospf",
	RTPROT_RIP:        "rip",
	RTPROT_EIGRP:      "eigrp",
}

// Scope
// from https://github.com/shemminger/iproute2/blob/e8e8f16ed155dfbe026ad3c22458d1277e17794e/lib/rt_names.c#L233-L239
const (
	RT_SCOPE_UNIVERSE = 0
	RT_SCOPE_SITE     = 200
	RT_SCOPE_LINK     = 253
	RT_SCOPE_HOST     = 254
	RT_SCOPE_NOWHERE  = 255
)

var ScopeTable = map[int]string{
	RT_SCOPE_UNIVERSE: "global",
	RT_SCOPE_SITE:     "nowhere",
	RT_SCOPE_LINK:     "link",
	RT_SCOPE_HOST:     "host",
	RT_SCOPE_NOWHERE:  "site",
}
