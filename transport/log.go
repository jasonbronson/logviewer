package transport

import (
	"github.com/gin-gonic/gin"
	"log"
	"io/ioutil"
	"net/http"
)

func GetLogs(g *gin.Context) {
	files, err := ioutil.ReadDir("./frontend/dist")
    if err != nil {
        log.Fatal(err)
	}
	
	var fileNameArr Logs

    for _, file := range files {
		log := Log{
			Name: file.Name(),
		}
		fileNameArr.Logs = append(fileNameArr.Logs, log)
    }
	g.Writer.Header().Set("Content-Type", "application/json")
	g.JSON(http.StatusOK, fileNameArr)
}

type Logs struct {
	Logs	[]Log `json:"logs"`
}

type Log struct {
	Name	string	`json:"Name"`
}
