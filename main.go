package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"sheremet-o/GO_spam_masker_v2.git/masker"
	"sheremet-o/GO_spam_masker_v2.git/service"
)

func main() {
	app := &cli.App{
		Name:  "spam-masker",
		Usage: "Удалить ссылки в тексте",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "",
				Usage: "Путь к исходному файлу",
			},
			&cli.StringFlag{
				Name:  "output",
				Value: "output.txt",
				Usage: "Путь к обработанному файлу",
			},
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",
				Usage: "Уровень логирования (trace, debug, info, warn, error, fatal, panic)",
			},
		},
		Action: func(c *cli.Context) error {
			filePath := c.String("file")
			if filePath == "" {
				fmt.Println("Укажите путь к файлу")
				return nil
			}

			logLevel, err := logrus.ParseLevel(c.String("log-level"))
			if err != nil {
				return err
			}
			logrus.SetLevel(logLevel)

			producer := &service.FileProducer{FilePath: filePath}
			presenter := &service.FileWriterPresenter{FilePath: c.String("output")}

			maskerService := masker.NewMaskingService(producer, presenter)
			err = maskerService.RunConcurrently()
			if err != nil {
				logrus.Errorf("Error running masking service: %v", err)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("Ошибка: %v", err)
	}
}
