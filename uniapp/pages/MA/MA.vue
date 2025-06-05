<template>
	
  <tabler-header v-if="!fulldis" @updatedcount="get_user"></tabler-header>
  <FullCalendar ref="fullCalendar" :options='calendarOptions' >
   <!-- <template v-slot:eventContent='arg'>
      <b>{{ arg.timeText }}</b>
      <i>{{ arg.event.title }}</i>
    </template> -->
  </FullCalendar>
  <tabler-footer v-if="!fulldis"></tabler-footer>


 
  
<div v-if="set_item_flag" class="modal modal-blur fade show" style="display: block;" >
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">{{event_id==0?"正在添加日程":"正在修改 "+event_username+" 创建的日程"}}</h5>
            <button type="button" class="btn-close" @click="set_item_close"></button>
          </div>
          <div class="modal-body">
            
            <div class="mb-3">

			
			<view class="date-seles">
				<!-- <datetime-picker v-model="range" type="daterange" /> -->
				<datatime-picker-for-fullCalendar :value="range" :title="title" :color_sele="color_sele" @update="update_date"></datatime-picker-for-fullCalendar>
				
				
			</view>
			
			<div>
				
				<uni-easyinput class="input" type="text" placeholder="输入日程内容" v-model="title" ></uni-easyinput>
				
			</div>
              <div>
				  
				<radio-group class="color_box" @change="change_color">
					<div class="color_box_item" v-for="(item, index) in radio_colors" :key="item">
					  <label class="form-colorinput">
						<radio :value="item.color" :checked="item.color === this.color_sele" :color="item.color" :background-color="item.color"/>
						{{item.text}}
					  </label>
					</div>
						  
				</radio-group>  
               
              </div>
            </div>
           
          </div>
          <div class="modal-footer">
            <button type="button" class="btn me-auto" @click="set_item_close">Close</button>
			<button type="button" class="btn" @click="copy_item">复制</button>
			<button type="button" class="btn" :disabled="!copy_color" @click="pise_item">粘贴</button>
			
            <button v-if="event_id==0" type="button" class="btn btn-primary" @click="add_even">添加日程</button>
			
			<button v-if="event_id!=0" type="button" class="btn btn-danger" @click="del_even">删除日程</button>
			<button v-if="event_id!=0" type="button" class="btn btn-primary" @click="change_even">修改日程</button>
			
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
</template>

<script>
import { public_func } from '@/public_js.js';

import FullCalendar from '@fullcalendar/vue3'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from "@fullcalendar/interaction"//拖动插件 需要用npm安装
import screenfull from 'screenfull'

