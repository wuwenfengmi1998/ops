export const my_network_func = {
	host: "",
	port: 0,
	head_path: "/api/v1",

	post_json(path, json,callback) {
		var re_data = {}
		uni.request({
			header: {
				'Content-Type': 'application/json'
			},
			url: this.head_path + path,
			method: 'POST',
			data: json,
			timeout: 10000,
			success(res) {

				re_data["statusCode"] = res.statusCode
				re_data["data"] = res.data
				callback(re_data)
				
			},
			fail() {
				re_data["statusCode"] = -1
				callback(re_data)
				
			}

		});
	},
}