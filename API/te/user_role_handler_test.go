package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"reflect"
// 	"testing"

// 	"github.com/goEventListing/API/delivery/http/handler"
// 	"github.com/goEventListing/API/entity"
// 	"github.com/goEventListing/API/user/repository"
// 	"github.com/goEventListing/API/user/services"
// 	"github.com/julienschmidt/httprouter"
// )

// func TestRole(t *testing.T) {

// 	userRoleRepo := repository.NewRoleGormRepo(nil)
// 	userRoleService := services.NewRoleServiceImpl(userRoleRepo)
// 	userRoleHandler := handler.NewUserRoleHandler(userRoleService)

// 	mux := httprouter.New()
// 	mux.GET("/el/role/role/:id", userRoleHandler.Role)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	url := ts.URL

// 	resp, err := tc.Get(url + "/el/role/role/1")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var mockRole entity.Role
// 	var Role entity.Role
// 	_ = json.Unmarshal(body, &Role)
// 	mockRole = mockRole
// 	fmt.Println(mockRole)
// 	fmt.Println(Role)
// 	if !reflect.DeepEqual(mockRole, Role) {
// 		// t.Errorf("want body to contain \n%q, but\n%q", mockUsers, users)
// 		t.Errorf("not expected result")
// 	}

// }
