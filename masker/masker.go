package masker

import (
	"sheremet-o/GO_spam_masker_v2.git/service"
	"sync"

	"github.com/sirupsen/logrus"
)

type MaskingService struct {
	producer  service.Producer
	presenter service.Presenter
}

func NewMaskingService(producer service.Producer, presenter service.Presenter) *MaskingService {
	return &MaskingService{
		producer:  producer,
		presenter: presenter,
	}
}

func (ms *MaskingService) RunConcurrently() error {
	data, err := ms.producer.Produce()
	if err != nil {
		return err
	}

	ch := make(chan string)
	var wg sync.WaitGroup
	limit := 10
	semaphore := make(chan struct{}, limit)

	for _, message := range data {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(msg string) {
			defer wg.Done()
			maskedMessage := ms.Masker(msg)
			ch <- maskedMessage
			<-semaphore
		}(message)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var resultData []string
	for maskedMessage := range ch {
		resultData = append(resultData, maskedMessage)
	}

	err = ms.presenter.Present(resultData)
	if err != nil {
		logrus.Errorf("Ошибка представленных данных: %v", err)
		return err
	}

	return nil
}

func (ms *MaskingService) Masker(message string) string {
	buffer := []byte(message)
	linkHttp := []byte("http://")

	for i := 0; i < len(buffer)-len(linkHttp); i++ {
		if string(buffer[i:i+len(linkHttp)]) == string(linkHttp) {
			j := i + len(linkHttp)
			for j < len(buffer) && buffer[j] != ' ' {
				buffer[j] = '*'
				j++
			}
			i = j
		}
	}
	return string(buffer)
}
