package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/entities"
	"github.com/configAnalyzer/internal/parser"
	"github.com/configAnalyzer/internal/rules"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {

	// флаги
	var (
		silent         bool
		fromStdin      bool
		checkPermition bool
	)

	cmd := &cobra.Command{
		Use:          "configAnalyzer [file]",
		Short:        "Анализатор безопасности конфигурационных файлов",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			// подготовим анализатор
			analyzer := analyzer.NewAnalyzer()

			var (
				cfg  map[string]any // распарсенный конфиг
				err  error
				path string // путь к файлу
			)

			if fromStdin {
				data, readErr := io.ReadAll(os.Stdin)
				if readErr != nil {
					return fmt.Errorf("чтение stdin: %w", readErr)
				}

				// создадим парсер для stdin
				dataParser := parser.NewDataParser()
				cfg, err = dataParser.Run(data)
				if err != nil {
					return fmt.Errorf("не удалось распарсить stdin: %w", err)
				}

			} else {
				if len(args) == 0 {
					return fmt.Errorf("не указан путь к файлу")
				}
				path = args[0]

				// создадим парсер для файла
				fileParser := parser.NewFileParser()
				cfg, err = fileParser.Run(path)
				if err != nil {
					return fmt.Errorf("ошибка парсинга: %w", err)
				}

				if checkPermition {
					// добавим новое правило в анализатор
					analyzer.AddRule(&rules.FilePermissionRule{Path: path})
				}
			}

			totalIssues := analyzer.Run(cfg)

			if len(totalIssues) > 0 {

				for _, i := range totalIssues {
					fmt.Println(entities.GetIssueInfo(i))
				}

				if !silent {
					return fmt.Errorf("проблем обнаружено: %d", len(totalIssues))
				}

			} else {
				fmt.Println("Проблем не обнаружено")
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&silent, "silent", "s", false, "не завершаться с ошибкой при наличии проблем")
	cmd.Flags().BoolVar(&fromStdin, "stdin", false, "читать конфигурацию из stdin")
	cmd.Flags().BoolVar(&checkPermition, "chp", false, "проверять права доступа к файлу")

	// подключаем возможные команды
	cmd.AddCommand(dirCmd())
	cmd.AddCommand(serveCmd())

	return cmd
}
