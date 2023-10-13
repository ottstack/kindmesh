package main

import (
	"github.com/ottstack/kindmesh/internal/configapi/watchclient"
	"github.com/ottstack/kindmesh/internal/mechaproxy/callback"
	"github.com/ottstack/kindmesh/internal/mechaproxy/egress"
	"github.com/ottstack/kindmesh/internal/mechaproxy/ingress"
)

func main() {
	watchclient.InitWatcher()

	// ingress: forword request to local container or send to queue (then wait for callback)
	// consumer: consume request and send it to local container, callback to producer if need
	// callback: find origin request and send response
	go ingress.WatchConfig()
	go ingress.Serve()
	go callback.Serve()

	// egress: find target namespace and deployment, forward request to one of downstream
	go egress.WatchDomain()
	go egress.WatchConfig()
	egress.Serve()
}
