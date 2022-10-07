package utils

import (
	"fmt"
	"strings"
	"os/exec"
	"github.com/TwiN/go-color"
)

// Guess context from project name
func GuessContext(param string) string {
	var context = "cluster_stg_2"

	if strings.Contains(param, "prod") {
		context = "cluster_prod_1"
	}

	if (strings.Contains(param, "stg") || strings.Contains(param, "preprod")) {
		context = "cluster_stg_2"
	}

	if strings.Contains(param, "dev") {
		context = "cluster_dev_1"
	}

	return context
}

// Return DB pass
func GetDBPass(param string) string {
	fmt.Println("Fetching DB pass for: " + color.Green + _boldText(param) + color.Reset)
	var context = GuessContext(param)
	dbPass := _fetchDBPass(param, context)

	return _boldText("DB Pass: ") + string(dbPass)
}


// Helpers
func _fetchDBPass(param string, context string) string {
	var command = "kubectl-moco credential -u moco-admin " + param + "-magento --context " + context + " -n " + param

	out, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		fmt.Printf("%s", err)
		return "Error while fetch DB pass"
	}

	return string(out)
}

func _boldText(text string) string {
	return color.Bold + text + color.Reset
}