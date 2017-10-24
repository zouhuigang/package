package main

/*
#pragma comment(lib, "libzkfp/x64lib/libzkfp.lib")
#include "winsock2.h"
#include "windows.h"
#include "libzkfp/include/zkinterface.h"
#include "libzkfp/include/libzkfperrdef.h"
#include "libzkfp/include/libzkfptype.h"
#include "libzkfp/include/libzkfp.h"
*/
import "C"
import (
	"fmt"
)

func main() {

	//s := C.ZKFPM_Init()
	C.ZKFPM_Init()

	fmt.Println("version is", C.FP_MTHRESHOLD_CODE, C.ZKFP_ERR_OK)
}
