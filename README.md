## A simple URL Shortner in Go

[![goreport][goreport-img]][goreport]

Build this app to test my knowledge and challenge myself in Go. Fell free to suggest changes and improvements.

### Phase 1:
 - Implemented a local DB using maps in Go
 - It runs only on localhost and has not been configured to deploy anywhere else
 - It randomly assigns a six character short code to any url

### Sample Usage
#### To Shorten
```
curl localhost:8080/shorten?url=<URL_TO_SHORT>
```
#### Using the Shortened URL
Access `localhost:8080/<SHORTENED_URL>`
#### Show All Urls
```
curl localhost:8080/display
```

### TODOS
 - Deploy on kubernetes (AKS)
 - Use a distributed DB instead of Go maps
 - Add basic UI (Someday)
