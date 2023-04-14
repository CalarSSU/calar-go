package tracto

import (
	"calar-go/parser"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseJson(result *Schedule, request parser.Request) {
	rsp, err := http.Get(fmt.Sprintf(
		"%s/%s/%s/%s", scribaToken, request.Education, request.Department,
		request.Group))

	if err != nil {
		log.Fatalln(err)
	}
	jsonString, err := io.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(jsonString, &result)
	if err != nil {
		fmt.Println(err)
	}

}
