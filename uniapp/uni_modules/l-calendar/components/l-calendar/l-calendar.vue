<template>
	<view @touchmove.stop.prevent>
		<view class="l-calendar-box" :class="{'calendar-box-show': value}">
			<view class="calendar-top">
				<view>{{title}}</view>
				<view class="close l-icons icon-shanchu" hover-class="l-opacity" :hover-stay-time="150" @tap="hide">
				</view>
			</view>

			<view class="date-box">
				<view class=" date-arrowleft l-icons icon-youjiantou" :style="{ color: yearChangeColor }"
					hover-class="l-opacity" :hover-stay-time="150"  @tap="changeYear(false)">
				</view>
				<view class=" date-arrowleft l-icons  icon-zuojiantou" :style="{ color: monthChangeColor }"
					hover-class="l-opacity" :hover-stay-time="150" @tap="changeMonth(false)"></view>
				<view class="title-datetime">{{ showTitle }}</view>
				<view class=" date-arrowright l-icons icon-youjiantou2" :style="{ color: monthChangeColor }"
					hover-class="l-opacity" :hover-stay-time="150" @tap="changeMonth(true)"></view>
				<view class=" date-arrowright l-icons  icon-youjiantou1" :style="{ color: yearChangeColor }"
					hover-class="l-opacity" :hover-stay-time="150"  @tap="changeYear(true)"></view>
			</view>
			<view class="date-weekday">
				<view class="date-weekday-item">日</view>
				<view class="date-weekday-item">一</view>
				<view class="date-weekday-item">二</view>
				<view class="date-weekday-item">三</view>
				<view class="date-weekday-item">四</view>
				<view class="date-weekday-item">五</view>
				<view class="date-weekday-item">六</view>
			</view>
			<view class="date-content" :style="{ height: dateHeight * 6 + 'px' }">
				<block v-for="(item, index) in weekdayArr" :key="index">
					<view class="date-weekday-item"></view>
				</block>
				<view class="date-weekday-item" :class="{
						'l-opacity': isDisable(year, month, index + 1),
						'start-date': (isRange && startDate == `${year}-${month}-${index + 1}`) || !isRange,
						'end-date': (isRange && endDate == `${year}-${month}-${index + 1}`) || !isRange
					}" :style="{ backgroundColor: getColor(index, 1), height: dateHeight + 'px',padding:0}"
					v-for="(item, index) in daysArr" :key="index" @tap="dateClick(index)">
					<view class="date-content-item" :style="{ color: getColor(index, 2) }">
						<view>{{ index + 1 }}</view>
						<!-- 农历 -->
						<view class="custom-desc">
							{{ getText(index, startDate, endDate) }}
						</view>
					</view>
					<view class="date-content-item-desc" :style="{ color: activeColor }"
						v-if="!lunar && isRange && startDate == `${year}-${month}-${index + 1}` && startDate != endDate">
						{{ startText }}
					</view>
					<view class="date-content-item-desc" :style="{ color: activeColor }"
						v-if="!lunar && isRange && endDate == `${year}-${month}-${index + 1}`">{{ endText }}</view>
				</view>
				<view class="bg-mounth">{{ month }}</view>
			</view>

			<view class="calendar-text">
				<view class="calendar-result">
					<text>{{ !isRange ? activeDate : startDate }}</text>
					<text v-if="endDate">至{{ endDate }}</text>
				</view>
				<view class="calendar-btn">
					<button :style="{opacity:disabled ? '.5' : '1'}" :size="28" :disabled="disabled"
						@click="confireBtnClick(false)">确定
					</button>
				</view>
			</view>
		</view>

		<view class="mask" :class="[value ? 'mask-show' : '']" @tap="hide"></view>
	</view>
