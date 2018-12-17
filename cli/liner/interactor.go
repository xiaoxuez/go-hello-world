package main

import (
	"fmt"
	"github.com/peterh/liner"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	state        *liner.State
	defaultModel liner.ModeApplier
	rawModel     liner.ModeApplier
	history_f    = filepath.Join(os.TempDir(), ".liner_example_history")
	command      = []string{"hello", "hi", "who are you", "what can u do"}
	support      = true
)

/**
  将历史记录的命令读入内存
*/
func beforeNew() {
	if state != nil {
		if content, err := ioutil.ReadFile(history_f); err == nil {
			state.ReadHistory(strings.NewReader(strings.Join(strings.Split(string(content), "\n"), "\n")))
		}
	}
}

func main() {
	beforeNew()
	newLiner()

	for {
		input, err := prompt("> ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				continue
			}
		}
		if input == "hello" {
			fmt.Println("hi")
		} else if input == "who are you" {
			fmt.Println("a robot")
		} else if input == "what can u do" {
			fmt.Println("emmmmmmm...  nothing")
		} else if input == "exit" {
			return
		}
		state.AppendHistory(input)
	}
	close()

}

func newLiner() {
	//当前模式 -- 默认模式
	defaultModel, _ = liner.TerminalMode()
	//newLiner,会返回一个State实例，并且会修改terminal的模式为rawModel
	state = liner.NewLiner()
	r, err := liner.TerminalMode()
	log.Printf("default Model is %v \n", defaultModel)
	log.Printf("after new liner, Model change to raw model: %v \n", rawModel)

	if err != nil {
		log.Printf("can not terminal Mode is %v:", err)
		support = false
	} else {
		rawModel = r
		defaultModel.ApplyMode()
	}
	state.SetCtrlCAborts(true)
	state.SetCompleter(complete)
}

func prompt(p string) (string, error) {
	if support {
		rawModel.ApplyMode()
		defer defaultModel.ApplyMode()
	} else {
		fmt.Print(p)
		p = ""
		defer fmt.Println()
	}
	return state.Prompt(p)
}

func close() {
	defaultModel.ApplyMode()
	state.Close()
}

func complete(line string) (s []string) {
	if line == "" {
		return command
	} else {
		for _, c := range command {
			if strings.HasPrefix(c, strings.ToLower(line)) {
				s = append(s, c)
			}
		}
	}
	return
}
