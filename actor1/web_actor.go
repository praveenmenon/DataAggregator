package actor1

import (
	"encoding/json"
	"fmt"
	"github.com/praveenmenon/DataAggregator/models"
	"math/rand"
	"time"
)
var (
	Web1     = fakeSearch(models.Cars{Id: 1, Name: "Model X", Description: "This an suv", Range: 351, Battery: 88, AWD: true})
	Web2     = fakeSearch(models.Cars{Id: 2, Name: "Model Y", Description: "This is a compact sport utility vehicle", Range: 316, Battery: 68, AWD: true})
	Web3     = fakeSearch(models.Cars{Id: 3, Name: "Model S", Description: "This is a luxury vehicle", Range: 348, Battery: 68, AWD: true})
	jsonData []byte
)

// This function is trying to mock the search data and print it in pretty format
func fakeSearch(kind models.Cars) models.Search {
	return func(query string) models.Result {
		jsonData,_ := json.Marshal(kind)
		//I am sleeping the function for a random duration for better results
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return models.Result(fmt.Sprintf("Website results for %q from Web%d : %s\nWeb timestamp: %d\n", query, kind.Id, string(jsonData), time.Now().UnixNano() / int64(time.Millisecond)))
	}
}
