package service

import (
	"io/ioutil"
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/goEventListing/client/entity"
)


var baseReviewURL = "http://localhost:8181/el/review/"
//	MakeReviewAndRating(r *entity.Review) (*entity.Review, []error)
//	EventReviews(id uint) ([]entity.Review, []error)
//	DeleteReview(id uint) (*entity.Review, []error)
//	UpdateReview(r *entity.Review) (*entity.Review, []error)



//MakeReviewAndRating ...request on "/el/review/make"
func MakeReviewAndRating(review *entity.Review)(*entity.Review,error){
	fmt.Println(review)
	output,err := json.MarshalIndent(review,"","\t\t")
	client :=&http.Client{}
	URL := fmt.Sprintf("%s%s",baseReviewURL,"make")

	req,_ := http.NewRequest("POST",URL,bytes.NewBuffer(output))
	res,err := client.Do(req)

	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return nil, err
	}
	revview := &entity.Review{}
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,revview)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return revview,nil
}
//EventReviews ... request on "/el/review/event/:id"
func EventReviews(id uint) (*[]entity.Review, error){
	client := &http.Client{}
	URL := fmt.Sprintf("%s%s%d",baseReviewURL,"event/",id)
	req,_ := http.NewRequest("GET",URL,nil)

	//DO return an http response
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	reviewData :=&[]entity.Review{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,reviewData)
	if err != nil{
		return nil,err
	}
	return reviewData,nil
}

//DeleteReview ... request on "/el/review/delete/:id"
func DeleteReview(id uint) (*entity.Review, error){
	client := &http.Client{}

	URL := fmt.Sprintf("%s%s/%d",baseReviewURL,"delete",id)
	req,_ := http.NewRequest("POST",URL,nil)

	//DO return an http responce
	res,err := client.Do(req)

	if err != nil {
		return nil, err
	}
	reviewData := &entity.Review{}
	body,err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,reviewData)
	if err != nil{
		return nil,err
	}
	return reviewData,nil
}
//UpdateReview ... request on "/el/review/edit"
func UpdateReview(revieww *entity.Review) (*entity.Review, error){
	ouput,err:= json.MarshalIndent(revieww,"","\t\t")

	client := &http.Client{}
	URL := fmt.Sprintf("%s%s/%d",baseReviewURL,"edit",revieww.ID)
	req,_ := http.NewRequest("PUT",URL,bytes.NewBuffer(ouput))

	//DO return an http response
	res,err := client.Do(req)
	
	if err != nil {
		return nil,err
	}
	review := &entity.Review{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body,review)
	if err != nil{
		return nil,err
	}
	return review,nil

}
