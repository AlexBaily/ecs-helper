package cmd

var rootCmd = &cobra.Command{
	Use:   "ecs-helper",
	Short: "ecs-helper is a command line utility for ECS maintenance",
	Long: `ecs-helper is a command line utility for ECS maintenance,
				  troubleshooting and other activies.`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
  }
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}