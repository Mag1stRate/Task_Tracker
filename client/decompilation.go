package client
import(
	"encoding/json"
	"TASK_TRACKER/model"
	"fmt"
)

func UnMarshalStruct(data string)*model.Task{
	var cacheStruct model.Task
	err := json.Unmarshal([]byte(data),&cacheStruct)
	if err != nil{
		fmt.Println("反序列化失败",err)
		return &cacheStruct
	}
	return &cacheStruct
}