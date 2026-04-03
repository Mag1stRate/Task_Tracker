package Task_alter
import(
	"TASK_TRACKER/client"
	"os"
	"fmt"
	"bufio"
	"io"
	"TASK_TRACKER/model"
)

func ShowAllInProgress (){
	filePath := "F:/TASK_TRACKER/TASK.json"
	//以读写方式读取json文件
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil{
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	//创建一个读取器读取json文件
	reader := bufio.NewReader(file)
	//创建一个model.Task文件用于存储读出的结构体
	var alldata []model.Task
	//循环读取json文件,
	//将读取出的存入到切片中
	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			fmt.Println("读取结束")
			break
		}
		var cacheStruct *model.Task
		cacheStruct = client.UnMarshalStruct(line)
		alldata = append(alldata, *cacheStruct)
	}
	fmt.Println("所有正在进行的任务")
	for i := 0; i < len(alldata); i++{
		if alldata.Status[i] == "in_progress"{
			fmt.Println(alldata[i].Title)
		}
	}
}