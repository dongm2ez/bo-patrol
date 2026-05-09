<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <div class="brand-logo">
          <span class="logo-icon">🛡️</span>
          <h1 class="brand-name">博巡</h1>
          <span class="brand-english">BoPatrol</span>
        </div>
        <p class="brand-slogan">博物馆智能巡检，开源守护文化</p>
        <p class="subtitle">管理后台登录</p>
      </div>

      <el-form
        ref="loginFormRef"
        :model="loginForm"
        :rules="rules"
        class="login-form"
        @keyup.enter="handleLogin"
      >
        <el-form-item prop="username">
          <el-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            size="large"
            prefix-icon="User"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            size="large"
            prefix-icon="Lock"
            show-password
          />
        </el-form-item>

        <el-form-item>
          <el-button
            type="primary"
            size="large"
            class="login-btn"
            :loading="loading"
            @click="handleLogin"
          >
            登 录
          </el-button>
        </el-form-item>
      </el-form>

      <div class="login-tips">
        <p>测试账号：</p>
        <p>管理员：admin / admin123</p>
        <p>巡检员：inspector / 123456</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { login, setUserInfo, type LoginRequest } from '@/api/auth'

const router = useRouter()
const loginFormRef = ref()
const loading = ref(false)

const loginForm = reactive<LoginRequest>({
  username: '',
  password: ''
})

const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  await loginFormRef.value.validate()

  loading.value = true
  try {
    const res = await login(loginForm)

    localStorage.setItem('token', res.data.token)
    setUserInfo(res.data.user)

    ElMessage.success('登录成功')

    router.push('/')
  } catch (err: any) {
    console.error('登录失败:', err)
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-box {
  width: 420px;
  padding: 40px;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.brand-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin-bottom: 8px;
}

.logo-icon {
  font-size: 40px;
}

.brand-name {
  font-size: 36px;
  font-weight: 700;
  color: #303133;
  margin: 0;
  letter-spacing: 4px;
}

.brand-english {
  font-size: 16px;
  color: #909399;
  font-weight: 500;
  letter-spacing: 2px;
}

.brand-slogan {
  font-size: 15px;
  color: #606266;
  margin: 12px 0 4px 0;
  font-weight: 500;
  letter-spacing: 1px;
}

.subtitle {
  font-size: 13px;
  color: #999;
  margin: 0;
}

.login-form {
  margin-top: 20px;
}

.login-btn {
  width: 100%;
  height: 48px;
  font-size: 16px;
  letter-spacing: 4px;
}

.login-tips {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
  text-align: center;
}

.login-tips p {
  margin: 4px 0;
  font-size: 13px;
  color: #999;
}
</style>
