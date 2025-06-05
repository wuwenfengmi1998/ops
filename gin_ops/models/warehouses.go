package models

func Warehouse_get_total_pages() int64 {
	var all_page int64 = 0

	DB.Model(&Warehouse{}).Count(&all_page)
	var repos_per_page = int64(Configs["warehouses"].(map[string]interface{})["repos_per_page"].(int))
	return (all_page / repos_per_page) + 1

}

func Warehouse_get_warehouses(page int64) []Warehouse {

	if page == 0 {
		return nil
	}

	var pageSize = int(Configs["warehouses"].(map[string]interface{})["repos_per_page"].(int))
	var Warehouses []Warehouse
	offset := int((int(page) - 1) * pageSize)
	DB.Model(Warehouse{}).
		Order("id DESC"). // 必须排序保证分页稳定
		Offset(offset).
		Limit(pageSize).
		Find(&Warehouses)

	return Warehouses
}

func Warehouse_get_items_from_whid(wh_id uint) []WarehouseItem {
	var seachf []WarehouseItem
	var seach WarehouseItem
	seach.WarehouseID = wh_id
	DB.Where(&seach).Order("id DESC").Find(&seachf)
	return seachf
}
