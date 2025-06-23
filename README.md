# gophp-utils

gophp 相关项目的通用 utils 包集合。

虽然此库与业务并不强耦合，但仍不推荐将此库用于其他项目，建议使用其他公开库(例如: [github.com/samber/lo](https://github.com/samber/lo))

## 子包简介

- `ascii`: ASCII 相关的函数库。(类比 c 语言中 ctype.h)
- `la`: 类型语言特性补丁的函数库，替代其他编程语言中常见但在 golang 中没有的语言特性.(例如: 布尔异或、三元表达式、错误断言等)
- `xmaps`: 标准库 `maps` 的补充
- `xslices`: 标准库 `slices` 的补充
- `xstrings`: 标准库 `strings` 的补充
