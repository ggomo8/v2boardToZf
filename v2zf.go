package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"v2boardToZf/tools"
)

var ZhongZhuanUrl = tools.Cfg.Section("").Key("ZHONG_ZHUAN_URL").MustString("127.0.0.1")
var ZhongZhuanName = tools.Cfg.Section("").Key("ZHONG_ZHUAN_NAME").MustString("自定义中转名称")
var SHOW = tools.Cfg.Section("").Key("SHOW").MustInt(0)

func init() {

}

func main() {

	//
Loop:
	var str string
	fmt.Printf(`
1. 获取所有直连节点批量添加的模板
2. 每个节点添加 conifg/ini 中配置的中转,提前在当前目录下放好咸蛋面板导出的json文件,只能有一个json文件
3. 根据关键词删除中转节点(默认为配置文件的中转名称)
请选择选项:`)
	fmt.Scanln(&str)
	if str == "1" {
		getAllRealUrls()
		goto Loop
	} else if str == "2" {
		readJson()
		goto Loop
	} else if str == "3" {
		var str string
		fmt.Printf(`请输入关键词(默认请回车):`)
		fmt.Scanln(&str)
		if str == "" {
			delNodeOfZz(ZhongZhuanName)
		} else {
			delNodeOfZz(str)
		}
		goto Loop
	}

	//getAllRealUrls()
	//readJson()
	//delNodeOfZz("香港02测试")
}

// delNodeOfZz 只删除带关键词的中转节点,会自动过滤直连节点
func delNodeOfZz(str string) {
	db := tools.GetDB()
	//这里Delete函数需要传递一个空的模型变量指针，主要用于获取模型变量绑定的表名
	rows1 := db.Where("name like ? and parent_id is not null ", "%"+str+"%").Delete(&tools.V2ServerV2Ray{})
	rows2 := db.Where("name like ? and parent_id is not null ", "%"+str+"%").Delete(&tools.V2ServerShadowsocks{})

	fmt.Println("共删除 ", rows1.RowsAffected+rows2.RowsAffected, "条")
}

// readJson 读取json文件并添加中转
func readJson() {
	files, _ := filepath.Glob("*")
	for _, file := range files {
		if strings.Contains(file, ".json") {
			filePtr, err := os.Open(file)
			if err != nil {
				fmt.Println("文件打开失败 [Err:%s]", err.Error())
				return
			}
			defer filePtr.Close()
			var info tools.ForWard
			// 创建json解码器
			decoder := json.NewDecoder(filePtr)
			err = decoder.Decode(&info)
			if err != nil {
				fmt.Println("解码失败", err.Error())
			} else {
				fmt.Println("解码成功")
				hadleJson(info)
			}
			return
		}
	}

}

