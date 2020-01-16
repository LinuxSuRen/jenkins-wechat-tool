package cmd

import (
	"github.com/linuxsuren/jenkins-wechat-tool/cmd/common"
	"github.com/spf13/cobra"
	"github.com/linuxsuren/jenkins-wechat-tool/cmd/article"
	"os"
)

var rootCmd = &cobra.Command {
	Use:   "jwt",
	Short: "Jenkins WeChat tool",
}

var commonOptions = common.Options{}
func init() {
	rootCmd.AddCommand(article.NewArticleCommand(&commonOptions))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(err)
		os.Exit(1)
	}
}