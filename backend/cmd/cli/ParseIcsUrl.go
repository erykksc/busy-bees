package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	ical "github.com/arran4/golang-ical"
)

func main() {
	url := "https://p165-caldav.icloud.com/published/2/MjAzOTgyMDc3ODkyMDM5OIN_TWj1mjkOGkci4pOOV-UVCZvlevYEg0_8ep-72gLB2IFC2W4M_jtCCaXTAkCL1wVCISU5fgeqt6i-YcffZDI"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching calendar:", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	cal, err := ical.ParseCalendar(string(data))
	if err != nil {
		fmt.Println("Error parsing calendar:", err)
		return
	}

	for _, event := range cal.Events() {
		summary := event.GetProperty("SUMMARY").Value
		start := event.GetProperty("DTSTART").Value
		fmt.Printf("%s %s\n", summary, start)
	}

}

// func main(){
// 	cal, err := ics.ParseCalendar(os.Stdin)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error parsing calendar: %v\n", err)
// 		os.Exit(1)
// 	}
//
// 	for _, event := range cal.Events() {
// 		b, _ := json.MarshalIndent(event, "", "\t")
// 		fmt.Println(string(b))
// 	}
//
// 	cal := ics.NewCalendar()
// 	cal.SetMethod(ics.MethodRequest)
// 	event := cal.AddEvent("id@domain")
// 	event.SetCreatedTime(time.Now())
// 	event.SetDtStampTime(time.Now())
// 	event.SetModifiedAt(time.Now())
// 	event.SetStartAt(time.Now())
// 	event.SetEndAt(time.Now())
// 	event.SetSummary("Summary")
// 	event.SetLocation("Address")
// 	event.SetDescription("Description")
// 	event.SetURL("https://URL/")
// 	event.AddRrule(fmt.Sprintf("FREQ=YEARLY;BYMONTH=%d;BYMONTHDAY=%d", time.Now().Month(), time.Now().Day()))
// 	event.SetOrganizer("sender@domain", ics.WithCN("This Machine"))
// 	event.AddAttendee("reciever or participant", ics.CalendarUserTypeIndividual, ics.ParticipationStatusNeedsAction, ics.ParticipationRoleReqParticipant, ics.WithRSVP(true))
// 	fmt.Print(cal.Serialize())
// }
