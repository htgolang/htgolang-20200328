package util

//根据权限返回权限对应的字符串
func Role(num int) string{
	switch num {
	case 0:
		return "普通用户"
	case 1:
		return "管理员"
	case 2:
		return "超级管理员"
	}
	return ""
}

//根据字符串返回对应的权限int
func GetRole(role string) int{
	switch role {
	case "普通用户":
		return 0
	case "管理员":
		return 1
	case "超级管理员":
		return 2
	}
	return 0
}
