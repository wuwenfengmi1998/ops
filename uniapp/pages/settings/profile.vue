<template>
	<tabler-header ref="the_heard" @updatedcount="get_user" :user_updata="user"></tabler-header>
	<div class="page-wrapper">
	        <!-- Page header -->
	        <div class="page-header d-print-none">
	          <div class="container-xl">
	            <div class="row g-2 align-items-center">
	              <div class="col">
	                <h2 class="page-title">
	                  设置
	                </h2>
	              </div>
	            </div>
	          </div>
	        </div>
	        <!-- Page body -->
	        <div class="page-body">
	          <div class="container-xl">
	            <div class="card">
	              <div class="row g-0">
	                <div class="col-12 col-md-3 border-end">
						<business-settings />
	                  
	                </div>
	                <div class="col-12 col-md-9 d-flex flex-column">
	                  <div class="card-body">
	                    <h2 class="mb-4">我的账号</h2>
						<div class="row g-3">
							<div class="col-md">
							  <div class="form-label">ID: {{user_id}}</div>
							  
							</div>
							<div class="col-md">
							<div class="form-label">账号: {{user_acc}}</div>
							
							</div>			  
						</div>
	                    <h3 class="card-title">Profile Details</h3>
	                    <div class="row align-items-center">
	                      <div class="col-auto"><span class="avatar avatar-xl" style="background-image: url(./static/avatars/000m.jpg)"></span>
	                      </div>
	                      <div class="col-auto"><a class="btn">
	                          Change avatar
	                        </a></div>
	                      <div class="col-auto"><a class="btn btn-ghost-danger">
	                          Delete avatar
	                        </a></div>
	                    </div>
	                    <h3 class="card-title mt-4">Business Profile</h3>
	                    
	                    <h3 class="card-title mt-4">Email</h3>
	                    <p class="card-subtitle">This contact will be shown to others publicly, so choose it carefully.</p>
	                    <div>
	                      <div class="row g-2">
	                        <div class="col-auto">
	                          <uni-easyinput v-model="input_email" class="input" type="text" placeholder="输入邮箱" maxlength="128"></uni-easyinput>
	                        </div>
	                        <div class="col-auto">
								<a class="btn" @click="change_email">Change</a></div>
	                      </div>
	                    </div>
	                    <h3 class="card-title mt-4">Password</h3>
	                    <p class="card-subtitle">You can set a permanent password if you don't want to use temporary login codes.</p>
	                    <div>
							<div class="col-auto">
							  <uni-easyinput v-model="input_oldpass" class="input mb-3" type="password" placeholder="输入旧密码" maxlength="64"></uni-easyinput>
							  <uni-easyinput v-model="input_newpass" class="input mb-3" type="password" placeholder="输入新密码" maxlength="64"></uni-easyinput>
							  <uni-easyinput v-model="input_cofpass" class="input mb-3" type="password" placeholder="确认新密码" maxlength="64"></uni-easyinput>
							</div>
	                      <a class="btn" @click="change_pass">Set new password</a>
	                    </div>
	                   
	                  </div>
	                  
	                </div>
	              </div>
	            </div>
	          </div>
	        </div>
	        <footer class="footer footer-transparent d-print-none">
	          <div class="container-xl">
	            <div class="row text-center align-items-center flex-row-reverse">
	              <div class="col-lg-auto ms-lg-auto">
	                <ul class="list-inline list-inline-dots mb-0">
	                  <li class="list-inline-item"><a href="https://tabler.io/docs" target="_blank" class="link-secondary" rel="noopener">Documentation</a></li>
	                  <li class="list-inline-item"><a href="./license.html" class="link-secondary">License</a></li>
	                  <li class="list-inline-item"><a href="https://github.com/tabler/tabler" target="_blank" class="link-secondary" rel="noopener">Source code</a></li>
	                  <li class="list-inline-item">
	                    <a href="https://github.com/sponsors/codecalm" target="_blank" class="link-secondary" rel="noopener">
	                      <!-- Download SVG icon from http://tabler-icons.io/i/heart -->
	                      <svg xmlns="http://www.w3.org/2000/svg" class="icon text-pink icon-filled icon-inline" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"></path></svg>
	                      Sponsor
	                    </a>
	                  </li>
	                </ul>
	              </div>
	              <div class="col-12 col-lg-auto mt-3 mt-lg-0">
	                <ul class="list-inline list-inline-dots mb-0">
	                  <li class="list-inline-item">
	                    Copyright © 2023
	                    <a href="." class="link-secondary">Tabler</a>.
	                    All rights reserved.
	                  </li>
	                  <li class="list-inline-item">
	                    <a href="./changelog.html" class="link-secondary" rel="noopener">
	                      v1.0.0-beta20
	                    </a>
	                  </li>
	                </ul>
	              </div>
	            </div>
	          </div>
	        </footer>
	      </div>
		  
		  <view>
		  	<!-- 提示信息弹窗 -->
		  	<uni-popup ref="message" type="message">
		  		<uni-popup-message :type="msgType" :message="messageText" :duration="2000"></uni-popup-message>
		  	</uni-popup>
		  </view>
