package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/onelazyugy/golang-microservice/services"
	"github.com/onelazyugy/golang-microservice/types"
	"github.com/onelazyugy/golang-microservice/util"
)

// OrderBubbleTeaHandler allows you to order bubble tea drink
func OrderBubbleTeaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bubbleTeaResponse := &types.BubbleTeaResponse{}
	bubbleTeaRequest, unMarshallError := unMarshallToBubbleTeaRequestType(r)
	if unMarshallError != nil {
		error := fmt.Errorf("Error: %v", unMarshallError)
		fmt.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}
	bubbleTeaResponse, serviceError := services.OrderBubbleTea(bubbleTeaRequest)
	if serviceError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(serviceError.Error()))
		return
	}
	createBubbleTeaResponse(bubbleTeaResponse, bubbleTeaResponse.Message, bubbleTeaResponse.Success, http.StatusOK)
	bubbleTeaResponseBytes, marshallBubbleTeaResponseError := json.Marshal(bubbleTeaResponse)
	if marshallBubbleTeaResponseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(marshallBubbleTeaResponseError.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bubbleTeaResponseBytes)
}

// RetrieveOrderedBubbleTeaHandler allows you to retreive all ordered bubble teas
func RetrieveOrderedBubbleTeaHandler(w http.ResponseWriter, r *http.Request) { //(types.RetrieveBubbleTeaResponse, error)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res, err := services.RetrieveOrderedBubbleTea()
	if err != nil {
		//handle error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(res.BubbleTeaResponse.StatusCd)
	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func createBubbleTeaResponse(bubbleTeaResponse *types.BubbleTeaResponse, message string, success bool, statusCd int) {
	bubbleTeaResponse.Message = message
	bubbleTeaResponse.Success = success
	bubbleTeaResponse.StatusCd = statusCd
}

func unMarshallToBubbleTeaRequestType(r *http.Request) (types.BubbleTeaRequest, error) {
	var bubbleTea types.BubbleTeaRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return bubbleTea, fmt.Errorf("Unable to read request body: %w", err)
	}
	if err := r.Body.Close(); err != nil {
		return bubbleTea, fmt.Errorf("Unable to close request body: %w", err)
	}
	if err := json.Unmarshal(body, &bubbleTea); err != nil {
		return bubbleTea, fmt.Errorf("Unable to convert request to BubbleTea: %w", err)
	}
	return bubbleTea, nil
}

// TODO: enhance this by using interface
// func marshall(bubbleTeaResponse *types.BubbleTeaResponse) ([]byte, error) {
// 	bubbleTeaResponseBytes, err := json.Marshal(bubbleTeaResponse)
// 	if err != nil {
// 		return bubbleTeaResponseBytes, errors.New("Unable to convert response to bytes")
// 	}
// 	return bubbleTeaResponseBytes, nil
// }

// HealthCheck return start time and uptime
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var healthStatus types.Health
	startTime := time.Now().String()
	healthStatus.StartTime = startTime
	uptime := util.GetUpTime()
	healthStatus.UpTime = uptime

	json, err := json.Marshal(healthStatus)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
