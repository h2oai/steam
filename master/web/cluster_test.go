package web

import (
	"fmt"
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

var clusterTests = []struct {
	in   string
	out  int64
	pass bool
	err  error
}{
	{in: "localhost:54321", out: 1, pass: true},
}

var readClusterTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
	{1, 5},
}

func TestSQLiteExternalCluster(t *testing.T) {
	svc, pz, temp := testSetup("cluster", "sqlite3")
	defer os.RemoveAll(temp)

	if !test_h2o {
		t.Skip("skipping cluster tests: requires h2o")
	} else if !pingExternal(cluster_url) {
		t.Fatal("unable to reach h2o")
	}

	t.Logf("Testing %d case(s)", len(clusterTests))
	// -- C --
	if ok := t.Run("Create", testExternalClusterCreate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testClusterRead(pz, svc)); !ok {
		t.FailNow()
	}
	// -- U --

	// -- D --
	if ok := t.Run("Delete", testExternalClusterDelete(pz, svc)); !ok {
		t.FailNow()
	}

}

func testExternalClusterCreate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range clusterTests {
			in, out := test.in, test.out
			id, err := svc.RegisterCluster(pz, in)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error registering cluster: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect cluster id: expected %d, got %d", in, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error registering cluster", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testYarnClusterCreate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range clusterTests {
			fmt.Println(test)
			// in, out := test.in, test.out
			// id, err := svc.StartClusterOnYarn(pz, clusterName, engineId, size, memory, keytab)
		}
	}
}

func testClusterRead(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		var totPass uint
		for _, test := range clusterTests {
			in, out := test.in, test.out
			cluster, err := svc.GetCluster(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading cluster: %+v", out, err)
				} else if cluster.Address != in {
					t.Errorf("Read(%+v): incorrect address: expected %s, got %s", out, in, cluster.Address)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading cluster", out)
				}
			}
		}

		for _, get := range readClusterTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			clusters, err := svc.GetClusters(pz, get.offset, get.limit)
			if err != nil {
				t.Errorf("Read(%+v): unexpected error reading clusters: %+v", get, err)
			} else if len(clusters) != count {
				t.Errorf("Read(%+v): incorrect number of clusters read: expected %d, got %d", get, count, len(clusters))
			} else if len(clusters) > 0 && clusters[0].Id-1 != int64(get.offset) {
				t.Errorf("Read(%+v): incorrect offset: expected %d, got %d)", get, get.offset, clusters[0].Id-1)
			}
		}
	}
}

func testClusterUpdate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
	}
}

func testExternalClusterDelete(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range clusterTests {
			out := test.out
			err := svc.UnregisterCluster(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Delete(%+v): unexpected error deleting cluster: %+v", out, err)
				}
			} else {
				if err == nil {
					t.Errorf("Delete(%+v): expected error deleting cluster", out)
				}
			}

			clusters, _ := svc.GetClusters(pz, 0, 1)
			if len(clusters) > 1 {
				t.Errorf("Delete: at least one cluster was not deleted")
			}
		}
	}
}
