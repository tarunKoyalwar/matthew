package restapi

import (
	"encoding/json"
	"fmt"
	"log"
)

//pretty print the checklist

var Intend string = "   "

func PrettyPrint(raw []byte) {
	// fmt.Println(string(raw))
	data := map[string][]string{}
	err := json.Unmarshal(raw, &data)
	if err != nil {
		fmt.Println(string(raw))
		log.Panic(err)
	}

	//root nodes
	// prefix := ""
	for _, v := range data[""] {
		//draw root branches
		if len(data[v]) == 0 {
			//this is not a branch
			fmt.Printf("-%v\n", v)
		} else {
			//this is a branch
			PrintBranch("", v, data[v], data)
		}
	}

}

func PrintBranch(prefix string, branch string, childs []string, alldata map[string][]string) {
	fmt.Printf("%v+%v\n", prefix, branch)
	for _, v := range childs {
		if len(alldata[v]) == 0 {
			//this is not a branch
			fmt.Printf("%v%v-%v\n", prefix, Intend, v)
		} else {
			//this is a branch
			PrintBranch(prefix+Intend, v, alldata[v], alldata)
		}
	}
}
