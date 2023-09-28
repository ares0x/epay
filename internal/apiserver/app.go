package apiserver

import (
	"epay/internal/apiserver/config"
	"epay/internal/apiserver/options"
	"epay/pkg/app"
	"epay/pkg/log"
)

const commandDesc = `Link2Web3 Background API Server`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("link2web3-background",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
