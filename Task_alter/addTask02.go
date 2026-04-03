package Task_alter
import(
	"fmt"
	"os"
	"TASK_TRACKER/model"
	"time"
	"bufio"
	"TASK_TRACKER/client"
	"io"
)

func AddTask(){
	filePaht := "TASK.json"
	file, err := os.OpenFile(filePaht, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("操作出错",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var alldata []model.Task
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
	var idTheNew int
	for i := 0; i < len(alldata); i++{
		fmt.Println(alldata[i].ID)
		idTheNew = alldata[i].ID
	}
	fmt.Println("\t\t进入到添加任务界面")
	var id int
	var tatle string
	var status = "todo"
	var now = time.Now()
	var updatetime = time.Now()
	id = idTheNew + 1
	fmt.Println("请输入任务内容")
	fmt.Scanln(&tatle)
	tdmodel := model.Task{
		ID : id,
		Title : tatle,
		Status : status,
		CreatedAt : now,
		UpdateAt : updatetime,
	}
	data := client.MarshalStruct(tdmodel)
	writer := bufio.NewWriter(file)
	_, err = writer.Write (append(data, '\n'))
	if err != nil{
		fmt.Println("写入失败")
		return
	}
	
	writer.Flush()

}