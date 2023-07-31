package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	// if start.Before(end) {
	// 	return !check.Before(start) && !check.After(end)
	// }
	// if start.Equal(end) {
	// 	return check.Equal(start)
	// }
	// return !start.After(check) || !end.Before(check)

	if check.Before(start) {
		return false
	}
	if check.After(end) {
		return false
	}
	return true
}

func main() {
	url := "http://192.168.20.24:9001/web_services/CompanyDetails.php?docid=080PXX80.XX80.150408181702.V5K5"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	type companystruct struct {
		Hop string
	}

	type daytimings struct {
		Starttime time.Time
		Endtime   time.Time
	}

	isvalid := json.Valid(body)
	if isvalid {
		fmt.Println("proper json------")
		// json.Unmarshal(body, &newres)
		// fmt.Printf("%#v\n", newres)
		var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)
		//json.Unmarshal(body, &result)
		p := result["080PXX80.XX80.150408181702.V5K5"]
		//m := p.(map[string]companystruct)
		m := p.(map[string]interface{})
		hopres := m["hop"]
		hopday := hopres.(map[string]interface{})
		hopdayres := hopday["val"]
		fmt.Println(hopdayres)
		curday := (time.Time.Weekday(time.Now())).String()
		curday = strings.ToLower(curday)
		curday = curday[0:3]
		fmt.Println(curday)
		type k map[string]interface{}
		hopdayfilter := hopdayres.(map[string]interface{})
		//hopcurday := hopdayfilter[curday]
		hopcurday := hopdayfilter["sat"]
		count := hopcurday.([]interface{})
		countitr := count[0].([]interface{})
		count2 := make([]string, len(countitr))
		for i, v := range countitr {
			count2[i] = v.(string)
		}
		starttime := count2[0]
		enditme := count2[1]
		newformat := "15:04"
		curtime := time.Now().Format("15:04")
		check, _ := time.Parse(newformat, curtime)
		start, _ := time.Parse(newformat, starttime)
		end, _ := time.Parse(newformat, enditme)
		// start, _ := time.Parse(newLayout, t.start)
		log.Println(starttime, enditme)
		log.Println("curent time")
		log.Println(check)

		output := inTimeSpan(start, end, check)
		fmt.Println(output)
	}
}
