<template>
  <view class="container">
    <view class="record-list">
      <view class="record-card" v-for="record in list" :key="record.id">
        <view class="record-header">
          <text class="task-name">{{ record.taskName }}</text>
          <view :class="['status-tag', record.status]">
            <text v-if="record.status === 'normal'">正常</text>
            <text v-else-if="record.status === 'abnormal'">异常</text>
            <text v-else>{{ record.status }}</text>
          </view>
        </view>
        <view class="record-info">
          <view class="info-row">
            <text class="label">点位：</text>
            <text class="value">{{ record.pointName }}</text>
          </view>
          <view class="info-row">
            <text class="label">巡检员：</text>
            <text class="value">{{ record.inspectorName }}</text>
          </view>
          <view class="info-row">
            <text class="label">打卡时间：</text>
            <text class="value">{{ record.checkTime }}</text>
          </view>
          <view class="info-row" v-if="record.checkMethod">
            <text class="label">打卡方式：</text>
            <text class="value">{{ record.checkMethod === 'qr_code' ? '扫码' : record.checkMethod }}</text>
          </view>
        </view>
        <view class="ticket-btn" v-if="record.status === 'abnormal' && !record.ticketId">
          <button size="mini" @click="createTicket(record)">创建工单</button>
        </view>
      </view>
    </view>
    
    <view class="empty" v-if="list.length === 0">
      <text>暂无记录</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      list: [],
      page: 1,
      pageSize: 10
    }
  },
  onShow() {
    this.loadRecords()
  },
  methods: {
    loadRecords() {
      uni.request({
        url: `/api/v1/records?page=${this.page}&pageSize=${this.pageSize}`,
        header: { Authorization: `Bearer ${uni.getStorageSync('token')}` },
        success: (res) => {
          if (res.data.code === 200) {
            this.list = res.data.data.list || []
          }
        }
      })
    },
    createTicket(record) {
      uni.navigateTo({ url: '/pages/ticket/create?recordId=' + record.id })
    }
  }
}
</script>

<style scoped>
.container {
  min-height: 100vh;
  background: #f5f7fa;
}

.record-list {
  padding: 30rpx;
  gap: 20rpx;
  display: flex;
  flex-direction: column;
}

.record-card {
  background: #fff;
  border-radius: 16rpx;
  padding: 30rpx;
}

.record-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20rpx;
}

.task-name {
  font-size: 28rpx;
  font-weight: 500;
  color: #333;
}

.status-tag {
  padding: 6rpx 16rpx;
  border-radius: 8rpx;
  font-size: 22rpx;
}

.status-tag.normal { background: #f6ffed; color: #52c41a; }
.status-tag.abnormal { background: #fff2f0; color: #ff4d4f; }

.record-info {
  margin-bottom: 16rpx;
}

.info-row {
  display: flex;
  margin-bottom: 12rpx;
}

.info-row:last-child {
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

.ticket-btn {
  padding-top: 16rpx;
  border-top: 1rpx solid #f0f0f0;
}

.ticket-btn button {
  background: #ff4d4f;
  color: #fff;
  border: none;
  font-size: 24rpx;
  height: 56rpx;
  line-height: 56rpx;
}

.empty {
  text-align: center;
  padding: 100rpx 0;
  color: #999;
  font-size: 28rpx;
}
</style>
