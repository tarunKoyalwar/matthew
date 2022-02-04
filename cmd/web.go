package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/matthew/restapi"
	"github.com/tarunKoyalwar/matthew/stdin"
)

var wchecklist bool

var wtoolname string

var working bool

var sub string

var web cobra.Command = cobra.Command{
	Use:   "web",
	Short: "Web Checklists Page",
	Long:  "Interacting with Web Checklists page",
	Run: func(cmd *cobra.Command, args []string) {
		if Proxy != "" {
			restapi.Proxy = Proxy
		}
		var err error
		R, err = restapi.NewRestapi(URL)
		HandleError(err)
		R.Alive()
		if working {
			dat, err := R.Workingsub()
			if dat == "default" {
				fmt.Println("Working Subdomain Not Set in UI")
				return
			} else {
				PrintAndHandleError("Working Subdomain is", dat, err)
			}
			return
		}
		if wchecklist {
			//print checklist
			bin, err := R.GetCheckList(true)
			HandleError(err)
			restapi.PrettyPrint(bin)
			return
		}
		if sub != "" && wtoolname != "" {
			if Post {
				func(chx chan stdin.Receive) {
					start := true
					for {
						dat, ok := <-ch
						if !ok {
							break
						} else {
							if start {
								R.PostWebToolOutput(sub, wtoolname, dat.ByteData, append)
								fmt.Print(dat.StringData)
								start = false
							} else {
								R.PostWebToolOutput(sub, wtoolname, dat.ByteData, true)
								fmt.Print(dat.StringData)
							}

						}

					}
					R.SavetoDB()
				}(ch)

			} else {
				dat, err := R.GetWebToolOutput(sub, wtoolname)
				HandleError(err)
				fmt.Println(dat)
			}
		} else {
			fmt.Println("Subdomain Not Mentioned or Toolname not mentioned")
		}

	},
}

func init() {
	web.Flags().BoolVarP(&wchecklist, "checklist", "L", false, "Just Print Available CheckList")
	web.Flags().StringVarP(&wtoolname, "tool", "t", "", "CheckList Item name to Use(Substring Works)")
	web.Flags().BoolVar(&working, "working", false, "Only Print Working Subdomain Name")
	web.Flags().StringVarP(&sub, "sub", "s", "", "Use Checklist for this subdomain")

	web.Example = `
	matthew web -L  // Prints Your CheckList
	matthew web -s abc.com -t nmap // Fetches nmap results of abc.com subdomain
	asserfinder abc.com | matthew web -s abc.com -t assetfinder --post // Posts assetfinder results to abc.com assetfinder
	matthew web --working // Fetches Subdomain that is currently Used 
	`

}
