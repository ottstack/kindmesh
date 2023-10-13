package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ottstack/kindmesh/internal/proxy/envoy"
	"github.com/ottstack/kindmesh/internal/spec"
	"github.com/ottstack/kindmesh/internal/watch/k8s"
	"github.com/ottstack/kindmesh/internal/watch/netdevice"
	"github.com/ottstack/kindmesh/internal/watch/processor"
	_ "k8s.io/dns/pkg/netif"
)

func Run() {
	// watch crd/pods
	netdevice.EnsureDevice("kindmesh")
	// netdevice.EnsureDevice("bridge0")
	netdevice.AddAddr(spec.DNS_BIND_IP)
	netdevice.AddAddr(spec.ENVOY_CONTROL_IP)

	processor.Emitor = func(dns *spec.DNSRequest, router *spec.RouterRequest) {
		buf, _ := json.Marshal(dns)
		body := bytes.NewBuffer(buf)
		resp, err := http.Post("http://"+spec.DNS_BIND_IP, "application/json", body)
		if err != nil {
			// TODO: retry
			log.Println("set dns error", err)
			return
		} else {
			log.Println("set dns succ", string(buf))
			resp.Body.Close()
		}

		if err := envoy.GenerateSnapshot(router); err != nil {
			log.Println("set envoy router error", err)
			return
		} else {
			log.Println("set envoy router succ")
		}
	}

	k8s.Watch()
	// start envoy control api
	envoy.ServeAPI()
}
