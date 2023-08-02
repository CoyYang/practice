package main

import "fmt"

func Delete[everything any](section []everything, index int) []everything {
	//检查索引是否有值，小于或者超出
	if index < 0 || index >= len(section) {
		return section
	}

	//返回个切片，长度为原始的长度
	ret_Section := make([]everything, len(section))
	//copy一个出来，做删除后的切片
	copy(ret_Section, section)

	//删除指定索引元素
	ret_Section = append(ret_Section[:index], ret_Section[index+1:]...)

	//判断缩容情况
	if len(ret_Section) <= cap(section)/2 {
		ret_Section = ret_Section[:len(ret_Section):cap(ret_Section)]
	}

	fmt.Println(ret_Section)
	return ret_Section
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	Delete(data, 2)

}
