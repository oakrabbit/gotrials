package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)


var (
	// used for flags
	cfgFile	string
	userLicense	string
	rootCmd = &cobra.Command{
		Use: "gotrials",
		Short: "this is just a trial",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("executed command!")
		},
	}
)

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init(){
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config","","config File is default($HOME/.gotrials)")
	rootCmd.PersistentFlags().StringP("author","a", "Alex", "author name")
	viper.BindPFlag("author",rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author","Alex . email@email.com")


}

func initConfig(){
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}else {
		//find home directory
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		//Search config in home directory with name ".gotrials" no extension
		viper.AddConfigPath(home)
		viper.SetConfigName(".gotrials")

	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using Config file:", viper.ConfigFileUsed())
	}
}