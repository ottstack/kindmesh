package state

import (
	"net"
	"sync/atomic"

	"github.com/ottstack/kindmesh/internal/configapi/watchclient"
	"github.com/ottstack/kindmesh/internal/meta"
	"github.com/ottstack/kindmesh/internal/pkg/innerip"
	"github.com/ottstack/kindmesh/internal/spec"
)

var allDomains = atomic.Value{}
var sourceIP2Search = atomic.Value{}
var egressIP = net.ParseIP(spec.EGRESS_IP)

func init() {
	allDomains.Store(&meta.AllDomainList{})
	sourceIP2Search.Store(&meta.SourceIP2Search{})
}

// GetHijackIP returns dns ip by domain and clientIP
func GetHijackIP(domain, clientIP string) net.IP {
	mm := allDomains.Load().(*meta.AllDomainList)
	domains := *mm
	if domains[domain] {
		return egressIP
	}

	m := sourceIP2Search.Load().(*meta.SourceIP2Search)
	searches := (*m)[clientIP]
	for _, s := range searches {
		if domains[domain+s] {
			return egressIP
		}
	}
	return nil
}

func WatchConfig() {
	sourceSearchKey := meta.SourceIP2SearchKeyPrefix + innerip.Get()
	go func() {
		for val := range watchclient.Watch(sourceSearchKey, &meta.SourceIP2Search{}) {
			sourceIP2Search.Store(val)
		}
	}()
	go func() {
		for val := range watchclient.Watch(meta.AllDomainListKey, &meta.AllDomainList{}) {
			allDomains.Store(val)
		}
	}()
}
