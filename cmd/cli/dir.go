package cli

import (
	"fmt"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/internal/rules"
	"github.com/configAnalyzer/internal/service"
	"github.com/spf13/cobra"
)

func dirCmd() *cobra.Command {

	var (
		silent         bool
		checkPermition bool
	)

	cmd := &cobra.Command{
		Use:          "dir <directory>",
		Short:        "Рекурсивно анализировать директорию с конфигами",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, args []string) error {

			if len(args) == 0 {
				return fmt.Errorf("не указан путь к директории")
			}
			path := args[0]

			// подготовим файловый анализатор
			fileAnalyer := analyzer.NewAnalyzer()
			if checkPermition {
				// добавим новое правило в анализатор
				fileAnalyer.AddRule(&rules.FilePermissionRule{Path: path})
			}

			// подготовим анализатор директории
			analyzer := service.NewDirAnalyzer(fileAnalyer)

			fileIssues, err := analyzer.Run(path)
			if err != nil {
				return err
			}

			if len(fileIssues) > 0 {

				totalIssues := 0
				for _, f := range fileIssues {
					fmt.Println(entities.GetFileIssuesInfo(f))
					totalIssues += len(f.Issues)
				}

				if !silent {
					return fmt.Errorf("проблем обнаружено: %d", totalIssues)
				}

			} else {
				fmt.Println("Проблем не обнаружено")
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&silent, "silent", "s", false, "не завершаться с ошибкой при наличии проблем")
	cmd.Flags().BoolVar(&checkPermition, "chp", false, "проверять права файлов")
	return cmd
}
