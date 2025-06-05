<template>
	
  <view class="container p-4">
	<!-- 新增悬浮导航 -->
  <view class="floating-nav">
    <button class="nav-toggle" @click="showNav = !showNav">
      <text class="icon">📖</text>
    </button>
    
    <view v-if="showNav" class="nav-panel">
      <view class="panel-header">
        <text>题目导航</text>
        <button @click="showNav = false" class="close-btn">×</button>
      </view>
      <view class="question-grid">
        <view 
          v-for="(q, index) in questions"
          :key="index"
          class="grid-item"
          :class="{ 
            'current': current === index,
            'answered': answers[index] !== undefined
          }"
          @click="jumpTo(index)"
        >
          {{ index + 1 }}
        </view>
      </view>
    </view>
  </view>
		  
    <!-- 进度条 -->
    <view class="card mb-4">
      <view class="card-body">
        <view class="d-flex align-items-center">
			<view class="flex-grow-1 me-3 position-relative">
			  <progress 
				class="progress w-100" 
				:value="current + 1" 
				:max="questions.length"
				style="height: 8px;"
			  ></progress>
			</view>
          <text class="text-muted">{{ current + 1 }}/{{ questions.length }}</text>
        </view>
      </view>
    </view>

    <!-- 题目区域 -->
    <view class="card mb-4">
      <view class="card-body">
        <h2 class="card-title mb-4 text-primary">{{ currentQuestion.question }}</h2>
        
        <view class="list-group list-group-flush">
          <view 
            v-for="(option, index) in currentQuestion.options" 
            :key="index"
            class="list-group-item list-group-item-action cursor-pointer"
            :class="{ 'active': selectedAnswer === index }"
            @click="handleAnswer(index)"
          >
            <view class="form-check">
              <view class="form-check-input" :class="{ 'checked': selectedAnswer === index }"></view>
              <span class="form-check-label">{{ option }}</span>
            </view>
          </view>
        </view>
      </view>
    </view>

    <!-- 操作按钮 -->
    <view class="row g-2">
      <view class="col">
        <button 
          class="btn btn-outline-primary w-100"
          @click="prevQuestion" 
          :disabled="current === 0"
        >
          ← 上一题
        </button>
      </view>
      <view class="col">
        <button 
          v-if="current < questions.length - 1"
          class="btn btn-primary w-100"
          @click="nextQuestion" 
          :disabled="selectedAnswer === null"
        >
          下一题 →
        </button>
        <button 
          v-else 
          class="btn btn-success w-100"
          @click="submit"
        >
          提交答卷
        </button>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
		showNav: false,
      current: 0,
      selectedAnswer: null,
      answers: [],
 questions: [
        {
          question: "uniapp的全局样式文件是？",
          options: ["App.vue", "main.js", "uni.css", "global.sass"],
          correct: 0
        },
        {
          question: "下列哪个API可以跳转页面？",
          options: ["uni.request", "uni.navigateTo", "uni.showToast", "uni.getStorage"],
          correct: 1
        },
        {
          question: "uniapp支持哪种小程序语法？",
          options: ["微信小程序", "支付宝小程序", "百度小程序", "全部支持"],
          correct: 3
        },
        {
          question: "如何获取设备信息？",
          options: ["uni.getSystemInfo", "wx.getDeviceInfo", "device.platform", "navigator.userAgent"],
          correct: 0
        },
        {
          question: "下列哪个不是uniapp项目类型？",
          options: ["HBuilderX", "微信开发者工具", "Visual Studio Code", "Android Studio"],
          correct: 3
        }
      ],
    };
  },
  computed: {
    currentQuestion() {
      return this.questions[this.current];
    }
  },
  methods: {
  jumpTo(index) {
		if(index >=0 && index < this.questions.length) {
		  this.current = index
		  this.selectedAnswer = this.answers[index] ?? null
		  this.showNav = false
		}
	  },
    handleAnswer(index) {
      this.selectedAnswer = index;
      // 震动反馈（真机生效）
      uni.vibrateShort();
    },
    prevQuestion() {
      if (this.current > 0) {
        this.current--;
        this.selectedAnswer = this.answers[this.current] ?? null;
      }
    },
    nextQuestion() {
      if (this.selectedAnswer !== null) {
        this.answers[this.current] = this.selectedAnswer;
        this.current++;
        this.selectedAnswer = this.answers[this.current] ?? null;
      }
    },
    submit() {
      const score = this.questions.reduce((acc, q, index) => {
        return acc + (this.answers[index] === q.correct ? 1 : 0);
      }, 0);
      
      uni.showModal({
        title: '答题完成',
        content: `正确率：${ ((score / this.questions.length) * 100).toFixed(0) }%`,
        showCancel: false
      });
    }
  }
};
</script>

