package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlages struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlag() *cmdFlages {
	cf := cmdFlages{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo item")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo item by index")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo item by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo item by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo items")

	flag.Parse()
	return &cf
}

func (cf *cmdFlages) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid formate for edit, please use id:new_tittle")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error, invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("No valid command provided")

	}
}
