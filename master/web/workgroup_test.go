package web

import (
	"testing"
)

func TestWorkgroupCRUD(tt *testing.T) {
	t := newTest(tt)

	const name1 = "group1"
	const description1 = "description1"

	id, err := t.svc.CreateWorkgroup(t.su, name1, description1)
	t.nil(err)

	group, err := t.svc.GetWorkgroup(t.su, id)
	t.nil(err)

	t.ok(name1 == group.Name, "name")
	t.ok(description1 == group.Description, "description")

	group, err = t.svc.GetWorkgroupByName(t.su, name1)
	t.nil(err)
	t.ok(name1 == group.Name, "name")
	t.ok(description1 == group.Description, "description")

	groups, err := t.svc.GetWorkgroups(t.su, 0, 1000)
	t.nil(err)

	t.ok(len(groups) == 1, "group count")

	group = groups[0]
	t.ok(name1 == group.Name, "name")
	t.ok(description1 == group.Description, "description")

	const name2 = "group2"
	const description2 = "description2"

	err = t.svc.UpdateWorkgroup(t.su, id, name2, description2)
	t.nil(err)

	group, err = t.svc.GetWorkgroup(t.su, id)
	t.nil(err)

	t.ok(name2 == group.Name, "name")
	t.ok(description2 == group.Description, "description")

	group, err = t.svc.GetWorkgroupByName(t.su, name2)
	t.nil(err)
	t.ok(name2 == group.Name, "name")
	t.ok(description2 == group.Description, "description")

	groups, err = t.svc.GetWorkgroups(t.su, 0, 2000)
	t.nil(err)

	t.ok(len(groups) == 1, "group count")

	group = groups[0]
	t.ok(name2 == group.Name, "name")
	t.ok(description2 == group.Description, "description")

	err = t.svc.DeleteWorkgroup(t.su, id)
	t.nil(err)

}

func TestWorkgroupDeletion(tt *testing.T) {
	t := newTest(tt)

	const name1 = "group1"
	const description1 = "description1"

	id, err := t.svc.CreateWorkgroup(t.su, name1, description1)
	t.nil(err)

	err = t.svc.DeleteWorkgroup(t.su, id)
	t.nil(err)

	_, err = t.svc.GetWorkgroup(t.su, id)
	t.notnil(err)

	_, err = t.svc.GetWorkgroupByName(t.su, name1)
	t.notnil(err)

	groups, err := t.svc.GetWorkgroups(t.su, 0, 1000)
	t.nil(err)
	t.ok(len(groups) == 0, "group count")

	// err = t.svc.DeleteWorkgroup(t.su, id) // should fail on a duplicate attempt
	// t.notnil(err)
}
