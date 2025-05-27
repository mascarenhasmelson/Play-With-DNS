package routes

import(
"encoding/json"
"log"
powerdns "github.com/joeig/go-powerdns/v3"
"dns/model"
"fmt"
// "time"
"net/http"
"context"
)


 var domain="example.org"

var pdns = powerdns.New("http://localhost:8081", "localhost", powerdns.WithAPIKey("changeme"))

func AddRecordHandler(w http.ResponseWriter, r *http.Request){
  var RR model.RecordResponse
  ctx:=context.Background()

if err:=json.NewDecoder(r.Body).Decode(&RR);err!=nil{
    http.Error(w,"Invalid R body",http.StatusBadRequest)
    return

}

if RR.Subdomain == "" || RR.Type == "" || len(RR.Values) == 0 {
      http.Error(w, "Invalid record data", http.StatusBadRequest)
      return
  }
existingZone, err := pdns.Zones.List(ctx)
  if err != nil {
      log.Printf("Error fetching zones: %v", err)
      http.Error(w, "Failed to fetch zones", http.StatusInternalServerError)
      return
  }
  zoneExists := false
 for _, z := range existingZone {
     if *z.Name == domain+"." {
         zoneExists = true
         break
     }
 }
// var records []model.RecordResponse

// rrType := powerdns.RRType(RR.Type)
// fmt.Println(domain, RR.Subdomain, rrType, RR.TTL, RR.Values)
// for _,me:=range records{
//   fmt.Println("hello worked me",me)
// }
// zone, err := pdns.Zones.AddNative(ctx, domain, false, "", false, "", "", true, []string{"localhost."})
// if err != nil {
// 	log.Fatalf("%v", err)
// }

// o, _ := json.MarshalIndent(zone, "", "\t")
// log.Printf("Zone: %s\n\n", o)
if !zoneExists {
        zone, err := pdns.Zones.AddNative(ctx, domain, false, "", false, "", "", true, []string{"localhost."})
        if err != nil {
            log.Printf("Error creating zone: %v", err)
            http.Error(w, "Failed to create zone", http.StatusInternalServerError)
            return
        }
        log.Printf("Zone created: %s\n", zone.Name)
    }

    // subdomain := RR.Subdomain
    var subdomain string
     fmt.Println("RR.Subdomain",RR.Subdomain)
     fmt.Println("domain",domain)
    if !isFullyQualified(RR.Subdomain, domain) {
      subdomain = RR.Subdomain + "." + domain + "."
  } else {
      subdomain = RR.Subdomain
  }

    rrType := powerdns.RRType(RR.Type)
 //    subdomain := RR.Subdomain + "." + domain + "."
 // rrType := powerdns.RRType(RR.Type)
 //
 // subdomain := RR.Subdomain
 //   if !subdomainHasDot(subdomain) {
 //       subdomain += "." + domain + "."
 //   }

// if err := pdns.Records.Add(ctx, domain, RR.Subdomain, rrType, RR.TTL, RR.Values); err != nil {
//   fmt.Println(domain, RR.Subdomain, rrType, RR.TTL, RR.Values)
//       log.Printf("Error adding record: %v", err)
//   		http.Error(w, "Failed to add record", http.StatusInternalServerError)
//   		return
//   	}
//     w.WriteHeader(http.StatusCreated)
//     	w.Write([]byte("Record added successfully"))
err = pdns.Records.Add(ctx, domain+".",  subdomain, rrType, RR.TTL, RR.Values)

fmt.Println("printing im here ",domain, subdomain, rrType, RR.TTL, RR.Values)
    if err != nil {
        log.Printf("Error adding record: %v", err)
        http.Error(w, "Failed to add record", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Record added successfully...."))
}

func subdomainHasDot(sub string) bool {
    return len(sub) > 0 && sub[len(sub)-1] == '.'
}

func isFullyQualified(sub string, domain string) bool {
    return len(sub) > 0 && (sub[len(sub)-1] == '.' || sub == domain || sub == domain+".")
}
