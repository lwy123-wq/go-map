package cou

type Course struct {
	id int
	name string
}

type Cou struct {
	/*couStu map[string] [10]int
	couTea map[string] [10]int*/
}
//couTea1:=make(map[string]string,10)
func getName(this *Course) string {
	return this.name

}


func ShowStu(name string)(string,string,string){
	couStu:=make(map[string] func()(string,string,string) ,10)
	couStu["语文"]= func()(string,string,string){
		return "小明", "小红", "张三 "
		}
	couStu["数学"]= func()(string,string,string){
		return "小红", "张三 ","李四"
	}
	couStu["英语"]= func()(string,string,string){
		return "王五", "赵六 ","莉莉"
	}
	var j string
	var j1 string
	var j2 string
	switch n:=name;{
	case n =="语文":
		j, j1, j2 =couStu["语文"]()
	case n=="数学":
		j, j1, j2=couStu["数学"]()
	case n=="英语":

		j, j1, j2 =couStu["英语"]()
	}
	return j,j1,j2
	/*for k,v:=range couStu{
		fmt.Println("课程名字"+k+"学生姓名"+strconv.Itoa(v))
	}*/
		//fmt.Println("课程名字"+k+"学生姓名"+v)

}
func ShowTea(name string)(string,string){
	couTea:=make(map[string] func()(string,string),10)
	couTea["语文"]= func() (string, string) {
		return "老刘","老王"
	}
	couTea["数学"]= func() (string, string) {
		return "老赵","老李"
	}
	couTea["英语"]= func() (string, string) {
		return "老徐","老张"
	}
	var j string
	var j1 string
	switch n := name; {
	case n == "语文":
		j, j1 =couTea["语文"]()
	case n=="数学":
		j, j1=couTea["数学"]()
	case n=="英语":

		j, j1=couTea["英语"]()
	}
	return j,j1
}
/*func Couu(name string) list.List {
	 l:=list.List{
	 }
	 l.PushFront(name)
	return l
}*/