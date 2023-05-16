// ready

package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

// writerHook is a hook that writes logs of specified LogLevels to specified Writer
// Эта секция определяет тип writerHook, который представляет собой крючок для логирования, 
// который записывает логи указанных LogLevels в указанные Writer.
type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
// Эти функции определяют реализацию крючка для логирования. Функция Fire форматирует 
// запись лога в строку и записывает ее в каждый Writer в списке Writer. 
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}
	return err
}

// Levels define on which log levels this hook would trigger
// Функция Levels возвращает список уровней логирования, на которых крючок будет вызываться.
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}


// Этот код определяет структуру Logger, которая содержит объект *logrus.Entry. 
// Функция GetLogger возвращает объект Logger, который содержит этот объект *logrus.Entry. 
// Функция GetLoggerWithField возвращает новый объект Logger, который содержит тот же объект *logrus.Entry, 
// но с добавленным полем k со значением v.
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{l.WithField(k, v)}
}

// Эта функция `Init` инициализирует объект логгера `logrus`. Она настраивает форматирование логов и указывает, 
// что логи должны быть отправлены в файл `logs/all.log` и на стандартный вывод, используя крючок `writerHook`. 
// Функция также устанавливает уровень логирования на `TraceLevel` и инициализирует объект `*logrus.Entry` для 
// использования в `Logger`.
// В целом, этот код предоставляет простой способ настройки и использования логирования в проектах на Golang, 
// который можно использовать для записи различных сообщений в файлы и/или на консоль.

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0755)

	if err != nil || os.IsExist(err) {
		panic("can't create log dir. no configured logging to files")
	} else {
		allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		if err != nil {
			panic(fmt.Sprintf("[Error]: %s", err))
		}

		l.SetOutput(os.Stdout) // Send all logs to nowhere by default

		l.AddHook(&writerHook{
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		})
	}

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}