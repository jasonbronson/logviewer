package transport

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/logreader/config"
)

func GetLogContent(g *gin.Context) {
	fileName := g.Param("filename")
	limit := g.Query("pageLimit")
	offset := g.Query("pageOffset")
	//var logFile Log
	//var res []string

	// get lastest log line
	// cmd := exec.Command("tail", "-n", string(limit), "./logs/"+fileName)
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Printf("cmd.Run() failed with %s\n", err)
	// 	return
	// }
	// content = strings.Split(string(out), "\n")

	limitLine, _ := strconv.Atoi(limit)
	pageOffset, _ := strconv.Atoi(offset)
	res := getLogData(fileName, "", limitLine, pageOffset)
	// if len(res) >= 20 {
	// 	res = res[len(res)-limitLine:]
	// }

	response(g, http.StatusOK, len(res), fileName, res)
}

func GetLogFiles(g *gin.Context) {
	//read folder to get filename
	var files []fs.FileInfo
	var err error
	if config.Cfg.LogDirectory != "" {
		files, err = ioutil.ReadDir(config.Cfg.LogDirectory)
	} else {
		files, err = ioutil.ReadDir("./logs/")
	}
	log.Println("Log files found ", len(files))
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
	offset := g.Query("pageOffset")
	searchKey := g.Query("searchKey")

	limitLine, _ := strconv.Atoi(limit)
	pageOffset, _ := strconv.Atoi(offset)
	res := getLogData(fileName, searchKey, limitLine, pageOffset)
	// if len(res) >= 20 {
	// 	res = res[len(res)-limitLine:]
	// }

	response(g, http.StatusOK, len(res), fileName, res)
}

func getLogData(fileName, searchKey string, pageLimit int, pageOffset int) []string {
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

	var pos, seekPosition int
	totalLines, _ := lineCounter(fileName)
	if pageOffset-totalLines > pageLimit {
		return nil
	} else if pageOffset > totalLines && pageOffset-totalLines < pageLimit {
		seekPosition = -1
		pageLimit = totalLines - pageOffset + pageLimit
	} else {
		seekPosition = totalLines - pageOffset
	}
	log.Println("seek", seekPosition, "count ", totalLines)
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
	cursor := 0
	for scanner.Scan() {
		results := strings.ToLower(scanner.Text())
		searchKey = strings.ToLower(searchKey)
		if len(searchKey) > 0 && strings.Contains(results, searchKey) {
			result = append(result, scanner.Text())
		}
		if len(searchKey) <= 0 {
			//result = append(result, LogLineFormat(scanner.Text(), "json"))
			result = append(result, scanner.Text())
		}
		if cursor >= pageLimit {
			break
		}
		cursor++
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

func response(g *gin.Context, code, totalrows int, fileName string, data []string) {
	g.Writer.Header().Set("Content-Type", "application/json")
	output := Log{
		Name:    fileName,
		Content: data,
		Total:   totalrows,
	}
	g.JSON(code, output)
}

type Logs struct {
	Logs []Log `json:"logs"`
}

type Log struct {
	Name    string   `json:"Name"`
	Content []string `json:"Content`
	Total   int      `json:"Total"`
}
