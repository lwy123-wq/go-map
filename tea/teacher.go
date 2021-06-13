package tea

import (
	"awesome/cou"
)

type Teacher struct {
	name   string
	course cou.Course
	id     int
}
func getName(this *Teacher) string {
	return this.name

}
/*type Tea struct {
	teaStu map[string] string
l.PushFront("小明")
}*/

func ShowStudents(name string) (string,string,string) {
	 teaStu:=make(map[string]func()(string,string,string))
	 teaStu["老刘"]= func() (string, string, string) {
		 return "小明","小红","张三"
	 }
	teaStu["老王"]= func() (string, string, string) {
		return "小明","小红","张三"
	}
	teaStu["老赵"]= func() (string, string, string) {
		return "张三","李四","王五"
	}
	teaStu["老李"]= func() (string, string, string) {
		return "李四","王五","赵六"
	}
	teaStu["老徐"]= func() (string, string, string) {
		return "王五","赵六","莉莉"
	}
	teaStu["老张"]= func() (string, string, string) {
		return "王五","赵六","莉莉"
	}
	var b string
	var b1 string
	var b2 string
	switch n:=name ;{
	case n=="老刘":
		b,b1,b2=teaStu["老刘"]()
	case n=="老王":
		b,b1,b2=teaStu["老王"]()
	case n=="老赵":
		b,b1,b2=teaStu["老赵"]()
	case n=="老李":
		b,b1,b2=teaStu["老李"]()
	case n=="老徐":
		b,b1,b2=teaStu["老徐"]()
	case n=="老张":
		b,b1,b2=teaStu["老张"]()
	}

	return b,b1,b2
}
func ShowCourse(name string) string {
	teaCou:=make(map[string]string)
	teaCou["老刘"]="语文"
	teaCou["老王"]="语文"
	teaCou["老赵"]="数学"
	teaCou["老李"]="数学"
	teaCou["老徐"]="英语"
	teaCou["老张"]="英语"
	var a string
	switch n := name; {
	case n=="老刘":
		a=teaCou["老刘"]
	case n=="老王":
		a=teaCou["老王"]
	case n=="老赵":
		a=teaCou["老赵"]
	case n=="老李":
		a=teaCou["老李"]
	case n=="老徐":
		a=teaCou["老徐"]
	case n=="老张":
		a=teaCou["老张"]
	}
	return a

}
/*func Teaa(name string)  {
	l:=list.List{}
	l.PushFront(name)
}*/