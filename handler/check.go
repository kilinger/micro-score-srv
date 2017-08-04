package handler

import (
	proto "code.xxxxxx.com/micro/proto/scores"
)

func CheckExchangeRequest(req *proto.ExchangeRequest) (string, bool) {
	var ( msg string; ok bool )
	msg, ok = CheckUserId(req.GetUserId()); if !ok { return msg, false }
	msg, ok = CheckScore(req.GetScore()); if !ok { return msg, false }
	msg, ok = CheckGiftId(req.GetGiftId()); if !ok { return msg, false }
	msg, ok = CheckGiftName(req.GetGiftName()); if !ok { return msg, false }
	return "", true
}

func CheckStatRequest(req *proto.StatRequest) (string, bool) {
	msg, ok := CheckUserId(req.GetUserId()); if !ok { return msg, false }
	return "", true
}

func CheckListRequest(req *proto.ListRequest) (string, bool) {
	var ( msg string; ok bool )
	msg, ok = CheckPage(req.GetPage()); if !ok { return msg, false }
	msg, ok = CheckCount(req.GetCount()); if !ok { return msg, false }
	return "", true
}

func CheckCreateRequest(req *proto.CreateRequest) (string, bool) {
	var ( msg string; ok bool )
	msg, ok = CheckUserId(req.GetUserId()); if !ok { return msg, false }
	msg, ok = CheckScore(req.GetScore()); if !ok { return msg, false }
	msg, ok = CheckNote(req.GetNote()); if !ok { return msg, false }
	return "", true
}

func CheckUserId(user_id int64) (string, bool) {
	if user_id == 0 {
		msg := "'user_id' is required"
		return msg, false
	}
	return "", true
}

func CheckScore(score int32) (string, bool) {
	if score == 0 {
		msg := "'score' is required"
		return msg, false
	}
	return "", true
}

func CheckNote(note string) (string, bool) {
	if note == "" {
		msg := "'note' is required"
		return msg, false
	}
	return "", true
}

func CheckPage(page int32) (string, bool) {
	if page <= 0 {
		msg := "'page' must more than zero"
		return msg, false
	}
	return "", true
}

func CheckCount(count int32) (string, bool) {
	if count <= 0 {
		msg := "'count' must more than zero"
		return msg, false
	}
	return "", true
}

func CheckGiftId(gift_id int64) (string, bool) {
	if gift_id == 0 {
		msg := "'gift_id' is required"
		return msg, false
	}
	return "", true
}

func CheckGiftName(gift_name string) (string, bool) {
	if gift_name == "" {
		msg := "'gift_name' is required"
		return msg, false
	}
	return "", true
}
