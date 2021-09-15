package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("*** backblaze b2 key updater")
	// Get the GITHUB_USERNAME environment variable
	applicationID, _ := os.LookupEnv("APPLICATION_ID")
  applicationKey, _ := os.LookupEnv("APPLICATION_KEY")
	bucketId, _ := os.LookupEnv("BUCKET_ID")
	nginxConfPath, _ := os.LookupEnv("NGINX_CONF_PATH")
	apiUrl, _ := os.LookupEnv("B2_API_URL")
	// login to account
	loginToken := authAccount(applicationID, applicationKey, apiUrl)

	// get 1 week token
	downloadToken := getAuthorization(loginToken, bucketId, apiUrl)

	// update vhost config with new key
	updateNginxConfig(nginxConfPath, downloadToken)

	// reload nginx to start using new key
	reloadNginx()
}
