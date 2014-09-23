package servers

import "github.com/rackspace/gophercloud"

func getListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("servers")
}

func getDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

func getServerURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

func getActionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
