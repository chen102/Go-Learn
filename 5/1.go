package main

import "fmt"
import "errors"
import "strconv"
import "regexp"
import "encoding/json"
import "os"
import "io"
import "bufio"

//erro接口的使用
func test01() {
	err1 := fmt.Errorf("%s", "waring")
	fmt.Println(err1)
	err2 := errors.New("WARING")
	fmt.Println(err2)
	result, err := test01_1(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
func test01_1(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return

}

//panic(致命错误处理)
func test02() {
	fmt.Println("ohhhhhh")
	//panic("panic test")
	fmt.Println("ohhhhh")
	//test02_1(55) //数组越界,自动掉panic函数
}
func test02_1(x int) {
	var a [10]int
	a[x] = 222
}

//recover恢复程序
func test03() {
	fmt.Println("ohhhhhh")
	test03_1(55) //但是看不到panic的错误信息了
	fmt.Println("ohhhhh")
	test03_2(5) //但是如果没有错，也会显示错误信息<nil>
	fmt.Println("ohhhhh")
	test03_3(5) //加个判断
	fmt.Println("ohhhhh")

}
func test03_1(x int) {
	defer func() {
		recover()
	}()
	var a [10]int
	a[x] = 222
}
func test03_2(x int) {
	defer func() {
		fmt.Println(recover())
	}()
	var a [10]int
	a[x] = 222
}
func test03_3(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(recover())
		}
	}()
	var a [10]int
	a[x] = 222
}

//文本和文件处理-字符串转换
func test04() {
	//转换为字符串后追加到字节数组
	s := make([]byte, 0, 1024)
	s = strconv.AppendBool(s, true)
	s = strconv.AppendInt(s, 1234, 10)
	s = strconv.AppendQuote(s, "cdajkghb")
	fmt.Println(string(s))
	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println(str)
	//整型转字符串
	str = strconv.Itoa(642642)
	fmt.Println(str)
	//字符串转其他类型
	var err error
	var flag bool
	flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println(flag)
	} else {
		fmt.Println(err)
	}
	//将字符串转换为整型
	a, _ := strconv.Atoi("316316")
	fmt.Println(a)
	b := strconv.Itoa(5655)
	fmt.Println(b)
}

//正则表达式
func test05() {
	t := "abc atc aic a6c a0c ajc a5c"
	t1 := "3.51 .53 15.53 53. 315.53 3.1 135"
	reg1 := regexp.MustCompile(`a[0-9]c`) //解释规则，他会解析正则表达式，如果成功返回解释器
	if reg1 == nil {
		fmt.Println(reg1)
	}
	reg2 := regexp.MustCompile(`\d+\.\d+`) //+匹配前一个的一次或多次
	if reg2 == nil {
		fmt.Println(reg2)
	}
	result := reg1.FindAllStringSubmatch(t, -1) //根据规则提取关键信息-1:匹配所有
	fmt.Println(result)
	result1 := reg2.FindAllStringSubmatch(t1, -1)
	fmt.Println(result1)
}

//JSON处理
//通过结构体生产JSON
type IT struct {
	Company  string   `json:"com"` //二次编码
	Subjects []string `json:"sub"`
	IsOk     bool     `json:",string"` //字符串输出
	Price    float64  `json:"-"`       //此字段不会输出
}

func test06() {
	s := IT{"cccc", []string{"gg", "ccc", "gcd"}, true, 666.666}
	t, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(t))
	t1, err1 := json.MarshalIndent(s, "", "") //格式化编码
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(t1))
	//通过map生成JSON
	m := make(map[string]interface{}, 4)
	m["company"] = "asdgh"
	m["jubject"] = []string{"et", "asgt", "asdgjh"}
	m["isok"] = true
	m["price"] = 66.66
	r, err2 := json.Marshal(m)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(r))

	r1, err3 := json.MarshalIndent(m, "", "") //格式化编码
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println(string(r1))

	//	解析到结构体
	jsonTest := `{
		"com": "cccc",
		"sub": [
				"gg",
				"ccc",
				"gcd"
		],
		"IsOk": "true"
		}
`
	var ss IT
	err4 := json.Unmarshal([]byte(jsonTest), &ss)
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	fmt.Println(ss)
	fmt.Println(ss.Company)
	//只需要解析一部分
	type IT1 struct {
		Subjects []string `json:"sub"`
	}
	var ss1 IT1
	err5 := json.Unmarshal([]byte(jsonTest), &ss1)
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	fmt.Println(ss1)

	//JSON解析到map
	m2 := make(map[string]interface{}, 4)
	err6 := json.Unmarshal([]byte(jsonTest), &m2)
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	fmt.Println(m2)

}

//磁盘文件的使用
func test07() {
	path := "./abc.txt"
	WriteFile(path)
	//ReadFile(path)
	//ReadFileline(path)

}
func WriteFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err1 := f.WriteString("啊这")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

}
func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	s := make([]byte, 1024*2) //2k大小
	n, err1 := f.Read(s)
	if err1 != nil && err1 != io.EOF { //文件出错，同时没有到结尾
		fmt.Println("err1:", err1)
		return
	}
	fmt.Println(string(s[:n]))

}

//读每一行的内容
func ReadFileline(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//新建一个缓存区，把内容放到缓存区
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err:", err)
		}
		fmt.Println(string(s))
	}
}

//my cp
func test08() {
	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage:xxx srcFile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目标文件不能相同")
		return
	}
	//只读方式打开源文件
	sF, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//新建目的文件
	dF, err1 := os.Create(dstFileName)
	if err1 != nil {
		fmt.Println("err:", err1)
		return
	}

	defer sF.Close()
	defer dF.Close()

	b := make([]byte, 4*1024)
	for {
		n, err := sF.Read(b)
		if err != nil {
			fmt.Println("err:", err)
			if err == io.EOF {
				break
			}
		}
		//往目的文件写，读多少写多少
		dF.Write(b[:n])
	}

}
