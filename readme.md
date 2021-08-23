# qbhy/chinese
golang 常见汉字词库

## 安装 - installation
```bash
go get github.com/qbhy/chinese
```

## 使用 - usage
```go
package tests

import (
	"fmt"
	"github.com/qbhy/chinese"
	"testing"
)

func TestFindWords(t *testing.T) {
	fmt.Println(chinese.RelatedWords("走吧"))
	// map[吧:[你看着办吧 吧主 吧务 吧台 吧女 好吧 安吧尔 歇了吧 泡吧 爆吧 百度贴吧 省省吧 网吧 贴吧 酒吧] 走:[一条道走到黑 东奔西走 临走 人往高处走水往低处流 偷走 出走 勾走 卷走 取走 咪走堂 夺走 奔走 奔走相告 带走 开走 往前走 抢走 拿走 掉头就走 排走 搬走 放走 暴走 暴走漫画 暴走鞋 正步走 游走 溜走 漂走 瓦萨利走廊 看走眼 离家出走 行走 被赶走 贩夫走卒 走人 走低 走俏 走光 走入 走出 走动 走势 走南闯北 走去 走向 走味的 走回 走在 走地盘 走地鸡 走失 走廊 走开 走扁带 走投无路 走时 走来走去 走桃花运 走漏 走火 走狗 走神 走秀 走私 走红 走肾 走访 走读 走调 走路 走过 走过场 走近 走进 走遍 走道 走避 走错 走门路 走马看花 赶走 轰走 送走 逃走 飞走]]
}
```

https://github.com/qbhy/chinese  
qbhy0715@qq.com
