package main

import (
	"chatgpt-adapter/internal/common"
	handler "chatgpt-adapter/internal/gin.handler"
	"chatgpt-adapter/internal/vars"
	"chatgpt-adapter/logger"
	"chatgpt-adapter/pkg"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var (
	version  = "v2.1.0"
	logLevel = "info"
	logPath  = "log"
	vms      bool

	cmd = &cobra.Command{
		Use:   "ChatGPT-Adapter",
		Short: "GPT接口适配器",
		Long: "GPT接口适配器。统一适配接口规范，集成了bing、claude-2，gemini...\n" +
			"项目地址：https://github.com/bincooo/chatgpt-adapter",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			if vms {
				fmt.Println("模型可用列表:")
				for _, model := range handler.GlobalExtension.Models() {
					fmt.Println("- " + model.Id)
				}
				return
			}

			banner()

			logger.InitLogger(logPath, switchLogLevel())

			pkg.InitConfig()
			common.InitCommon()

			// 获取 PORT 环境变量
			port := os.Getenv("PORT")
			if port == "" {
				port = "8080" // 默认端口
			}

			// 将端口字符串转换为整数
			portInt, err := strconv.Atoi(port)
			if err != nil {
				logrus.Fatalf("无效的端口号: %s", port)
			}

			handler.Bind(portInt, version, vars.Proxies)
		},
	}
)

func banner() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("\n █████╗ ██████╗  █████╗ ██████╗ ████████╗███████╗██████╗ \n██╔══██╗██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝██╔══██╗\n███████║██║  ██║███████║██████╔╝   ██║   █████╗  ██████╔╝\n██╔══██║██║  ██║██╔══██║██╔═══╝    ██║   ██╔══╝  ██╔══██╗\n██║  ██║██████╔╝██║  ██║██║        ██║   ███████╗██║  ██║\n╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝        ╚═╝   ╚══════╝╚═╝  ╚═╝\n                                                         \n")
	fmt.Println("PROJECT: https://github.com/bincooo/chatgpt-adapter")
	fmt.Println("VERSION:", version)
	fmt.Println("欢迎STAR！   Welcome STAR！")
	fmt.Print("-----------------------------------------------------------\n\n\n")
}

func main() {
	cmd.PersistentFlags().StringVar(&vars.Proxies, "proxies", "", "本地代理 proxies")
	cmd.PersistentFlags().StringVar(&logLevel, "log", logLevel, "日志级别: trace|debug|info|warn|error")
	cmd.PersistentFlags().StringVar(&logPath, "log-path", logPath, "日志路径")
	cmd.PersistentFlags().BoolVar(&vms, "models", false, "查看所有模型")
	_ = cmd.Execute()
}

func switchLogLevel() logrus.Level {
	switch logLevel {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
