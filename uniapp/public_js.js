
export const public_func={
	baseUrl:"http://test.ems.local/",

	formatDate(date) {
	    var year = date.getFullYear();
	    var month = (date.getMonth() + 1).toString().padStart(2, '0');
	    var day = date.getDate().toString().padStart(2, '0');
	    return `${year}-${month}-${day}`;
	},
	test(tttt)
	{
		console.log(tttt)
	},
	isValidEmail(email) 
	{ 
		// 定义邮箱的正则表达式 
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/; 
		// 使用正则表达式测试邮箱 
		return emailRegex.test(email); 
	}
}