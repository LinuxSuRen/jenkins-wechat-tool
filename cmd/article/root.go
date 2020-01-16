package article

import (
	"github.com/linuxsuren/jenkins-wechat-tool/cmd/common"
	"github.com/spf13/cobra"
)

func NewArticleCommand(commonOpts *common.Options) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "article",
	}

	cmd.AddCommand(NewCheckCommand(commonOpts))
	return
}

type Article struct {
	Title       string
	Description string
	Date        string
	Poster      string
	Author      string
	Translator  string
	Original    string
	Toc         bool
}
