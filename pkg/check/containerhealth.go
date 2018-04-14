package check

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/samirprakash/boom/pkg/handle"
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
// make specified number of attempts with a gap as per specified duration
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

// IfDockerComposeResponds is called from docker compose command with a comma seperated string of exposed health check ports
// which are used to generate a well formed URL and sent to checkURLStatus() to verify if the URL is accessible or not
func IfDockerComposeResponds(healthcheckPorts string) {
	ports := strings.Split(healthcheckPorts, ",")
	for i := range ports {
		// generate the URL
		url := "http://localhost:" + ports[i] + "/health"

		// check URL status - 5XX, 4XX, 200 Ok
		err := checkURLStatus(url)
		if err != nil {
			handle.Info("not able to check the status for %v", url)
			break
		}
	}
}

func checkURLStatus(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("unable to make request: %s", err)
	}

	return retry(60, time.Second*3, func() error {
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
			handle.Info(url + " is verified to be up and running")
			return nil
		}
	})
}
