package cmd

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type stop struct {
	error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func retry(attempts int, sleep time.Duration, f func() error) error {
	err := f()
	if err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2

			time.Sleep(sleep)
			return retry(attempts, sleep, f)
		}
		return err
	}
	return nil
}

func healthcheck(healthcheckPorts string) {
	ports := strings.Split(healthcheckPorts, ",")
	for i := range ports {
		url := "http://localhost:" + ports[i] + "/health"
		checkURLStatus(url)
	}
}

func checkURLStatus(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("unable to make request: %s", err)
	}

	return retry(10, time.Second*5, func() error {
		// Build the request
		resp, err := http.DefaultClient.Do(req)
		fmt.Println("call made to ", url)

		if err != nil {
			// Return error to retry
			return err
		}
		defer resp.Body.Close()

		s := resp.StatusCode
		switch {
		case s >= 500:
			return fmt.Errorf("server error: %v", s)
		case s >= 400:
			return stop{fmt.Errorf("client error: %v", s)}
		default:
			// Happy
			fmt.Println(url + " is up and running")
			return nil
		}
	})
}
