package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	//logrusExample()
	//logPrintMsg()
	//logFatal()
	//logToFile()
	multiLogToFile()
}

func logrusExample() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
}

func myDive(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("被除数不能为0")
	}
	return x / y, nil
}

func logPrintMsg() {
	var x float64 = 128.2
	var y float64

	res, err := myDive(x, y)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(res)
}

func logFatal() {
	var x float64 = 128.2
	var y float64

	res, err := myDive(x, y)
	if err != nil {
		log.Fatal(err) // 在调用Print后就会调用 os.Exit(1)
	}
	// 不会输出了
	fmt.Println(res)
}

func logToFile() {
	var x float64 = 128.2
	var y float64

	res, exception := myDive(x, y)
	file, err := os.OpenFile("info1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Print(exception)
	log.Println(res)
}

func multiLogToFile() {
	taskNum := 10
	var wg sync.WaitGroup

	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		// init log file
		file, err := os.OpenFile("task-"+strconv.Itoa(i)+"-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("打开文件失败:", err)
		}
		defer file.Close()
		logger := log.New(file, "Logger:", log.Ldate|log.Ltime|log.Lshortfile)
		_, _, date := time.Now().Date()
		taskId := strconv.Itoa(i) + "-taskinfo-" + strconv.Itoa(date)
		// goruntine
		go taskContent(logger, taskId, &wg)
	}
	wg.Wait()
	fmt.Println("finished!!")
}

func taskContent(logger *log.Logger, taskId string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	time.Sleep(5000 * time.Millisecond)
	logger.Println("开始执行任务:", taskId)
	logger.Println("正在执行任务:", taskId, " .......")
	logger.Println("正在执行任务:", taskId, " .......")
	logger.Println("正在执行任务:", taskId, " .......")
	logger.Println("完成任务:", taskId)

}
