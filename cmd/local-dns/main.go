package main

import (
	// blank imports to make sure the plugin code is pulled in from vendor when building node-cache image

	"log"
	"os"

	_ "github.com/coredns/coredns/plugin/bind"
	_ "github.com/coredns/coredns/plugin/bufsize"
	_ "github.com/coredns/coredns/plugin/cache"
	_ "github.com/coredns/coredns/plugin/debug"
	_ "github.com/coredns/coredns/plugin/dns64"
	_ "github.com/coredns/coredns/plugin/errors"
	_ "github.com/coredns/coredns/plugin/forward"
	_ "github.com/coredns/coredns/plugin/health"
	_ "github.com/coredns/coredns/plugin/hosts"
	_ "github.com/coredns/coredns/plugin/loadbalance"
	_ "github.com/coredns/coredns/plugin/log"
	_ "github.com/coredns/coredns/plugin/loop"
	_ "github.com/coredns/coredns/plugin/metrics"
	_ "github.com/coredns/coredns/plugin/pprof"
	_ "github.com/coredns/coredns/plugin/reload"
	_ "github.com/coredns/coredns/plugin/rewrite"
	_ "github.com/coredns/coredns/plugin/template"
	_ "github.com/coredns/coredns/plugin/trace"
	_ "github.com/coredns/coredns/plugin/whoami"
	"github.com/ottstack/kindmesh/internal/configapi/watchclient"
	_ "github.com/ottstack/kindmesh/internal/dns/hijack"
	"github.com/ottstack/kindmesh/internal/dns/state"
	"github.com/ottstack/kindmesh/internal/pkg/netdevice"
	"github.com/ottstack/kindmesh/internal/spec"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

func main() {
	netdevice.EnsureDevice("bridge0")
	netdevice.AddAddr(spec.DNS_BIND_IP)

	os.Setenv("DNS_BIND_IP", spec.DNS_BIND_IP)

	watchclient.InitWatcher()
	state.WatchConfig()

	log.Println("Serving dns on", spec.DNS_BIND_IP+":53")

	dnsserver.Directives = append([]string{"hijack"}, dnsserver.Directives...)
	coremain.Run()
}
