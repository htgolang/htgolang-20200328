func name(parameters) returns {
    return args
}


a := func() {

}

func() {

}()

sort.Sort(slice, func(i, j) {return slice[i] > slice[j]})


func() {
    defer func() {
        fmt.Println("defer")
    }()


    1 / 0
}


github.com/imsilence/仓库名称(项目名称)

strutil


git@github.com:imsilence/strutil.git


在调用的时候
导入: github.com/imsilence/strutil => mod name

va.b.c


md5 hash => 不可逆
string (n:1)=> md5

md5 => string

彩虹表(暴力破解) MD5表 => string



A => md5 密码iamkk => abcdef
B => md5 密码iamkk => abcdef

某一天A被脱裤了 abcdef ==> iamkk

加盐

md5(iamkk + salt) => salt + defg


A => md5 密码iamkk => abcdef => iamkk + xxx => hash
B => md5 密码iamkk saltxxx => xyz
xyz =>


单元测试 => 功能性测试

测试覆盖率

go test -coverprofile=cover.out