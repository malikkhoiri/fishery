# fishery
> API Library Wrapper for [eFishery](https://efishery.com/en/). Helps to interact with external API ([Stein](https://docs.steinhq.com/introduction))

### Library Structure
```
.
├── README.md
├── fishery.go      
├── go.mod
├── go.sum
├── record.go
├── request.go
├── response.go
└── sheet.go
```
### Installation
```
go get -u github.com/malikkhoiri/fishery
```
### Usage
#### Initalize
```go
func main() {
	// Create new client
	client, err := fishery.NewClient(BASE_URL, API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	// Get working sheet
	sheet := client.GetSheet(fishery.List)
  
  ...
}
```
#### Add Records
```go
// Add new records
// records := &Records{}
res, err := sheet.AddRecords(records)
```
#### Get Record by ID
```go
// Get record by id
res, err := sheet.GetByID(id)
```
#### Get All Record
```go
// Get all record from sheet
res, err := sheet.GetAll()
```
#### Update Records
```go
// Update by id of record
// record := &Record{}
res, err := sheet.UpdateRecords(record)
```
#### Delete Records
```go
// Delete record by id
res, err := sheet.DeleteRecords(id)
```
#### Get All by Commudity
```go
// Get all record by commodity
res, err := sheet.GetAllByCommodity(commodity)
```
#### Get All by Area
```go
// Get all record by area
// area := &FilterArea{}
res, err := sheet.GetAllByArea(area)
```
### Get Max Price by Weeks
```go
// Get max price by weeks
res, err := sheet.GetAllByArea(week)
```
### Get Max Price by Commodity
```go
// Get max price by commodity
res, err := sheet.GetMaxPriceByCommudity(commodity)
```
