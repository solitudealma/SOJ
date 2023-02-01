package result

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/solitudealma/SOJ/backend/common/validate"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if ok := xerr.IsCustomizeErr(causeErr); ok { //自定义错误类型
			//自定义CodeError
			e, _ := causeErr.(*xerr.CodeError)
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusOK, Error(errcode, errmsg))
	}
}

// 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {

	var errMsg string
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		var builder strings.Builder
		trans := validate.Trans("zh")

		for idx, value := range ve {
			builder.WriteString(value.Translate(trans))
			if idx < len(ve)-1 {
				builder.WriteString("，")
			}
		}

		errMsg = fmt.Sprintf("%s：%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), builder.String())
	} else {
		errMsg = fmt.Sprintf("%s：%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	}

	httpx.WriteJson(w, http.StatusOK, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}
