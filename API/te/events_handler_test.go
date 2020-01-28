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
// 	eventRepo "github.com/goEventListing/API/event/repository"
// 	eventService "github.com/goEventListing/API/event/services"

// 	"github.com/julienschmidt/httprouter"
// )

// func TestEvents(t *testing.T) {

// 	eventRepo := eventRepo.NewEventRepoImp(nil)
// 	eventService := eventService.NewEventServiceImpl(eventRepo)
// 	eventHandler := handler.NewEventHandler(eventService)

// 	mux := httprouter.New()
// 	mux.GET("/el/event/allevents", eventHandler.AllEvents)
// 	ts := httptest.NewTLSServer(mux)
// 	defer ts.Close()

// 	tc := ts.Client()
// 	url := ts.URL

// 	resp, err := tc.Get(url + "/el/event/allevents")
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
// 	var mockEvents []entity.Event
// 	var Events []entity.Event
// 	_ = json.Unmarshal(body, &Events)
// 	mockEvents = append(mockEvents, entity.MockEvent)
// 	fmt.Println(mockEvents)
// 	fmt.Println(Events)
// 	if !reflect.DeepEqual(mockEvents, Events) {
// 		// t.Errorf("want body to contain \n%q, but\n%q", mockUsers, users)
// 		t.Errorf("not expected result")
// 	}

// }
