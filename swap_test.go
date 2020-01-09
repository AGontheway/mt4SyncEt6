package mt4SyncEt6

import (
	"encoding/json"
	"fmt"
	"mt4SyncEt6/swap"
	"testing"
)

func TestSwap(t *testing.T) {
	str, _ := swap.GetSwap()
	var swaps []SwapInfo
	json.Unmarshal([]byte(str), &swaps)

	engine := GetEngine()
	for _, v := range swaps {
		sql := "update source set swap_long = ?,swap_short = ?,source_cn = ?,swap_3_day = ? where source = ?"
		_, err := engine.Exec(sql, v.SwapLong, v.SwapShort, v.SourceCN,v.Swap3Day, v.Symbol)
		if err != nil {
			fmt.Println(err)
		}
	}
}