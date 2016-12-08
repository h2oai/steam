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
		t.Errorf("registering cluster: %+v", err)
		t.FailNow()
	}
	t.Log("Registered cluster with ID:", clusterId)

	if err = svc.UnregisterCluster(pz, clusterId); err != nil {
		t.Errorf("unregistering cluster: %+v", err)
		t.FailNow()
	}
	t.Log("Unregistreed clsuter with ID:", clusterId)
}
