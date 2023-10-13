package state

import (
	"testing"

	"github.com/ottstack/kindmesh/internal/meta"
	"github.com/ottstack/kindmesh/internal/spec"
	"github.com/stretchr/testify/assert"
)

func TestHijack(t *testing.T) {

	allDomains.Store(&meta.AllDomainList{
		"abc.default.svc.cluster.local.": true,
	})
	sourceIP2Search.Store(&meta.SourceIP2Search{
		"172.1.2.3": []string{"default.svc.cluster.local.", "svc.cluster.local."},
	})

	clusterDomain := ".svc.cluster.local"
	pod := "172.1.2.3"
	ns := "default"
	domain := "abc."

	ip := GetHijackIP(domain, pod)
	assert.Equal(t, spec.EGRESS_IP, ip.String())

	ip = GetHijackIP(domain+ns+".", pod)
	assert.Equal(t, spec.EGRESS_IP, ip.String())

	ip = GetHijackIP(domain+ns+clusterDomain+".", pod)
	assert.Equal(t, spec.EGRESS_IP, ip.String())

	ip = GetHijackIP(domain, "127.0.0.1")
	assert.Nil(t, ip)

	ip = GetHijackIP(domain+".abc.", pod)
	assert.Nil(t, ip)
}
