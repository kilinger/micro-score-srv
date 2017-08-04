package handler


func ValidatePage(page int32) int32 {
	default_page := 1
	if page > 1 {
		return page
	}
	return int32(default_page)
}

func ValidatetCount(count int32) int32 {
	default_count := 20
	if count > 0 {
		return count
	}
	return int32(default_count)
}

func GetLimitOffset(page int32, count int32, total int32) (int32, int32) {
	page = ValidatePage(page)
	count = ValidatetCount(count)
	offset := (page - 1) * count
    limit := count
	return limit, offset
}