<script>
	// Import Tabler CSS
	import '@tabler/core/dist/css/tabler.min.css';
	
	// Import Tabler JavaScript (if needed for components that require JavaScript)
	import '@tabler/core/dist/js/tabler.min.js';
	
	
	
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style>
	/*每个页面公共css */
	 /* #ifdef H5 */ 
	    uni-page-head { display: none; } 
	    /* #endif */
</style>
