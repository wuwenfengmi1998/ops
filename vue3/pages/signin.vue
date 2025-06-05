<template>
	
	<view class="d-flex flex-column">
		<tabler-header type="mini"></tabler-header>
		
		<div class="page page-center">
			
			<div class="container container-normal py-4">
				<div class="row align-items-center g-4">
					<div class="col-lg">
						<div class="container-tight">
							<div class="text-center mb-4">
								<div class="navbar-brand navbar-brand-autodark"><img src="/static/logo.svg"
										height="36" alt=""></div>
							</div>
							<div class="card card-md">
								<div class="card-body">
									<h2 class="h2 text-center mb-4">登录你的账号</h2>
									

									<div class="mb-3">
										<label class="form-label">用户名</label>
										<input type="text" class="my-input-field" placeholder="输入你的用户名" autocomplete="off" maxlength="25">
									</div>
									<div class="mb-2">
										<label class="form-label">
											密码
										</label>
										
										<input type="password" class="my-input-field" password="true" placeholder="输入你的密码" autocomplete="off" maxlength="100">
										
									</div>
									<div class="mb-2">
										<label class="d-flex">
											<checkbox />
											<span class="form-check-label">保持登录</span>
										</label>
									</div>
									<div class="form-footer">
										<button class="btn btn-primary w-100">登录</button>
									</div>

								</div>


							</div>
							<div class="text-center text-secondary mt-3">
								还没账号？
								<a href="/sign-up/" tabindex="-1">点击注册</a>
								<span class="form-label-description">
									<a href="./forgot-password.html">忘记密码？</a>
								</span>
							</div>
						</div>
					</div>
					<div class="col-lg d-none d-lg-block">
						<img src="/static/illustrations/undraw_secure_login_pdn4.svg" height="300"
							class="d-block mx-auto" alt="">
					</div>
				</div>
			</div>
		</div>


	</view>
</template>

<script>
	export default {
		data() {
			return {

			}
		},
		methods: {

		}
	}
</script>

<style>
	.my123{
		  /* 设置边框：宽度、样式、颜色 */
		  border: 1px solid #e3e3e3; /* 2像素宽的红色实线边框 */
		  
		  /* 设置圆角，值越大，角越圆 */
		  border-radius: 3px; /* 也可以使用百分比，例如50%会变成圆形 */
		  
		  font-size: 24px;
		
		
		
		 
		 
		 
	}

</style>