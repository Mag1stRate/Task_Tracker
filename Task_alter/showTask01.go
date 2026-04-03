package Task_alter
import(
	"TASK_TRACKER/model"
	"TASK_TRACKER/client"
	"os"
	"bufio"
	"io"
	"fmt"

)


func ShowTask(){
	filePath := "F:/TASK_TRACKER/TASK.json"
	file, err := os.Open(filePath)
	if err != nil{
		fmt.Println("文件打开失败")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	fmt.Println("\t\t显示所有任务")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		var cacheStruct *model.Task
		cacheStruct = client.UnMarshalStruct(line)
		fmt.Printf("任务ID\t%d\n",cacheStruct.ID)
		fmt.Printf("任务简报\t%v\n",cacheStruct.Title)
		fmt.Printf("任务状态\t%v\n",cacheStruct.Status)
		fmt.Printf("任务创建时间\t%v\n",cacheStruct.CreatedAt)
		fmt.Printf("任务上次更新时间\t%v\n",cacheStruct.UpdateAt)
		
	}
}