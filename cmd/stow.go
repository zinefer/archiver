package cmd

import (
	"fmt"
	"time"
	"os"

	"github.com/zinefer/archiver/pkg/archutil"
	"github.com/spf13/cobra"
)

// stowCmd represents the stow command
var stowCmd = &cobra.Command{
	Use:   "stow",
	Aliases: []string{"s"},
	Short: "Stow items older than maxAge",
	Long: `Move any item in targetDir to the stowDir if the latest timestamp 
associated with it is older than maxAge`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Archiver is looking for files to stow...")

		threshold := time.Now().Unix() - int64(maxAge)		
		files := archutil.ListDirectoryByModifiedTimeAsc(targetDir)
		for _, file := range files {
			if (file.ChildModTime().Unix() < threshold) {
				fileStowDir := stowDir

				if (yearDirs) {
					fileStowDir = stowDir + "/" + fmt.Sprint(file.ChildModTime().Year())
				}
				
				filePath    := targetDir + "/" + file.Name()
				destPath := fileStowDir + "/" + file.Name()

				if _, err := os.Stat(fileStowDir); os.IsNotExist(err) {
					if (!dry) {
						fmt.Println("Creating", fileStowDir)
						err := os.MkdirAll(fileStowDir, os.ModePerm)
						if err != nil {
							fmt.Println(err)
						}
					} else {
						fmt.Println("Would have created", fileStowDir)
					}
				}
				
				if (!dry) {
					archutil.Move(filePath, destPath)
					fmt.Println("Moved", filePath, "to", fileStowDir)
				} else {
					fmt.Println("Would have moved", filePath, "to", fileStowDir)
				}
			} else { 
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(stowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}