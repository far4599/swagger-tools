package commands

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/far4599/swagger-openapiv2-merge/pkg/merger"
)

const (
	swaggerDefaultVersion = "2.0"
)

func NewMergeCommand() *cobra.Command {
	var (
		infoFile, specVersion, outputFile, filterExt string
		withSubdir                                   bool
	)

	mergeCmd := &cobra.Command{
		Use:   "merge [flags] path",
		Short: "Merge swagger json files in path directory",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				fromDir = args[0]
			)

			specsToJoin, err := merger.LoadSpecsFromDir(fromDir, filterExt, withSubdir)
			if err != nil {
				return errors.Wrap(err, "failed to load specs to merge")
			}

			if len(specsToJoin) == 0 {
				return fmt.Errorf("no valid specs found in directory '%s' with extension '%s'", fromDir, filterExt)
			}

			m := merger.NewMerger(specVersion)
			for i := range specsToJoin {
				m.MergeSpec(specsToJoin[i])
			}

			err = merger.WriteSpecToFile(m.Result(), outputFile)
			if err != nil {
				return errors.Wrapf(err, "failed to write merged spec to file '%s'", outputFile)
			}

			if len(infoFile) > 0 {
				infoFileSpec, err := merger.GetSpecFromFile(infoFile)
				if err != nil {
					return errors.Wrapf(err, "failed to load info spec from file '%s'", infoFile)
				}

				err = m.MergeInfo(infoFileSpec, false)
				if err != nil {
					return errors.Wrap(err, "failed to add info spec")
				}
			}

			fmt.Printf("%d files merged", len(specsToJoin))

			return nil
		},
	}

	mergeCmd.PersistentFlags().StringVarP(&infoFile, "info", "i", "", "File with swagger spec to use as a source of info and host properties")
	mergeCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "./swagger.json", "File to write merged specs to")
	mergeCmd.PersistentFlags().StringVar(&specVersion, "spec-version", swaggerDefaultVersion, "Version of swagger specification")
	mergeCmd.PersistentFlags().StringVar(&filterExt, "filter-ext", ".swagger.json", "Merge only files with this extension in their names")
	mergeCmd.PersistentFlags().BoolVar(&withSubdir, "with-subdir", false, "Look for spec files in sub directories")

	return mergeCmd
}
