package e

import "net/http"

var MsgFlag = map[int]string{
	http.StatusOK:                  "ok",
	http.StatusInternalServerError: "fail",
	http.StatusBadRequest:          "参数错误",
	// 用户操作
	30001: "用户名已存在",
	30002: "密码加密失败",
	30003: "用户不存在",
	30004: "密码错误",
	30005: "token下发失败",
	30006: "用户数据更新失败",
	30007: "token认证失败",
	30008: "token过期",
	30009: "用户头像上传失败",
	30010: "token不存在",

	// 邮件发送
	40001: "邮件发送失败",

	// 商品信息
	50001: "上传商品封面失败",
	50002: "上传商品图片失败",
}

func GetMSG(code int) string {
	msg, ok := MsgFlag[code]
	if !ok {
		return MsgFlag[http.StatusInternalServerError]
	}
	return msg
}
