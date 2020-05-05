package main

import (
	"fmt"
)

func main() {
	fmt.Println("*** backblaze b2 key updater")

	// login to account
	loginToken := authAccount(applicationID, applicationKey)

	// get 1 week token
	downloadToken := getAuthorization(loginToken, bucketId)

	// update vhost config with new key
	updateNginxConfig(nginxConfPath, downloadToken)

	// reload nginx to start using new key
	reloadNginx()
}
