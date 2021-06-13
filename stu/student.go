package stu

type Student struct {

	id int
	name string
	age int
}
/*type Stu struct {
	stuTea map[string] string
	stuCou map[string] string
}*/
func GetName(this *Student) string {
return this.name

}

func ShowTeachers(name string)(string,string){
	 stuTea:=make(map[string]func()(string,string))
	 stuTea["小明"]= func() (string, string) {
		 return "老刘","老王"
	 }
	stuTea["小红"]= func() (string, string) {
		return "老刘","老王"
	}
	stuTea["张三"]= func() (string, string) {
		return "老刘","老王"
	}
	stuTea["李四"]= func() (string, string) {
		return "老赵","老李"
	}
	stuTea["王五"]= func() (string, string) {
		return "老赵","老李"
	}
	stuTea["莉莉"]= func() (string, string) {
		return "老徐","老张"
	}
	stuTea["赵六"]= func() (string, string) {
		return "老徐","老张"
	}
	var a string
	var a1 string
	switch n := name; {
	case n=="小明":
		a, a1=stuTea["小明"]()
	case n=="小红":
		a,a1=stuTea["小红"]()
	case n=="张三":
		a,a1=stuTea["张三"]()
	case n=="李四":
		a,a1=stuTea["李四"]()
	case n=="王五":
		a,a1=stuTea["王五"]()
	case n=="赵六":
		a,a1=stuTea["赵六"]()
	case n=="莉莉":
		a,a1=stuTea["莉莉"]()

	}
	return a,a1
}
func ShowCourses(name string) (string,string) {
	 stuCou:=make(map[string]func()(string,string))
	 stuCou["小明"]= func() (string, string) {
		 return "语文","数学"
	 }
	stuCou["小红"]= func() (string, string) {
		return "语文","数学"
	}
	stuCou["张三"]= func() (string, string) {
		return "语文","数学"
	}
	stuCou["李四"]= func() (string, string) {
		return "语文","英语"
	}
	stuCou["王五"]= func() (string, string) {
		return "语文","英语"
	}
	stuCou["赵六"]= func() (string, string) {
		return "数学","英语"
	}
	stuCou["莉莉"]= func() (string, string) {
		return "数学","英语"
	}
	var a string
	var a1 string
	switch n := name; {
	case n=="小明":
		a,a1=stuCou["小明"]()
	case n=="小红":
		a,a1=stuCou["小红"]()
	case n=="张三":
		a,a1=stuCou["张三"]()
	case n=="李四":
		a,a1=stuCou["李四"]()
	case n=="王五":
		a,a1=stuCou["王五"]()
	case n=="赵六":
		a,a1=stuCou["赵六"]()
	case n=="莉莉":
		a,a1=stuCou["莉莉"]()

	}
	return a,a1

}
/*func Stud(name string) {
	l:=list.List{}
	l.PushFront(name)
	l.PushFront("小明")
	l.PushFront("小红")
	l.PushFront("张三")
	l.PushFront("李四")
	l.PushFront("王五")
	l.PushFront("赵六")
	l.PushFront("莉莉")
	l.PushFront("小刚")
}*/