/*
Package vlantransparent provides the ability to retrieve and manage networks
with the vlan-transparent extension through the Neutron API.

Example of Listing Networks with the vlan-transparent extension

    iTrue := true
    networkListOpts := networks.ListOpts{}
    listOpts := vlantransparent.ListOptsExt{
        ListOptsBuilder: networkListOpts,
        VLANTransparent: &iTrue,
    }

    type NetworkWithVLANTransparentExt struct {
        networks.Network
        vlantransparent.NetworkVLANTransparentExt
    }

    var allNetworks []NetworkWithVLANTransparentExt

    allPages, err := networks.List(networkClient, listOpts).AllPages()
    if err != nil {
        panic(err)
    }

    err = networks.ExtractNetworksInto(allPages, &allNetworks)
    if err != nil {
        panic(err)
    }

    for _, network := range allNetworks {
        fmt.Println("%+v\n", network)
	}

Example of Getting a Network with the vlan-transparent extension

	var network struct {
		networks.Network
		vlantransparent.TransparentExt
	}

	err := networks.Get(networkClient, "db193ab3-96e3-4cb3-8fc5-05f4296d0324").ExtractInto(&network)
	if err != nil {
		panic(err)
	}

	fmt.Println("%+v\n", network)
*/
package vlantransparent
