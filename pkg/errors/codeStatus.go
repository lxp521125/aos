package errors

// CodeTypeInt ...
type CodeTypeInt int

//U为更新 F为失败  A为添加  SUCC 为成功 C 复制 G查找
const (
	SUCCESSSTATUS = 11999999
	SYSERR        = 11100001
	DBERR         = 11100002
	PARAMETERERR  = 11100003
	NOMOREDATA    = 11100004
	NODATA        = 11100005
)

// INFO ...
var INFO = GetInfo()

// GetInfo ...
func GetInfo() map[int]string {
	info := make(map[int]string)
	info[SUCCESSSTATUS] = "请求成功"
	info[SYSERR] = "系统错误"
	info[DBERR] = "DB 错误"
	info[NOMOREDATA] = "没有更多数据"
	info[PARAMETERERR] = "参数错误"
	info[NODATA] = "无数据"
	return info
}
