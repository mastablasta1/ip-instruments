package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

const (
	appToken = "12398uhoh97t8"
	url      = "http://ip-api.com/json/"
)

/*
проверка получаемых данных и отправка данных в обработку, обратно возвращаем чать данных полученных
*/
func GetIpData(w http.ResponseWriter, r *http.Request) {

	//	Enter data
	query := r.URL.Query()

	//	return struct
	var responceStruct ResponceStruct

	responceStruct.ErrorInfo = true
	responceStruct.ErrorName = "All good!"

	//	Токен приложения проверим
	if query.Get("app_token") == "" {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "app_token is zero"
		json.NewEncoder(w).Encode(responceStruct)
		return

	} else if appToken != query.Get("app_token") {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "app_token is not correct"
		json.NewEncoder(w).Encode(responceStruct)
		return

	}

	//	Проверим полчили мы IP или нет
	if query.Get("ip") == "" {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "ip is zero"
		json.NewEncoder(w).Encode(responceStruct)
		return

	}

	//	проверка валидности IP адреса
	ip := net.ParseIP(query.Get("ip"))
	if ip == nil {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "Incorrect IP"
		responceStruct.Query = query.Get("ip")
		json.NewEncoder(w).Encode(responceStruct)
		return

	}

	//	Получаем данные по IP
	resp, err := http.Get(url + query.Get("ip"))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Распаковка JSON-ответа в структуру
	err = json.Unmarshal(body, &responceStruct)
	if err != nil {
		fmt.Println("Ошибка при распаковке JSON:", err)
		return
	}

	//	Проверим что нам выдали.
	if responceStruct.Status == "fail" {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "Fail from API-service"
		json.NewEncoder(w).Encode(responceStruct)
		return

	} else if responceStruct.Query != query.Get("ip") {

		responceStruct.ErrorInfo = false
		responceStruct.ErrorName = "Request IP != Responce IP"
		json.NewEncoder(w).Encode(responceStruct)
		return

	}

	//	return data
	json.NewEncoder(w).Encode(responceStruct)
	return
}
