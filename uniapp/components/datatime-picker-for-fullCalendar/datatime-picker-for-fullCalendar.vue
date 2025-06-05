<template>
	<view class="dddppppfffff">
		<view class="datashow" @click="showclick">
			
			<view class="this_icon">
				<!-- icon -->
				<span>
					<svg  xmlns="http://www.w3.org/2000/svg"  width="32"  height="32"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-calendar-week"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M4 7a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12z" /><path d="M16 3v4" /><path d="M8 3v4" /><path d="M4 11h16" /><path d="M7 14h.013" /><path d="M10.01 14h.005" /><path d="M13.01 14h.005" /><path d="M16.015 14h.005" /><path d="M13.015 17h.005" /><path d="M7.01 17h.005" /><path d="M10.01 17h.005" /></svg>
				</span>
			</view>
			<view class="startdate">
				<!-- start data -->
				
				{{this.startdatestr}}
				
			</view>
			<view>
				<!-- to flag -->
				至
			</view>
			<view>
				<!-- end data -->
				{{this.enddatestr}}
			</view>
			<view>
				<!-- clear button -->
				<span>
					<svg  xmlns="http://www.w3.org/2000/svg"  width="24"  height="24"  viewBox="0 0 24 24"  fill="none"  stroke="currentColor"  stroke-width="2"  stroke-linecap="round"  stroke-linejoin="round"  class="icon icon-tabler icons-tabler-outline icon-tabler-x"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>
				</span>
			</view>
		</view>
		<view v-if="showclander" class="fclander">
			<FullCalendar ref="fullCalendar" :options='calendarOptions' >
			</FullCalendar>
		</view>		
	</view>

</template>

