# testify 框架

## 常用函数
```golang
//判断是否相等, 成立成OK, 不成立则报错
func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
//判断是否不相等, 成立成OK, 不成立则报错
func NotEqual(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

//判断是否为nil, 成立成OK, 不成立则报错
func Nil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
//判断是否不为nil, 成立成OK, 不成立则报错
func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool


//判断是否为空, 成立成OK, 不成立则报错
func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
//判断是否不为空, 成立成OK, 不成立则报错
func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool


//判断是否无错, 成立成OK, 不成立则报错
func NoError(t TestingT, err error, msgAndArgs ...interface{}) bool
//判断是否是错, 成立成OK, 不成立则报错
func Error(t TestingT, err error, msgAndArgs ...interface{}) bool

//判断是否等于0, 成立成OK, 不成立则报错
func Zero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool
//判断是否不等于0, 成立成OK, 不成立则报错
func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool


//判断是否为true, 成立成OK, 不成立则报错
func True(t TestingT, value bool, msgAndArgs ...interface{}) bool
//判断是否为false, 成立成OK, 不成立则报错
func False(t TestingT, value bool, msgAndArgs ...interface{}) bool

//判断是否长度, 成立成OK, 不成立则报错
func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) bool


//判断是否包含, 成立成OK, 不成立则报错
func Contains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool
//判断是否不包含, 成立成OK, 不成立则报错
func NotContains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool

//判断子切片是否是list的真子集
func Subset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool)
//判断子切片是否不是list的真子集
func NotSubset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool)

//判断文件是否存在
func FileExists(t TestingT, path string, msgAndArgs ...interface{}) bool
//判断目录是否存在
func DirExists(t TestingT, path string, msgAndArgs ...interface{}) bool
```