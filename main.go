package main

import (
	"awesome/cou"
	_ "awesome/person"
	"awesome/stu"
	_ "awesome/stu"
	"awesome/tea"
	"fmt"
)

func showMenu() {
	fmt.Println("-----------------------------------------------")
	fmt.Println(
		`
        1. 通过学生的姓名查找该学生的课程
        2. 通过学生的姓名查找该学生的老师
        3. 通过老师的姓名查找该老师的课程
        4. 通过老师的姓名查找该老师的学生
        5. 通过课程查找该课程学生
        6. 通过课程查找该课程老师`)



}
func main()  {

	/*type test struct {
		stu.Student
	}
	//var ss = new(test)
	//var name string
	//var a string

	fmt.Println(cou.ShowStu("语文"))
*/
	/*var a string
	fmt.Scanf("%S",a)*/
showMenu()
	var ss string

	//fmt.Scanf("%s",&ss)
	//fmt.Println(ss)
	var choice int
	for {
		// 等待用户输入
		fmt.Println("请选择要进行的操作：")
		_, _ = fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("请输入学生的姓名查找该学生的课程")
			fmt.Scanf("%s",&ss)
			fmt.Println(stu.ShowCourses(ss))
		case 2:
			fmt.Println("请输入学生的姓名查找该学生的老师")
			fmt.Scanf("%s",&ss)
			fmt.Println(stu.ShowTeachers(ss))
		case 3:
			fmt.Println("请输入老师的姓名查找该老师的课程")
			fmt.Scanf("%s",&ss)
			fmt.Println(tea.ShowCourse(ss))
		case 4:
			fmt.Println("请输入老师的姓名查找该老师的学生")
			fmt.Scanf("%s",&ss)
			fmt.Println(tea.ShowStudents(ss))
		case 5:
			fmt.Println("请输入课程查找该课程学生")
			fmt.Scanf("%s",&ss)
			fmt.Println(cou.ShowStu(ss))
		case 6:
			fmt.Println("请输入课程查找该课程老师")
			fmt.Scanf("%s",&ss)
			fmt.Println(cou.ShowTea(ss))
		}

	}

}
