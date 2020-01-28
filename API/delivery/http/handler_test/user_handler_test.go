package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/goEventListing/API/delivery/http/handler"
	"github.com/goEventListing/API/entity"
	userRepo "github.com/goEventListing/API/user/repository"
	userServ "github.com/goEventListing/API/user/services"
	"github.com/julienschmidt/httprouter"
)

func TestUserss(t *testing.T) {

	userRepo := userRepo.NewMockUserGormRepo(nil)
	userServ := userServ.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userServ)

	mux := httprouter.New()
	mux.GET("/el/user", userHandler.GetUsers)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/el/user")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	var mockUsers []entity.User
	var users []entity.User
	_ = json.Unmarshal(body, &users)
	mockUsers = append(mockUsers, entity.MockUser)
	fmt.Println(mockUsers)
	fmt.Println(users)
	if !reflect.DeepEqual(mockUsers, users) {
		// t.Errorf("want body to contain \n%q, but\n%q", mockUsers, users)
		t.Errorf("not expected result")
	}

}
