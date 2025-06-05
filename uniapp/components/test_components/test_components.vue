<template>
	<view>
		<button type="primary" @click="sendMsg">给父组件传值</button>
	</view>
</template>

<script>
	export default {
		data () {
			return {
				status: 'hello uni-app'
			}
		},
		mounted() {
			//this.$emit('myEvent',this.status)  // 通过$emit触发事件，第二个参数就是传递的参数
		},
		methods: {
			sendMsg () {
				//this.$emit('myEvent',this.status)  // 通过$emit触发事件，第二个参数就是传递的参数
			}
		}
	}
</script>
