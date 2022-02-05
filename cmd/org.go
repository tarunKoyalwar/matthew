package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/matthew/restapi"
	"github.com/tarunKoyalwar/matthew/stdin"
)

var ochecklist bool

var otoolname string

var org cobra.Command = cobra.Command{
	Use:   "org",
	Short: "Organization Checklist Page",
	Long:  "Interacting with Organization Checklist page",
	Run: func(cmd *cobra.Command, args []string) {
		ch = stdin.GetStdinPipe()
		if Proxy != "" {
			restapi.Proxy = Proxy
		}
		var err error
		R, err = restapi.NewRestapi(URL)
		HandleError(err)
		R.Alive()
		if ochecklist {
			bin, err := R.GetCheckList(false)
			HandleError(err)
			restapi.PrettyPrint(bin)
			return
		} else if otoolname != "" {
			if Post {
				//post data here
				func(chx chan stdin.Receive) {
					start := true
					for {
						dat, ok := <-ch
						if !ok {
							break
						} else {
							if start {
								R.PostOrgToolOutput(otoolname, dat.ByteData, append)
								fmt.Print(dat.StringData)
								start = false
							} else {
								R.PostOrgToolOutput(otoolname, dat.ByteData, true)
								fmt.Print(dat.StringData)
							}

						}

					}
					R.SavetoDB()
				}(ch)
			} else {
				dat, err := R.GetOrgToolOutput(otoolname)
				HandleError(err)
				fmt.Println(dat)
			}
		} else {
			fmt.Println("Toolname not mentioned")
		}

	},
}

func init() {
	org.Flags().BoolVarP(&ochecklist, "checklist", "L", false, "Just Print Available CheckList")
	org.Flags().StringVarP(&otoolname, "tool", "t", "", "CheckList Item name to Use(Substring Works)")

	org.Example = `
	matthew org -L  // Prints Your CheckList
	matthew org -s abc.com -t sublist3r // Fetches sublist3r results of abc.com subdomain
	asserfinder abc.com | matthew org -t assetfinder --post // Posts assetfinder results to checklist entry
	that contains keyword assetfinder
	`
}
