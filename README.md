GoLang Basic Project Skelton

Running the server locally:
  - Clone the repo.
  - Set the GOPATH: export GOPATH=PATH_TO/summary/
  - Go to src dir: cd PATH_TO/src
  - glide install
  - Go to project root dir: CD PATH_TO
  - set env vars: source ./src/config/env.sh
  - Start App: go run src/main.go



Learning:
  - pkg manager (glide)
  - config
  - env config
  - Print API routing at server start
  - Logging struct as JSON string
  - REST API
  - Request Method restriction
  - DB configuration
  - CRUD API
  - ES configuration
  - ES basic ops (create, search, filter)
  - Docker deployment
  
Learning:
  - https://tour.golang.org
  - https://golang.org/doc/
  - https://godoc.org/github.com/spf13/viper
  - https://gowebexamples.com/
  - https://echo.labstack.com/guide
  - http://gorm.io/docs/
  - https://gopkg.in/go-playground/validator.v9
  - https://github.com/asaskevich/govalidator
  
CRUD API's:
  - ```javascript
    curl -X GET http://localhost:8090/user/1 
    ```
  - ```javascript
    curl -X POST \
      http://localhost:8090/user \
      -H 'content-type: application/json' \
      -d '{
    	"name": "Tom",
    	"team": "Sales",
    	"member": "RT",
    	"x": [1, 2, 3],
    	"y": {
    		"a": {"b": 1},
    		"c": [3, 4, 5]
    	}
    }'
    ``` 
     
