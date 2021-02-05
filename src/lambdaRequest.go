package src

import(
	"encoding/json"
	"../structs"
	"log"
)

func ParseReqBody(body []byte) (structs.StanRequest, error)  {
	var buf structs.StanRequest = structs.StanRequest{}
	err := json.Unmarshal([]byte(body), &buf)
	if err != nil {
		log.Println(err)
		return structs.StanRequest{}, err
	}
	return buf, nil
}