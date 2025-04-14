package sakura

import (
	"fmt"
	"os"

	"github.com/s-chernyavskiy/sakura/internal/sakura/config"
	"github.com/s-chernyavskiy/sakura/internal/sakura/errors"
	"github.com/s-chernyavskiy/sakura/internal/sakura/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sakura",
	Short: "sakura short",
	Long:  `sakura long`,
	Run:   runApp,
}

var (
	verbose bool
	cfgFile string
)

const APPVER = "0.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display version",
	Long:  `display version on the screen`,
	Run: func(cmd *cobra.Command, args []string) {
		//  TODO: automated build version to be added
		fmt.Println("Version", APPVER)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "", "configuration file")
	rootCmd.Flags().StringP("host", "", "127.0.0.1", "host for running application")
	rootCmd.Flags().IntP("port", "p", 6969, "port for running application")
	rootCmd.Flags().IntP("maxClients", "", 10000, "max connections can be handled")
	rootCmd.Flags().IntP("maxTimeout", "", 120, "max timeout for clients(in seconds)")

	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("host", rootCmd.Flags().Lookup("host"))
	viper.BindPFlag("maxClients", rootCmd.Flags().Lookup("maxClients"))
	viper.BindPFlag("maxTimeout", rootCmd.Flags().Lookup("maxTimeout"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	loadingDefault := false
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		loadingDefault = true
		viper.SetConfigName("config.yaml")
		viper.AddConfigPath("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		if !loadingDefault {
			fmt.Fprintf(os.Stderr, "Error reading config from %s : %s\n", viper.ConfigFileUsed(), err)
			os.Exit(1)
		}

		fmt.Printf("%T\n", err)

		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			fmt.Fprintf(os.Stderr, "Error reading config file from config directory\nLoading with application defaults...\n")
			break

		default:
			fmt.Fprintf(os.Stderr, "Error reading config in %s : %s\n", viper.ConfigFileUsed(), err)
			os.Exit(1)
		}
	}
}

func Execute() {
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runApp(cmd *cobra.Command, args []string) {
	var cfg config.AppConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		errors.PrintErrorAndExit(err, 2)
	}

	server.Start(cfg)
}
