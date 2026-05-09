<template>
  <view class="container">
    <view class="form-card">
      <view class="form-item">
        <text class="label">标题 <text class="required">*</text></text>
        <input v-model="form.title" placeholder="请输入工单标题" class="input" />
      </view>
      
      <view class="form-item">
        <text class="label">类型</text>
        <picker :value="typeIndex" :range="typeOptions" range-key="label" @change="onTypeChange">
          <view class="picker">
            <text>{{ typeOptions[typeIndex].label }}</text>
            <text class="arrow">›</text>
          </view>
        </picker>
      </view>
      
      <view class="form-item">
        <text class="label">优先级</text>
        <picker :value="priorityIndex" :range="priorityOptions" range-key="label" @change="onPriorityChange">
          <view class="picker">
            <text>{{ priorityOptions[priorityIndex].label }}</text>
            <text class="arrow">›</text>
          </view>
        </picker>
      </view>
      
      <view class="form-item">
        <text class="label">位置</text>
        <input v-model="form.location" placeholder="请输入问题位置" class="input" />
      </view>
      
      <view class="form-item">
        <text class="label">描述 <text class="required">*</text></text>
        <textarea v-model="form.description" placeholder="请详细描述问题情况" class="textarea" />
      </view>
      
      <view class="form-item">
        <text class="label">图片</text>
        <view class="upload-area" @click="chooseImage">
          <text class="upload-icon">+</text>
          <text class="upload-text">上传图片</text>
        </view>
      </view>
    </view>
    
    <view class="footer">
      <button class="submit-btn" @click="submit">提交工单</button>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      form: {
        title: '',
        type: 'equipment',
        priority: 'medium',
        description: '',
        location: '',
        images: ''
      },
      typeOptions: [
        { label: '设备故障', value: 'equipment' },
        { label: '安全隐患', value: 'safety' },
        { label: '环境问题', value: 'environment' },
        { label: '其他问题', value: 'other' }
      ],
      priorityOptions: [
        { label: '紧急', value: 'urgent' },
        { label: '高', value: 'high' },
        { label: '中', value: 'medium' },
        { label: '低', value: 'low' }
      ],
      typeIndex: 0,
      priorityIndex: 2
    }
  },
  onLoad(options) {
    if (options.recordId) {
      this.form.recordId = parseInt(options.recordId)
    }
  },
  methods: {
    onTypeChange(e) {
      this.typeIndex = e.detail.value
      this.form.type = this.typeOptions[this.typeIndex].value
    },
    onPriorityChange(e) {
      this.priorityIndex = e.detail.value
      this.form.priority = this.priorityOptions[this.priorityIndex].value
    },
    chooseImage() {
      uni.chooseImage({
        count: 3,
        success: (res) => {
          this.form.images = res.tempFilePaths.join(',')
          uni.showToast({ title: '上传成功', icon: 'success' })
        }
      })
    },
    submit() {
      if (!this.form.title || !this.form.description) {
        uni.showToast({ title: '请填写必填项', icon: 'none' })
        return
      }
      
      uni.request({
        url: '/api/v1/tickets',
        method: 'POST',
        header: { Authorization: `Bearer ${uni.getStorageSync('token')}` },
        data: this.form,
        success: (res) => {
          if (res.data.code === 200) {
            uni.showToast({ title: '提交成功', icon: 'success' })
            setTimeout(() => {
              uni.navigateBack()
            }, 1000)
          } else {
            uni.showToast({ title: res.data.message, icon: 'none' })
          }
        },
        fail: () => {
          uni.showToast({ title: '提交失败', icon: 'none' })
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
  padding-bottom: 120rpx;
}

.form-card {
  background: #fff;
  margin: 30rpx;
  border-radius: 16rpx;
  padding: 30rpx;
}

.form-item {
  margin-bottom: 40rpx;
}

.form-item:last-child {
  margin-bottom: 0;
}

.label {
  display: block;
  font-size: 28rpx;
  color: #333;
  margin-bottom: 16rpx;
}

.required {
  color: #ff4d4f;
}

.input {
  width: 100%;
  height: 80rpx;
  border: 2rpx solid #e5e7eb;
  border-radius: 12rpx;
  padding: 0 24rpx;
  font-size: 28rpx;
  box-sizing: border-box;
}

.textarea {
  width: 100%;
  min-height: 200rpx;
  border: 2rpx solid #e5e7eb;
  border-radius: 12rpx;
  padding: 20rpx 24rpx;
  font-size: 28rpx;
  box-sizing: border-box;
}

.picker {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 80rpx;
  border: 2rpx solid #e5e7eb;
  border-radius: 12rpx;
  padding: 0 24rpx;
  font-size: 28rpx;
  color: #333;
}

.arrow {
  color: #999;
  font-size: 32rpx;
}

.upload-area {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 200rpx;
  height: 200rpx;
  border: 2rpx dashed #d9d9d9;
  border-radius: 12rpx;
  background: #fafafa;
}

.upload-icon {
  font-size: 48rpx;
  color: #d9d9d9;
  margin-bottom: 8rpx;
}

.upload-text {
  font-size: 24rpx;
  color: #999;
}

.footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #fff;
  padding: 20rpx 30rpx;
  box-shadow: 0 -4rpx 20rpx rgba(0, 0, 0, 0.05);
}

.submit-btn {
  width: 100%;
  height: 88rpx;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 12rpx;
  font-size: 32rpx;
  font-weight: 500;
}
</style>
