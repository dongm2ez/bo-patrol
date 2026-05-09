<template>
  <view class="container">
    <view class="header">
      <view class="user-info">
        <view class="avatar">{{ userInfo?.name?.charAt(0) || 'U' }}</view>
        <view class="user-text">
          <text class="greeting">欢迎回来</text>
          <text class="name">{{ userInfo?.name || '用户' }}</text>
        </view>
      </view>
      <view class="role-tag" :class="userInfo?.role">
        {{ getRoleName(userInfo?.role) }}
      </view>
    </view>
    
    <view class="stats-grid">
      <view class="stat-item">
        <text class="stat-value">{{ stats.totalTasks || 0 }}</text>
        <text class="stat-label">总任务数</text>
      </view>
      <view class="stat-item">
        <text class="stat-value">{{ stats.completedTasks || 0 }}</text>
        <text class="stat-label">已完成</text>
      </view>
      <view class="stat-item">
        <text class="stat-value">{{ stats.pendingTickets || 0 }}</text>
        <text class="stat-label">待处理工单</text>
      </view>
    </view>
    
    <view class="section">
      <view class="section-header">
        <text class="section-title">快捷操作</text>
      </view>
      <view class="action-grid">
        <view class="action-item" @click="goToTask">
          <view class="icon task-icon">
            <text>📋</text>
          </view>
          <text class="action-label">我的任务</text>
        </view>
        <view class="action-item" @click="goToRecord">
          <view class="icon record-icon">
            <text>📝</text>
          </view>
          <text class="action-label">巡检记录</text>
        </view>
        <view class="action-item" @click="createTicket">
          <view class="icon ticket-icon">
            <text>⚠️</text>
          </view>
          <text class="action-label">上报问题</text>
        </view>
        <view class="action-item" @click="doScan">
          <view class="icon scan-icon">
            <text>📷</text>
          </view>
          <text class="action-label">扫码打卡</text>
        </view>
      </view>
    </view>
    
    <view class="section">
      <view class="section-header">
        <text class="section-title">最近任务</text>
        <text class="more" @click="goToTask">查看全部</text>
      </view>
      <view class="task-list" v-if="recentTasks.length > 0">
        <view class="task-item" v-for="task in recentTasks" :key="task.id">
          <view class="task-info">
            <text class="task-name">{{ task.name }}</text>
            <text class="task-desc">{{ task.routeName || '-' }}</text>
          </view>
          <view :class="['status-tag', task.status]">
            <text v-if="task.status === 'pending'">待开始</text>
            <text v-else-if="task.status === 'running'">进行中</text>
            <text v-else-if="task.status === 'completed'">已完成</text>
            <text v-else>{{ task.status }}</text>
          </view>
        </view>
      </view>
      <view class="empty-tasks" v-else>
        <text class="empty-text">暂无任务</text>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getUserInfo, type UserInfo } from '../../api/auth'

const userInfo = ref<UserInfo | null>(getUserInfo())

const stats = ref({
  totalTasks: 0,
  completedTasks: 0,
  pendingTickets: 0
})

const recentTasks = ref<any[]>([])

const getRoleName = (role?: string) => {
  const roleMap: Record<string, string> = {
    'admin': '管理员',
    'inspector': '巡检员',
    'repair': '维修员'
  }
  return roleMap[role || ''] || role
}

const goToTask = () => {
  uni.switchTab({ url: '/pages/task/list' })
}

const goToRecord = () => {
  uni.switchTab({ url: '/pages/record/list' })
}

const createTicket = () => {
  uni.navigateTo({ url: '/pages/ticket/create' })
}

const doScan = () => {
  uni.scanCode({
    success: (res) => {
      console.log('扫码结果:', res.result)
      uni.showToast({ title: '扫码成功', icon: 'success' })
    },
    fail: () => {
      uni.showToast({ title: '扫码失败', icon: 'none' })
    }
  })
}

onMounted(() => {
  userInfo.value = getUserInfo()
})
</script>

<style scoped>
.container {
  min-height: 100vh;
  background: #f5f7fa;
  padding: 40rpx 30rpx;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40rpx;
  padding: 30rpx;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 24rpx;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 24rpx;
}

.avatar {
  width: 80rpx;
  height: 80rpx;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 36rpx;
  font-weight: 500;
}

.user-text {
  display: flex;
  flex-direction: column;
  gap: 8rpx;
}

.greeting {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
}

.name {
  font-size: 32rpx;
  font-weight: 500;
  color: #fff;
}

.role-tag {
  padding: 8rpx 24rpx;
  border-radius: 30rpx;
  font-size: 24rpx;
  background: rgba(255, 255, 255, 0.2);
  color: #fff;
}

.stats-grid {
  display: flex;
  background: #fff;
  border-radius: 24rpx;
  padding: 40rpx 20rpx;
  margin-bottom: 40rpx;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
}

.stat-item {
  flex: 1;
  text-align: center;
}

.stat-value {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: #667eea;
  margin-bottom: 12rpx;
}

.stat-label {
  display: block;
  font-size: 24rpx;
  color: #999;
}

.section {
  background: #fff;
  border-radius: 24rpx;
  padding: 40rpx 30rpx;
  margin-bottom: 30rpx;
  box-shadow: 0 4rpx 20rpx rgba(0, 0, 0, 0.05);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30rpx;
}

.section-title {
  font-size: 32rpx;
  font-weight: 500;
  color: #333;
}

.more {
  font-size: 26rpx;
  color: #667eea;
}

.action-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  gap: 30rpx;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16rpx;
}

.icon {
  width: 100rpx;
  height: 100rpx;
  border-radius: 24rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 40rpx;
}

.task-icon { background: #e8f3ff; }
.record-icon { background: #e6fffb; }
.ticket-icon { background: #fff7e6; }
.scan-icon { background: #f9f0ff; }

.action-label {
  font-size: 24rpx;
  color: #666;
}

.task-list {
  display: flex;
  flex-direction: column;
  gap: 20rpx;
}

.task-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24rpx 0;
  border-bottom: 1rpx solid #f0f0f0;
}

.task-item:last-child {
  border-bottom: none;
}

.task-info {
  flex: 1;
}

.task-name {
  display: block;
  font-size: 28rpx;
  font-weight: 500;
  color: #333;
  margin-bottom: 8rpx;
}

.task-desc {
  display: block;
  font-size: 24rpx;
  color: #999;
}

.status-tag {
  padding: 8rpx 20rpx;
  border-radius: 8rpx;
  font-size: 22rpx;
}

.status-tag.pending { background: #f5f5f5; color: #999; }
.status-tag.running { background: #e6f7ff; color: #1890ff; }
.status-tag.completed { background: #f6ffed; color: #52c41a; }

.empty-tasks {
  text-align: center;
  padding: 60rpx 0;
}

.empty-text {
  font-size: 28rpx;
  color: #999;
}
</style>
