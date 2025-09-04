package constant

const (
	Reply  = "reply"
	Inline = "inline"
)

const (
	DefaultDatePageLimit = 7
	DefaultTimePageLimit = 12
)

const (
	New    = "new"
	Change = "change"
)

const (
	GetUpdatesMethod        = "/getUpdates"
	SendMessageMethod       = "/sendMessage"
	EditMessageMethod       = "/editMessage"
	EditInlineMessageMethod = "/editMessageReplyMarkup"
	AnswerCallBackMethod    = "/answerCallbackQuery"
	ContentType             = "application/json"
	ConfigFilePath          = "app-config.yaml"
)

const (
	ChooseDateLabel      = "✨Выбери день для записи из списка"
	NextPageLabel        = "🔽 Дальше"
	PreviousCommandLabel = "◀️ Назад"
	ShowMyLessonsLabel   = "Посмотреть мои занятия 📅"
	RegToLessonLabel     = "Записаться на занятие ✍️"
	ChooseAction         = "Выбери действие👇"
)
