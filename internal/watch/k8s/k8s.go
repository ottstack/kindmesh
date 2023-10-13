package k8s

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ottstack/kindmesh/internal/watch/processor"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Watch() {
	go watchCRDs()
	go watchPods()
	go processor.Init()
}

func getRestConfig() *rest.Config {
	configPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		return config
	}
	//Load kubernetes config
	cfg, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
