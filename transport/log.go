package transport

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLogContent(g *gin.Context) {
	fileName := g.Param("filename")
	limit := g.Query("pageLimit")
	var logFile Log
	//var res []string

	// get lastest log line
	// cmd := exec.Command("tail", "-n", string(limit), "./logs/"+fileName)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Printf("cmd.Run() failed with %s\n", err)
	// 	return
	// }
	// content = strings.Split(string(out), "\n")

	log.Print(limit)
	limitLine, _ := strconv.Atoi(limit)
	res := getLogData(fileName, "", true, limitLine)
	// if len(res) >= 20 {
	// 	res = res[len(res)-limitLine:]
	// }
	logFile = Log{
		Name:    fileName,
		Content: res,
	}

	g.Writer.Header().Set("Content-Type", "application/json")
	g.JSON(http.StatusOK, logFile)
}

func GetLogs(g *gin.Context) {
	//read folder to get filename
	files, err := ioutil.ReadDir("./logs/")
	if err != nil {
		log.Fatal(err)
	}
	var logs Logs
	for _, file := range files {
		logFile := Log{
			Name: file.Name(),
		}
		logs.Logs = append(logs.Logs, logFile)
	}
	g.Writer.Header().Set("Content-Type", "application/json")
	g.JSON(http.StatusOK, logs)
}

func SearchLog(g *gin.Context) {
	fileName := g.Param("filename")
	limit := g.Query("pageLimit")
	searchKey := g.Query("searchKey")

	limitLine, _ := strconv.Atoi(limit)
	res := getLogData(fileName, searchKey, false, limitLine)
	if len(res) >= 20 {
		res = res[len(res)-limitLine:]
	}

	g.JSON(http.StatusOK, res)
}

func getLogData(fileName, searchKey string, last1K bool, pageLimit int) []string {
	file, err := os.Open("./logs/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var result []string
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if last1K {
		var pos int
		totalLines, _ := lineCounter(fileName)
		seekPosition := totalLines - pageLimit

		scanner.Split(func(data []byte, atEof bool) (advance int, token []byte, err error) {
			advance, token, err = bufio.ScanLines(data, atEof)
			pos += advance
			return
		})

		for i := 0; i <= seekPosition; i++ {
			if !scanner.Scan() {
				log.Println("EOF")
			}
		}
	}
	for scanner.Scan() {
		if len(searchKey) > 0 && strings.Contains(scanner.Text(), searchKey) {
			result = append(result, scanner.Text())
		} 
		if len(searchKey) <= 0 {
			//result = append(result, LogLineFormat(scanner.Text(), "json"))
			result = append(result, scanner.Text())
		}
	}
	return result
}

func LogLineFormat(line string, formatType string) string {
	var result string
	if formatType == "json" {
		out, _ := json.Marshal(line)
		result = string(out)
	} else {
		result = line
	}
	return result
}

func lineCounter(fileName string) (int, error) {
	r, err := os.Open("./logs/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

type Logs struct {
	Logs []Log `json:"logs"`
}

type Log struct {
	Name    string   `json:"Name"`
	Content []string `json:"content`
}
