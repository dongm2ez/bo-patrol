<template>
  <view class="login-container">
    <view class="login-header">
      <view class="logo-icon">
        <text>🏛️</text>
      </view>
      <text class="title">博物馆智能巡检</text>
      <text class="subtitle">移动端</text>
    </view>
    
    <view class="login-form">
      <view class="form-item">
        <input
          v-model="loginForm.username"
          placeholder="请输入用户名"
          class="input"
          type="text"
        />
      </view>
      
      <view class="form-item">
        <input
          v-model="loginForm.password"
          placeholder="请输入密码"
          class="input"
          type="password"
          password
        />
      </view>
      
      <button
        class="login-btn"
        :loading="loading"
        @click="handleLogin"
      >
        登 录
      </button>
    </view>
    
    <view class="login-tips">
      <text class="tip-title">测试账号：</text>
      <text class="tip-text">管理员：admin / admin123</text>
      <text class="tip-text">巡检员：inspector / 123456</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { login, setUserInfo, type LoginRequest } from '../../api/auth'

const loading = ref(false)

const loginForm = reactive<LoginRequest>({
  username: '',
  password: ''
})

const handleLogin = async () => {
  if (!loginForm.username.trim()) {
    uni.showToast({ title: '请输入用户名', icon: 'none' })
    return
  }
  
  if (!loginForm.password.trim()) {
    uni.showToast({ title: '请输入密码', icon: 'none' })
    return
  }
  
  loading.value = true
  
  try {
    const res = await login(loginForm)
    
    uni.setStorageSync('token', res.data.token)
    setUserInfo(res.data.user)
    
    uni.showToast({ title: '登录成功', icon: 'success' })
    
    setTimeout(() => {
      uni.switchTab({ url: '/pages/index/index' })
    }, 1000)
  } catch (err: any) {
    console.error('登录失败:', err)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  flex-direction: column;
  padding: 80rpx 60rpx;
}

.login-header {
  text-align: center;
  margin-bottom: 80rpx;
}

.logo-icon {
  font-size: 120rpx;
  margin-bottom: 20rpx;
}

.title {
  display: block;
  font-size: 48rpx;
  font-weight: bold;
  color: #fff;
  margin-bottom: 16rpx;
}

.subtitle {
  display: block;
  font-size: 28rpx;
  color: rgba(255, 255, 255, 0.8);
}

.login-form {
  background: #fff;
  border-radius: 24rpx;
  padding: 60rpx 40rpx;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.form-item {
  margin-bottom: 40rpx;
}

.input {
  width: 100%;
  height: 88rpx;
  border: 2rpx solid #e5e7eb;
  border-radius: 16rpx;
  padding: 0 32rpx;
  font-size: 28rpx;
  box-sizing: border-box;
  background: #fafafa;
}

.input:focus {
  border-color: #667eea;
  background: #fff;
}

.login-btn {
  width: 100%;
  height: 88rpx;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: #fff;
  border: none;
  border-radius: 16rpx;
  font-size: 32rpx;
  font-weight: 500;
  letter-spacing: 8rpx;
  margin-top: 20rpx;
}

.login-btn::after {
  border: none;
}

.login-tips {
  margin-top: 60rpx;
  padding: 40rpx;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 16rpx;
}

.tip-title {
  display: block;
  font-size: 28rpx;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.9);
  margin-bottom: 20rpx;
}

.tip-text {
  display: block;
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 12rpx;
  line-height: 1.5;
}

.tip-text:last-child {
  margin-bottom: 0;
}
</style>
