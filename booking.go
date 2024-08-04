package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

func init() {

}

const (
	selfMail = "obitoquilt@qq.com"

	temp = `
<table>
  <tr>
    <th>日期</th>
    <th>时间段</th>
    <th>费用</th>
    <th>总共</th>
    <th>剩余</th>
    <th>时间范围</th>
  </tr>
{{ range .Result }}
  <tr>
    <td>{{.Date}}</td>
    <td>{{.TimeName}}</td>
    <td>{{.TreatFee}}</td>
    <td>{{.RegTotalCount}}</td>
    <td>{{.RegLeaveCount}}</td>
    <td>{{.StartTime}}-{{.EndTime}}</td>
  </tr>
{{ end }}
</table>`
)

var (
	otherMail = []string{"1278413867@qq.com", "1054602988@qq.com", "3235992830@qq.com"}
)

func NewMailMessage(t *template.Template, items Items) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", selfMail)
	m.SetHeader("To", selfMail)
	m.SetHeader("Cc", otherMail...)
	m.SetHeader("Subject", fmt.Sprintf("已成功抢「%s」号，快去支付！！！", doctorName))

	var buf bytes.Buffer
	t.Execute(&buf, items)
	m.SetBody("text/html", buf.String())
	return m
}

func NewMailDialer() *gomail.Dialer {
	return gomail.NewDialer("smtp.qq.com", 587, selfMail, "qximpfhkwepadffi")
}

var (
	oneDay = true
	day    = "2024-08-08"

	//host       = "https://mobile.pku-hit.com"
	host      = "https://zocwbyypt.gzzoc.com"
	sessionId = "JSESSIONID=AAE52F6AF38609CEBE7975FE7BF75E2A"
	deptId    = 659 // 眼眶病
	deptName  = "眼眶病与眼肿瘤门诊"

	doctorCode = 2998 // 杨华胜
	doctorName = "杨华胜"

	//doctorName = "卢蓉"
	//doctorCode = 3260

	//doctorCode = 7312
	//doctorName = "肖伟"

	//doctorName = "毛羽翔"
	//doctorCode = 3004

	memberId  = "ABA2DCBBFDD511EC8856FA163E04CEFF"
	patientId = 5685273
	userJKK   = "L80800179"
)

func main() {
	t, err := template.New("health").Parse(temp)
	if err != nil {
		panic(err)
	}

	dialer := NewMailDialer()
	schedulePath := fmt.Sprintf("/MedicalMobile/client/register/doctor/schedule?deptId=%d&doctorCode=%d&patientId=&sourceDeptId=", deptId, doctorCode)
	scheduleURL := fmt.Sprintf("%s%s", host, schedulePath)
	for {
		scheduleDetails, err := doScheduleResult(scheduleURL)
		if err != nil {
			fmt.Println(err)
			continue
			//panic(err)
		}
		results := processScheduleDetails(scheduleDetails)
		//fmt.Printf("%+v", results)
		if len(results) == 0 {
			time.Sleep(5 * time.Second)
			continue
		}
		orderPath := "/MedicalMobile/client/register/order/save"
		orderURL := fmt.Sprintf("%s%s", host, orderPath)

		res := results[len(results)/2]
		err = doOrderResult(orderURL, res)
		if err == nil {
			if err = dialer.DialAndSend(NewMailMessage(t, Items{[]Result{res}})); err != nil {
				fmt.Printf("邮箱发送失败，错误：%v\n", err)
				continue
			}
			fmt.Println("已成功抢号，快去支付！！！")
			return
		}
		fmt.Println(err)
		for _, res := range results {
			err = doOrderResult(orderURL, res)
			if err != nil {
				//panic(err)
				fmt.Println(err)
				continue
			}
			if err = dialer.DialAndSend(NewMailMessage(t, Items{[]Result{res}})); err != nil {
				fmt.Printf("邮箱发送失败，错误：%v\n", err)
				continue
			}
			fmt.Println("已成功抢号，快去支付！！！")
			return
		}
		//time.Sleep(100 * time.Millisecond)
		time.Sleep(3 * time.Second)
	}
}

