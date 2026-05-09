<template>
  <view class="container">
    <view class="tab-bar">
      <view :class="['tab-item', activeTab === 'pending' ? 'active' : '']" @click="switchTab('pending')">待开始</view>
      <view :class="['tab-item', activeTab === 'running' ? 'active' : '']" @click="switchTab('running')">进行中</view>
      <view :class="['tab-item', activeTab === 'completed' ? 'active' : '']" @click="switchTab('completed')">已完成</view>
    </view>
    
    <view class="task-list">
      <view class="task-card" v-for="task in list" :key="task.id" @click="taskClick(task)">
        <view class="task-header">
          <text class="task-name">{{ task.name }}</text>
          <view :class="['status-tag', task.status]">
            <text v-if="task.status === 'pending'">待开始</text>
            <text v-else-if="task.status === 'running'">进行中</text>
            <text v-else-if="task.status === 'completed'">已完成</text>
            <text v-else>{{ task.status }}</text>
          </view>
        </view>
        <view class="task-info">
          <view class="info-item">
            <text class="label">路线：</text>
            <text class="value">{{ task.routeName }}</text>
          </view>
          <view class="info-item">
            <text class="label">执行人：</text>
            <text class="value">{{ task.assigneeName }}</text>
          </view>
          <view class="info-item">
            <text class="label">计划时间：</text>
            <text class="value">{{ task.planTime }}</text>
          </view>
        </view>
        <view class="task-actions" v-if="task.status !== 'completed'">
          <button v-if="task.status === 'pending'" class="action-btn start-btn" @click.stop="startTask(task)">开始任务</button>
          <button v-if="task.status === 'running'" class="action-btn scan-btn" @click.stop="doScan(task)">扫码打卡</button>
        </view>
      </view>
    </view>
    
    <view class="empty" v-if="list.length === 0">
      <text>暂无任务</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      activeTab: 'pending',
      list: [],
      page: 1,
      pageSize: 10
    }
  },
  onShow() {
    this.loadTasks()
  },
  methods: {
    switchTab(tab) {
      this.activeTab = tab
      this.page = 1
      this.loadTasks()
    },
    loadTasks() {
      uni.request({
        url: `/api/v1/tasks?page=${this.page}&pageSize=${this.pageSize}&status=${this.activeTab}`,
        header: { Authorization: `Bearer ${uni.getStorageSync('token')}` },
        success: (res) => {
          if (res.data.code === 200) {
            this.list = res.data.data.list || []
          }
        }
      })
    },
    taskClick(task) {
      console.log('task click', task)
    },
    startTask(task) {
      uni.request({
        url: `/api/v1/tasks/${task.id}/start`,
        method: 'PUT',
        header: { Authorization: `Bearer ${uni.getStorageSync('token')}` },
        success: (res) => {
          if (res.data.code === 200) {
            uni.showToast({ title: '任务已开始', icon: 'success' })
            this.loadTasks()
          } else {
            uni.showToast({ title: res.data.message, icon: 'none' })
          }
        }
      })
    },
    doScan(task) {
      uni.scanCode({
        success: (res) => {
          const qrCode = res.result
          uni.setStorageSync('scanQRCode', qrCode)
          
          uni.request({
            url: '/api/v1/records',
            method: 'POST',
            header: { Authorization: `Bearer ${uni.getStorageSync('token')}` },
            data: {
              taskId: task.id,
              pointId: 1,
              qrCode: qrCode,
              status: 'normal'
            },
            success: (res) => {
              if (res.data.code === 200) {
                uni.showToast({ title: '打卡成功', icon: 'success' })
              } else {
                uni.showToast({ title: res.data.message, icon: 'none' })
              }
            },
            fail: () => {
              uni.showToast({ title: '打卡失败', icon: 'none' })
            }
          })
        },
        fail: () => {
          uni.showToast({ title: '扫码失败', icon: 'none' })
        }
      })
    }
  }
}
</script>

<style scoped>
.container {
  min-height: 100vh;
  background: #f5f7fa;
}

.tab-bar {
  display: flex;
  background: #fff;
  padding: 0 40rpx;
  border-bottom: 1rpx solid #eee;
}

.tab-item {
  flex: 1;
  text-align: center;
  padding: 30rpx 0;
  font-size: 28rpx;
  color: #666;
  position: relative;
}

.tab-item.active {
  color: #667eea;
  font-weight: 500;
}

.tab-item.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 60rpx;
  height: 4rpx;
  background: #667eea;
  border-radius: 2rpx;
}

.task-list {
  padding: 30rpx;
  gap: 20rpx;
  display: flex;
  flex-direction: column;
}

.task-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 30rpx;
}

.task-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.task-name {
  font-size: 32rpx;
  font-weight: 500;
  color: #333;
}

.status-tag {
  padding: 8rpx 16rpx;
  border-radius: 8rpx;
  font-size: 22rpx;
}

.status-tag.pending { background: #f5f5f5; color: #999; }
.status-tag.running { background: #e6f7ff; color: #1890ff; }
.status-tag.completed { background: #f6ffed; color: #52c41a; }

.task-info {
  margin-bottom: 20rpx;
}

.info-item {
  display: flex;
  margin-bottom: 12rpx;
}

.info-item:last-child {
  margin-bottom: 0;
}

.label {
  font-size: 26rpx;
  color: #999;
}

.value {
  font-size: 26rpx;
  color: #333;
}

.task-actions {
  display: flex;
  gap: 20rpx;
  padding-top: 20rpx;
  border-top: 1rpx solid #f0f0f0;
}

.action-btn {
  flex: 1;
  height: 72rpx;
  border-radius: 12rpx;
  font-size: 26rpx;
  border: none;
}

.start-btn {
  background: #667eea;
  color: #fff;
}

.scan-btn {
  background: #1890ff;
  color: #fff;
}

.empty {
  text-align: center;
  padding: 100rpx 0;
  color: #999;
  font-size: 28rpx;
}
</style>
