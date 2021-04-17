

# Go Homework

Go homework for IBM.

## Prerequisites
- Go 1.15.7 or above

## Usage

- Base URL: http://localhost:8080

**Get Epoch**
----
  Returns latest Epoch timestamp.

* **URL**

  /api/v1/getEpoch


* **Method:**

  `GET`
  
* **URL Params**

  None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `1618662464`
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `No data available.`

**Save Epoch**
----
  Saves a valid epoch timestamp.

* **URL**

  /api/v1/saveEpoch


* **Method:**

  `POST`
  
* **URL Params**

  None

* **Data Params**

  * **Raw Body must be a valid Unix timestamp:** <br />
    **Example:** `1618662464`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
 
* **Error Response:**

  * **Code:** 404 NOT FOUND <br />
    **Content:** `No data available.`
    
    OR
    
  * **Code:** 500 INTERNAL SERVER ERROR <br />
    **Content:** `strconv.ParseInt: parsing "WRONGINPUT": invalid syntax\n`




### Using the application
By running the application with:
```
go run main.go
```

### Running the tests
By running the application with:
```
go test
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
