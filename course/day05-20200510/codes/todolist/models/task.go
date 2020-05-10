package models

import "time"

/*
结构体名称首字母小写，在包外不能访问
*/
type taskv1 struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	user      string
}

/**
结构体属性明首字母小写，在包外不能访问
*/
type Taskv2 struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	user      string
}

// 对name只有写权限，没有读权限

type Taskv3 struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

type taskv4 struct {
	Id        int
	Name      string
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

// 命名嵌入
// 结构体包外不可见
type wrapperv1 struct {
	taskv1 taskv1
	Taskv2 Taskv2
	Taskv3 Taskv3
	taskv4 taskv4
}

// 结构体包外可见
type Wrapperv2 struct {
	taskv1 taskv1
	Taskv2 Taskv2 // 属性包外可见
	Taskv3 Taskv3 // 属性包外可见
	taskv4 taskv4
}

// 包外不可见
type anonyWrapperv1 struct {
	taskv1
}

type anonyWrapperv2 struct {
	Taskv2
}

type anonyWrapperv3 struct {
	Taskv3
}

type anonyWrapperv4 struct {
	taskv4
}

// 包外可见的
type AnonyWrapperv1 struct {
	taskv1 //属性 包外不可见
}

type AnonyWrapperv2 struct {
	Taskv2 // 属性包外不可见
}

type AnonyWrapperv3 struct {
	Taskv3 // 属性包外可见
}

type AnonyWrapperv4 struct {
	User   string
	taskv4 // taskv4中的User属性无法在包外访问
}
