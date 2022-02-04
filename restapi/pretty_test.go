package restapi_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tarunKoyalwar/matthew/restapi"
)

func Test_prettychecklist(t *testing.T) {
	var tssemplate = map[string][]string{
		"":                      {"Information Gathering", "Scanning", "Enumeration"},
		"Information Gathering": {"Nmap", "Whois", "Whatweb", "Nuclei"},
		"Scanning":              {"Nikito", "JSFinder"},
		"Nuclei":                {"CVE", "Tech-Detect", "Others"},
	}

	bin, _ := json.Marshal(tssemplate)

	fmt.Printf("\n\n")
	restapi.PrettyPrint(bin)
	fmt.Printf("\n\n")
}
