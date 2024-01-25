package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	_, err := strconv.ParseFloat(" 1234256562147422", 64)

	if err == nil {
		fmt.Println("纯数字")
	} else {
		fmt.Println("非纯数字")
	}

	text := "!@#"

	re := regexp.MustCompile(`^\p{P}+$`)
	if re.MatchString(text) {
		fmt.Println("包含标点符号")
	} else {
		fmt.Println("包含非标点符号")
	}

	text = "整结果为\"河内市清春区光定区武统潘街125号KOJI体育股份公司\"。\\n</process>\\n<result>河内市清春区光定区武统潘街125号KOJI体育股份公司</result>'\n"

}