</template>
<script>
	import calendar from './calendar.js';
	export default {
		name: 'lCalendar',
		model: {
			prop: 'value',
			event: 'input'
		},
		props: {
			//双向绑定的值 用于展示/关闭日历
			value: {
				type: Boolean,
				default: false
			},
      //是否选择范围 true是 false选择单个日期
      isRange: {
      	type: Boolean,
      	default: false
      },
			//可切换最大年份
			maxYear: {
				type: Number,
				default: 2100
			},
			//可切换最小年份
			minYear: {
				type: Number,
				default: 1920
			},
			//最小可选日期 不在范围内日期禁选
			minDate: {
				type: String,
				default: '1920-01-01'
			},
			//最大可选日期
			maxDate: {
				type: String,
				default: '2100-1-1'
			},
      //组件标题
			title: {
				type: String,
				default: '日期选择'
			},
			//月份切换箭头颜色
			monthChangeColor: {
				type: String,
				default: '#999'
			},
			//年份切换箭头颜色
			yearChangeColor: {
				type: String,
				default: '#bfbfbf'
			},
			//默认日期字体颜色
			color: {
				type: String,
				default: '#333'
			},

			//选中日期字体颜色
			activeColor: {
				type: String,
				default: '#fff'
			},
			//选中日期背景色
			activeBgColor: {
				type: String,
				default: '#55BBF9'
			},
			//范围内日期背景色
			rangeBgColor: {
				type: String,
				default: 'rgba(85, 187, 249, 0.1)'
			},
			//范围内日期字体颜色
			rangeColor: {
				type: String,
				default: '#55BBF9'
			},

			//范围选择时生效 开始日期自定义文字
			startText: {
				type: String,
				default: '开始'
			},
			//范围选择时生效 结束日期自定义文字
			endText: {
				type: String,
				default: '结束'
			},
			//是否显示农历
			lunar: {
				type: Boolean,
				default: true
			},
			//初始化开始选中日期 格式： 2020-06-06 或 2020/06/06
			initStartDate: {
				type: String,
				default: ''
			},
			//初始化结束日期 格式： 2020-06-06 或 2020/06/06
			initEndDate: {
				type: String,
				default: ''
			}
		},
		data() {
			return {
				weekday: 1, // 星期几,值为1-7
				weekdayArr: [],
				days: 0, //当前月有多少天
				daysArr: [],
				showTitle: '',//当前年月标题
				year: 2020,
				month: 0,
				day: 0,
				startYear: 0,
				startMonth: 0,
				startDay: 0,
				endYear: 0,
				endMonth: 0,
				endDay: 0,
				today: '', //今天的日期
				activeDate: '', //当前选中日期
				startDate: '', //范围选择时的选中开始日期
				endDate: '', //范围选择时的选中结束日期
				isStart: true,
				min: null,
				max: null,
				dateHeight: 20
			};
		},
		computed: {
			dataChange() {
				return `${this.type}-${this.minDate}-${this.maxDate}-${this.initStartDate}-${this.initEndDate}`;
			},
			disabled() {
				return this.isRange && (!this.startDate || !this.endDate)
			}
		},
		watch: {
			dataChange(val) {
				this.init();
			},
		},
		created() {
			this.init();
		},
		methods: {
			getColor(index, type) {
				let color = type == 1 ? '' : this.color;
				let day = index + 1;
				let date = `${this.year}-${this.month}-${day}`;
				let timestamp = new Date(date.replace(/\-/g, '/')).getTime();
				let start = this.startDate.replace(/\-/g, '/');
				let end = this.endDate.replace(/\-/g, '/');
				if ((this.activeDate == date) || this.startDate == date || this.endDate == date) {
					color = type == 1 ? this.activeBgColor : this.activeColor;
				} else if (this.endDate && timestamp > new Date(start).getTime() && timestamp < new Date(end).getTime()) {
					color = type == 1 ? this.rangeBgColor : this.rangeColor;
				}
				return color;
			},
			getText(index, startDate, endDate) {
				let text = this.lunar ? this.getLunar(this.year, this.month, index + 1) : '';
				if (this.isRange) {
					if (this.lunar) {
						let date = `${this.year}-${this.month}-${index + 1}`;
						if (startDate == date && startDate != endDate) {
							text = this.startText;
						} else if (endDate == date) {
							text = this.endText;
						}
					}
				}
				return text;
			},
			getLunar(year, month, day) {
				let obj = calendar.solar2lunar(year, month, day);
				if (obj.IDayCn == '初一') {
					return obj.IMonthCn
				}
				return obj.IDayCn;
			},
			init() {
				this.dateHeight = uni.getSystemInfoSync().windowWidth / 7;
				let now = new Date();
				this.year = now.getFullYear();
				this.month = now.getMonth() + 1;
				this.day = now.getDate();
				this.today = `${now.getFullYear()}-${now.getMonth() + 1}-${now.getDate()}`;
				this.activeDate = this.today;
				this.min = this.initDate(this.minDate);
				this.max = this.initDate(this.maxDate);
				if (this.isDisable(this.year, this.month, this.day)) {
					this.year = this.min.year;
					this.month = this.min.month;
					this.day = this.min.day;
					this.activeDate = `${this.min.year}-${this.min.month}-${this.min.day}`;
					this.max = this.initDate(this.maxDate || this.minDate);
				}
				this.startDate = '';
				this.startYear = 0;
				this.startMonth = 0;
				this.startDay = 0;
				if (this.initStartDate) {
					let start = new Date(this.initStartDate.replace(/\-/g, '/'));
					if (!this.isRange) {
						this.year = start.getFullYear();
						this.month = start.getMonth() + 1;
						this.day = start.getDate();
						this.activeDate = `${start.getFullYear()}-${start.getMonth() + 1}-${start.getDate()}`;
					} else {
						this.startDate = `${start.getFullYear()}-${start.getMonth() + 1}-${start.getDate()}`;
						this.startYear = start.getFullYear();
						this.startMonth = start.getMonth() + 1;
						this.startDay = start.getDate();
						this.activeDate = '';
					}

				}
				this.endYear = 0;
				this.endMonth = 0;
				this.endDay = 0;
				this.endDate = '';
				if (this.initEndDate && this.isRange) {
					let end = new Date(this.initEndDate.replace(/\-/g, '/'));
					this.endDate = `${end.getFullYear()}-${end.getMonth() + 1}-${end.getDate()}`;
					this.endYear = end.getFullYear();
					this.endMonth = end.getMonth() + 1;
					this.endDay = end.getDate();
					this.activeDate = '';
					this.year = end.getFullYear();
					this.month = end.getMonth() + 1;
					this.day = end.getDate();
				}
				this.isStart = true;
				this.changeData();
			},
			//日期处理
			initDate(date) {
				let dateArr = date.split('-');
				return {
					year: Number(dateArr[0] || 1920),
					month: Number(dateArr[1] || 1),
					day: Number(dateArr[2] || 1)
				};
			},
			isDisable(year, month, day) {
				let bool = true;
				let date = `${year}/${month}/${day}`;
				let min = `${this.min.year}/${this.min.month}/${this.min.day}`;
				let max = `${this.max.year}/${this.max.month}/${this.max.day}`;
				let timestamp = new Date(date).getTime();
				if (timestamp >= new Date(min).getTime() && timestamp <= new Date(max).getTime()) {
					bool = false;
				}
				return bool;
			},
			generateArray(start, end) {
				return Array.from(new Array(end + 1).keys()).slice(start);
			},
			formatNum(num) {
				return num < 10 ? '0' + num : num + '';
			},
			//一个月有多少天
			getMonthDay(year, month) {
				let days = new Date(year, month, 0).getDate();
				return days;
			},
			// 获取当前日期是星期几
			getWeekday(year, month) {
				let date = new Date(`${year}/${month}/01 00:00:00`);
				return date.getDay();
			},
			changeMonth(isAdd) {
				if (isAdd) {
					let month = this.month + 1;
					let year = month > 12 ? this.year + 1 : this.year;
					if (year > this.minYear || year < this.maxYear) {
						this.month = month > 12 ? 1 : month;
						this.year = year;
						this.changeData();
					}
				} else {
					let month = this.month - 1;
					let year = month < 1 ? this.year - 1 : this.year;
					if (year > this.minYear || year < this.maxYear) {
						this.month = month < 1 ? 12 : month;
						this.year = year;
						this.changeData();
					}
				}
			},
			changeYear(isAdd) {
				let year = isAdd ? this.year + 1 : this.year - 1;
				if (year > this.minYear || year < this.maxYear) {
					this.year = year;
					this.changeData();
				}
			},
			changeData() {
				this.days = this.getMonthDay(this.year, this.month);
				this.daysArr = this.generateArray(1, this.days);
				this.weekday = this.getWeekday(this.year, this.month);
				this.weekdayArr = this.generateArray(1, this.weekday);
				this.showTitle = `${this.year}年${this.month}月`;
				if (!this.isRange) {
					this.confireBtnClick(true);
				}
			},
			dateClick: function(day) {
				day += 1;
				if (!this.isDisable(this.year, this.month, day)) {
					this.day = day;
					let date = `${this.year}-${this.month}-${day}`;
					if (!this.isRange) {
						this.activeDate = date;
					} else {
						let compare = new Date(date.replace(/\-/g, '/')).getTime() < new Date(this.startDate.replace(
							/\-/g, '/')).getTime();
						if (this.isStart || compare) {
							this.startDate = date;
							this.startYear = this.year;
							this.startMonth = this.month;
							this.startDay = this.day;
							this.endYear = 0;
							this.endMonth = 0;
							this.endDay = 0;
							this.endDate = '';
							this.activeDate = '';
							this.isStart = false;
						} else {
							this.endDate = date;
							this.endYear = this.year;
							this.endMonth = this.month;
							this.endDay = this.day;
							this.isStart = true;
						}
					}
				}
			},
			hide() {
				this.$emit('input', false)
			},
			getWeekText(date) {
				date = new Date(`${date.replace(/\-/g, '/')} 00:00:00`);
				let week = date.getDay();
				return '星期' + ['日', '一', '二', '三', '四', '五', '六'][week];
			},
			confireBtnClick(show) {
				if (!show) {
					this.hide();
				}
				if (!this.isRange) {
					let arr = this.activeDate.split('-');
					let year = +arr[0];
					let month = +arr[1];
					let day = +arr[2];
					//当前月有多少天
					let days = this.getMonthDay(year, month);
					let result = `${year}-${this.formatNum(month)}-${this.formatNum(day)}`;
					let weekText = this.getWeekText(result);
					let isToday = false;
					if (`${year}-${month}-${day}` == this.today) {
						//今天
						isToday = true;
					}
					let lunar = calendar.solar2lunar(year, month, day);
					this.$emit('change', {
						year: year,
						month: month,
						day: day,
						days: days,
						result: result,
						week: weekText,
						isToday: isToday,
						lunar: lunar
					});
				} else {
					if (!this.startDate || !this.endDate) return;
					let startMonth = this.formatNum(this.startMonth);
					let startDay = this.formatNum(this.startDay);
					let startDate = `${this.startYear}-${startMonth}-${startDay}`;
					let startWeek = this.getWeekText(startDate);
					let startLunar = calendar.solar2lunar(this.startYear, startMonth, startDay);

					let endMonth = this.formatNum(this.endMonth);
					let endDay = this.formatNum(this.endDay);
					let endDate = `${this.endYear}-${endMonth}-${endDay}`;
					let endWeek = this.getWeekText(endDate);
					let endLunar = calendar.solar2lunar(this.endYear, endMonth, endDay);
					this.$emit('change', {
						startYear: this.startYear,
						startMonth: this.startMonth,
						startDay: this.startDay,
						startDate: startDate,
						startWeek: startWeek,
						startLunar: startLunar,
						endYear: this.endYear,
						endMonth: this.endMonth,
						endDay: this.endDay,
						endDate: endDate,
						endWeek: endWeek,
						endLunar: endLunar
					});
				}
			}
		}
	};
