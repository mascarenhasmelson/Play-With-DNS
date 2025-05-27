curl -v -X POST \
  -H "Content-Type: application/json" \
  -H "X-API-Key: not-a-secret" \
  -d '{
        "name": "0xmm.in.",
        "kind": "Native",
        "masters": [],
        "nameservers": ["ns1.0xmm.in.", "ns2.0xmm.in."]
      }' \
  http://localhost:8081/api/v1/servers/localhost/zones
