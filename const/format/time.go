package format

// Go 语言真的是脑洞打开，一开始并不知道 20060102150405 这样的数字有什么意义，以为只要按照特定格式写个样例时间，go 就知道格式化成什么样了
// 实际并不行，你给的样例时间，必须按照指定下面这样的格式
//
// [简记]
// 2006：年|01：月|02：日|15：时|04：分|05：秒
//
// [详情]
// 年　 06,2006
// 月　 1,01,Jan,January
// 日　 2,02,_2
// 时　 3,03,15,PM,pm,AM,am
// 分　 4,04
// 秒　 5,05
//
// 周　 Mon,Monday
// 时区 -07,-0700,Z0700,Z07:00,-07:00,MST
//
// [特殊：小时]
// 3    表示用12小时制表示，去掉前导0
// 03   表示用12小时制表示，保留前导0
// 15   表示用24小时制表示，保留前导0
// 03pm 表示用24小时制am/pm表示上下午表示，保留前导0
// 3pm  表示用24小时制am/pm表示上下午表示，去掉前导0
//
// 总结：本质上是最初说的以一个时间格式为样本，但是这个时间一定得是 go 语言诞生的时间
// 当然，这样肯定是有些不合理的，所以可以从 1~6 来表示
const (
	FmtDateTime    = "2006-01-02 15:04:05"
	FmtDateTime_   = "2006-01-02_15:04:05"
	FmtDate        = "2006-01-02"
	FmtTime        = "15:04:05"
	FmtDateTimeCN  = "2006年01月02日 15时04分05秒"
	FmtDateTimeCN_ = "2006年01月02日_15时04分05秒"
	FmtDateCN      = "2006年01月02日"
	FmtTimeCN      = "15时04分05秒"
)
