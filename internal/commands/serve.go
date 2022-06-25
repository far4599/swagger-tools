package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/far4599/swagger-openapiv2-merge/pkg/merger"
	"github.com/far4599/swagger-openapiv2-merge/pkg/stoplight_elements"
)

func NewServeCommand() *cobra.Command {
	var (
		serverHost, serverPort string
		specHostnameOverwrite  string
		withOpenUrl            bool
	)

	serveCmd := &cobra.Command{
		Use:   "serve path",
		Short: "Serve web UI for swagger specification",
		Long:  "It starts the web server with pretty UI for showing the swagger specification. Reload the page to show the changes made to the spec.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				filePath = args[0]
			)

			_, err := merger.GetSpecFromFile(filePath)
			if err != nil {
				return errors.Wrapf(err, "failed to get spec from file '%s'", filePath)
			}

			// validate spec here

			return stoplight_elements.
				NewServer(filePath).
				WithServerHost(serverHost).
				WithServerPort(serverPort).
				WithHostname(specHostnameOverwrite).
				WithOpenURL(withOpenUrl).
				Run()
		},
	}

	serveCmd.PersistentFlags().StringVarP(&serverHost, "host", "", "127.0.0.1", "Server host to serve UI on")
	serveCmd.PersistentFlags().StringVarP(&serverPort, "port", "p", "8080", "Server port to serve UI on")
	serveCmd.PersistentFlags().StringVarP(&specHostnameOverwrite, "hostname", "", "", "A new hostname to overwrite one in the spec")
	serveCmd.PersistentFlags().BoolVarP(&withOpenUrl, "open", "o", false, "Open URL in browser after server is started")

	return serveCmd
}
