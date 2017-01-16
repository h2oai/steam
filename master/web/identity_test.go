package web

import (
	"os"
	"testing"

	"github.com/h2oai/steam/master/az"
)

type identityIn struct {
	name string
	pasw string
}

var identityTests = []struct {
	in   identityIn
	out  int64
	pass bool
	err  error
}{
	{in: identityIn{name: "user1", pasw: "password1"}, out: 2, pass: true},
	{in: identityIn{name: "user2", pasw: "password1"}, out: 3, pass: true},
}

var readIdentityTests = []struct {
	offset uint
	limit  uint
}{
	{0, 10},
}

var updateIdentityTests = map[int64][]struct {
	in   identityIn
	pass bool
	err  error
}{
	2: {
		{in: identityIn{pasw: "pass_change"}, pass: true},
	},
}

func TestSQLiteIdentity(t *testing.T) {
	svc, pz, temp := testSetup("identity", "sqlite3")
	defer os.RemoveAll(temp)

	t.Logf("Testing %d case(s)", len(identityTests))
	// -- C --
	if ok := t.Run("Create", testIdentityCreate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- R --
	if ok := t.Run("Read", testIdentityRead(pz, svc)); !ok {
		t.FailNow()
	}

	// -- U --
	if ok := t.Run("Update", testIdentityUpdate(pz, svc)); !ok {
		t.FailNow()
	}

	// -- D --
	// Identities never get deleted
}

func testIdentityCreate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for _, test := range identityTests {
			in, out := test.in, test.out
			id, err := svc.CreateIdentity(pz, in.name, in.pasw)
			if test.pass {
				if err != nil {
					t.Errorf("Create(%+v): unexpected error creating identity: %+v", in, err)
				} else if id != out {
					t.Errorf("Create(%+v): incorrect cluster id: expected %d, got %d", out, out, id)
				}
			} else {
				if err == nil {
					t.Errorf("Create(%+v): expected error creating identity", in)
				} else if err.Error() != test.err.Error() {
					t.Errorf("Create(%+v): incorrect error: expected %q, got %q", in, test.err, err)
				}
			}
		}
	}
}

func testIdentityRead(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		var totPass uint = 1
		for _, test := range identityTests {
			in, out := test.in, test.out
			identity, err := svc.GetIdentity(pz, out)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading identity: %+v", out, err)
				} else if in.name != identity.Name {
					t.Errorf("Read(%+v): incorrect identity name: expected %s, got %s", out, in.name, identity.Name)
				}
				totPass++
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading identity", out)
				}
			}

			identity, err = svc.GetIdentityByName(pz, in.name)
			if test.pass {
				if err != nil {
					t.Errorf("Read(%+v): unexpected error reading identity: %+v", in.name, err)
				} else if out != identity.Id {
					t.Errorf("Read(%+v): incorrect identity id: expected %s, got %s", in.name, out, identity.Id)
				}
			} else {
				if err == nil {
					t.Errorf("Read(%+v): expected error reading identity", in.name)
				}
			}

		}

		for _, get := range readIdentityTests {
			var count int
			if totPass-get.offset < get.limit {
				count = int(totPass - get.offset)
			} else {
				count = int(get.limit)
			}
			identities, err := svc.GetIdentities(pz, get.offset, get.limit)
			if err != nil {
				t.Errorf("Read(%+v): unexpected error reading identities: %+v", get, err)
			} else if len(identities) != count {
				t.Errorf("Read(%+v): incorrect number of identities read: expected %d, got %d", get, count, len(identities))
			} else if len(identities) > 0 && identities[0].Id-1 != int64(get.offset) {
				t.Errorf("Read(%+v): incorrect offset: expected %d, got %d)", get, get.offset, identities[0].Id-1)
			}
		}
	}
}

func testIdentityUpdate(pz az.Principal, svc *Service) func(t *testing.T) {
	return func(t *testing.T) {
		for id, upds := range updateIdentityTests {
			for _, test := range upds {
				in := test.in
				err := svc.UpdateIdentity(pz, id, in.pasw)
				if test.pass {
					if err != nil {
						t.Errorf("Update(%d:%+v): unexpected error updating identity: %+v", id, in, err)
					}
				} else {
					if err == nil {
						t.Errorf("Update(%d:%+v): expected error updating identity: %+v", id, in, err)
					} else if err.Error() != test.err.Error() {
						t.Errorf("Update(%d:%+v): incorrect error: expected %q, got %q", id, in, test.err, err)
					}
				}
			}
		}
	}
}