<style>
/* 引入Tabler核心样式 */


.container {
  background-color: var(--tblr-bg-surface);
  min-height: 100vh;
}

.card {
  box-shadow: var(--tblr-box-shadow-sm);
  border-radius: var(--tblr-border-radius);
}

.list-group-item {
  transition: all 0.2s ease;
  border-left: 0;
  border-right: 0;
  cursor: pointer;
  -webkit-tap-highlight-color: transparent;
}

.list-group-item.active {
  background-color: var(--tblr-primary-bg-subtle);
  border-color: var(--tblr-primary);
  z-index: 2;
}

.form-check {
  display: flex;
  align-items: center;
}

.form-check-input {
  position: relative;
  width: 1.2em;
  height: 1.2em;
  border: 2px solid var(--tblr-border-color);
  border-radius: 50%;
  margin-right: 0.5em;
  transition: all 0.2s ease;
}

.form-check-input.checked {
  border-color: var(--tblr-primary);
}

.form-check-input.checked::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0.6em;
  height: 0.6em;
  background: var(--tblr-primary);
  border-radius: 50%;
  transform: translate(-50%, -50%);
}

.btn-success {
  background-color: var(--tblr-success);
  border-color: var(--tblr-success);
}

.card-title {
  font-size: 1.25rem;
  font-weight: 600;
}

/* 添加进度条样式 */
progress {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  overflow: hidden;
  background: var(--tblr-bg-surface-secondary);
}

progress::-webkit-progress-bar {
  background: var(--tblr-bg-surface-secondary);
}

progress::-webkit-progress-value {
  background: var(--tblr-primary);
  transition: all 0.3s ease;
}

progress::-moz-progress-bar {
  background: var(--tblr-primary);
}
/* 新增的进度条样式 */
progress {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  border-radius: 8px;
  height: 8px !important;
}

progress::-webkit-progress-bar {
  background: #f3f3f3;
  border-radius: 8px;
}

progress::-webkit-progress-value {
  background: var(--tblr-primary);
  border-radius: 8px;
  transition: width 0.3s ease;
}

progress::-moz-progress-bar {
  background: var(--tblr-primary);
  border-radius: 8px;
}

/* 悬浮导航样式 */
.floating-nav {
  position: fixed;
  right: 20rpx;
  bottom: 120rpx;
  z-index: 999;
}

.nav-toggle {
  width: 100rpx;
  height: 100rpx;
  border-radius: 50%;
  background: #007AFF;
  box-shadow: 0 5rpx 20rpx rgba(0,0,0,0.2);
}

.icon {
  font-size: 50rpx;
  color: white;
}

.nav-panel {
  position: absolute;
  right: 0;
  bottom: 120rpx;
  width: 600rpx;
  background: white;
  border-radius: 20rpx;
  box-shadow: 0 10rpx 40rpx rgba(0,0,0,0.15);
  padding: 20rpx;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 20rpx;
  border-bottom: 1rpx solid #eee;
}

.close-btn {
  font-size: 40rpx;
  color: #999;
  padding: 0 20rpx;
}

.question-grid {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 15rpx;
  margin-top: 20rpx;
}

.grid-item {
  width: 100rpx;
  height: 100rpx;
  border-radius: 15rpx;
  background: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 34rpx;
  transition: all 0.2s;
}

.grid-item.answered {
  background: #e3f2fd;
  color: #007AFF;
}

.grid-item.current {
  transform: scale(1.1);
  border: 2rpx solid #007AFF;
  box-shadow: 0 5rpx 15rpx rgba(0,122,255,0.3);
}

</style>