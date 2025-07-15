package routers

import (
	"fmt"
	"saas/models"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// 定义接收JSON数据的结构体
type Add_item_json struct {
	// 使用结构体标签指定JSON字段名
	WarehouseID uint    `json:"warehouse_id"`
	Name        string  `json:"item_name"`
	Info        string  `json:"item_info"`
	SN          string  `json:"item_sn"`
	Who         string  `json:"item_who"`
	Int         int     `json:"item_int"`
	Consts      float32 `json:"item_consts"`
}

func V1_warehouses_api(r *gin.RouterGroup) {

	r.POST("/create", func(ctx *gin.Context) {

		//先判断是否已经登录
		//获取中间件处理的结果

		user_info, is_login := ctx.Get("user_info")
		if is_login == true {

			//转换传进来的数据
			var jsonData map[string]interface{}
			data, is_have_data := ctx.Get("data")
			if is_have_data {
				if err := mapstructure.Decode(data, &jsonData); err == nil {
					//fmt.Println(jsonData)
					if jsonData["warehouses_name"].(string) != "" {
						warehouses_data := models.Warehouse{
							Name:         jsonData["warehouses_name"].(string),
							Info:         jsonData["warehouses_info"].(string),
							CreatorID:    user_info.(*models.User_info).UserID,
							Capacity:     0,
							UsedCapacity: 0,
							Location:     "local",
						}
						models.DB.Create(&warehouses_data) // 传入指针
						//fmt.Println(dberr.Error)

						Return_json(ctx, "api_ok", nil)

					} else {
						Return_json(ctx, "warehouses_name_err", nil)

					}
				} else {
					Return_json(ctx, "json_error", nil)

				}
			} else {
				Return_json(ctx, "json_error", nil)
			}

		} else {
			//fmt.Println("no loged")
			Return_json(ctx, "user_no_sign", nil)

		}

	})

	r.POST("/add_item", func(ctx *gin.Context) {

		//先判断是否已经登录
		//获取中间件处理的结果
		user_info_, is_login := ctx.Get("user_info")

		if is_login == true {
			user_info := user_info_.(*models.User_info)

			data, is_have_data := ctx.Get("data")
			if is_have_data {
				//转换传进来的数据
				var item Add_item_json
				if err := mapstructure.Decode(data, &item); err == nil {
					//先判断是否有权限

					//（还没弄）

					//后插入数据

					var add_item_temp models.WarehouseItem
					add_item_temp.CreatedByID = user_info.UserID
					add_item_temp.WarehouseID = item.WarehouseID
					add_item_temp.Name = item.Name
					add_item_temp.SerialNumber = item.SN
					add_item_temp.Description = item.Info
					add_item_temp.Destiny = item.Who
					add_item_temp.Quantity = item.Int
					add_item_temp.ItemValue = int(item.Consts * 100)

					fmt.Println(data)
					// //插入一条数据
					// models.DB.Create(&add_item_temp)
					// //更新仓库信息
					// var seach_wh models.Warehouse
					// seach_wh.ID = item.WarehouseID
					// var out_wh models.Warehouse
					// models.DB.Where(&seach_wh).First(&out_wh)
					// out_wh.UsedCapacity += 1
					// models.DB.Where(&seach_wh).Updates(&out_wh)

					Return_json(ctx, "api_ok", nil)

				} else {
					Return_json(ctx, "json_error", nil)

				}

			} else {
				Return_json(ctx, "json_error", nil)
			}

		} else {
			//fmt.Println("no loged")
			Return_json(ctx, "user_no_sign", nil)

		}

	})

	r.GET("/get_items/:wh_id", func(ctx *gin.Context) {

		// var seachf []models.WarehouseItem
		// //先判断是否已经登录
		// //获取中间件处理的结果
		// is_login, _ := ctx.Get("is_login")
		// if is_login == true {
		// 	id := ctx.Param("wh_id")
		// 	id_int, err := strconv.ParseInt(id, 10, 0)
		// 	if err == nil {
		// 		if id_int > 0 {

		// 			seachf = models.Warehouse_get_items_from_whid(uint(id_int))
		// 			//fmt.Println(seachf)
		// 			Return_json(ctx, "api_ok", nil)

		// 		} else {
		// 			Return_json(ctx, "warehouses_id_err", nil)

		// 		}

		// 	} else {
		// 		Return_json(ctx, "warehouses_id_err", nil)

		// 	}

		// } else {
		// 	//fmt.Println("no loged")
		// 	Return_json(ctx, "user_no_sign", nil)

		// }

	})

}
