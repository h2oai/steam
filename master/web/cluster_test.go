package web

import (
	"os"
	"testing"
)

func TestExternalClusterCRUD(t *testing.T) {
	svc, pz, temp := testSetup("cluster", "sqlite3")
	defer os.RemoveAll(temp)

	clusterId, err := svc.RegisterCluster(pz, "localhost:54321")
	if err != nil {
		t.Errorf("registering cluster: %s", err)
		t.FailNow()
	}
	t.Log("Registered cluster with ID:", clusterId)

	cluster, err := svc.GetCluster(pz, clusterId)
	if err != nil {
		t.Errorf("getting cluster: %s", err)
		t.FailNow()
	}
	t.Logf("Returned cluster with ID: %d and Name: %s", cluster.Id, cluster.Name)

	clusterStat, err := svc.GetClusterStatus(pz, clusterId)
	if err != nil {
		t.Errorf("getting cluster status: %s", err)
		t.FailNow()
	}
	t.Logf("Returned cluster stats: %+v", clusterStat)

	if err = svc.UnregisterCluster(pz, clusterId); err != nil {
		t.Errorf("unregistering cluster: %s", err)
		t.FailNow()
	}
	t.Log("Unregistreed cluster with ID:", clusterId)
}
