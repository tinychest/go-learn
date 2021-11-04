package excel

/*
【参考】
https://www.bookstack.cn/read/excelize-v2.0/spilt.1.spilt.1.8.md

【概念】
DataExcel（主要用于数据展示的表格文件）

Sheet（表格页）
 - SheetName

SheetRow
 - RowEmpty
 - RowAWord
 - RowTable
 - RowTableHead
 - RowTableBody

Cell（单元格）

【长宽】
横长是列的属性，故应加在表格头上
纵宽是行的属性，故应加在行上

列采用表头、表体，最大宽度最为宽度（由上至下，即越下面的优先级越高）
所有高度、宽度字段，代表设置时的值
所有 0，代表自适应

【样式】
- 类库对单元格样式的定义（一个单元格有着 数据 和 样式 两个属性）
- 字体；基于基础 Excel 文件的话，样式是无法修改的，需要 页面布局 → 主题 → 字体

[github.com/xuri/excelize/v2@v2.4.1/styles.go:2504]
Border 边框
- Type：left、right、top、bottom、diagonalUp、diagonalDown（后边两个是对角线边框）
- Color：#颜色的十六进制码
- Style：https://www.bookstack.cn/read/Excelize-2.4-zh/8c0c3cb5d0250978.md#边框

[github.com/xuri/excelize/v2@v2.4.1/styles.go:2370]
Fill 背景填充
- Type：gradient（渐变）、pattern（纯色）
- Pattern：（Type=pattern）填充样式 https://www.bookstack.cn/read/Excelize-2.4-zh/8c0c3cb5d0250978.md#图案填充
- Color：（Type=pattern）指定颜色；（Type=gradient）渐变色范围
- Shading：（Type=gradient）https://www.bookstack.cn/read/excelize-v2.0/spilt.2.spilt.1.8.md

[github.com/xuri/excelize/v2@v2.4.1/styles.go:2181]
Font 字体
- Bold：加粗
- Italic：斜体
- Underline：下划线 single=单线、double=双线
- Family：字型（微软雅黑、Times New Roman、Berlin Sans FB Demi）
- Size：大小
- Strike：中划线
- Color：颜色

[https://www.bookstack.cn/read/excelize-v2.0/spilt.4.spilt.1.8.md]
Alignment 居中
- Horizontal 水平居中 见链接
- Indent 缩进
- JustifyLastLine ？
- ReadingOrder ？
- RelativeIndent ？
- ShrinkToFit 可以理解
- TextRotation 顺时针旋转度数
- Vertical 竖直居中 见链接
- WrapText 是否换行

Protection
- Hidden 隐藏
- Locked 锁定

-- 以下属性实在无使用场景 --

[github.com/xuri/excelize/v2@v2.4.1/styles.go:31]
NumFmt：数字展示格式（和 lang 存在关联关系）
CustomNumFmt：NumFmt 不满足需求的话，自定义数字展示格式

DecimalPlaces：浮点数四舍五入展示的小数位

[github.com/xuri/excelize/v2@v2.4.1/styles.go:68]
Lang

NegRed：？

【pane】
窗格，窗格冻结可以将特定的表格行或者列始终保持在特定的位置
水平方向、数值方向 都可以进行咯分
因为，内部没有暴露结构体，所以只能通过 json 串的形式进行配置

(*excelize.File) SetPanes(sheet, panes string) error
例：{"freeze": true, "y_split": 1, "top_left_cell": "A2"}
（"y_split": 1）：以纵轴为 1 的线，将表格划分成上下两个窗格
（"freeze": true）：将上窗格悬浮起来
（"top_left_cell": "A2"）：未知，在当前场景中，不能没有，且要保证行号大于 y_split，否则导出的表格非法
*/

// TODO 定制化样式（友好的设置样式）

// TODO Margin 设计

// TODO 复杂表头