func doOrderResult(url string, result Result) error {
	client := &http.Client{}
	var req *http.Request
	data := fmt.Sprintf("memberId=%s&deptCode=%d&deptName=%s&"+
		"doctorCode=%d&doctorName=%s&date=%s&time=%s&timeName=%s&beginTime=%s&endTime=%s&patientId=%d&regType=1&fee=0.00&treatFee=%s&firstPatientTypeId=0&userJKK=%s",
		memberId, deptId, deptName, doctorCode, doctorName, result.Date, result.Time, result.TimeName, result.StartTime, result.EndTime, patientId, result.TreatFee, userJKK)
	//data := "memberId=ABA2DCBBFDD511EC8856FA163E04CEFF&deptCode=659&deptName=%E7%9C%BC%E7%9C%B6%E7%97%85%E4%B8%8E%E7%9C%BC%E8%82%BF%E7%98%A4%E9%97%A8%E8%AF%8A&doctorCode=7312&doctorName=%E8%82%96%E4%BC%9F&date=2023-05-09&time=1&timeName=%E4%B8%8B%E5%8D%88&beginTime=16%3A30&endTime=17%3A00&fee=0.00&treatFee=20.00&patientId=5685273&patientYLZ=&firstPatientTypeId=0&firstPatientTypeName=&oldPatientTypeId=1480&oldPatientTypeName=%E8%87%AA%E8%B4%B9&title=%E5%89%AF%E4%B8%BB%E4%BB%BB%E5%8C%BB%E5%B8%88&registerType=%E5%89%AF%E6%95%99%E6%8E%88&regFlag=false&userJKK=L80800179&canStudySample=1&regType=1"
	req, _ = http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Add("Cookie", sessionId)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/6.8.0(0x16080000) MacWechat/3.7.1(0x13070110) XWEB/30419 Flue")
	req.Header.Add("Referer", "http://mobile.pku-hit.com/MedicalMobile/dist/?")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("doOrderResult error", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("doOrderResult error", err)
		return err
	}
	if resp.StatusCode != 200 {
		fmt.Println("error not 200", resp.StatusCode)
		return errors.New("order response error not 200")
	}
	var ret OrderResult
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return err
	}
	if ret.Code != "0" {
		return errors.New("error, not 0")
	}
	return nil
}

func doScheduleResult(url string) (*ScheduleDetails, error) {
	client := &http.Client{}
	var req *http.Request
	req, _ = http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", sessionId)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 NetType/WIFI MicroMessenger/6.8.0(0x16080000) MacWechat/3.7.1(0x13070110) XWEB/30419 Flue")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("doScheduleResult error", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("doScheduleResult error", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		fmt.Println("error not 200", resp.StatusCode)
		return nil, errors.New("schedule response error not 200")
	}
	//fmt.Println(string(body))
	var scheduleDetails ScheduleDetails
	err = json.Unmarshal(body, &scheduleDetails)
	if err != nil {
		fmt.Println("doScheduleResult error", err)
		return nil, err
	}
	if scheduleDetails.Code != "0" {
		fmt.Println("doScheduleResult error, response code not 0")
		return nil, errors.New("response code not 0")
	}

	return &scheduleDetails, nil
}

func processScheduleDetails(details *ScheduleDetails) []Result {
	results := make([]Result, 0)
	for _, reg := range details.RegList {
		if reg.Morning != nil {
			res := processDetail(reg.Morning, reg.Date, "0", "上午")
			results = append(results, res...)
		}
		if reg.Afternoon != nil {
			res := processDetail(reg.Afternoon, reg.Date, "1", "下午")
			results = append(results, res...)
		}
	}

	return results
}

func processDetail(detail *Detail, date string, time, timeName string) []Result {
	if oneDay {
		if date != day {
			return nil
		}
	}
	results := make([]Result, 0)
	if detail.RegLeaveCount != "0" {
		for _, timesolt := range detail.TimesoltList {
			if timesolt.RegLeaveCount != "0" {
				results = append(results, Result{
					Date:          date,
					Time:          time,
					TimeName:      timeName,
					TreatFee:      detail.TreatFee,
					RegTotalCount: timesolt.RegTotalCount,
					RegLeaveCount: timesolt.RegLeaveCount,
					StartTime:     timesolt.StartTime,
					EndTime:       timesolt.EndTime,
				})
			}
		}
	}
	return results
}

type ScheduleDetails struct {
	Code       string
	Msg        string
	DoctorCode string
	RegDate    string
	RegList    []Reg `json:"regList"`
}

type Reg struct {
	Date      string
	Morning   *Detail `json:"morning"`
	Afternoon *Detail `json:"afternoon"`
}

type Detail struct {
	Date          string
	RegLeaveCount string
	IsTimeDiv     string // 0: morning, 1: afternoon
	TimeName      string
	TreatFee      string
	RegTotalCount string
	Time          string
	TimesoltList  []Timesolt `json:"timesoltList"`
}

type Timesolt struct {
	RegTotalCount      string
	RegLeaveCount      string
	StartTime, EndTime string
}

type Result struct {
	Date               string
	Time               string
	TimeName           string
	TreatFee           string
	RegTotalCount      string
	RegLeaveCount      string
	StartTime, EndTime string
}

type OrderResult struct {
	Code string
}

type Items struct {
	Result []Result `json:"result"`
}
