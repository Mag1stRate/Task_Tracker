package Task_alter
import(
	"fmt"
	"os"
	"bufio"
	"TASK_TRACKER/model"
	"io"
	"TASK_TRACKER/client"
	"time"
)

func Alter(){
	//获取存储的json文件地址
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
	fmt.Println("现存所有任务")
	for i := 0; i < len(alldata); i++{
		fmt.Println(alldata[i].Title)
	}
	var sel int
	for{
		fmt.Println("请选择您要更改状态的ID")
		fmt.Scanln(&sel)
		if sel > len(alldata) - 1{
			fmt.Println("所选无效")
		}else{
			break
		}
	}
	fmt.Println("请选择您要更改的状态")
	fmt.Println("0.todo;1.in_progress;2.done")
	var cacheStatus int
	for{
		var isCorect bool
		fmt.Scanln(&cacheStatus)
		switch cacheStatus{
		case 0:
			alldata[sel].Status = "todo"
			isCorect = true
		case 1:
			alldata[sel].Status = "in_progress"
			isCorect = true
		case 2:
			alldata[sel].Status = "done"
			isCorect = true
		default:
			fmt.Println("输入有误")
		}
		alldata[sel].UpdateAt = time.Now()
		if isCorect {
			break
		}
	}
	filewrite, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0666)
	defer filewrite.Close()
	for _, v := range alldata {
		writer := bufio.NewWriter(filewrite)
		data := client.MarshalStruct(v)
		_, err := writer.Write(append(data,'\n'))
		if err != nil{
			fmt.Println("写入失败")
		}
		writer.Flush()
	} 
	
}