package api

import (
	"os"
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

const testProxyLBName = "test_libsakuracloud_proxylb"

func TestProxyLBCreate(t *testing.T) {
	defer initProxyLB()()

	serverIP := os.Getenv("LIBSACLOUD_PROXYLB_SERVER")
	if serverIP == "" {
		t.Skipf("environment variable LIBSACLOUD_PROXYLB_SERVER is empty. skip")
	}
	sorryIP := os.Getenv("LIBSACLOUD_PROXYLB_SORRY_SERVER")
	if serverIP == "" {
		t.Skipf("environment variable LIBSACLOUD_PROXYLB_SORRY_SERVER is empty. skip")
	}

	currentRegion := client.Zone
	defer func() { client.Zone = currentRegion }()
	client.Zone = "is1a"

	item := client.ProxyLB.New(testProxyLBName)
	assert.Equal(t, item.Name, testProxyLBName)

	item.SetPlan(sacloud.ProxyLBPlan1000)
	item.SetSorryServer(sorryIP, 80)
	item.SetHTTPHealthCheck("libsacloud.com", "/", 0)
	item.AddBindPort("http", 80, false, false)
	item.AddServer(serverIP, 80, true)

	item, err := client.ProxyLB.Create(item)

	assert.NoError(t, err)

	assert.Equal(t, item.GetPlan(), sacloud.ProxyLBPlan1000)

	assert.Equal(t, item.Settings.ProxyLB.SorryServer.IPAddress, sorryIP)
	assert.Equal(t, *item.Settings.ProxyLB.SorryServer.Port, 80)

	assert.Equal(t, item.Settings.ProxyLB.HealthCheck.Protocol, "http")
	assert.Equal(t, item.Settings.ProxyLB.HealthCheck.Host, "libsacloud.com")
	assert.Equal(t, item.Settings.ProxyLB.HealthCheck.Path, "/")
	assert.Equal(t, item.Settings.ProxyLB.HealthCheck.DelayLoop, 10)

	assert.Len(t, item.Settings.ProxyLB.BindPorts, 1)
	assert.Len(t, item.Settings.ProxyLB.Servers, 1)

	assert.Equal(t, item.Settings.ProxyLB.BindPorts[0].ProxyMode, "http")
	assert.Equal(t, item.Settings.ProxyLB.BindPorts[0].Port, 80)

	assert.Equal(t, item.Settings.ProxyLB.Servers[0].IPAddress, serverIP)
	assert.Equal(t, item.Settings.ProxyLB.Servers[0].Port, 80)
	assert.Equal(t, item.Settings.ProxyLB.Servers[0].Enabled, true)

	newPlan, err := client.ProxyLB.ChangePlan(item.ID, sacloud.ProxyLBPlan5000)

	assert.NoError(t, err)
	assert.NotEqual(t, newPlan.ID, item.ID)

	client.ProxyLB.Delete(newPlan.ID)
}

func initProxyLB() func() {
	cleanupProxyLB()
	return cleanupProxyLB
}

func cleanupProxyLB() {
	items, _ := client.ProxyLB.Reset().WithNameLike(testProxyLBName).Find()

	for _, item := range items.CommonServiceProxyLBItems {
		client.ProxyLB.Delete(item.ID)
	}
}
