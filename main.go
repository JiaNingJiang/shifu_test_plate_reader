package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var rotaionCycle = 3 * time.Second

func main() {
	targetUrl := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
	//targetUrl := "http://deviceshifu-thermometer.deviceshifu.svc.cluster.local/read_value"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	for {

		fmt.Printf("Time Now: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		s := string(body)

		numbers := strings.Fields(s)

		var sum float64
		count := 0

		for _, numStr := range numbers {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Printf("Error converting %s to float: %v\n", numStr, err)
				continue
			}
			sum += num
			count++
		}

		fmt.Println("测试数据: \n", string(body))
		avg := sum / float64(count)
		fmt.Printf("Average of numbers: %.2f\n", avg)
		fmt.Println()

		res.Body.Close()
		time.Sleep(rotaionCycle)
	}
}
