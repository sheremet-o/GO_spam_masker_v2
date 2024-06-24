package service

import (
	"bufio"
	"os"
	"strconv"
)

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type FileProducer struct {
	FilePath string
}

func (fp *FileProducer) Produce() ([]string, error) {
	// Чтение данных из файла и возврат в виде []string
	file, err := os.Open(fp.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

type FileWriterPresenter struct {
	FilePath string
}

func (fwp *FileWriterPresenter) Present(data []string) error {
	// Запись данных в файл
	file, err := os.Create(fwp.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, message := range data {
		file.WriteString(strconv.Itoa(i+1) + ". " + message + "\n")
	}
	return nil
}
