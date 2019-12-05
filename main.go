package main

import (
	"log"
	"os"

	"github.com/Microsoft/hcsshim/hcn"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	networkName := os.Args[1]

	existingNetwork, err := hcn.GetNetworkByName(networkName)
	if err != nil {
		log.Fatalf("Unable to find HNS Network with name %q", networkName)
	}

	for _, policy := range existingNetwork.Policies {
		if policy.Type == hcn.HostRoute {
			log.Fatalf("Network %v already has a HostRoute policy, aborting", networkName)
		}
	}

	if err := existingNetwork.AddPolicy(hcn.PolicyNetworkRequest{
		Policies: []hcn.NetworkPolicy{
			{
				Type:     hcn.HostRoute,
				Settings: []byte("{}"),
			},
		},
	}); err != nil {
		log.Fatalf("Could not apply HostRoute policy for local host to local pod connectivity. This policy requires Windows 18321.1000.19h1_release.190117-1502 or newer")
	}

	log.Printf("Successfully applied HostRoute policy for network %v. Enjoy your newfound ability to route to containers from the host!!!", networkName)
}
