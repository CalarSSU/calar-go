package tracto

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ParseJson(result *Schedule, tokenForm, tokenDepartment, tokenGroup string) {
	rsp, err := http.Get(fmt.Sprintf(
		"%s/%s/%s/%s", scribaToken, tokenForm, tokenDepartment, tokenGroup))

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
