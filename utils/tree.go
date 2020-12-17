package utils

import (
	"encoding/json"
)

type NavListJson struct {
	Id       int            `json:"_key"`
	Name     string         `json:"_name"`
	NavType  string         `json:"_type"`
	Text     string         `json:"text"`
	Icon     string         `json:"icon"`
	ParentId int            `json:"dataGroup" `
	Router   string         `json:"to"`
	Exact    bool           `json:"exact"`
	Target   string         `json:"target"`
	Children []*NavListJson `json:"_children"`
}

var Jsondata []byte //存储json数据

func NavListTree(navList []*NavListJson) (res []*NavListJson) {
	Data := navList
	var result []*NavListJson
	for _, item := range Data {
		if item.ParentId == 0 {
			Anode := item         //父节点
			makeTree(Data, Anode) //调用生成tree
			transformJson(Anode)  //转化为json
			jsontoTree(Jsondata)  //json 转为struct
			result = append(result, Anode)
		}
	}
	return result
}

func makeTree(Allnode []*NavListJson, node *NavListJson) { //参数为父节点，添加父节点的子节点指针切片
	childs, _ := haveChild(Allnode, node) //判断节点是否有子节点并返回
	if childs != nil {
		node.Children = append(node.Children, childs[0:]...) //添加子节点
		for _, v := range childs {                           //查询子节点的子节点，并添加到子节点
			_, has := haveChild(Allnode, v)
			if has {
				makeTree(Allnode, v) //递归添加节点
			}
		}
	}
}

func haveChild(Allnode []*NavListJson, node *NavListJson) (childs []*NavListJson, yes bool) {
	for _, v := range Allnode {
		if v.ParentId == node.Id {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}

func transformJson(Data *NavListJson) { //转为json

	Jsondata, _ = json.Marshal(Data)
}

func jsontoTree(jsondata []byte) { //json转struct
	var tree_node *NavListJson
	json.Unmarshal(jsondata, &tree_node)
}
