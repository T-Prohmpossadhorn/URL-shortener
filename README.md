# URL-shortener
A super fast URL shortener and redirect tool written in Go

### Prerequisites

- Docker
- docker-compose

### Running the tool

```sh
docker-compose up -d
```

### Function

- API Client can send a url and be returned a shortened URL.
- API Client can specify an expiration time for URLs, expired URLs must return HTTP 410
- Input URL should be validated and respond with error if not a valid URL
- Visiting the Shortened URLs must redirect to the original URL with a HTTP 302 redirect, 404 if not found.
- Hit counter for shortened URLs (increment with every hit)
- Admin api (requiring token) to list
    - Short Code
    - Full Url
    - Expiry (if any)
    - Number of hits
- Above list can filter by Short Code
- Admin api to delete a URL (after deletion shortened URLs must return HTTP 410)

### Using the tool

- get the token by using the post method
  `http://localhost:5000/login`
  use an email, password in the config.json and add to the body
  <img width="1015" alt="Screen Shot 2564-08-27 at 12 51 53" src="https://user-images.githubusercontent.com/49223359/131078903-f670355a-c66b-4124-90b2-19a9098fa708.png">
- test the received token
  `http://localhost:5000/admin/test`
  <img width="1007" alt="Screen Shot 2564-08-27 at 12 57 52" src="https://user-images.githubusercontent.com/49223359/131079320-8a8dbcc8-55d0-4ab9-b817-bb6d286314bb.png">
- post the url that wants to be shorten using the recieved token
  `http://localhost:5000/admin`
  <img width="1018" alt="Screen Shot 2564-08-27 at 13 02 12" src="https://user-images.githubusercontent.com/49223359/131079667-0bb59046-f207-4c34-b731-74b991b96950.png">
  the expire can be blank if the shortened url will not expire, if not use this format:`Mon, 02 Jan 2006 15:04:05 MST`
- get the shorten url info using received token
  `http://localhost:5000/admin/:shortlink`
  <img width="1000" alt="Screen Shot 2564-08-27 at 13 08 10" src="https://user-images.githubusercontent.com/49223359/131080288-4464a8ae-ab1f-4cc6-9f70-a2941f5bb62b.png">
- delete the shorten url by changing the method from get to delete
  `http://localhost:5000/admin/:shortlink`
  <img width="1010" alt="Screen Shot 2564-08-27 at 13 16 06" src="https://user-images.githubusercontent.com/49223359/131081056-bc19ee5a-4b96-47db-96a5-0108de134b22.png">
  
### Using the shorten link
- use the url provided by url which received when posting the shorten url (does not require the admin token)
  `http://localhost:5000/:shortlink`
  <img width="1009" alt="Screen Shot 2564-08-27 at 13 21 27" src="https://user-images.githubusercontent.com/49223359/131081581-cda32058-a554-4f2b-8daa-5a31025eb3d1.png">



  

