<template>

	<tabler-header type="mini" ref="heard"></tabler-header>
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

						<setting-menu></setting-menu>

						<div class="col-12 col-md-9 d-flex flex-column">
							<div class="card-body">

								<h2 class="mb-4">信息设置</h2>
								<!-- <h3 class="card-title">Profile Details</h3> -->
								<div class="row align-items-center">
									<div class="col-auto">

										<avatar size="xl" :url="user_info.AvatarPath"></avatar>
									</div>
									<div class="col-auto"><a class="btn" @click="showPopup">更改头像 </a></div>
									<!-- <div class="col-auto"><a href="#" class="btn btn-ghost-danger">
	                        删除头像
	                      </a></div> -->
								</div>
								<!-- <h3 class="card-title mt-4">Business Profile</h3> -->
								<div class="row g-3">
									<div class="col-md">
										<div class="form-label">名字</div>
										<input type="text" class="my_input_field" v-model="user_info.Username"
											maxlength="30">
									</div>

									<div class="col-md">
										<div class="form-label">备注</div>
										<input type="text" class="my_input_field" v-model="user_info.FirstName"
											maxlength="50">
									</div>
								</div>

								<div class="mb-3">
									<label class="form-label">生日</label>
									<div class="row g-2">
										<div class="col-5">

											<uni-datetime-picker type="date" :clear-icon="false" v-model="birthdate" />
										</div>
									</div>
								</div>


							</div>

							<div class="card-footer bg-transparent mt-auto">
								<div class="btn-list justify-content-end">

									<a href="#" class="btn btn-primary">
										提交
									</a>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

	</div>

	<tabler-footer></tabler-footer>

	<uni-popup ref="popup" type="center">
		<view class="popup-content">
			这里是弹框内容
			<button @click="closePopup">关闭</button>
		</view>
	</uni-popup>


</template>

<script>
	export default {
		data() {
			return {
				user_info: {
					AvatarPath: "",
					Username: "",
					FirstName: "",
					Birthdate: "",
				},
				birthdate: "",
			}
		},
		methods: {
			showPopup() {
				this.$refs.popup.open();
			},
			closePopup() {
				this.$refs.popup.close();
			}
		},

		mounted() {

			if (this.$refs.heard.is_login) {
				this.user_info = this.$refs.heard.user_info
				this.birthdate = this.user_info.Birthdate.substring(0, 10)
			} else {

			}



		}
	}
</script>

<style>

</style>