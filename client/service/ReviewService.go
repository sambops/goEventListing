package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/goEventListing/client/entity"
)

// var baseURL = "http://localhost:8081/el/"

//Reviews ... returns all Reviews  /Reviews/
func Reviews() ([]entity.Review, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s", baseURL, "reviews")
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	ReviewData := []entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ReviewData)
	if err != nil {
		return nil, err
	}
	return ReviewData, nil
}

//Review ... returns all Reviews  /Reviews/
func Review(id uint) (*entity.Review, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "review", "single", id)
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	ReviewData := entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ReviewData)
	if err != nil {
		return nil, err
	}
	return &ReviewData, nil
}

//GetMyReviews ... returns all Reviews  /Reviews/
func GetMyReviews(id uint) ([]entity.Review, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "user", "review", id)
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	ReviewData := []entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ReviewData)
	if err != nil {
		return nil, err
	}
	return ReviewData, nil
}

//EventReviews ... returns all Reviews  /Reviews/
func EventReviews(id uint) ([]entity.Review, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "event", "reviews", id)
	req, _ := http.NewRequest("GET", URL, nil)
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	ReviewData := []entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ReviewData)
	if err != nil {
		return nil, err
	}
	return ReviewData, nil
}

// DeleteReview ..
func DeleteReview(id uint) (*entity.Review, error) {
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%s/%d", baseURL, "review", "delete", id)
	req, _ := http.NewRequest("Delete", URL, nil)

	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	ReviewData := &entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, ReviewData)
	if err != nil {
		return nil, err
	}
	return ReviewData, nil

}

//MakeReview ... request on baseUserURL/register
func MakeReview(review *entity.Review) (*entity.Review, error) {
	ouput, err := json.MarshalIndent(review, "", "\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%s", baseURL, "review", "make")

	//we use bytes.NewBuffer which gives us a bytes buffer based on our bytes slice.
	// This buffer is both readable and writable.
	// It’s “readable” part satisfies the io.Reader interface and serves our purpose.
	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(ouput))
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	Review := &entity.Review{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, Review)
	if err != nil {
		return nil, err
	}
	return Review, nil
}

//EditReview ...
func EditReview(review *entity.Review) (*entity.Review, error) {
	ouput, err := json.MarshalIndent(review, "", "\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%s/%d", baseUserURL, "review", "edit", review.ID)
	req, _ := http.NewRequest("PUT", URL, bytes.NewBuffer(ouput))
	//DO return an http responce
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	rev := &entity.Review{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, rev)
	if err != nil {
		return nil, err
	}
	return rev, nil
}
