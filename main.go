package main

import (
	"TgBot/cmd/app/output/persist"
	"TgBot/cmd/app/output/persist/repository"
	"TgBot/cmd/app/output/rest"
	"TgBot/cmd/core/service/impl"
	"TgBot/cmd/core/tools"
	"context"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	err := container.Provide(tools.InitConfig)
	err = container.Provide(rest.InitConfig)
	err = container.Provide(rest.NewTelegramIntegrationService)
	err = container.Provide(persist.InitConfig)
	err = container.Provide(persist.InitDB)
	err = container.Provide(persist.NewTransactionManager)
	err = container.Provide(repository.NewBaseRepository)
	err = container.Provide(repository.NewOffsetRepository)
	err = container.Provide(repository.NewLessonRepository)
	err = container.Provide(repository.NewTeacherRepository)
	err = container.Provide(repository.NewStudentRepository)
	err = container.Provide(repository.NewPlaceholderRepository)
	err = container.Provide(impl.NewOffsetService)
	err = container.Provide(impl.NewTelegramService)
	err = container.Provide(impl.NewRequestHandler)
	err = container.Provide(impl.NewTelegramMessageConsumer)
	err = container.Provide(impl.NewResponseBuilder)
	err = container.Invoke(func(consumer *impl.TelegramMessageConsumer) {
		err = consumer.ConsumeMessages(context.Background())
	})
	if err != nil {
		panic(err)
	}
}
