package util

import (
	"fmt"
	"strconv"
	"strings"
)

// compare version
func VersionCompare(version1, version2 string) int {
	a := strings.Split(version1, ".")
	b := strings.Split(version2, ".")
	flag := 1
	if len(a) > len(b) {
		a, b = b, a
		flag = -1
	}
	for i := range a {
		x, _ := strconv.Atoi(a[i])
		y, _ := strconv.Atoi(b[i])
		if x < y {
			return -1 * flag
		} else if x > y {
			return 1 * flag
		}
	}
	for _, v := range b[len(a):] {
		y, _ := strconv.Atoi(v)
		if y > 0 {
			return -1 * flag
		}
	}
	return 0
}

func PrintLogo() {
	fmt.Print(`
	               _ooOoo_
                  o8888888o
                  88" . "88
                  (| -_- |)
                  O\  =  /O
               ____/'---'\____
             .'  \\|     |//  '.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||-  \
           |   | \\\  -  /// |   |
           | \_|  ''\-/''  |   |
           \  .-\__  '-'  ___/-. /
         ___'. .'  /-.-\  '. . __
      ."" '<  '.___\_<|>_/___.'  >'"".
     | | :  '- \'.;'\ _ /';.'/ - ' : | |
     \  \ '-.   \_ __\ /__ _/   .-' /  /
======'-.____'-.___\_____/___.-'____.-'======
                   '=-='

`)
}
