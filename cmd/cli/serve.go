package cli

import (
	"net/http"

	"github.com/configAnalyzer/internal/analyzer"
	"github.com/configAnalyzer/internal/parser"
	"github.com/configAnalyzer/internal/server/httpServer"
	"github.com/configAnalyzer/internal/server/httpServer/handlers"
	"github.com/spf13/cobra"
)

func serveCmd() *cobra.Command {

	var (
		httpAddr string
	)

	cmd := &cobra.Command{
		Use:          "serve <:port>",
		Short:        "Запустить HTTP сервер",
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) error {

			parser := parser.NewDataParser()
			analyzer := analyzer.NewAnalyzer()
			errCh := make(chan error, 1)

			// запускаем http сервер
			if httpAddr != "" {

				mux := http.NewServeMux()
				mux.HandleFunc("POST /analyze", handlers.NewAnalyzerHandler(parser, analyzer).Analyze())
				mux.HandleFunc("GET /health", handlers.NewHealthCheckHandler().CheckHealth())
				server := httpServer.NewHTTPServer(httpAddr, mux)

				go func() {
					errCh <- server.Run()
				}()
			}
			return <-errCh
		},
	}
	cmd.Flags().StringVar(&httpAddr, "http", "", "порт HTTP сервера")

	return cmd
}
