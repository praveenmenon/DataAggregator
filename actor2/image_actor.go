package actor2

import (
	"encoding/json"
	"fmt"
	"github.com/praveenmenon/DataAggregator/models"
	"math/rand"
	"time"
)

var (
	Image1 = fakeSearch(models.Thumbnail{Id: 1, Name: "Model X", CarImageUrl: []string{"https://tinyurl.com/sccs8zk", "https://tinyurl.com/rz8242y"}})
	Image2 = fakeSearch(models.Thumbnail{Id: 2, Name: "Model Y", CarImageUrl: []string{"https://tinyurl.com/uc438zj", "https://tinyurl.com/urk58y3"}})
	Image3 = fakeSearch(models.Thumbnail{Id: 3, Name: "Model S", CarImageUrl: []string{"https://tinyurl.com/s6wbu7c", "https://tinyurl.com/qwllp8p"}})
	jsonData []byte
)

func fakeSearch(kind models.Thumbnail) models.Search {
	return func(query string) models.Result {
		jsonData,_ := json.Marshal(kind)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return models.Result(fmt.Sprintf("Image results for %q from Image%d : %s\nImage timestamp: %d\n", query, kind.Id, string(jsonData), time.Now().UnixNano() / int64(time.Millisecond)))
	}
}