export default {
  components: {
    FullCalendar // make the <FullCalendar> tag available
  },
  onReady() {
  	console.log("onReady")
  },
  mounted() {
  	console.log("mounted")
	//实例被创建
	//console.log(FullCalendar)
	
	this.get_evensts();

	

	//
	
  },
	onLoad() {
		console.log("onLoad");
		clearTimeout(this.auto_get_event_time);
		this.auto_get_event_time=null;
		this.auto_get_event_time=setInterval(this.get_evensts, 5000);
		
	},
	onUnload() {
		console.log("onUnload");
		clearTimeout(this.auto_get_event_time);
		this.auto_get_event_time=null;
	},
	onHide() {
		console.log("onHide");
		
		
	},
  onShow() {
  	console.log("onshow")
  },
  methods:{
	  // 进入全屏
	      enterFullscreen(element)
		  {
	        screenfull.request();
	      },
	      
	      // 退出全屏
	      exitFullscreen() 
		  {
	       screenfull.exit();
	      },
		  
  	  test(){
		  
	  },
	  handleDatesRender(arg)
	  {
		console.log(arg.currentStart)
		console.log(arg.currentEnd)
	  },
	  pise_item(){
		  this.color_sele=this.copy_color;
		  this.title=this.copy_title;
	  },
	  copy_item(){
		  this.copy_color=this.color_sele;
		  this.copy_title=this.title;
	  },
	  change_even(){
	  	var that = this;
		
	  	uni.request({
	  		header: {'Content-Type': 'application/x-www-form-urlencoded'},
	  		url:public_func.baseUrl+"api/?ac=ma&do=change",
	  		method:'POST',
	  		data:{
	  			user_id:this.user.id,
	  			user_pass:this.user.pass,
				id:this.event_id,
	  			stad_date:this.range[0],
				end_date:this.range[1],
	  			title:this.title,
	  			color_sele:this.color_sele,
	  			type:"events",
	  			allday:1,
	  			
	  		},
	  		timeout:10000,
	  		success(res) {
	  			console.log(res.data);
	  			if(res.data['error_code']==0)
	  			{
	  				that.set_item_close();
	  				that.messageToggle('success',"已修改加日程");
	  				that.get_evensts();
	  		
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
	  del_even(){
		  var that = this;
		  uni.request({
		  	header: {'Content-Type': 'application/x-www-form-urlencoded'},
		  	url:public_func.baseUrl+"api/?ac=ma&do=del",
		  	method:'POST',
		  	data:{
		  		user_id:this.user.id,
		  		user_pass:this.user.pass,
		  		id:this.event_id,
		  		
		  		
		  	},
		  	timeout:10000,
		  	success(res) {
		  		console.log(res.data);
		  		if(res.data['error_code']==0)
		  		{
		  			that.set_item_close();
		  			that.messageToggle('success',"已删除");
		  			that.get_evensts();
		  	
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
	  get_evensts(){
		  //console.log("get_evensts");
		  var user=this.user;
		  if(user==null)
		  {
			  user={
				  id:"0",
				  pass:"0"
			  };
			  
		  }
		  var that = this;
		  uni.request({
		  	header: {'Content-Type': 'application/x-www-form-urlencoded'},
		  	url:public_func.baseUrl+"api/?ac=ma&do=show",
		  	method:'POST',
		  	data:{
		  		user_id:user.id,
		  		user_pass:user.pass,
		  	},
		  	timeout:10000,
		  	success(res) {
		  		console.log(res.data);
		  		if(res.data['error_code']==0)
		  		{
		  			/* var events=res.data['events'].map(item => {
					return { ...item, recurrence: ['RRULE:FREQ=DAILY'] };//给每一项插入一个标签 不过好像没什么用
					});
		  			that.calendarOptions.events=events; */
					
					that.calendarOptions.events=res.data['events'];
		  	
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
	  updateCurrentMonth(info) { 
		  const currentDate = info.start; 
		  const currentMonth = currentDate.getMonth() + 1; // 月份从0开始，所以需要加1 
		  const currentYear = currentDate.getFullYear(); 
		  this.currentMonth = `当前显示的月份是: ${currentYear}年${currentMonth}月`; 
		},
	  messageToggle(type,msg) {
	  	this.msgType=type;
	  	this.messageText=msg;
	  	this.$refs.message.open();
	  },
	  add_even()
	  {
		var that = this;
		

		
		uni.request({
			header: {'Content-Type': 'application/x-www-form-urlencoded'},
			url:public_func.baseUrl+"api/?ac=ma&do=add",
			method:'POST',
			data:{
				user_id:this.user.id,
				user_pass:this.user.pass,
				stad_date:this.range[0],
				end_date:this.range[1],
				title:this.title,
				color_sele:this.color_sele,
				type:"events",
				allday:1,
				
			},
			timeout:10000,
			success(res) {
				console.log(res.data);
				if(res.data['error_code']==0)
				{
					that.set_item_close();
					that.messageToggle('success',"已添加日程");
					that.get_evensts();
	
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
	  edit_item(e){
		  
	  },
	  set_item_show(){
		
		this.set_item_flag=true;
	  },
	  set_item_close(){
		this.set_item_flag=false;
	  },
	  change_color(ev)
	  {
		  console.log(ev.detail.value)
		  this.color_sele=ev.detail.value;
	  },
	  get_user(user)
	  {
			console.log(user)
			this.user=user;
	  },
	  update_date(date)
	  {
		  console.log(date)
		  this.range[0]=date[0];
		  this.range[1]=date[1];
	  }
	  
	
	  
  },
  data() {
	  var that=this;
    return {
		fulldis:false,
		click_time:0,
		set_item_flag:false,
		color_sele:"#066FD1",
		user:null,
		range:[0,0],
		range2:[0,0],
		stad_date:null,
		end_date:null,
		title:"",
		event_username:"",
		event_id:0,
		msgType:"success",
		message:"这是一条成功提示",
		messageText:"",
		show_mon:0,
		auto_get_event_time:null,
		copy_title:"",
		copy_color:null,
		radio_colors:[
			{color:"#066FD1",text:"工作"},
			{color:"#09d119",text:"值班"},	
			{color:"#ff00ff",text:"考试"},
			{color:"#ffff00",text:"备用"},
			{color:"#d16c13",text:"个人假期"},
			{color:"#d10d21",text:"公众假期"},
		],
		this_mon_fastday:null,
		this_mon_endday:null,
		show_switch_flag:true,
      calendarOptions: {
		
		height:"auto",
		  
        plugins: [
			dayGridPlugin,
			timeGridPlugin,
			interactionPlugin//导入拖动插件
			],
		
        initialView: 'dayGridMonth',
		headerToolbar: {
		  left: 'prevYearCustom,prevMonthCustom,todayCustom,nextMonthCustom,nextYearCustom',
		  center: 'title',
		  right: 'quanping,dayGridMonth'//,timeGridWeek,timeGridDay'
		},
		fixedWeekCount:false,
		weekNumbers:true,
		businessHours: {
		    dow: [ 1, 2, 3, 4 ,5], // 周一 - 周四
		
		    start: '9:00', // 上午10点开始
		    end: '18:00', // 下午18点结束
		},
		//weekNumbersWithinDays:true,
        weekends: true,//是否显示周末


	  // 自定义按钮
	  customButtons: {
		prevYearCustom: {
		  text: '上一年',
		  click: function() {
			that.show_switch_flag=true;
			that.$refs.fullCalendar.getApi().prevYear();
		  }
		},
		nextYearCustom: {
		  text: '下一年',
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
		},
		quanping: {
		  text: '全屏',
		  click: function() {
			if(that.fulldis)
			{
				that.fulldis=false;
				that.exitFullscreen();
			}else
			{
				that.fulldis=true;
				
				that.enterFullscreen();
			}
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
				//info.el.style.backgroundColor = '#ffff7f'; 
			}
			
			info.el.style.border = '1px solid #4b4b4b'; // 浅蓝色边框
			
			
			
		},
		selectable: true,
		dateClick(info) {
			var click_time;
			const now = new Date();
			const nextime=now.getTime()-that.click_time
			
			if(nextime<400)
			{
				console.log(nextime)//双击时间
				
				var click_date={
					date:info.date,
					datestr:info.dateStr,
					
				}
				console.log(click_date)
				
				if(that.user==null)
				{
					uni.redirectTo({
						url:'/pages/login/login'
					});
				}else
				{
					that.event_id=0;
					that.title="";
					that.color_sele=that.radio_colors[0].color;
					that.range[0]=click_date.datestr;
					
					var next_day=new Date(click_date.datestr);
					next_day.setDate(next_day.getDate() + 1);	
					
					that.range[1]=public_func.formatDate(next_day);
					
					that.set_item_show();
				}
				
			}
			
			that.click_time=now.getTime()

		},
		select(info) {
			if((info.end-info.start)>86400000)
			{
				var select_date={
					start:info.start,
					end:info.end,
					startStr:info.startStr,
					endStr:info.endStr,
				}
				console.log(select_date)
				
				if(that.user!=null)
				{
					that.event_id=0;
					that.title="";
					that.color_sele=that.radio_colors[0].color;
					that.range[0]=public_func.formatDate(select_date.start);
					that.range[1]=public_func.formatDate(select_date.end);
					that.set_item_show();
					
				}
				
				
			}
			
		},
		//placeholder:true,
		eventClick(info) {
			
			console.log(info)
			
			const now = new Date();
			const nextime=now.getTime()-this.click_time
			
			if(nextime<400)
			{
				console.log(nextime)//双击时间
				
				var ev={
					id:info.el.fcSeg.eventRange.def.publicId,
					title:info.el.fcSeg.eventRange.def.title,
					allDay:info.el.fcSeg.eventRange.def.allDay,
					start:info.el.fcSeg.eventRange.range.start,
					end:info.el.fcSeg.eventRange.range.end,
					bg_color:info.el.fcSeg.eventRange.ui.backgroundColor,
					edit:info.el.fcSeg.eventRange.ui.startEditable,
					username:info.el.fcSeg.eventRange.def.extendedProps.event_username,
				}
				
				
				
				if(that.user==null)
				{
					uni.redirectTo({
						url:'/pages/login/login'
					});
				}else
				{
					if(ev.edit)
					{
						that.event_id=ev.id;
						that.title=ev.title;
						that.color_sele=ev.bg_color;
						that.range[0]=public_func.formatDate(ev.start);
						that.range[1]=public_func.formatDate(ev.end);
						
						that.event_username=ev.username;
						
						that.set_item_show();
					}else
					{
						that.messageToggle('warn',"这个不是您的日程");
					}
					
				}
				
				

			}
			
			this.click_time=now.getTime()
			

		
		    // change the border color just for fun
		    //info.el.style.borderColor = 'red';
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
			console.log(info.event)
			
			//var user=that.user;
			
			//var that2 = this;
			uni.request({
				header: {'Content-Type': 'application/x-www-form-urlencoded'},
				url:public_func.baseUrl+"api/?ac=ma&do=event_droup",
				method:'POST',
				data:{
					user_id:that.user.id,
					user_pass:that.user.pass,
					id:ev.id,
					start:public_func.formatDate(ev.start),
					end:public_func.formatDate(ev.end),
					
					
				},
				timeout:10000,
				success(res) {
					console.log(res.data);
					if(res.data['error_code']==0)
					{
						
						//that.messageToggle('success', public_func.formatDate(ev.start));
						that.get_evensts();
				
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
		
		//weekNumbers:true,//是否在日历中显示周次(一年中的第几周)
		firstDay:1,
        events: [],
      }
    }
  },

}
</script>

<style lang="scss">
	
	.color_box
	{
		display: flex;
		//justify-content: space-between; /* 调整子元素之间的间距 */
		.color_box_item
		{
			
			border-radius: 5px;
			box-shadow: 0 0 10px 5px rgba(207, 207, 207, 0.5); /* 外部阴影效果 */
			margin-left: 20rpx;
		}
	}
	
	
	.input{
		padding-bottom: 30rpx;
	}
	.date-seles{
		display: flex;
		flex-wrap: wrap; /* 允许换行 */
		padding-bottom: 10rpx;
		
	}
	// 相应的CSS样式
	.fc-event-centered {
	    display: flex;
	    align-items: center;
	    justify-content: center;
	    text-align: center;
	    height: 100%;
	}
	
	.fc-highlight {
	  background: #55aaff !important;
	}
</style>
