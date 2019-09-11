/*
golang http请求
*/
package zhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	Ok = "00000"

	// Generic error definition, 10000-10999
	ErrInternal       = "10000"
	ErrInMaintain     = "10001"
	ErrRequestFormat  = "10002"
	ErrApiDeprecated  = "10003" // Api is deprectated, need to update app
	ErrSessionExpired = "10004"
	ErrInBlacklist    = "10005"
	ErrTempBlock      = "10006" // 暂时被禁止访问
	ErrAuth           = "10007"
	ErrUnimplemented  = "10008"
	ErrUnreachable    = "10009" // 第三方平台不可达

	// User/Auth error definition, 11000-11099
	ErrNicknameExists          = "11000"
	ErrUserNotFound            = "11001"
	ErrUserExists              = "11002"
	ErrPasswordWrong           = "11003"
	ErrUserBlacked             = "11004"
	ErrMobileUnmodified        = "11005"
	ErrMobileFormat            = "11006"
	ErrMobileNotSet            = "11007"
	ErrPasswordLength          = "11008"
	ErrPasswordFormat          = "11009"
	ErrNicknameStopWord        = "11010" // 包含违禁词
	ErrTokenNotFound           = "11011"
	ErrTokenFormat             = "11012"
	ErrTokenExpired            = "11013"
	ErrTokenRefreshExpired     = "11014"
	ErrRoleFormat              = "11015"
	ErrRoleWrong               = "11016"
	ErrNicknameFormat          = "11017"
	ErrUserBindWeChatForbidden = "11018"
	ErrUserBindMobileForbidden = "11019"

	//activity error definition,11200-11299
	ErrEnrollExists      = "11201"
	ErrConsultExists     = "11202"
	ErrActivityNotExists = "11203"
	ErrActivityEnrolled  = "11204"

	// Verification error definition, 11100-11119
	ErrSmsNotFound   = "11100"
	ErrWrongVeriCode = "11101"
	ErrSmsFrequent   = "11102"

	// diary and story, 11300-11399
	ErrStoryNotExist   = "11300"
	ErrLikedAlready    = "11301"
	ErrCommentYourself = "11302"

	// version control, 11400-11499
	ErrVersionNotFound        = "11400"
	ErrVersionDuplicate       = "11401"
	ErrVersionChannelNotFound = "11402"

	//statistic 11500-11599
	ErrStatisticSignIllegal = "11500"

	//config 11600-11699
	ErrConfigDuplicate     = "11600"
	ErrConfigScopeNotFound = "11601"
	ErrConfigNoScopeExist  = "11602"

	//scm 11700-11799
	ErrScmDuplicate = "11700"

	//confidant 11800-11899
	ErrConfidantActiveCodeNotExist    = "11800"
	ErrConfidantActiveCodeAlreadyUsed = "11801"
	ErrConfidantInviterExist          = "11802"
	ErrConfidantInviteDuplicate       = "11803"
	ErrConfidantInviteMax             = "11804"
	ErrConfidantInviteForbidden       = "11805"
	ErrConfidantAcceptForbidden       = "11806"
	ErrConfidantInviteeNotExist       = "11807"

	// hospital 11900-11999
	ErrHospitalDuplicate    = "11900"
	ErrHospitalNotFound     = "11901"
	ErrHospotalStaffNotBind = "11902"

	//goods 12000-12099
	ErrCollectedAlready      = "12000"
	ErrGoodsNotExist         = "12001"
	ErrGoodsStockEmpty       = "12002"
	ErrGoodsStatusNotMatched = "12003"

	// media 12100-12199
	ErrFileUploadFail = "12100"

	// order 12200-12299
	ErrOrderNotExist               = "12200"
	ErrOrderSourceNotMatched       = "12201"
	ErrOrderOwnerNotMatched        = "12202"
	ErrOrderStatusNotMatched       = "12203"
	ErrOrderHasAddedOn             = "12204"
	ErrAmountNotCorrect            = "12205"
	ErrOrderClosed                 = "12206"
	ErrOrderComplated              = "12207"
	ErrOrderSettleStatusNotMatched = "12208"

	// account 12300-12399
	ErrWithdrawTimesExceed   = "12300"
	ErrWithdrawAmountExceed  = "12301"
	ErrWithdrawTimeYetToCome = "12302"
	ErrAccountMissAlipay     = "12303"
	ErrAccountMissCash       = "12304"
	ErrInsufficientBalance   = "12305"

	// payment 12400-12499
	ErrCurrencyNotSupported = "12400"
	ErrUnknownHookType      = "12401"

	//easemob 12500-12599
	ErrEasemobUserTokenNone = "12501"
	ErrEasemobRegisterUser  = "12502"

	// coupon 12600-12699
	ErrCouponSoldOut  = "12600"
	ErrCouponUserSent = "12601"

	// leaflet 12700-12799
	ErrLeafletCounselorDuplicate = "12700"

	// proposal 12800-12899
	ErrDesireNotEnd       = "12800"
	ErrDesireScopeIllegal = "12801"

	// wechat 12900-12999
	ErrWeChatTemplateParams     = "12900"
	ErrWeChatNoSuchTemplate     = "12901"
	ErrWeChatUnAuth             = "12902"
	ErrWeChatUnionIdNotExist    = "12903"
	ErrWeChatQRCodeDupDuplicate = "12904"
)

var client = &http.Client{}

func GET(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		log.Printf("The response data is error1============ : %s", err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, errs.NewRpcError(errs.ErrUnreachable, "status code %d", resp.StatusCode)
		log.Printf("The response data is error2============ : %s", err.Error())
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		log.Printf("The response data is error3============ : %s", err.Error())
		return nil, err
	}

	log.Printf("The response data is : %s", data)

	return data, nil
}

//GET请求，返回json
func GETWithUnmarshal(url string, resp interface{}) error {
	data, err := GET(url)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, resp); err != nil {
		//return errs.NewRpcError(errs.ErrInternal, err.Error())
		return err
	}
	return nil
}

func POST(url string, data []byte) ([]byte, error) {
	resp, err := client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, errs.NewRpcError(errs.ErrUnreachable, "status code %d", resp.StatusCode)
		return nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		return nil, err
	}

	return data, nil
}

func POSTContentType(url string, contentType string, data []byte) ([]byte, error) {
	resp, err := client.Post(url, contentType, bytes.NewReader(data))
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//return nil, errs.NewRpcError(errs.ErrUnreachable, "status code %d", resp.StatusCode)
		return nil, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//return nil, errs.NewRpcError(errs.ErrInternal, err.Error())
		return nil, err
	}

	return data, nil
}

func POSTWithUnmarshal(url string, req interface{}, resp interface{}) error {
	var data []byte
	var err error

	switch req.(type) {
	case []byte:
		data = req.([]byte)
	default:
		data, err = json.Marshal(req)
		if err != nil {
			//return errs.NewRpcError(errs.ErrInternal, err.Error())
			return err
		}
	}

	data, err = POST(url, data)
	if err != nil {
		return err
	}

	if resp == nil {
		return nil
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		//return errs.NewRpcError(errs.ErrInternal, err.Error())
		return err
	}
	return nil
}
