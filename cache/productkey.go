package cache

import (
	"fmt"
)

const (
	RankKey = "rank_key"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("product_view_number：%d", id)
}
