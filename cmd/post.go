package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"

	"./arg"
	"./request"
	"./unjson"
	"github.com/spf13/cobra"
)

// postCmd - send POST request to API
var postCmd = &cobra.Command{
	Aliases: []string{"create"},
	Use:     "add [resourse]",
	Short:   "Send POST request to gitlab api.",
	//PreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		gla.Type, gla.Subj, err = arg.Parse(args)
		if err != nil {
			cmd.Help()
			os.Exit(0)
		}
		resp := POST(gla.Type, gla.Subj)
		respPretty, _ := unjson.Pretty([]byte(resp))
		fmt.Printf("%s\n", string(respPretty))
	},
}

func init() {
	pflag.Parse()
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().StringVar(&gla.Args, "args", "", "")
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// POST func
func POST(t string, s string) string {
	var locationID string
	location, name := arg.Split(s, "/")
	if name != "" {
		if location != "" {
			locationJSON := GET("groups", arg.Cook(location))
			var locationInfo unjson.BaseJSON = unjson.Get(locationJSON)
			locationID = strconv.Itoa(locationInfo.ID)
		}
	}

	// Make full URL
	locationType := "parent_id"
	if t == "projects" {
		locationType = "namespace_id"
	}
	uri := "/api/v4/" + t

	req := request.Conf{
		Host: gla.Host,
		URI:  uri,
		Type: "POST",
	}

	req.Headers = request.Params{
		"Private-Token": gla.Token,
	}

	req.Forms = request.Params{
		"name":       name,
		"path":       name,
		locationType: locationID,
	}

	if gla.Args != "" {
		req.Forms = request.Split(req.Forms, gla.Args, ",")
	}

	resp, err := request.Send(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp

}
