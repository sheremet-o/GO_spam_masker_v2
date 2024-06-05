package main

import (
	"fmt"
	"os"
	"sheremet-o/GO_spam_masker_v2.git/masker"
	"sheremet-o/GO_spam_masker_v2.git/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите путь к файлу в аргументах запуска")
		return
	}

	filePath := os.Args[1]

	producer := &service.FileProducer{FilePath: filePath}
	presenter := &service.FileWriterPresenter{FilePath: "output.txt"}

	maskingService := masker.NewMaskingService(producer, presenter)
	maskingService.RunConcurrently()

}