func hadleJson(info tools.ForWard) {
	forwards := info.Forwards
	db := tools.GetDB()
	for _, forward := range forwards {
		zhong_zhuan_port := forward.InternetPort
		remotePort := forward.RemotePort
		real_url := forward.RemoteHost
		//查询直连节点过滤掉非本面板的v2ray节点
		var v2ServerV2Rays []tools.V2ServerV2Ray
		db.Where("parent_id is null and host =? and port = ?", real_url, remotePort).Find(&v2ServerV2Rays)
		//查询直连节点过滤掉非本面板的ss中转
		var v2ServerSs []tools.V2ServerShadowsocks
		db.Where("parent_id is null and host =? and port = ?", real_url, remotePort).Find(&v2ServerSs)
		if len(v2ServerV2Rays) == 0 && len(v2ServerSs) == 0 {
			fmt.Println("过滤非本面板的中转配置一条,对应的真实url: ", real_url, ":", remotePort)
			continue
		}
		if len(v2ServerV2Rays) > 0 {
			//开始复制 v2ray 类型节点
			//先判断是否有重复添加的中转，重复则跳过
			var v2ServerV2RaysZZ []tools.V2ServerV2Ray
			db.Where("parent_id is not null and port = ? and server_port = ? and host = ? ",
				zhong_zhuan_port, remotePort, ZhongZhuanUrl).Find(&v2ServerV2RaysZZ)
			if len(v2ServerV2RaysZZ) > 0 {
				fmt.Println("中转已添加,重复，跳过 ", ZhongZhuanUrl, ":", zhong_zhuan_port)
			} else {
				fmt.Println("添加中转 ", ZhongZhuanUrl, ":", zhong_zhuan_port)
				//查询对应的直连节点
				ray := v2ServerV2Rays[0]
				data, _ := json.Marshal(ray) //使用json方式 深拷贝对象
				var newNode tools.V2ServerV2Ray
				json.Unmarshal(data, &newNode)

				newNode.Id = 0
				newNode.Name = newNode.Name + ZhongZhuanName
				newNode.Host = ZhongZhuanUrl
				newNode.Port = strconv.Itoa(zhong_zhuan_port)
				newNode.Sort = sql.NullInt32{Int32: 222 + ray.Sort.Int32, Valid: true}
				newNode.ParentId = sql.NullInt32{Int32: int32(ray.Id), Valid: true}
				newNode.Show = SHOW
				settings := ray.NetworkSettings
				strNetworkSettings := settings.String
				if strNetworkSettings != "" && ray.Network == "ws" {
					var data tools.WsConfig
					err := json.Unmarshal([]byte(strNetworkSettings), &data)
					if err != nil {
						panic("解析wsConfig错误")
					} else {
						data.Headers.Host = ZhongZhuanUrl
					}
					marshal, err := json.Marshal(data)
					if err != nil {
						return
					}
					newStrNetworkSettings := string(marshal)
					newNode.NetworkSettings = sql.NullString{String: newStrNetworkSettings, Valid: true}
				}
				db.Create(&newNode)
			}
		}

		if len(v2ServerSs) > 0 {
			//开始复制 ss 类型节点
			//先判断是否有重复添加的中转，重复则跳过
			var v2ServerShadowsocksZz []tools.V2ServerShadowsocks
			db.Where("parent_id is not null and port = ? and server_port = ? and host = ? ",
				zhong_zhuan_port, remotePort, ZhongZhuanUrl).Find(&v2ServerShadowsocksZz)

			if len(v2ServerShadowsocksZz) > 0 {
				fmt.Println("中转已添加,重复，跳过 ", ZhongZhuanUrl, ":", zhong_zhuan_port)
			} else {
				fmt.Println("添加中转 ", ZhongZhuanUrl, ":", zhong_zhuan_port)
				//查询对应的直连节点
				ss := v2ServerSs[0]
				data, _ := json.Marshal(ss) //使用json方式 深拷贝对象
				var newNode tools.V2ServerShadowsocks
				json.Unmarshal(data, &newNode)
				newNode.Id = 0
				newNode.Name = newNode.Name + ZhongZhuanName
				newNode.Host = ZhongZhuanUrl
				newNode.Port = strconv.Itoa(zhong_zhuan_port)
				newNode.Sort = sql.NullInt32{Int32: 222 + ss.Sort.Int32, Valid: true}
				newNode.ParentId = sql.NullInt32{Int32: int32(ss.Id), Valid: true}
				newNode.Show = SHOW
				db.Create(&newNode)
			}

		}

	}
}

// getAllRealUrls 打印输出批量添加的模板
func getAllRealUrls() {
	db := tools.GetDB()
	var v2ServerV2Rays []tools.V2ServerV2Ray
	db.Where("parent_id is null").Find(&v2ServerV2Rays)
	var builder strings.Builder
	for _, serverV2Ray := range v2ServerV2Rays {
		A := serverV2Ray.Host
		B := serverV2Ray.Port
		builder.WriteString(A)
		builder.WriteString(":")
		builder.WriteString(B)
		builder.WriteString("\n")
	}
	var v2ServerSs []tools.V2ServerShadowsocks
	db.Where("parent_id is null and port!=0").Find(&v2ServerSs)
	for _, ss := range v2ServerSs {
		A := ss.Host
		B := ss.Port
		builder.WriteString(A)
		builder.WriteString(":")
		builder.WriteString(B)
		builder.WriteString("\n")
	}

	println(builder.String())

}
