package article

import (
	"fmt"
	"github.com/linuxsuren/jenkins-wechat-tool/cmd/common"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func NewCheckCommand(commonOpts *common.Options) (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "check",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mkFile := args[0]

			var data []byte
			if data, err = ioutil.ReadFile(mkFile); err != nil {
				return
			}

			var article Article
			if err = yaml.Unmarshal(data, &article); err != nil {
				return
			}

			// check this mandatory fields
			if article.Title == "" ||
				article.Description == "" ||
				article.Author == "" ||
				article.Date == "" ||
				article.Poster == "" {
				err = fmt.Errorf("title, description, author, date, poster are  mandatory fields, party of them are missing.")
				return
			}

			if article.Title == article.Description {
				err = fmt.Errorf("description cannot be same with title")
				return
			}

			return
		},
	}
	return
}
