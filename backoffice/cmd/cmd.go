/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Terence1105/LineBot/db"
	_ "github.com/Terence1105/LineBot/viper"
	"github.com/spf13/cobra"
)

var userId string

var rootCmd = &cobra.Command{
	Use: "",
}

var GetMessageLogByUserIdCmd = &cobra.Command{
	Use:   "GetMessageLogByUserIdCmd",
	Short: "Get specify user message logs.",
	Long:  `Get specify user message logs.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		logs := db.NewDB().GetMessageLogByUserId(ctx, userId)

		for _, log := range logs {
			fmt.Println(log.Message)
		}
	},
}

var GetAllMessageLogCmd = &cobra.Command{
	Use:   "GetAllMessageLogCmd",
	Short: "Get all user message logs.",
	Long:  `Get all user message logs.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		logs := db.NewDB().GetAllMessageLog(ctx)

		for _, log := range logs {
			fmt.Println(log.Message)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(GetMessageLogByUserIdCmd)
	rootCmd.AddCommand(GetAllMessageLogCmd)
	GetMessageLogByUserIdCmd.Flags().StringVarP(&userId, "userId", "u", "", "user id")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
