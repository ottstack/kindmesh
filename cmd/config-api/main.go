package main

import (
	_ "github.com/ottstack/kindmesh/internal/configapi/handler"
	"github.com/ottstack/kindmesh/internal/configapi/watcher"
	"github.com/ottstack/kindmesh/internal/meta"

	"github.com/ottstack/gofunc"
	"github.com/ottstack/gofunc/pkg/middleware"
)

func main() {
	// mock
	// ip := "192.168.2.124"
	// watcher.SetValue(meta.IngressKeyPrefix+ip, &meta.IngressConfig{
	// 	HostInfo: map[string]*meta.IngressHostInfo{
	// 		"127.0.0.1:8080": {
	// 			Addr:             "127.0.0.1:8080",
	// 			ConcurrencyLimit: 2,
	// 			QueueSource:      "local://",
	// 		},
	// 	},
	// })

	// watcher.SetValue(meta.EgressConfigKeyPrefix+ip, &meta.EgressConfig{
	// 	// DomainList:    map[string]bool{"abc": true},
	// 	SourceNamespace: map[string]string{"127.0.0.1": "default"},
	// })

	// watcher.SetValue(meta.DomainConfigKeyPrefix+"abc.default", &meta.DomainConfig{
	// 	Domain: "abc.default",
	// 	IsZero: false,
	// 	DownStreams: map[int]*meta.DownStreamInfo{
	// 		0: {
	// 			Addr:        "127.0.0.1:8080",
	// 			IngressAddr: "127.0.0.1:17000",
	// 		},
	// 	},
	// })

	watcher.SetValue(meta.AllDomainListKey, &meta.AllDomainList{
		"abc.default.svc.cluster.local.": true,
	})

	watcher.SetValue(meta.SourceIP2SearchKeyPrefix+"192.168.2.124", &meta.SourceIP2Search{
		"169.254.99.1": []string{"default.svc.cluster.local.", "svc.cluster.local."},
	})

	// processor.Emitor = func(dns *spec.DNSRequest, router *spec.RouterRequest) {
	// 	// TODO
	// }

	// k8s.Watch()

	gofunc.Use(middleware.Recover).Use(middleware.Validator)
	gofunc.Serve()
}
