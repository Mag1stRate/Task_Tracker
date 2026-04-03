package Task_alter
import(
	"encoding/json"
	"TASK_TRACKER/model"
	"os"
	"bufio"
	"io"
	"fmt"

)

func UnMarshalStruct(data string)model.Task{
	var cacheStruct model.Task
	err := json.Unmarshal([]byte(data),&cacheStruct)
	if err != nil{
		fmt.Println("反序列化失败")
		return cacheStruct
	}
	return cacheStruct
}
func ShowTask(){
	filePath := "F:/TASK_TRACKER/TASK.json"
	file, err := os.Open(filePath)
	if err != nil{
		fmt.Println("文件打开失败")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println("\t\t显示所有任务")
	for {
		line, err := reader.ReadString('\n')
		var cacheStruct model.Task
		cacheStruct = UnMarshalStruct(line)
		fmt.Printf("任务ID\t%d\n",cacheStruct.ID)
		fmt.Printf("任务简报\t%v\n",cacheStruct.Title)
		fmt.Printf("任务状态\t%v\n",cacheStruct.Status)
		fmt.Printf("任务创建时间\t%v\n",cacheStruct.CreatedAt)
		fmt.Printf("任务上次更新时间\t%v\n",cacheStruct.UpdateAt)
		if err == io.EOF{
			break
		}
	}
}