package hijack

import (
	"context"
	"strings"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"github.com/ottstack/kindmesh/internal/dns/state"
)

const name = "hijack"

type Hijacking struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface to hijack dns record.
func (it Hijacking) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	req := request.Request{W: w, Req: r}

	clientIP := strings.SplitN(req.RemoteAddr(), ":", 2)[0]
	ip := state.GetHijackIP(req.Name(), clientIP)
	if ip == nil {
		return plugin.NextOrFailure(it.Name(), it.Next, ctx, w, r)
	}
	// Returns empty result for IPV6
	answer := []dns.RR{}
	if req.Type() == "A" {
		answer = []dns.RR{
			&dns.A{
				Hdr: dns.RR_Header{
					Name:   req.Name(),
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    0,
				},
				A: ip,
			},
		}
	}

	dnsMsg := dns.Msg{
		MsgHdr: dns.MsgHdr{
			Authoritative: true,
		},
		Answer: answer,
	}
	dnsMsg.SetReply(r)
	err := w.WriteMsg(&dnsMsg)
	return dns.RcodeSuccess, err
}

// Name implements the Handler interface.
func (it Hijacking) Name() string { return name }
