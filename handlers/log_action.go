package handlers

import (
	"github.com/labstack/echo"
	"github.com/thimunri/logtest/handlers/responses"
	"github.com/satori/go.uuid"
	"fmt"
	"os"
	"net/http"
	"time"
	"math/rand"
)

const MAX_MOCK_USERS = 50

type LogHandler struct {
	MockUsers map[int]string
	LogPath string
}

func (l *LogHandler) LogAction(c echo.Context) error {

	dirPath := fmt.Sprintf("%s/%s", l.LogPath, c.Request().URL.Path)
	logPath := fmt.Sprintf("%s/access.log", dirPath)
	CheckDir(dirPath)

	fmt.Println(logPath)
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ErrorMessage{
			Code:http.StatusInternalServerError,
			Message:err.Error()})
	}
	defer file.Close()

	currentDate := time.Now()
	userId := l.MockUsers[rand.Intn(MAX_MOCK_USERS)]
	logRow := fmt.Sprintf("%s - - [%s] \"%s %s HTTP/1.1\" 200 - \"userid=%s\"\n",
		c.Request().RemoteAddr,
		currentDate.Format("02/Jan/2006:15:04:05 -0700"),
		c.Request().Method,
		c.Request().URL.Path,
		userId,
	)

	file.WriteString(logRow)

	return c.String(http.StatusOK, "OK")
}

func (l *LogHandler) GenerateMockUsers() {
	mockUsers := make(map[int]string)
	for i:= 0; i< MAX_MOCK_USERS; i++ {
		uuid := uuid.NewV4()
		mockUsers[i] = fmt.Sprintf("%s", uuid)
	}
	l.MockUsers = mockUsers
}

func CheckDir(path string) {
	if _, error := os.Stat(path); os.IsNotExist(error) {
		os.Mkdir(path, 0555)
	}
}
