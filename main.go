package main
//
import (
// 	"github.com/joeig/go-powerdns/v3"
//   "encoding/json"
    "dns/routes"
    "net/http"
//   "fmt"
  "log"
//   // "math/rand"
// 	"time"
// 	"context"
)
//
//
// func main() {
//   // domain := fmt.Sprintf("%d.example.org.", rand.Int())
// domain:="example.org"
// // client := powerdns.NewClient("http://localhost:8081", "172.17.0.2", nil)
// pdns := powerdns.New("http://localhost:8081", "localhost", powerdns.WithAPIKey("changeme"))
// // client.APIKey = "changeme"
// ctx := context.Background()
//
// // Create the zone
// 	zone, err := pdns.Zones.AddNative(ctx, domain, false, "", false, "", "", true, []string{"localhost."})
// 	if err != nil {
//     if err.Error()=="Conflict" || err.Error() == "Not Found" {
//       log.Printf("Zone %s already exists or cannot be created, proceeding...\n", domain)
// 		} else {
// 			log.Fatalf("Failed to create zone: %v", err)
// 		}
// 	} else {
// 		o, _ := json.MarshalIndent(zone, "", "\t")
// 		log.Printf("Zone created: %s\n\n", o)
// 	}
//
//
//   // deffer()
//
// 	o, _ := json.MarshalIndent(zone, "", "\t")
// 	log.Printf("Zone created: %s\n\n", o)
//
// 	// Add an A record for www.{domain}
// 	recordName := fmt.Sprintf("test.%s", domain)
// 	// if err := pdns.Records.Add(ctx, domain, recordName, powerdns.RRTypeA, 3600, []string{"1.1.1.3"}); err != nil {
// 	// 	log.Fatalf("Error adding record: %v", err)
// 	// }
//
//   // Add a MX record with multiple values
// // if err := pdns.Records.Add(ctx, domain, domain, powerdns.RRTypeMX, 1337, []string{"10 mx1.example.com.", "20 mx2.example.com."}); err != nil {
// // 	log.Fatalf("%v", err)
// // }
//
// // Add a TXT record
// if err := pdns.Records.Add(ctx, domain, recordName, powerdns.RRTypeTXT, 1337, []string{"\"foo1\""}); err != nil {
// 	log.Fatalf("%v", err)
// }
//   comment := powerdns.Comment{
// 	Content:    powerdns.String("Example comment"),
// 	Account:    powerdns.String("example account"),
// 	ModifiedAt: powerdns.Uint64(uint64(time.Now().Unix())),
// }
//   if err := pdns.Records.Change(ctx, domain,recordName , powerdns.RRTypeA, 42, []string{"1.1.1.2"}, powerdns.WithComments(comment)); err != nil {
// 	log.Fatalf("%v", err)
// }
//
// 	// Retrieve and print the updated zone
// 	changedZone, err := pdns.Zones.Get(ctx, domain)
// 	if err != nil {
// 		log.Fatalf("Error getting zone: %v", err)
// 	}
//
// 	o, _ = json.MarshalIndent(changedZone, "", "\t")
// 	fmt.Printf("Changed zone: %s\n\n", o)
// fmt.Printf("Account is %q and DNSsec is %t\n", powerdns.StringValue(changedZone.Account), powerdns.BoolValue(changedZone.DNSsec))
// }


func main(){
  http.HandleFunc("/dns/record", routes.AddRecordHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
