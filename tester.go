package main

//
//import (
//	"fmt"
//	"io"
//	"net/http"
//)
//
//func main() {
//
//	url := "http://localhost:8080/videos"
//	method := "GET"
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, nil)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
//
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer func(Body io.ReadCloser) {
//		err := Body.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(res.Body)
//
//	body, err := io.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(body))
//}
