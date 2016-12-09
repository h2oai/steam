package web

import (
	"os"
	"testing"
)

func TestProjectCRUD(t *testing.T) {
	svc, pz, temp := testSetup("project", "sqlite3")
	defer os.RemoveAll(temp)

	projectId, err := svc.CreateProject(pz, "project1", "project1", "binomial")
	if err != nil {
		t.Errorf("creating project: %v", err)
		t.FailNow()
	}
	t.Log("Created project with ID:", projectId)

	project, err := svc.GetProject(pz, projectId)
	if err != nil {
		t.Errorf("getting project: %v", err)
		t.FailNow()
	}
	t.Logf("Returned project with ID: %d and Name: %s", project.Id, project.Name)

	if err := svc.DeleteProject(pz, projectId); err != nil {
		t.Errorf("deleting project: %v", err)
		t.FailNow()
	}
	t.Log("Deleted project with ID:", projectId)
}
