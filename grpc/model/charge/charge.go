package charge

type (
	Request struct {
		UserID       int64  `json:"user_id"`        // 用户ID
		Token        string `json:"token"`          // 用户Token
		Session      string `json:"session"`        // 用户临时会话钥匙
		Amount       int64  `json:"amount"`         // 单位:分
		PayType      int64  `json:"pay_type"`       // 支付类型
		PayChannelID int64  `json:"pay_channel_id"` // 支付渠道ID
	}

	Response struct {
		Code   int64  `json:"code"`
		Msg    string  `json:"msg"`
		PayUrl string `json:"pay_url"`	// 支付链接
	}
)
