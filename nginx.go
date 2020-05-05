package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func updateNginxConfig(configPath string, downloadKey string) {
	// make sure both config and key even exist
	b2Config, err := fileExists(configPath)
	if err != nil {
		log.Fatal("!!! could not read nginx config or other error:", err)
	}
	haveKey := false
	if downloadKey != "" {
		haveKey = true
	}
	if b2Config && haveKey {
		fmt.Println("*** configuring new token")
		// read in file, buffer text to work on
		// https://stackoverflow.com/a/26153102
		input, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatal("!!! couldn't read nginx config, err:", err)
		}
		confText := strings.Split(string(input), "\n")
		// config string to replace with, with new key
		newKey := fmt.Sprintf("\t\tproxy_set_header Authorization '%s';", downloadKey)
		// iterate over the buffered text
		for i, line := range confText {
			// swap the key out here
			if strings.Contains(line, "Authorization") {
				confText[i] = newKey
			}
		}
		// write modified config out
		output := strings.Join(confText, "\n")
		err = ioutil.WriteFile(configPath, []byte(output), 0644)
		if err != nil {
			log.Fatal("!!! error writing nginx config, err:", err)
		}
	} else {
		log.Fatal("!!! nginx config or key not found, or other error occurred:", err)
	}
}

func reloadNginx() {
	fmt.Println("*** reloading nginx")
	command := exec.Command("/usr/sbin/nginx", "-s", "reload") // think this works
	err := command.Run()
	if err != nil {
		log.Fatal("!!! wasn't able to reload nginx, err:", err)
	}
}
