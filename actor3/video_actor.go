package actor3

import (
	"encoding/json"
	"fmt"
	"github.com/praveenmenon/DataAggregator/models"
	"math/rand"
	"time"
)

var (
	Video1 = fakeSearch(models.Video{Id: 1, Name: "Model S", Source: "youtube", VideoURL: "https://www.youtube.com/watch?v=ySMQ0Bfsrfg"})
	Video2 = fakeSearch(models.Video{Id: 2, Name: "Model X", Source: "dailymotion", VideoURL: "https://tinyurl.com/uqtouqp"})
	Video3 = fakeSearch(models.Video{Id: 3, Name: "Model S", Source: "metacafe", VideoURL: "https://tinyurl.com/tpmlde2"})
	jsonData []byte
)

func fakeSearch(kind models.Video) models.Search {
	return func(query string) models.Result {
		jsonData,_ := json.Marshal(kind)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return models.Result(fmt.Sprintf("Video results for %q from Video%d : %s\nVideo timestamp: %d\n", query, kind.Id, string(jsonData), time.Now().UnixNano() / int64(time.Millisecond)))
	}
}