<script>
	import { public_func } from '@/public_js.js';
	import FullCalendar from '@fullcalendar/vue3'
	import dayGridPlugin from '@fullcalendar/daygrid'
	import timeGridPlugin from '@fullcalendar/timegrid'
	import interactionPlugin from "@fullcalendar/interaction"//拖动插件 需要用npm安装
	export default {
		name:"datatime-picker-for-fullCalendar",
		components: {
		  FullCalendar // make the <FullCalendar> tag available
		},
		props: {
			value: {
				type: [Array, Date],
				default: null
			},
			title: {
				type: [String],
				default: ""
			},
			color_sele: {
				type: [String],
				default: ""
			},
		},
		mounted() 
		{
			//console.log("datapick");
			
			 var temp;
			if(this.value!=null)
			{
				temp=
					{
						"id": "0",
						"title": this.title,
						"start": this.value[0],
						"end": this.value[1],
						"backgroundColor": this.color_sele,
						"allDay": true,
						"editable": true
					}
				this.calendarOptions.events.push(temp)
			} 
			
			this.startdatestr=this.value[0];
			var next_day=new Date(this.value[1]);
			next_day.setDate(next_day.getDate() - 1);	
			this.enddatestr=public_func.formatDate(next_day);
		},
		data() {
			var that=this;
			return {
				startdatestr:"开始日期",
				enddatestr:"结束日期",
				showclander:false,
				tempdate1:null,
				tempdate2:null,
				tempdate_step:0,
				  calendarOptions: {
					
					height:"460px",
					  
					plugins: [
						dayGridPlugin,
						timeGridPlugin,
						interactionPlugin//导入拖动插件
						],
					
					initialView: 'dayGridMonth',
					headerToolbar: {
					  left: 'prevMonthCustom,todayCustom,nextMonthCustom',
					  center: 'title',
					  right: ''//,timeGridWeek,timeGridDay'
					},
					fixedWeekCount:false,
					weekNumbers:true,

					//weekNumbersWithinDays:true,
					weekends: true,//是否显示周末


				  // 自定义按钮
				  customButtons: {
					prevYearCustom: {
					  text: '上年',
					  click: function() {
						that.show_switch_flag=true;
						that.$refs.fullCalendar.getApi().prevYear();
					  }
					},
					nextYearCustom: {
					  text: '下年',
					  click: function() {
						that.show_switch_flag=true;
						that.$refs.fullCalendar.getApi().nextYear();
					  }
					},
					prevMonthCustom: {
					  text: '上月',
					  click: function() {
						that.show_switch_flag=true;
						that.$refs.fullCalendar.getApi().prev(); // 将日历后退一步（例如，一个月或一周）。
					  }
					},
					nextMonthCustom: {
					  text: '下月',
					  click: function() {
						  that.show_switch_flag=true;
						  that.$refs.fullCalendar.getApi().next(); // 将日历向前移动一步（例如，一个月或一周）。
						
					  }
					},
					todayCustom: {
					  text: '今天',
					  click: function() {
						that.show_switch_flag=true;
						that.$refs.fullCalendar.getApi().today();
					  }
					}
				  },
				  
				datesSet(info)
				{
					//console.log(info);
					that.this_mon_fastday=info.start;
					that.this_mon_endday=info.end;
					
				},
					dayCellDidMount(info) 
					{

						switch(info.dow)
						{
							case 0:
							info.el.style.backgroundColor = '#ffb5b5'; 
							break;
							case 6:
							info.el.style.backgroundColor = '#ffb5b5'; 
							break;
						}
						
						
						if(info.isToday)
						{
							info.el.style.backgroundColor = '#ffff7f'; 
						}
						
						info.el.style.border = '1px solid #4b4b4b'; // 浅蓝色边框
						
						//info.el.style.height="50px";
						//console.log(info.el.style);
						
						
						
					},
					selectable: true,
					dateClick(info) {
						
						switch(that.tempdate_step)
						{
							case 0:
								that.tempdate1=info.date;
								that.tempdate_step=1;
							break;
							case 1:
								that.tempdate2=info.date;
								
								var temp=[];
								
								if(that.tempdate1>that.tempdate2)
								{
									temp.push(public_func.formatDate(that.tempdate2));
									temp.push(public_func.formatDate(that.tempdate1));
								}else
								{
									temp.push(public_func.formatDate(that.tempdate1));
									temp.push(public_func.formatDate(that.tempdate2));
								}

								//console.log(temp);
								
								that.startdatestr=temp[0];
								that.enddatestr=temp[1];
								
								var next_day=new Date(temp[1]);
								next_day.setDate(next_day.getDate() + 1);	
								temp[1]=public_func.formatDate(next_day);
								
								that.calendarOptions.events[0].start=temp[0];
								that.calendarOptions.events[0].end=temp[1];
								
								that.tempdate_step=0;
								
								that.$emit('update',temp);  // 通过$emit触发事件，第二个参数就是传递的参数
							break;
							
						}
						
					


					},
					select(info) {
						if((info.end-info.start)>86400000)
						{
							
							var ev={
								start:info.start,
								end:info.end,
								startStr:info.startStr,
								endStr:info.endStr,
							}
							//console.log(ev)
							
							var temp=[];
							temp.push(public_func.formatDate(ev.start));
							temp.push(public_func.formatDate(ev.end));
							
							that.calendarOptions.events[0].start=temp[0];
							that.calendarOptions.events[0].end=temp[1];
							
							that.startdatestr=temp[0];
							var next_day=new Date(temp[1]);
							next_day.setDate(next_day.getDate() - 1);	
							that.enddatestr=public_func.formatDate(next_day);
							
							
							that.$emit('update',temp);  // 通过$emit触发事件，第二个参数就是传递的参数
							
							
						}
						
					},
					
					eventDrop(info){
						var olo_ev={
							id:info.oldEvent._def.publicId,
							title:info.oldEvent._def.title,
							allDay:info.oldEvent._def.allDay,
							start:info.oldEvent._instance.range.start,
							end:info.oldEvent._instance.range.end,
						}
						var ev={
							id:info.event._def.publicId,
							title:info.event._def.title,
							allDay:info.event._def.allDay,
							start:info.event._instance.range.start,
							end:info.event._instance.range.end,
						}
						//console.log(info)
						//console.log(olo_ev)
						//console.log(info.event)
						
						var temp=[];
						temp.push(public_func.formatDate(ev.start));
						temp.push(public_func.formatDate(ev.end));
						
						that.calendarOptions.events[0].start=temp[0];
						that.calendarOptions.events[0].end=temp[1];
						
						that.startdatestr=temp[0];
						var next_day=new Date(temp[1]);
						next_day.setDate(next_day.getDate() - 1);	
						that.enddatestr=public_func.formatDate(next_day);
						
						
						that.$emit('update',temp);  // 通过$emit触发事件，第二个参数就是传递的参数
					},


					
					//weekNumbers:true,//是否在日历中显示周次(一年中的第几周)
					firstDay:1,
				
					events: [],
				  }
		
			};
		},
		
		watch: {
		      color_sele: 
			  {
		        handler(newValue,oldValue) 
				{
		          this.calendarOptions.events[0].backgroundColor=newValue;
		        },
		        deep: true // 深度监听父组件传过来对象变化
		      },
			  title:
			  {
			    handler(newValue,oldValue) 
			  	{
			      this.calendarOptions.events[0].title=newValue;
			    },
			    deep: true // 深度监听父组件传过来对象变化
			  },
		    },
		
		
		methods:{
			showclick()
			{
				if(this.showclander)
				{
					this.showclander=false;
				}else
				{
					this.showclander=true;
				}
				
			}
		}
	}
</script>

<style lang="scss">
	.dddppppfffff
	{
		width: 100%;
		margin-bottom: 30rpx;
	}
	.datashow
	{
		
		height: 64rpx;
		background-color: #f3f3f3;
		border-radius: 5px;
		
		  box-shadow: 0 0 10px 5px rgba(207, 207, 207, 0.5); /* 外部阴影效果 */
		display: flex;
		  justify-content: space-between; /* 调整子元素之间的间距 */
		  align-items: center;
		  .this_icon
		  {
			  
		  }
		  .startdate{
			 
			  
		  }
	}
	.fclander
	{
		
		//background-color: #e3e3e3;
		
		width: 500px;
		height: 450px;
		
		margin-top: 5rpx;
		
	}
	
	.fc-header-toolbar
	{
		font-size: 20rpx;
		height: 30rpx;
	}
</style>