<template>
	<view>
		<text @click="openPopup()">{{selectorDate}} </text>

		<view v-if="showPopup" class="hl-popup" ref="popup" @tap="closedPopup()">
			<view class="shade-bg-view"></view>
			<view class="calendar" @tap.stop="" v-if="currentMonthItem != null">
				<view class="topHandleView">
					<button class="month-change-btn" type="default" @click="preMonthAction">上个月</button>
					<text>{{currentMonthItem.time}}</text>
					<button class="month-change-btn" @click="nextMonthAction">下个月</button>
				</view>
				<view class="weekdays-view">
					<view class="weekday" v-for="(weak,index) in weekDays" :key="index">
						<view class="block-item">
							{{weak}}
						</view>
					</view>
				</view>
				<view class="monthdays-view">
					<view class="monthday" v-for="(dayItem,index) in currentMonthItem.days" :key="index">
						<view class="block-item" :class="{'active': dayItem.enabled == true && dayItem.date == currentSelecteDate.date}" @click="didSelectorDate" :data-model="dayItem">
							<text class="sign" v-if="dayItem.status == 'today'">今日</text>
							{{dayItem.day}}
						</view>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>
<script>
	export default {
		data() {
			return {
				word: 'word',
				showPopup: false,
				currentMonthItem: null,
				currentSelecteDate : {},
				todayTimestamp : null,
				weekDays: [
					"日", "一", "二", "三", "四", "五", "六",
				]

			}
		},

		computed: {
			selectorDate() {
				return this.currentSelecteDate.date || '选择日期'
			},
		},

		mounted() {
			this.setBaseDateData()
		},

		methods: {

			setBaseDateData: function() {
				let todayDt = new Date()
				let today_y = todayDt.getFullYear()
				let today_m = todayDt.getMonth() + 1
				let today_d = todayDt.getDate()
				this.todayTimestamp = new Date(today_y, today_m - 1, today_d).valueOf()
				this.currentMonthItem = this.getCurrentMonthDayArray(today_y, today_m, this.todayTimestamp)
			},

			getCurrentMonthDayArray : function(y, m, todayTimestamp, ) {
				let days = new Date(y, m, 0).getDate()
				var dataArray = []
				for (let day = 1; day <= days; day++) {
					let dt = new Date(y, m - 1, day);
					let timesTamp = dt.valueOf()
					let wk = dt.getDay()
					var status = ""
					var enabled = true
					if (todayTimestamp > timesTamp) {
						status = 'past'
					} else if (todayTimestamp == timesTamp) {
						status = 'today'
					} else {
						status = 'future'
					}

					dataArray.push({
						day: day,
						month: m,
						year: y,
						enabled: enabled,
						weak: wk,
						status: status,
						timesTamp: timesTamp,
						date: y + '-' + m + '-' + day
					})
				}

				let model = dataArray[0]
				for (let begin = model.weak - 1; begin >= 0; begin--) {
					let newModel = {
						day: "",
						enabled: false,
						month: model.month,
						year: model.year,
						weak: begin,
						timesTamp: -1
					}
					dataArray.splice(0, 0, newModel)
				}
				return {
					days: dataArray,
					month: m,
					year : y,
					time: y + "-" + m
				}
			},
				
				
			openPopup : function(){
				this.showPopup = true
			},

			closedPopup: function() {
				this.showPopup = false
			},
			
			preMonthAction : function(){
				let currentMonth = this.currentMonthItem.month
				let currentYear = this.currentMonthItem.year
				if(currentMonth == 1){
					currentMonth = 12
					currentYear --
				}else{
					currentMonth --
				}
				this.currentMonthItem = this.getCurrentMonthDayArray(currentYear, currentMonth, this.todayTimestamp)
			},
			
			nextMonthAction : function(){
				let currentMonth = this.currentMonthItem.month
				let currentYear = this.currentMonthItem.year
				if(currentMonth == 12){
					currentMonth = 1
					currentYear ++
				}else{
					currentMonth ++
				}
				this.currentMonthItem = this.getCurrentMonthDayArray(currentYear, currentMonth, this.todayTimestamp)
			},
			
			didSelectorDate: function(e) {
				let model = e.currentTarget.dataset.model
				if(model.enabled == true){
					this.currentSelecteDate = model
					this.$emit("selectorDate", this.currentSelecteDate)
					this.closedPopup()
				}
			},
		}
	} 
</script>
<style lang="scss">

	$hl-calendar-primary: #007aff !default;
	button::after{ border: none;} 
	.hl-popup {
		position: fixed;
		z-index: 9999;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		background-color: transparent;

	}

	.shade-bg-view {
		position: fixed;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
		background-color: #000;
		opacity: 0.25;
	}

	.calendar {
		position: fixed;
		width: 100%;
		right: 0;
		bottom: 0;
		left: 0;
		background-color: #FFF;
	}
	
	.topHandleView{
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
		width: 100%;
		padding: 16px 0;
		font-size: 16px;
		.month-change-btn{
			display: flex;
			align-items: center;
			color: $hl-calendar-primary;
			background-color: transparent;
			height: 32px;
			font-size: 14px;
			margin-left: 0;
			margin-right: 0;

		}
	}
	
	.weekdays-view,.monthdays-view{
		display: grid;
		width: 100%;
		grid-template-columns: repeat(7,1fr);
		.weekday,.monthday{
			display: flex;
			position: relative;
			padding-bottom: 100%;
			height: 0;
			overflow: hidden;
			.block-item{
				display: flex;
				position: absolute;
				width: 100%;
				height: 100%;
				justify-content: center;
				align-items: center;
				font-size: 16px;
				.sign{
					position: absolute;
					color: #FF0000;
					top: 0;
					font-size: 10px;
				}
				&.active{
					background-color: $hl-calendar-primary;
					color: #FFF;
					border-radius: 4px;
				}
			}
		}
	}
	

</style>