</template>

<script>
	import { public_func } from '@/public_js.js';
	
	export default {
		
		
		data() {
			return {
				msgType:"success",
				message:"这是一条成功提示",
				messageText:"",
				user:null,
				user_id:0,
				user_acc:"",
				input_email:"",
				input_oldpass:"",
				input_newpass:"",
				input_cofpass:"",
			}
		},
		onReady() {
			console.log("onready")
		},
		mounted() {
			console.log("mounted")
		},
		methods: {
			messageToggle(type,msg) {
				this.msgType=type;
				this.messageText=msg;
				this.$refs.message.open();
			},
			get_user(user)
			{
				//console.log(user)
				if(user==null)//如果还没登录 跳转到登录
				{
					uni.redirectTo({
						url:'/pages/login/login'
					});
				}else
				{
					this.user=user;
					this.user_id=user.id;
					this.user_acc=user.acc;
					if(user.email!=null)
					{
						this.input_email=user.email;
					}
				}
				
			},
			change_email()
			{
				if(this.input_email=="")
				{
					this.messageToggle("error","请输入email");
					return;
				}
				if(public_func.isValidEmail(this.input_email))
				{
					//try change email
					var that = this;
					uni.request({
						header: {'Content-Type': 'application/x-www-form-urlencoded'},
						url:public_func.baseUrl+"api/?ac=users&do=change_email",
						method:'POST',
						data:{
							id:that.user.id,
							acc:that.user.acc,
							pass:that.user.pass,
							email:that.input_email,
						},
						timeout:10000,
						success(res) {
							console.log(res.data);
							if(res.data['error_code']==0)
							{
								that.user=res.data['user'];
								that.messageToggle('success',"邮箱修改成功");
								
								return ;
								
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
				}else
				{
					this.messageToggle("error","邮箱格式不正确");
				}
				
			},
			change_pass()
			{
				if(this.input_oldpass=="")
				{
					this.messageToggle("error","请输入旧密码");
					return;
				}
				if(this.input_newpass=="")
				{
					this.messageToggle("error","请输入新密码");
					return;
				}
				if(this.input_cofpass!=this.input_newpass)
				{
					this.messageToggle("error","确认密码错误");
					return;
				}
				var that = this;
				uni.request({
					header: {'Content-Type': 'application/x-www-form-urlencoded'},
					url:public_func.baseUrl+"api/?ac=users&do=change_pass",
					method:'POST',
					data:{
						id:that.user.id,
						acc:that.user.acc,
						pass:that.input_oldpass,
						newpass:that.input_cofpass
					},
					timeout:10000,
					success(res) {
						console.log(res.data);
						if(res.data['error_code']==0)
						{
							that.user=res.data['user'];
							that.messageToggle('success',"密码修改成功");
							that.input_oldpass="";
							that.input_newpass="";
							that.input_cofpass="";
							setTimeout(function() {
							    that.$refs.the_heard.login_out();
							}, 1000);
							
							return ;
							
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
				
				
			}
			
		}
	}
</script>

<style>

</style>
