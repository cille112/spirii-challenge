# Spirii Challenge

## Assumptions:
- Negative data is not valid and is therefore skipped
- Duplicate data is also skipped, since we already have it once in the database
- Duplicate data is when the same meterId sends data with the same timestamp.
- In this setup a consumer can have multiple meters, but different consumers can have meters with the same id as other consumers. (Just to mae things a little easier)
- If there is no meter-id we assume that it is not valid and is skipped


## To run it
Make sure you have go installed and SQLite.

Run `go mod tidy` to install packages used
Run `go run main.go` to run the code.

To access API documentation go to [Swagger](http://localhost:8000/swagger/index.html)

## Use API
The api is secured with basic authentication (the username and password is hardcoded) use `admin` as username and `password` as password when calling the api. If I were to create this on a real api I would not use predefined credentials like this, they shouldn't be hardcoded. Could consider using token based like JWT or OAuth. 

## Data stream
I have created a data stream that at random times adds a new reading from a meter. These readings are created using the rand package. A predefined list of meterId and consumerId is used to pick a random meter and consumer, in the meterId list there is an empty Id, which is to create anomalies. 

To create anomalies I sometimes send the same data twice and also sometimes make the reading from the meter negative. 

## Storage
I decided to go with sqlite because it doesn't require a lot of setup. There is no need for a relational DB, since we are only storing data in one table and there are not any relationships in the DB. 

## Test
To run the test `go test -v ./...`

## 30% consumer
I know the way I calculate this is not entirely correct since I just add consumptions until I reach 30% or more, but due to time limitations I decided to go with this approach. 