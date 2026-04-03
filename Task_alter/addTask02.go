package Task_alter
import(
	"fmt"
	"os"
	"TASK_TRACKER/model"
	"time"
	"encoding/json"
	"bufio"
)

func AddTask(){
	filePaht := "F:/TASK_TRACKER/TASK.json"
	file, err := os.OpenFile(filePaht, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("操作出错",err)
		return
	}
	defer file.Close()
	fmt.Println("\t\t进入到添加任务界面")
	var id int
	var tatle string
	var status = "todo"
	var now = time.Now()
	var updatetime = time.Now()
	fmt.Println("请输入id")
	fmt.Scanln(&id)
	fmt.Println("请输入任务内容")
	fmt.Scanln(&tatle)
	tdmodel := model.Task{
		ID : id,
		Title : tatle,
		Status : status,
		CreatedAt : now,
		UpdateAt : updatetime,
	}
	data, err := json.Marshal(&tdmodel)
	if err != nil{
		fmt.Println("json化失败")
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write (append(data, '\n'))
	if err != nil{
		fmt.Println("写入失败")
		return
	}
	
	writer.Flush()

}