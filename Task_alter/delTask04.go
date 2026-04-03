package Task_alter
import(
	"TASK_TRACKER/client"
	"os"
	"fmt"
	"bufio"
	"io"
	"TASK_TRACKER/model"
)

func DelTask (){
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
		fmt.Println("请选择您要删除的ID")
		fmt.Scanln(&sel)
		if sel > len(alldata) - 1{
			fmt.Println("所选无效")
		}else{
			break
		}
	}
	// 删除下标为 sel 的元素
	alldata = append(alldata[:sel], alldata[sel+1:]...)
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