</script>

<style lang="scss" scoped>
	@font-face {
		font-family: 'l-icons';
		src: url('data:font/ttf;charset=utf-8;base64,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') format('truetype');
		font-weight: normal;
		font-style: normal;
		font-display: swap;
	}

	.l-icons {
		font-family: 'l-icons';
		font-size: 38rpx;
		color: #333333;
		font-style: normal;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale
	}

	.icon-shanchu:before {
		content: "\e6f0";
	}

	.icon-youjiantou:before {
		content: "\e600";
	}

	.icon-zuojiantou:before {
		content: "\e623";
	}

	.icon-youjiantou1:before {
		content: "\e607";
	}

	.icon-youjiantou2:before {
		content: "\e61f";
	}
	
	.l-calendar-box {
		width: 100%;
		position: fixed;
		left: 0;
		right: 0;
		bottom: 0;
		z-index: 9999;
		visibility: hidden;
		transform: translate3d(0, 100%, 0);
		transform-origin: center;
		transition: all 0.3s ease-in-out;
		min-height: 20rpx;
		
		.calendar-top {
			width: 100%;
			height: 80rpx;
			padding: 0 40rpx;
			display: flex;
			justify-content: center;
			align-items: center;
			box-sizing: border-box;
			font-size: 30rpx;
			font-weight: bold;
			background-color: #fff;
			color: #333;
			position: relative;
			border-top-left-radius: 20rpx;
			border-top-right-radius: 20rpx;
			overflow: hidden;
			.close {
				position: absolute;
				right: 30rpx;
				top: 50%;
				transform: translateY(-50%);
				color: #999;
			
			}
		}
		
		.date-box {
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			padding: 20rpx 0 30rpx;
			background-color: #fff;
			.date-arrowleft {
				margin-right: 32rpx;
			}
			
			.date-arrowright {
				margin-left: 32rpx;
			}
			
			.title-datetime {
				padding: 0 16rpx;
				color: #333;
				font-size: 30rpx;
				line-height: 30rpx;
				font-weight: bold;
			}
		}
		
		.date-weekday {
			width: 100%;
			display: flex;
			align-items: center;
			background-color: #fff;
			font-size: 24rpx;
			line-height: 24rpx;
			color: #555;
			box-shadow: 0 15rpx 20rpx -15rpx #efefef;
			position: relative;
			z-index: 2;
			.date-weekday-item {
				width: 14.2857%;
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 12rpx 0;
				overflow: hidden;
				position: relative;
				z-index: 2;
			}
		}
		.date-content {
			width: 100%;
			display: flex;
			flex-wrap: wrap;
			padding: 12rpx 0;
			box-sizing: border-box;
			background-color: #fff;
			position: relative;
			align-content: flex-start;
			.date-weekday-item {
				width: 14.2857%;
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 12rpx 0;
				overflow: hidden;
				position: relative;
				z-index: 2;
				
				.date-content-item {
					width: 80rpx;
					height: 80rpx;
					display: flex;
					align-items: center;
					justify-content: center;
					flex-direction: column;
					font-size: 32rpx;
					line-height: 32rpx;
					position: relative;
					border-radius: 50%;
					background-color: transparent;
					.custom-desc {
						width: 100%;
						font-size: 24rpx;
						line-height: 24rpx;
						transform: scale(0.8);
						transform-origin: center center;
						text-align: center;
					}
					
				}
				.date-content-item-desc {
					width: 100%;
					font-size: 24rpx;
					line-height: 24rpx;
					position: absolute;
					left: 0;
					transform: scale(0.8);
					transform-origin: center center;
					text-align: center;
					bottom: 8rpx;
					z-index: 2;
				}
			}
			.start-date {
				border-top-left-radius: 8rpx;
				border-bottom-left-radius: 8rpx;
			}
			
			.end-date {
				border-top-right-radius: 8rpx;
				border-bottom-right-radius: 8rpx;
			}
			
			.bg-mounth {
				position: absolute;
				font-size: 260rpx;
				line-height: 260rpx;
				left: 50%;
				top: 50%;
				transform: translate(-50%, -50%);
				color: #f5f5f7;
				z-index: 1;
			}
		}
		
		
		.calendar-text {
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			flex-direction: column;
			background-color: #fff;
			padding: 0 42rpx 30rpx;
			box-sizing: border-box;
			font-size: 24rpx;
			color: #666;
			.calendar-result {
				height: 48rpx;
				transform: scale(0.9);
				transform-origin: center 100%;
			}
			
			.calendar-btn {
				width: 100%;
			
				button {
					background-color: #55BBF9;
					color: #fff;
					height: 72rpx;
					line-height: 72rpx;
					font-size: 32rpx
				}
			}
		}
		
		
		
	}
	.mask {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: rgba(0, 0, 0, 0.6);
		z-index: 9996;
		transition: all 0.3s ease-in-out;
		opacity: 0;
		visibility: hidden;
	}
	
	.mask-show {
		opacity: 1;
		visibility: visible;
	}

	/* @font-face {
		font-family: 'tuiDateFont';
		src: url(data:application/font-woff;charset=utf-8;base64,d09GRgABAAAAAAVgAA0AAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABGRlRNAAAFRAAAABoAAAAci0/w50dERUYAAAUkAAAAHgAAAB4AKQANT1MvMgAAAaAAAABDAAAAVjxuSNNjbWFwAAAB+AAAAEoAAAFS5iPQt2dhc3AAAAUcAAAACAAAAAj//wADZ2x5ZgAAAlQAAAFHAAABvPf29TBoZWFkAAABMAAAADAAAAA2GMsN3WhoZWEAAAFgAAAAHQAAACQHjAOFaG10eAAAAeQAAAATAAAAFgzQAPJsb2NhAAACRAAAABAAAAAQAOoBSG1heHAAAAGAAAAAHgAAACABEwA3bmFtZQAAA5wAAAFJAAACiCnmEVVwb3N0AAAE6AAAADQAAABLUwjqHHjaY2BkYGAAYp5Gj5/x/DZfGbhZGEDg1tUn7+F00P/LzOuY9YFcDgYmkCgAa0gNlHjaY2BkYGBu+N/AEMPCAALM6xgYGVABCwBT4AMaAAAAeNpjYGRgYGBn0GZgYgABEMkFhAwM/8F8BgANaAFLAAB42mNgZGFgnMDAysDA1Ml0hoGBoR9CM75mMGLkAIoysDIzYAUBaa4pDA7PGJ49ZG7438AQw9zA0AAUZgTJAQDrcAy8AHjaY2GAABYIDgLCBQx1AAcEAc8AeNpjYGBgZoBgGQZGBhDwAfIYwXwWBgMgzQGETAwMzxifcTx7+P8/kMUAYUkxS/6VVIXqAgNGNgY4lxGoB6QPBTAyDHsAADDkDYkAAAAAAAAAAAAAADQAagC2AN542m2QsU7DMBCG/Tt1bNPUiUnkSgiVtqKpxJAgVLVbeAa6MaK+B4JXgJWBjY21UtW5gpkdMTFX7dzApaJLhXU6n8+n//ttxtn458N79XJWZ8eMxS00C4wy9A1EP8PQncAlIQzS4WgsVtPpSmwzV3OFRqLetH5TSQMK939X61ptPZ2p2EAttNMLBRMrtschQblDeS34aY50cIkCzg/B2Y5C+VpyQxhFkRgu515O8jvU5mmPM2O0wJ5Z27vhX+yMsV437WvCdTM+GI40MgwKfuGammC0uURqeqFMfe9cxaJclkt5GMaB1hIR1VobOgpEiKq+sLZcIrJWhO3/Jw7qWlYj1Jf21FaCtmd5bevrlk28O/7A4spXTl4KTh9MTlqQ8PESBRstReic+sRj0Dni9fIqmNS/pXNWCvWOeYBmx5S9Bsn9Ah+5WtAAeNp9kD1OAzEQhZ/zByQSQiCoXVEA2vyUKRMp9Ailo0g23pBo1155nUg5AS0VB6DlGByAGyDRcgpelkmTImvt6PObmeexAZzjGwr/3yXuhBWO8ShcwREy4Sr1F+Ea+V24jhY+hRvUf4SbuFUD4RYu1BsdVO2Eu5vSbcsKZxgIV3CKJ+Eq9ZVwjfwqXMcVPoQb1L+EmxjjV7iFa2WpDOFhMEFgnEFjig3jAjEcLJIyBtahOfRmEsxMTzd6ETubOBso71dilwMeaDnngCntPbdmvkon/mDLgdSYbh4FS7YpjS4idCgbXyyc1d2oc7D9nu22tNi/a4E1x+xRDWzU/D3bM9JIbAyvkJI18jK3pBJTj2hrrPG7ZynW814IiU68y/SIx5o0dTr3bmniwOLn8owcfbS5kj33qBw+Y1kIeb/dTsQgil2GP5PYcRkAAAB42mNgYoAALjDJyIAO2MGiTIxMjMyMLIys7GmJeRmlmWZQ2pQ5OSORLaU0Mz2/FACDfwlbAAAAAf//AAIAAQAAAAwAAAAWAAAAAgABAAMABgABAAQAAAACAAAAAHjaY2BgYGQAgqtL1DlA9K2rT97DaABNlwiuAAA=) format('woff');
		font-weight: normal;
		font-style: normal;
	}

	.tui-iconfont {
		font-family: 'tuiDateFont' !important;
		font-size: 36rpx;
		font-style: normal;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	.tui-font-close:before {
		content: '\e608';
	}

	.tui-font-check:before {
		content: '\e6e1';
	}

	.date-arrowright:before {
		content: '\e600';
	}

	.date-arrowleft:before {
		content: '\e601';
	} */

	
	// .calendar-radius {
	// 	border-top-left-radius: 20rpx;
	// 	border-top-right-radius: 20rpx;
	// 	overflow: hidden;
	// }

	

	

	

	


	

	


	

	

	.tui-btn-calendar {
		padding: 16rpx;
		box-sizing: border-box;
		text-align: center;
		text-decoration: none;
	}

	.l-opacity {
		opacity: 0.5;
	}

	

	.calendar-box-show {
		transform: translate3d(0, 0, 0);
		visibility: visible;
	}

	

	



	

	// .tui-lunar-unshow {
	// 	position: absolute;
	// 	left: 0;
	// 	bottom: 8rpx;
	// 	z-index: 2;
	// }

	

	

	
</style>
