package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	var portAssign string

	if len(os.Args) >= 3 && (os.Args[1] == "-p" || os.Args[1] == "-P") {
		// main -p 8080
		// 尝试将第三个参数转化为整数类型
		port, err := strconv.Atoi(os.Args[2])
		if err != nil {
			port = 8080
			fmt.Println("Invalid Port \"%s\", use 8080 as default", os.Args[2])
		}
		portAssign = fmt.Sprintf(":%d", port)
	}

	r := chi.NewRouter()

	http.ListenAndServe(portAssign, r)
}
