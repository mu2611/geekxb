package controllers

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	DEBUG = true // 输出调试信息
	// DEBUG = false    // 不输出调试信息

	PreUrl = "" //上一级url 操作完成后跳转使用

)

const (
	ItemsPerPage = 8 // 列表页上每页显示文章数量
)

func Check(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

// 当DEBUG开关打开时，输出调试信息
// 使用方法：`debug('格式化字符串', 参数1, 参数2, ...)`
// 输出结果：`DEBUG: 提示信息`(会自动换行)
// 相当于加强版的`fmt.Println()`，支持格式化字符串，输出以`DEBUG: `开头
func Debug(infos ...interface{}) {
	if DEBUG {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

// 把用`, `间隔的字符串转换为字符串列表(例如关键字列表)
func Str2slice(str string) []string {
	return strings.Split(str, ", ")
}

// 把字符串转换为time.Time对象
func Str2date(timeStr string) (timeObj time.Time, err error) {
	layout := "2006-01-02"
	timeObj, err = time.Parse(layout, timeStr)
	//Check(err)
	return
}

// 判断字符串切片中是否包含某个字符串
func SliceContains(strSlice []string, str string) bool {
	strMap := make(map[string]bool)
	for _, strItem := range strSlice {
		strMap[strItem] = true
	}
	if _, ok := strMap[str]; ok {
		return true
	}
	return false
}

//后台部分分页代码
// itemCount数据总数  curPageNum当前页数
func ResponsePageDiv(itemCount int64, curPageNum int64) (pageDiv string) {

	// 获得总页数
	var maxPagenum, prePagenum, nextPagenum int64
	if itemCount/ItemsPerPage*ItemsPerPage < itemCount {
		maxPagenum = itemCount/ItemsPerPage + 1
	} else {
		maxPagenum = itemCount / ItemsPerPage
	}

	if curPageNum > maxPagenum { // 如果当前页码超出范围则重置
		curPageNum = maxPagenum
	} else if curPageNum <= 0 {
		curPageNum = 1
	}

	prePagenum--
	if (prePagenum - 1) <= 1 { //上一页
		prePagenum = 1
	}

	nextPagenum++
	if (nextPagenum + 1) >= maxPagenum { //下一页
		nextPagenum = maxPagenum
	}

	return fmt.Sprintf(`<div class="fg-toolbar ui-toolbar ui-widget-header ui-corner-bl ui-corner-br ui-helper-clearfix">
	                     	<div class="dataTables_paginate fg-buttonset ui-buttonset fg-buttonset-multi ui-buttonset-multi paging_full_numbers" id="DataTables_Table_0_paginate">
	                     	<label class="ui-corner-tl ui-corner-bl fg-button ui-button ui-state-default" >%d/%d</label>
	                         <a class="first ui-corner-tl ui-corner-bl fg-button ui-button ui-state-default" href="?page=1">首页</a>
	                         <a class="previous fg-button ui-button ui-state-default" href="?page=%d">上一页</a>
	                         <!--<span>
	                            <a class="fg-button ui-button ui-state-default">1</a>
	                            <a class="fg-button ui-button ui-state-default">2</a>
	                            <a class="fg-button ui-button ui-state-default">3</a>
	                            <a class="fg-button ui-button ui-state-default">4</a>
	                            <a class="fg-button ui-button ui-state-default">5</a>
	                         </span>-->
	                         <a class="next fg-button ui-button ui-state-default" href="?page=%d">下一页</a>
	                         <a class="last ui-corner-tr ui-corner-br fg-button ui-button ui-state-default" href="?page=%d">末页</a>
	                        </div>
                 		</div>`, curPageNum, maxPagenum, prePagenum, nextPagenum, maxPagenum)

}
