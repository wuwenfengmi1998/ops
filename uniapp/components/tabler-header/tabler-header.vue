<template>
	<header class="navbar navbar-expand-md d-print-none" >
	            <div class="container-xl">
					<button class="navbar-toggler collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu" aria-controls="navbar-menu" aria-expanded="false" aria-label="Toggle navigation">
					    <span class="navbar-toggler-icon"></span>
					</button>
	            <h1 class="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">        
	                <img src="/static/R.png" alt="EMS" class="navbar-brand-image">	                
	            </h1>
	            <div class="navbar-nav flex-row order-md-last">  
	                
	                <div v-if="user" class="nav-item dropdown">
	                    <a class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown" aria-label="Open user menu">
	                        <span class="avatar avatar-sm" ></span>
	                        <div class="d-none d-xl-block ps-2">
	                        <div>{{user.acc}}</div>
	                        <div class="mt-1 small text-secondary">{{user.name}}</div>
	                        </div>
	                    </a>
	                    <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
	                        <a class="dropdown-item" @click="goto_profile">账号</a>
	                        <div class="dropdown-divider"></div>
	                        <a class="dropdown-item">Settings</a>
	                        <a class="dropdown-item" @click="login_out">登出</a>
	                    </div>          
	                </div>
	                <div v-if="!user" class="nav-item d-md-flex me-3">
	                    <div class="btn-list">   
	                        <button class="btn btn-primary" @click="goto_login">
	                          <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"></path><path d="M8 7a4 4 0 1 0 8 0a4 4 0 0 0 -8 0"></path><path d="M6 21v-2a4 4 0 0 1 4 -4h4"></path><path d="M19 22v-6"></path><path d="M22 19l-3 -3l-3 3"></path></svg>
	                          登录
	                        </button>
	                    </div>
	                </div>
	            </div>
	            </div>
	        </header>
	        <header class="navbar-expand-md">
	        <div class="collapse navbar-collapse" id="navbar-menu">
	          <div class="navbar">
	            <div class="container-xl">
	              <div class="navbar-nav">
					  
	                <a class="nav-item" @click="goto_home">
	                    <div class="nav-link">
	                    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
	                      <svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M5 12l-2 0l9 -9l9 9l-2 0" /><path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" /><path d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6" /></svg>
	                    </span>
	                    <span class="nav-link-title">
	                      Home
	                    </span>
	                    </div>
	                </a>
					<a class="nav-item" @click="goto_ma">
					    <div class="nav-link">
					    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
					      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-calendar-month"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z" /><path d="M16 3v4" /><path d="M8 3v4" /><path d="M4 11h16" /><path d="M7 14h.013" /><path d="M10.01 14h.005" /><path d="M13.01 14h.005" /><path d="M16.015 14h.005" /><path d="M13.015 17h.005" /><path d="M7.01 17h.005" /><path d="M10.01 17h.005" /></svg>
						  </span>
					    <span class="nav-link-title">
					      日程管理
					    </span>
					    </div>
					</a>
					<a class="nav-item" >
					    <div class="nav-link">
					    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
					      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-building-warehouse"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M3 21v-13l9 -4l9 4v13" /><path d="M13 13h4v8h-10v-6h6" /><path d="M13 21v-9a1 1 0 0 0 -1 -1h-2a1 1 0 0 0 -1 1v3" /></svg>
						</span>
					    <span class="nav-link-title">
					      仓库管理
					    </span>
					    </div>
					</a>
					<a class="nav-item" >
					    <div class="nav-link">
					    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
					     <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-file-description"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M14 3v4a1 1 0 0 0 1 1h4" /><path d="M17 21h-10a2 2 0 0 1 -2 -2v-14a2 2 0 0 1 2 -2h7l5 5v11a2 2 0 0 1 -2 2z" /><path d="M9 17h6" /><path d="M9 13h6" /></svg>
						 </span>
					    <span class="nav-link-title">
					      工单管理
					    </span>
					    </div>
					</a>
					<a class="nav-item" >
					    <div class="nav-link">
					    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
					      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-shopping-cart"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M6 19m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" /><path d="M17 19m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" /><path d="M17 17h-11v-14h-2" /><path d="M6 5l14 1l-1 7h-13" /></svg>
						  </span>
					    <span class="nav-link-title">
					      采购管理
					    </span>
					    </div>
					</a>
					<a class="nav-item" @click="goto_qa">
					    <div class="nav-link">
					    <span class="nav-link-icon d-md-none d-lg-inline-block"><!-- Download SVG icon from http://tabler-icons.io/i/home -->
					      <svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-book"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M3 19a9 9 0 0 1 9 0a9 9 0 0 1 9 0" /><path d="M3 6a9 9 0 0 1 9 0a9 9 0 0 1 9 0" /><path d="M3 6l0 13" /><path d="M12 6l0 13" /><path d="M21 6l0 13" /></svg>
						  </span>
					    <span class="nav-link-title">
					      题库
					    </span>
					    </div>
					</a>
	              </div>
	              
	            </div>
	          </div>
	        </div>
	      </header>
</template>

<script>
	import { public_func } from '@/public_js.js';
	
	export default {
		name:"tabler-header",
		props: {
		      user_updata: {},
		    },
			watch: {
			    //haveTask 为接收参数的名称
			      user_updata: {
			        handler(newValue,oldValue) {
						this.updata_user(newValue);
			          //console.log(newValue)
					  //console.log(oldValue)			  
			        },
			        deep: true // 深度监听父组件传过来对象变化
			      }
			    },
		
		emits :{
		          updatedcount: null, 
		        },
		data() {
			return {
				
				user:null
			};
		},
			
		methods:{
			updata_user(user)
			{
				sessionStorage.setItem('user',JSON.stringify(user));
				var user2=localStorage.getItem('user');
				if(user2!=null)
				{
					localStorage.setItem('user',JSON.stringify(user));
				}
			},
			goto_profile()
			{
				uni.redirectTo({
					url:'/pages/settings/profile'
				});
			},
			
			goto_login()
			{
				uni.redirectTo({
					url:'/pages/login/login'
				});
			},
			goto_home()
			{
				uni.redirectTo({
					url:'/pages/index/index'
				});
			},
			goto_ma()
			{
				uni.redirectTo({
					url:'/pages/MA/MA'
				});
			},
			goto_qa()
			{
				uni.redirectTo({
					url:'/pages/qa/qa'
				});
			},
			
			login_out()
			{
				console.log("logout")
				
				sessionStorage.removeItem('user');
				localStorage.removeItem('user');
				this.user=null;
				
				
				this.$emit('updatedcount',this.user);  // 通过$emit触发事件，第二个参数就是传递的参数
				
			}
		},
		mounted() {
			console.log("mounted")
			var user1=sessionStorage.getItem('user');//先检查有没有登录
			var user2=localStorage.getItem('user');
			//console.log(user1)
			if(user1==null)
			{
				 console.log("nosession")
				
				if(user2==null)
				{
					console.log("nolocal")
					this.user=null;
					
				}else
				{
					sessionStorage.setItem('user',user2);
					this.user=JSON.parse(user2);
					
				} 
			}else
			{
				this.user=JSON.parse(user1);
				
			}
			
			this.$emit('updatedcount',this.user);  // 通过$emit触发事件，第二个参数就是传递的参数
		}
		
		
	}
</script>

<style>

</style>