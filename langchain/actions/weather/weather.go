package weather

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *WeatherAction) GetNotification() string {
	return fmt.Sprintf("🌤️ Getting forecast: **%s**", a.MustGetParam("location"))
}

func (a *WeatherAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	// Create a new HTTP client
	client := &http.Client{}

	// Define the API endpoint
	apiEndpoint := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", a.MustGetParam("location"), config.React.OpenWeatherMapAPIKey)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	// Send the HTTP request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	// Close the response body
	defer resp.Body.Close()

	if _, err := obj.Write(body); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}
