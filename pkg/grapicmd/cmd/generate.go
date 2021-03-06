package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/module"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/usecase"
)

func newGenerateCommand(cfg grapicmd.Config, ui module.UI, generator module.Generator, commandFactory module.CommandFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate GENERATOR",
		Short:   "Generate new code",
		Aliases: []string{"g", "gen"},
	}

	cmd.AddCommand(newGenerateServiceCommand(cfg, ui, generator, commandFactory))
	cmd.AddCommand(newGenerateScaffoldServiceCommand(cfg, ui, generator, commandFactory))
	cmd.AddCommand(newGenerateCommandCommand(cfg, generator))

	return cmd
}

func newGenerateServiceCommand(cfg grapicmd.Config, ui module.UI, generator module.ServiceGenerator, commandFactory module.CommandFactory) *cobra.Command {
	var (
		skipTest     bool
		resourceName string
	)

	cmd := &cobra.Command{
		Use:           "service NAME [METHODS...]",
		Short:         "Generate a new service",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cfg.IsInsideApp() {
				return errors.New("geneate command should execut inside a grapi applicaiton directory")
			}

			ui.Section("Generate service")
			genCfg := module.ServiceGenerationConfig{ResourceName: resourceName, Methods: args[1:], SkipTest: skipTest}
			err := errors.WithStack(generator.GenerateService(args[0], genCfg))
			if err != nil {
				return err
			}

			protocUsecase := usecase.NewExecuteProtocUsecase(cfg.ProtocConfig(), cfg.Fs(), ui, commandFactory, cfg.RootDir())
			return errors.WithStack(protocUsecase.Perform())
		},
	}

	cmd.PersistentFlags().BoolVarP(&skipTest, "skip-test", "T", false, "Skip test files")
	cmd.PersistentFlags().StringVar(&resourceName, "resource-name", "", "ResourceName to be used")

	return cmd
}

func newGenerateScaffoldServiceCommand(cfg grapicmd.Config, ui module.UI, generator module.ServiceGenerator, commandFactory module.CommandFactory) *cobra.Command {
	var (
		skipTest     bool
		resourceName string
	)

	cmd := &cobra.Command{
		Use:           "scaffold-service NAME",
		Short:         "Generate a new service with standard methods",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cfg.IsInsideApp() {
				return errors.New("geneate command should execut inside a grapi applicaiton directory")
			}

			ui.Section("Scaffold service")
			genCfg := module.ServiceGenerationConfig{ResourceName: resourceName, SkipTest: skipTest}
			err := errors.WithStack(generator.ScaffoldService(args[0], genCfg))
			if err != nil {
				return err
			}

			protocUsecase := usecase.NewExecuteProtocUsecase(cfg.ProtocConfig(), cfg.Fs(), ui, commandFactory, cfg.RootDir())
			return errors.WithStack(protocUsecase.Perform())
		},
	}

	cmd.PersistentFlags().BoolVarP(&skipTest, "skip-test", "T", false, "Skip test files")
	cmd.PersistentFlags().StringVar(&resourceName, "resource-name", "", "ResourceName to be used")

	return cmd
}

func newGenerateCommandCommand(cfg grapicmd.Config, generator module.CommandGenerator) *cobra.Command {
	return &cobra.Command{
		Use:           "command NAME",
		Short:         "Generate a new command",
		SilenceErrors: true,
		SilenceUsage:  true,
		Args:          cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !cfg.IsInsideApp() {
				return errors.New("geneate command should execut inside a grapi applicaiton directory")
			}

			return errors.WithStack(generator.GenerateCommand(args[0]))
		},
	}
}
