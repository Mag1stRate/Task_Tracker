/*
本项目目标实现一个任务跟踪器
仅要求在命令行界面实现:
1.添加,更新和删除任务
2.列出所有已经完成的任务
3.列出未完成的任务
4.列出所有正在进行的任务
*/
package main
import(
	"fmt"
	"TASK_TRACKER/Surface"
	"TASK_TRACKER/Task_alter"
)

func main(){
	for{
	//实现界面
	var key int
	key = mainface.MainFace()
	var loop bool
	//以上,用户在主界面做出选择
	switch key{
	case 1:
		Task_alter.ShowTask()
	case 2:
		Task_alter.AddTask()
	case 3:
		Task_alter.Alter()
	case 4:
		Task_alter.DelTask()
	case 5:
		Task_alter.ShowAllNotDone()
	case 6:
	case 7:
		fmt.Println("您确定要退出吗Y/N")
		var exit rune
		fmt.Scanf("%c",&exit)
		if exit == 'Y'{
			loop = true
		}
	default:
		fmt.Println("输入有误,请重新输入")
	}
	if loop {
		fmt.Println("程序退出")
		break
	}
	}
	
}