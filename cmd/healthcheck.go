package cmd

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// wrapped error struct to be returned when retry is not going to work i.e. 4XX error codes
type stop struct {
	error
}

// initialize random number generation with a default time based seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// recursively retry to connect to the URL
// attempt the call for as many attempts as specified in the code base
// with a gap between retries based on the sleep duration
// returning success or error based on the conditions checked in the calling method
func retry(attempts int, sleep time.Duration, f func() error) error {
	err := f()
	if err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, sleep, f)
		}
		return err
	}
	return nil
}

// healthcheck is called from docker compose command with a comma seperated string of exposed health check ports
// which are used to generate a well formed URL and sent to checkURLStatus() to verify if the URL is accessible or not
func healthcheck(healthcheckPorts string) {
	ports := strings.Split(healthcheckPorts, ",")
	for i := range ports {
		// generate the URL
		url := "http://localhost:" + ports[i] + "/health"
		// check URL status - 5XX, 4XX, 200 Ok
		checkURLStatus(url)
	}
}

func checkURLStatus(url string) error {
	// build the request based on url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// return err if not able to make a request
		return fmt.Errorf("unable to make request: %s", err)
	}

	return retry(10, time.Second*5, func() error {
		// execute the request
		resp, err := http.DefaultClient.Do(req)
		fmt.Println("Atrempting to connect to : ", url)
		if err != nil {
			// Return error to retry
			return err
		}
		defer resp.Body.Close()

		s := resp.StatusCode
		switch {
		case s >= 500:
			// 5XX results in retry on the URL
			return fmt.Errorf("server error: %v", s)
		case s >= 400:
			// 4XX results in exit condition as the error seems to be from the calling client
			return stop{fmt.Errorf("client error: %v", s)}
		default:
			// this condition is executed when there is no error and the response is as per expectation
			fmt.Println(url + " is verified to be up and running")
			return nil
		}
	})
}
