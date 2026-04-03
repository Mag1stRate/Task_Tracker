//此包用于表现任务主界面
package mainface
import(
	"fmt"
)

func MainFace()(key int){
	fmt.Println("\t\t任务跟踪器系统")
	fmt.Println("\t\t1.显示所有任务")
	fmt.Println("\t\t2.添加任务")
	fmt.Println("\t\t3.更新任务状态")
	fmt.Println("\t\t4.删除任务")
	fmt.Println("\t\t5.显示所有未完成任务")
	fmt.Println("\t\t6.显示所有正在进行的任务")
	fmt.Println("\t\t7.退出")
	fmt.Println("请输入(1-7):")
	fmt.Scanln(&key)
	return
}

func No1Face(){
	fmt.Println()
}