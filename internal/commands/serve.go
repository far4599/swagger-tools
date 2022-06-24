package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/far4599/swagger-openapiv2-merge/pkg/merger"
	"github.com/far4599/swagger-openapiv2-merge/pkg/stoplight_elements"
)

func NewServeCommand() *cobra.Command {
	var (
		httpPort              string
		specHostnameOverwrite string
		openUrl               bool
	)

	serveCmd := &cobra.Command{
		Use:   "serve path",
		Short: "Serve UI for swagger specification",
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
				WithPort(httpPort).
				WithHostname(specHostnameOverwrite).
				WithOpenURL(openUrl).
				Run()
		},
	}

	serveCmd.PersistentFlags().StringVarP(&httpPort, "port", "p", "8080", "Port to serve UI on")
	serveCmd.PersistentFlags().StringVarP(&specHostnameOverwrite, "hostname", "", "", "A new hostname to overwrite one in the spec")
	serveCmd.PersistentFlags().BoolVarP(&openUrl, "open", "o", false, "Open URL after server is loaded")

	return serveCmd
}
