# Gandi Live DNS api client

## Generated using golang swagger 

Using swagger spec file [on github](https://github.com/t0k4rt/gandi-livedns-openapi)

## Generate the client from swagger definition file

Init submodules: `git submodules init`

then update to get latest livedns open api definition: `git submodule update`

At the root of the project, generate then client: 

`GOPATH=/Your/Go/Path swagger generate client -f gandi-livedns-openapi/gandi_livedns.yml`

The client will be generated.

## Example usage

You can use the client like this:

```golang
package main

import (
    "log"
    gandiApi "github.com/t0k4rt/gandi-livedns-go/client"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
    auth := httptransport.APIKeyAuth("X-Api-Key", "header", "api key"),
    client := gandiAPi.Default
    
    param := domains.NewGetDomainsDomainRecordsRecordNameRecordTypeParams()
	param.SetDomain(domain.Hostname())
	param.SetRecordName("@")
    param.SetRecordType("A")
    domainResp, err := client.Domains.GetDomainsDomainRecordsRecordNameRecordType(param, auth)

	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Printf("Got result with record name %s", domainResp.Payload.RrsetName)
}

```

Code generated is consistent with naming in the openapi file. There is no doc but you can read the code easily.


