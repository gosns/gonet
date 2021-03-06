package ipc

import "misc/packet"
import . "types"

var Code = map[string]int16{
	"heart_beat_req":         0,    // 心跳包..
	"user_login_req":         1,    // 客户端发送用户登陆请求包
	"user_login_succeed_ack": 2,    // 登陆成功
	"user_login_faild_ack":   3,    // 登陆失败
	"talk_req":               1000, // talk给一个用户
	"talk_notify":            1001, // notify客户端
	"key_exchange_req":       2000, // Diffie-Hellman
	"key_exchange_ack":       2001, // Diffie-Hellman
}

var RCode = map[int16]string{
	0:    "heart_beat_req",         // 心跳包..
	1:    "user_login_req",         // 客户端发送用户登陆请求包
	2:    "user_login_succeed_ack", // 登陆成功
	3:    "user_login_faild_ack",   // 登陆失败
	1000: "talk_req",               // talk给一个用户
	1001: "talk_notify",            // notify客户端
	2000: "key_exchange_req",       // Diffie-Hellman
	2001: "key_exchange_ack",       // Diffie-Hellman
}

var ProtoHandler = map[uint16]func(*Session, *packet.Packet) []byte{
	0:    P_heart_beat_req,
	1:    P_user_login_req,
	1000: P_talk_req,
	2000: P_key_exchange_req,
}
