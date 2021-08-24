package common

import (
	"bufio"
	"bytes"
	"fmt"
	"fsclone/logs"
	"io"
	"os/exec"
	"time"
)

func ExecCommand2(strCommand string) ([]string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	var outText []string
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logs.GetLogger().Error(err)
		return outText, err
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return outText, err
	}
	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		outText = append(outText, string(output))
		if err == io.EOF {
			break
		}
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			return outText, err
		}
		fmt.Printf("%s\n", string(output))
	}
	return outText, err
}

func ExecCommand(strCommand string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func TimeStamp(fileTime string) (string, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(timeLayout, fileTime, loc)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	sr := theTime.Unix()
	return fmt.Sprint(sr), err
}

func BubbleSort(items []string) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] < items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}
