package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var (
	targetDir string
	stowDir   string
	maxAge    int
	yearDirs  bool
	dry       bool

	rootCmd = &cobra.Command{
		Use:   "archiver",
		//Short: "",
		Long: `🔍 Archiver helps keep a directory clean by moving old or unused items 
to an archive location`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		Args: cobra.ExactArgs(1),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.archiver.yaml)")
	rootCmd.PersistentFlags().StringVar(&targetDir, "targetDir", ".", "Directory to stow items from")
	rootCmd.PersistentFlags().StringVar(&stowDir, "stowDir", "./archive", "Directory to stow items to")
	rootCmd.PersistentFlags().IntVar(&maxAge, "age", 90, "Maximum age before stowing an item in days")
	rootCmd.PersistentFlags().BoolVar(&yearDirs, "years", true, "Organize the stow directory with annual folders")
	rootCmd.PersistentFlags().BoolVar(&dry, "dry", false, "Dry run prevents files from actually being moved")
	
	maxAge = maxAge * 24 * 60 * 60

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}