package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
Create a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
  POST /zones/:zone_identifier/dns_records
*/
func (api *API) CreateDNSRecord(zone string, rr DNSRecord) error {
	z, err := api.ListZones(zone)
	if err != nil {
		return err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
	uri := "/zones/" + zid + "/dns_records"
	res, err := api.makeRequest("POST", uri, rr)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return nil
}

/*
Fetches DNS records for a zone.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
  GET /zones/:zone_identifier/dns_records
*/
func (api *API) DNSRecords(zone string, rr DNSRecord) ([]DNSRecord, error) {
	z, err := api.ListZones(zone)
	if err != nil {
		return []DNSRecord{}, err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID

	// Construct a query string
	v := url.Values{}
	if rr.Name != "" {
		v.Set("name", rr.Name)
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}
	var query string
	if len(v) > 0 {
		query = "?" + v.Encode()
	}
	uri := "/zones/" + zid + "/dns_records" + query
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []DNSRecord{}, err
	}
	var r DNSListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []DNSRecord{}, err
	}
	return r.Result, nil
}

/*
Fetches a single DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
  GET /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) DNSRecord(zone, id string) (DNSRecord, error) {
	z, err := api.ListZones(zone)
	if err != nil {
		return DNSRecord{}, err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
	uri := "/zones/" + zid + "/dns_records/" + id
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, err
	}
	return r.Result, nil
}

/*
Change a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
  PUT /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) UpdateDNSRecord(zone, id string, rr DNSRecord) error {
	z, err := api.ListZones(zone)
	if err != nil {
		return err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
	rec, err := api.DNSRecord(zone, id)
	if err != nil {
		return err
	}
	rr.Name = rec.Name
	rr.Type = rec.Type
	uri := "/zones/" + zid + "/dns_records/" + id
	res, err := api.makeRequest("PUT", uri, rr)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return err
}

/*
Delete a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
  DELETE /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) DeleteDNSRecord(zone, id string) error {
	z, err := api.ListZones(zone)
	if err != nil {
		return err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
	uri := "/zones/" + zid + "/dns_records/" + id
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return nil
}
