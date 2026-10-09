package cli

import (
	"fmt"
	"os"

	del "github.com/pikpakidm/pikpakidm/cli/del"
	"github.com/pikpakidm/pikpakidm/cli/download"
	"github.com/pikpakidm/pikpakidm/cli/empty"
	"github.com/pikpakidm/pikpakidm/cli/list"
	"github.com/pikpakidm/pikpakidm/cli/move"
	"github.com/pikpakidm/pikpakidm/cli/new"
	"github.com/pikpakidm/pikpakidm/cli/quota"
	"github.com/pikpakidm/pikpakidm/cli/rename"
	"github.com/pikpakidm/pikpakidm/cli/rubbish"
	"github.com/pikpakidm/pikpakidm/cli/save"
	"github.com/pikpakidm/pikpakidm/cli/setup"
	"github.com/pikpakidm/pikpakidm/cli/share"
	"github.com/pikpakidm/pikpakidm/cli/upload"
	"github.com/pikpakidm/pikpakidm/conf"
	"github.com/pikpakidm/pikpakidm/internal/logx"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pikpakidm",
	Short: "PikPakIDM is a command line download manager for PikPak with aria2 support",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.CommandPath() == "pikpakidm setup" {
			logx.Init(debug, debugTopics)
			return
		}
		err := conf.InitConfig(configPath)
		if err != nil {
			fmt.Println("Init config failed")
			logx.Error(err)
			os.Exit(1)
		}
		logx.Init(debug, debugTopics)
	},
}

var configPath string
var debug bool
var debugTopics []string

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
	rootCmd.PersistentFlags().StringSliceVar(&debugTopics, "debug-topic", nil, "enable debug topics: api,session,transfer")
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "config.yml", "config file path")
	rootCmd.AddCommand(upload.UploadCmd)
	rootCmd.AddCommand(download.DownloadCmd)
	rootCmd.AddCommand(share.ShareCommand)
	rootCmd.AddCommand(new.NewCommand)
	rootCmd.AddCommand(quota.QuotaCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(del.DeleteCmd)
	rootCmd.AddCommand(empty.EmptyCmd)
	rootCmd.AddCommand(rubbish.RubbishCmd)
	rootCmd.AddCommand(rename.RenameCmd)
	rootCmd.AddCommand(move.MoveCmd)
	rootCmd.AddCommand(save.SaveCommand)
	rootCmd.AddCommand(setup.SetupCmd)
	rootCmd.AddCommand(shellCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
