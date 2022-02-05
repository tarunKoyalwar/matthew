package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/matthew/restapi"
)

var append bool
var chunk int
var Proxy string
var Post bool
var Debug bool
var URL string
var R *restapi.Restapi

var rootCmd = &cobra.Command{
	Use:               "matthew",
	Short:             "Client For Sandman",
	Long:              "An Awesome Client for Sandman that can easily be integrated to recon",
	DisableAutoGenTag: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&append, "append", "a", false, "Append to Existing Data")
	rootCmd.PersistentFlags().IntVarP(&chunk, "chunk", "c", 10, "Chunk Size While Writing to Network")
	rootCmd.PersistentFlags().StringVarP(&Proxy, "proxy", "x", "", "Proxy URL")
	rootCmd.PersistentFlags().BoolVar(&Post, "post", false, "Post Data to CheckList Item")
	rootCmd.PersistentFlags().StringVarP(&URL, "url", "u", "http://127.0.0.1:8088", "Rest API URL")
	rootCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "Debug")

	rootCmd.AddCommand(&web)
	rootCmd.AddCommand(&org)
	rootCmd.AddCommand(&Pages)

}

func HandleError(er error) {
	if er != nil {
		panic(er)
	}
}

func PrintAndHandleError(msg string, dat string, err error) {
	HandleError(err)
	fmt.Printf("%v :\n%v", msg, dat)
}
