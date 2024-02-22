package oauthcli

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetUser uses an initialized oauth2 client to get
// information about a user
func GetUser(client *http.Client) error {
	url := "https://api.github.com/user"

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code from", url, ":", resp.StatusCode)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}
	return nil
}
