package client
import(
	"encoding/json"
	"TASK_TRACKER/model"
	"fmt"
)

func MarshalStruct(cachestruct model.Task)(data []byte){
	data, err := json.Marshal(cachestruct)
	if err != nil{
		fmt.Println("编译失败")
		return
	}
	return 

}