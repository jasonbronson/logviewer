package transport

import (
	"github.com/gin-gonic/gin"
	"log"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func GetLogContent(g *gin.Context) {
	fileName := g.Param("filename")
	limit := g.Query("pageLimit")
	var logFile Log
	var content []string

	// get lastest log line
	cmd := exec.Command("tail", "-n", string(limit), "./logs/dir/" + fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
	   log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	content = strings.Split(string(out), "\n")

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

func SearchLog(g *gin.Context) {
	fileName := g.Param("filename")
	limit := g.Query("pageLimit")
	key := g.Query("searchKey")

	file, err := os.Open("./logs/dir/" + fileName)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	var result []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			result = append(result, scanner.Text())
		}
        // fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	limitLine, _ := strconv.Atoi(limit)

	res := result
	if (len(result) >= 20) {
		res = result[len(result) - limitLine:]
	}
	
	g.JSON(http.StatusOK, res)
}

type Logs struct {
	Logs	[]Log `json:"logs"`
}

type Log struct {
	Name		string		`json:"Name"`
	Content		[]string	`json:"content`
}
