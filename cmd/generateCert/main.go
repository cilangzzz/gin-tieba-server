/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/9/20
  @desc: //TODO
**/

package main

import . "ginFlutterBolg/util"

func main() {
	GenerateCert("cert", "www.cilang.buzz", "CN",
		"GuangZhou", "GuangZhou", "cilang.buzz", "cilang", "localhost")
}