package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alirezaarzehgar/ticketservice/api/handler"
	"github.com/alirezaarzehgar/ticketservice/model"
)

var (
	MOCK_ORG = map[string]any{
		"name":         "mockorg",
		"address":      "mockaddr",
		"phone_number": "09xxxxxxx",
	}
)

func TestCreateOrganization(t *testing.T) {
	body, _ := json.Marshal(MOCK_ORG)
	req := httptest.NewRequest(http.MethodPost, "/organization/new", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	slog.Debug("test create org", "data", string(body))
	if handler.CreateOrganization(c); rec.Code != http.StatusOK {
		t.Errorf("error on creating org")
	}

	nilBodyTest(t, handler.CreateOrganization, http.MethodPost, "/organization/new")
}

func TestGetAllOrganizations(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/organization/new", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if handler.GetAllOrganizations(c); rec.Code != http.StatusOK {
		t.Errorf("error on fetching orgs")
	}

	res := struct {
		Status bool
		Alert  string
		Data   []model.Organization
	}{}
	json.NewDecoder(rec.Body).Decode(&res)
	if len(res.Data) < 1 {
		slog.Debug("decoded response", "data", res)
		t.Errorf("lest one org registred on db. org len: %v", len(res.Data))
	}
}

func TestEditOrganization(t *testing.T) {
	var org model.Organization
	db.Select("id").Last(&org)

	newName := "updated name"
	MOCK_ORG["name"] = newName
	body, _ := json.Marshal(MOCK_ORG)
	req := httptest.NewRequest(http.MethodPut, "/organization/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(org.ID))

	if handler.EditOrganization(c); rec.Code != http.StatusOK {
		t.Errorf("error on udpate org: %v", rec.Code)
	}

	res := struct {
		Status bool
		Alert  string
		Data   model.Organization
	}{}
	json.NewDecoder(rec.Body).Decode(&res)
	if res.Data.Name != newName {
		t.Errorf("name doesn't changed: %v", org)
	}

	MOCK_ORG["id"] = res.Data.ID
}

func TestAssignAdminToOrganization(t *testing.T) {
	var org model.Organization
	db.Select("id").Last(&org)

	var user model.User
	db.Select("id").Last(&user)

	req := httptest.NewRequest(http.MethodPut, "/organization/hire-admin/1/1", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)

	c := e.NewContext(req, rec)
	c.SetParamNames("org_id", "user_id")
	c.SetParamValues(fmt.Sprint(org.ID), fmt.Sprint(user.ID))

	if handler.AssignAdminToOrganization(c); rec.Code != http.StatusOK {
		t.Errorf("error on hireing an admin: %v", rec.Code)
	}
}

func TestDeleteOrganization(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/organization/1", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+ADMIN_TOKEN)

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(MOCK_ORG["id"]))

	if handler.DeleteOrganization(c); rec.Code != http.StatusOK {
		t.Errorf("error on delete org: %v", rec.Code)
	}

	var org *model.Organization
	db.First(org, MOCK_ORG["id"])
	if org != nil {
		t.Errorf("this test removed org but it's already exists: %v", org)
	}
}
