### golang对接指纹识别仪(ZK4500)


首先导入win下的包

	#include "windows.h"

这个会处理一些类型错误,如HANDLE句柄错误，这个类型只有win下有用。



问题详情：

collect2.exe [Error] ld returned 1 exit status：


	 1.编译成功的例子在后台执行，有时一闪而过，如果再次build ，则会提示上述错误。

	解决方法：打开任务管理器，找到相应的exe进程，关闭即可；  或者直接关闭QtCreator。
	
	2.没有编译成功的情况下，最常见情况是程序本身需要include的头文件被遗漏了
	
	解决方法：细心查找基类所用的头文件，include之后即可。
	
	3..h文件中相关的槽函数在cpp文件中没有定义
	
	解决方法：查找遗漏的槽函数，根据需要，具体的定义。

 