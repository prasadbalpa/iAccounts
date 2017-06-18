package main

import (
	"fmt"
	"iAccounts/utils"
)

func main() {
	var strs string
	strs = "org-34343-3434-3434-3322"

	fmt.Println(utils.FindAndReplace(strs, "-", "_"))

}
