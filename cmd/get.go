package cmd

import (
	"fmt"
	"os"

	"./arg"
	"./request"
	"./unjson"
	"github.com/spf13/cobra"
)

// getCmd send get request
var getCmd = &cobra.Command{
	Use:   "get [type] [resource]",
	Short: "Send get request to gitlab api.",
	//PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		gla.Type, gla.Subj, err = arg.Parse(args)
		if err != nil {
			cmd.Help()
			os.Exit(0)
		}
		resp := GET(gla.Type, gla.Subj)
		respPretty, _ := unjson.Pretty([]byte(resp))
		fmt.Printf("%s\n", string(respPretty))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVar(&gla.Args, "args", "", "")
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// GET func
func GET(t string, s string) string {

	// Make full URL
	uri := "/api/v4/" + t + "/" + arg.Cook(s)

	req := request.Conf{
		Host: gla.Host,
		URI:  uri,
	}

	req.Headers = request.Params{
		"Private-Token": gla.Token,
	}

	resp, err := request.Send(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp

}
