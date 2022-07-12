package cmd

import (
	"io/fs"
	"os"
	"bufio"
	"fmt"

	"github.com/zinefer/archiver/pkg/archutil"
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore [item]",
	Aliases: []string{"r"},
	Short: "Restore a stowed item",
	Long: `Search for [item] in stowDir and offer to move it to targetDir`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Archiver is looking for your files to restore...")

		wantedFile := args[0]
		var stowedFile fs.FileInfo

		if (yearDirs) {
			yearDirs := archutil.ListDirectory(stowDir)
			for i := len(yearDirs)-1; i >= 0; i-- {
				yearDir := yearDirs[i]
				fmt.Println("🔍", yearDir.Name())
				
				stowedFile = SearchStowDirForFile(stowDir + "/" + yearDir.Name(), wantedFile)
				if (stowedFile != nil) {
					stowDir = stowDir + "/" + yearDir.Name()
					break
				}
			}
		} else {
			stowedFile = SearchStowDirForFile(stowDir, wantedFile)
		}

		if (stowedFile != nil) {
			filePath := stowDir + "/" + stowedFile.Name()
			destPath := targetDir + "/" + stowedFile.Name()

			if (!dry) {
				archutil.Move(filePath, destPath)
				fmt.Println("Moved", filePath, "to", targetDir)
			} else {
				fmt.Println("Would have moved", stowDir + "/" + stowedFile.Name(), "to", targetDir)	
			}
		} else {
			fmt.Println("No suitable items found")
		}
		
	},
}

func SearchStowDirForFile(path string, file string) fs.FileInfo {
	files := archutil.ListDirectoryByModifiedTimeAsc(path)
	for i := len(files)-1; i >= 0; i-- {
		name := files[i].Name()
		if (file == name) {

			if (files[i].IsDir()) {
				fmt.Println("Found Directory", path + "/" + name, "with contents:")
				fmt.Println(archutil.PrintDirectory(path + "/" + name))
			} else {
				fmt.Println("Found File", path + "/" + name)
			}

			if (UserConfirm("Restore this stowed item?")) {
				return files[i];
			} else {
				fmt.Println("Continuing to look for other items ...")
			}
		}
	}
	return nil
}

func UserConfirm(question string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + " y\\n ")
	text, _ := reader.ReadString('\n')
	return text == "y\n" || text == "Y\n"
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
