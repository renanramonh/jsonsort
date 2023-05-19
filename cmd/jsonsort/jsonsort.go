package jsonsort

import (
	"fmt"
	"os"

	"github.com/renanramonh/jsonsort/pkg/file"
	"github.com/renanramonh/jsonsort/pkg/json"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "jsonsort [file]",
	Short:   "Sort json file",
	Long:    `Sort a json file and save the output in a new file suffixed with _sorted`,
	Args:    cobra.MinimumNArgs(1),
	Example: "jsonsort ./myjson.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		source := args[0]
		if source == "" {
			return fmt.Errorf("a valid file should be informed")
		}
		if !file.Exists(source) {
			return fmt.Errorf("file not found")
		}

		content, err := os.ReadFile(source)
		if err != nil {
			return err
		}

		destinationFile, err := os.Create(fmt.Sprintf("%v_sorted.json", file.GetFileNameWithoutExtension(source)))
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		sorted, err := json.Sort(string(content))
		if err != nil {
			return err
		}

		_, err = destinationFile.Write([]byte(sorted))
		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
