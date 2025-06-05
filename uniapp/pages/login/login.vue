<template>
	<view>
		<!-- <tabler-header></tabler-header> -->
		
		<body  class=" d-flex flex-column">
		    
		    <div class="page page-center">
		      <div class="container container-normal py-4">
		        <div class="row align-items-center g-4">
		          <div class="col-lg">
		            <div class="container-tight">
		              <div class="text-center mb-4">
		                
		              </div>
		              <div class="card card-md">
		                <div class="card-body">
							<button class="btn" @click="goto_home">
								<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-reply" viewBox="0 0 16 16">
									<path d="M6.598 5.013a.144.144 0 0 1 .202.134V6.3a.5.5 0 0 0 .5.5c.667 0 2.013.005 3.3.822.984.624 1.99 1.76 2.595 3.876-1.02-.983-2.185-1.516-3.205-1.799a8.74 8.74 0 0 0-1.921-.306 7.404 7.404 0 0 0-.798.008h-.013l-.005.001h-.001L7.3 9.9l-.05-.498a.5.5 0 0 0-.45.498v1.153c0 .108-.11.176-.202.134L2.614 8.254a.503.503 0 0 0-.042-.028.147.147 0 0 1 0-.252.499.499 0 0 0 .042-.028l3.984-2.933zM7.8 10.386c.068 0 .143.003.223.006.434.02 1.034.086 1.7.271 1.326.368 2.896 1.202 3.94 3.08a.5.5 0 0 0 .933-.305c-.464-3.71-1.886-5.662-3.46-6.66-1.245-.79-2.527-.942-3.336-.971v-.66a1.144 1.144 0 0 0-1.767-.96l-3.994 2.94a1.147 1.147 0 0 0 0 1.946l3.994 2.94a1.144 1.144 0 0 0 1.767-.96v-.667z"/>
								</svg>
								返回
							</button>
		                  <h2 class="h2 text-center mb-4">登录您的账号</h2>
		                  
		                    <div class="mb-3">
								<label class="form-label">
									用户名
								</label>
		                      <uni-easyinput v-model="username" class="input" type="text" placeholder="输入用户名" maxlength="24"></uni-easyinput>
		                    </div>
		                    <div class="mb-2">
							  <label class="form-label">
								密码
							  </label>
							  <uni-easyinput v-model="password" class="input_pass" type="password" placeholder="输入密码" maxlength="64" @confirm="login"></uni-easyinput>
		                     
		                    </div>
		                    <div class="mb-2">
								<uni-fav :checked="autologin" :star="false" class="favBtn" @click="autologin_change" :content-text="contentText"/>
		                      
		                    </div>
		                    <div class="form-footer">
		                      <button @click="login" class="btn btn-primary w-100">注册或登录</button>
		                    </div>
		                  
		                </div>
		                
		              </div>
		              <div class="text-center text-secondary mt-3">
		                注意：如果账号未注册，将会自动注册后自动登录
		              </div>
		            </div>
		          </div>
		          <div class="col-lg d-none d-lg-block">
		            <img src="/static/illustrations/undraw_secure_login_pdn4.svg" height="300" class="d-block mx-auto" alt=""> 
		          </div>
		        </div>
		      </div>
		    </div>
		    
			<view>
				<!-- 提示信息弹窗 -->
				<uni-popup ref="message" type="message">
					<uni-popup-message :type="msgType" :message="messageText" :duration="2000"></uni-popup-message>
				</uni-popup>
			</view>
			
		  </body>
	</view>
	
	<tabler-footer></tabler-footer>
</template>

<script>


	import { public_func } from '@/public_js.js';
	export default {
		data() {
			return {	
				username:"",
				password:"",
				autologin: false,
				msgType:"success",
				message:"这是一条成功提示",
				messageText:"",
				contentText: {
					contentDefault: '自动登录',
					contentFav: '自动登录'
				}
			}
		},
		methods: {
			
			messageToggle(type,msg) {
				this.msgType=type;
				this.messageText=msg;
				this.$refs.message.open();
			},
			goto_home(){
				uni.redirectTo({
					url:"/pages/index/index"
				});
			},
			login(){
				console.log("try login")
				if(this.username==""||this.password=="")
				{
					//console.log(baseUrl);
					this.messageToggle('error',"账号或密码为空");
					return ;
				}		
				if(this.username.length<5)
				{
					this.messageToggle('error',"账号少于5位");
					return ;
				}
				
				//let user=public_func.login_user(this.username,this.password);
				//console.log(public_func.login_user(this.username,this.password));
				var that = this;
				uni.request({
					header: {'Content-Type': 'application/x-www-form-urlencoded'},
					url:public_func.baseUrl+"api/?ac=login&do=login",
					method:'POST',
					data:{
						acc:this.username,
						pass:this.password,
					},
					timeout:10000,
					success(res) {
						console.log(res.data);
						if(res.data['error_code']==0)
						{
							//login ok
							sessionStorage.setItem('user',JSON.stringify(res.data['user']));
							
							if(that.autologin)
							{
								localStorage.setItem('user',JSON.stringify(res.data['user']));
								
							}
												
							
							that.messageToggle('success',"登录成功");
							setTimeout(function(){
								that.goto_home();
							}, 1000);
					
							
						}else
						{
							that.messageToggle('error',res.data['error_msg']);
							return ;
						}
						
					},
					fail() {
						that.messageToggle('error',"网络错误");
					}
					
				});
				
				
			},
			autologin_change(){
				if(this.autologin)
				{
					this.autologin=false;
				}else
				{
					this.autologin=true;
				}
			}
		}
	}
</script>

<style lang="scss">

.input{
		

	}
.input_pass{
	
}
</style>
