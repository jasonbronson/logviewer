package transport

import (
	"github.com/gin-gonic/gin"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/pyama86/ptail/ptail"
)

func GetLog(g *gin.Context) {
	fileName := g.Param("filename")
	var logFile Log

	p := ptail.NewPtail("./logs/dir/" + fileName, 20)
	var content []string

	p.Use(func(l []byte) error {
		content = append(content, string(l))
		return nil
	})
	p.Execute()

	logFile = Log{
		Name: fileName,
		Content: content,
	}
	
	g.Writer.Header().Set("Content-Type", "application/json")
	g.JSON(http.StatusOK, logFile)
}

func GetLogs(g *gin.Context) {
	//read foler to get filename
	files, err := ioutil.ReadDir("./logs/dir")
    if err != nil {
        log.Fatal(err)
	}
	var fileNameArr Logs
    for _, file := range files {
		logFile := Log{
			Name: file.Name(),
		}
		fileNameArr.Logs = append(fileNameArr.Logs, logFile)
	}
	g.Writer.Header().Set("Content-Type", "application/json")
	g.JSON(http.StatusOK, fileNameArr)
}

type Logs struct {
	Logs	[]Log `json:"logs"`
}

type Log struct {
	Name	string	`json:"Name"`
	Content	[]string	`json:"content`
}
