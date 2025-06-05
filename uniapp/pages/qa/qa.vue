<template>
	<tabler-header @updatedcount="get_user"></tabler-header>
  <view class="container">
    <!-- 导航栏 -->
    <uni-nav-bar title="题库列表" :border="false" />
	

		
	<div class="qa_ye_box">
		
		<div class="qa_ye">
			<ul class="pagination ">
				<li class="page-item disabled">
					<a class="page-link" tabindex="-1" aria-disabled="true">
						<!-- Download SVG icon from http://tabler-icons.io/i/chevron-left -->
						<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M15 6l-6 6l6 6" /></svg>
						prev
					</a>
				</li>
				<li class="page-item active"><a class="page-link" >1</a></li>
				<li class="page-item"><a class="page-link" >2</a></li>
				<li class="page-item"><a class="page-link" >3</a></li>
				<li class="page-item"><a class="page-link" >4</a></li>
				<li class="page-item"><a class="page-link" >5</a></li>
				<li class="page-item">
					<a class="page-link" >
						next <!-- Download SVG icon from http://tabler-icons.io/i/chevron-right -->
						<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M9 6l6 6l-6 6" /></svg>
					</a>
				</li>
			</ul>
			
		</div>
		
		<div class="qa_ye_but">
			<a class="btn btn-primary w-100">添加题库</a>
		</div>
		

		
	</div>

    <!-- 内容区域 -->
    <scroll-view scroll-y class="content">
      <view class="qbank-list">
        <view 
          v-for="qbank in qbanks" 
          :key="qbank.qbank_id"
          class="qbank-card"
          @click="navigateToDetail(qbank.qbank_id)"
        >
          <!-- 卡片内容 -->
          <view class="card-content">
            <!-- 左侧图标 -->
            <image 
              :src="qbank.logo_path || '/static/logo-placeholder.png'"
              class="logo"
              mode="aspectFit"
            />

            <!-- 右侧信息 -->
            <view class="info">
              <text class="title">{{ qbank.name }}</text>
              <text class="description">{{ qbank.description }}</text>
              <view class="meta">
                <uni-icons type="calendar" size="14" color="#666" />
                <text class="time">{{ formatDate(qbank.created_at) }}</text>
              </view>
            </view>
          </view>

          <!-- 右侧箭头 -->
          <uni-icons type="arrowright" size="20" color="#999" />
        </view>
      </view>

      <!-- 空状态 -->
      <view v-if="qbanks.length === 0" class="empty">
        <image src="/static/empty.png" class="empty-img" />
        <text class="empty-text">暂无题库数据</text>
      </view>
    </scroll-view>
  </view>
  
  <tabler-footer></tabler-footer>
</template>

<script setup>
import { ref } from 'vue';

// 本地测试数据
const qbanks = ref([
  {
    qbank_id: 1,
    name: '前端开发题库',
    description: '涵盖HTML/CSS/JavaScript等前端核心知识',
    logo_path: '/static/fe-logo.png',
    created_at: '2024-03-15 09:00:00'
  },
  {
    qbank_id: 2,
    name: 'Java面试宝典',
    description: 'Java基础、框架与分布式系统面试题',
    logo_path: '',
    created_at: '2024-03-10 14:30:00'
  },
  {
    qbank_id: 3,
    name: '数据库专项练习',
    description: 'MySQL、Redis等数据库实战题库',
    logo_path: '/static/db-logo.png',
    created_at: '2024-03-05 16:45:00'
  }
]);

// 日期格式化
const formatDate = (dateStr) => {
  return dateStr.split(' ')[0].replace(/-/g, '/');
};

// 跳转详情页
const navigateToDetail = (id) => {
  uni.navigateTo({
    url: `/pages/qbank/detail?id=${id}`
  });
};
</script>

<style lang="scss" scoped>
	.qa_ye_box{
		display: flex;          /* 启用弹性布局 */
		  gap: 10px;              /* 子元素间距（可选） */
		  justify-content: start; /* 水平对齐方式 */
		  .qa_ye{
		  	width: 90%;
		  }
		  .qa_ye_but{
		  	//width: 10%;  
		  }

	}

.container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f6f8;
}

.content {
  flex: 1;
  padding: 20rpx 24rpx;
}

.qbank-list {
  display: flex;
  flex-direction: column;
  gap: 24rpx;
}

.qbank-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.05);
  transition: all 0.2s;

  &:active {
    transform: scale(0.98);
    background: #f8f9fa;
  }
}

.card-content {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 24rpx;
}

.logo {
  width: 120rpx;
  height: 120rpx;
  border-radius: 12rpx;
  background-color: #f1f2f3;
}

.info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12rpx;
}

.title {
  font-size: 32rpx;
  font-weight: 600;
  color: #333;
  line-height: 1.4;
}

.description {
  font-size: 26rpx;
  color: #666;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}

.meta {
  display: flex;
  align-items: center;
  gap: 8rpx;
  margin-top: 8rpx;
}

.time {
  font-size: 24rpx;
  color: #999;
}

.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 100rpx 0;
  
  &-img {
    width: 300rpx;
    height: 300rpx;
    opacity: 0.6;
  }
  
  &-text {
    font-size: 28rpx;
    color: #999;
    margin-top: 24rpx;
  }
}


</style>