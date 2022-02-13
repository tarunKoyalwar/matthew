package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/matthew/restapi"
	"github.com/tarunKoyalwar/matthew/stdin"
)

var Pname string

var Pages cobra.Command = cobra.Command{
	Use:   "pages",
	Short: "Interact with other pages",
	Long:  "Interacting With Other Pages like in-scope etc",
	Run: func(cmd *cobra.Command, args []string) {
		ch = stdin.GetStdinPipe()
		// fmt.Println("pages called")
		if Proxy != "" {
			DebugPrintWithArgs("Using Proxy %v", Proxy)
			restapi.Proxy = Proxy
		}
		var err error
		R, err = restapi.NewRestapi(URL)
		HandleError(err)
		R.Alive()
		if Pname != "" {
			if Post && file == "" {
				DebugPrint("Sending Data to Server")
				func(chx chan stdin.Receive) {
					start := true
					for {
						dat, ok := <-ch
						if !ok {
							break
						} else {
							if start {
								R.PostPage(Pname, dat.ByteData, append)
								fmt.Print(dat.StringData)
								start = false
							} else {
								R.PostPage(Pname, dat.ByteData, true)
								fmt.Print(dat.StringData)
							}
						}

					}
					R.SavetoDB()
				}(ch)

			} else if Post && file != "" {
				dat := GetFileData()
				R.PostPage(Pname, dat, append)
				R.SavetoDB()
				fmt.Printf("Data updated to server from %v\n", file)
			} else {
				dat, err := R.GetPage(Pname)
				HandleError(err)
				fmt.Println(dat)
			}
		}
	},
}

func init() {
	Pages.Flags().StringVarP(&Pname, "page", "p", "", "Page Name Ex:in-scope etc")

	Pages.Example = `
	matthew pages -p in-scope // Gets all Inscope Data
	bash subdomains.sh | matthew pages -p all-subs --post // adds all unique subs to all-subs page
	echo abc.com | gau | matthew pages -p all-urls --post // Adds all unique urls to all-urls page
	`
}
