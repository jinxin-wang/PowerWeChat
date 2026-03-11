package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseGetFreightTemplateList 获取运费模板列表响应
type ResponseGetFreightTemplateList struct {
	response.ResponseStore
	// 运费模板列表
	TemplateList []FreightTemplateListItem `json:"template_list"`
	// 分页游标
	NextKey string `json:"next_key"`
	// 是否还有更多
	HasMore bool `json:"has_more"`
}

// FreightTemplateListItem 运费模板列表项
type FreightTemplateListItem struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板名称
	Name string `json:"name"`
	// 计费类型：1-按件计费，2-按重量计费
	ValuationType int `json:"valuation_type"`
}
