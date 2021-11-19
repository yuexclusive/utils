/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/spf13/cobra"
	"github.com/yuexclusive/utils/cmd"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		publish()
	},
}

type Publish struct {
	ProjectName string
	Type        string
	AppName     string
	Version     string
	Host        string
}

var golangTemplate string = `
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
docker build . -t registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker push registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
ssh root@{{.Host}} "
docker pull registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker stop future.{{.Type}}.{{.AppName}}_1
docker rm future.{{.Type}}.{{.AppName}}_1
docker run -d --network=future_default --name=future.{{.Type}}.{{.AppName}}_1 registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
"
`

var vueTemplate string = `
npm install
npm run build
docker build . -t registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker push registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
ssh root@{{.Host}} "
docker stop future.{{.Type}}.admin_1
docker rm future.{{.Type}}.admin_1
docker pull registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
docker run -d -p 9090:80 -v /root/future/nginx.conf:/etc/nginx/nginx.conf --network=future_default --name=future.{{.Type}}.admin_1 registry.cn-shenzhen.aliyuncs.com/liuyuexclusive/{{.ProjectName}}.{{.Type}}.{{.AppName}}:{{.Version}}
"
`

var version, host string

func publish() {
	now := time.Now()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := path.Base(dir)

	reg := regexp.MustCompile(`(.+).(srv|web|front).(.+)`)

	if !reg.MatchString(baseDir) {
		panic("错误的项目名称，请使用[project_name].[srv|web|front].[app_name]的格式")
	}

	res := reg.FindAllStringSubmatch(baseDir, -1)

	var publish Publish

	publish.ProjectName = res[0][1]
	publish.Type = res[0][2]
	publish.AppName = res[0][3]
	publish.Version = version
	publish.Host = host

	_, err = os.Stat("./publish.txt")
	if err != nil {
		fileInfo, err := os.Create("publish.txt")

		if err != nil {
			log.Fatal(err)
		}
		defer fileInfo.Close()
		var content string
		switch publish.Type {
		case "srv", "web":
			content = golangTemplate
		case "front":
			content = vueTemplate

		}
		fileInfo.WriteString(content)
	}

	tem, err := template.ParseFiles("publish.txt")

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)

	tem.Execute(buf, &publish)

	command := buf.String()

	fmt.Println(command)

	cmd.Run("sh", "-c", command)

	fmt.Printf("总耗时:%f秒\n", time.Since(now).Seconds())
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	publishCmd.Flags().StringVar(&version, "v", "latest", "docker image version")
	publishCmd.Flags().StringVar(&host, "h", "49.232.166.55", "host")
}
