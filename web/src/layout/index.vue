<template>
  <div class="layout-container">
    <el-container>
      <el-aside :width="isCollapse ? '64px' : '240px'" class="sidebar">
        <div class="logo">
          <div v-if="!isCollapse" class="logo-full">
            <span class="logo-icon-text">🛡️</span>
            <span class="logo-text">博巡</span>
            <span class="logo-slogan">BoPatrol</span>
          </div>
          <div v-else class="logo-mini">
            <span class="logo-icon-text">🛡️</span>
          </div>
        </div>

        <el-menu
          :default-active="activeMenu"
          :collapse="isCollapse"
          :unique-opened="true"
          router
          background-color="#304156"
          text-color="#bfcbd9"
          active-text-color="#409eff"
        >
          <el-menu-item index="/dashboard">
            <el-icon><DataLine /></el-icon>
            <template #title>数据概览</template>
          </el-menu-item>

          <el-menu-item index="/department">
            <el-icon><OfficeBuilding /></el-icon>
            <template #title>科室管理</template>
          </el-menu-item>

          <el-menu-item index="/equipment">
            <el-icon><Tools /></el-icon>
            <template #title>设备台账</template>
          </el-menu-item>

          <el-menu-item index="/space">
            <el-icon><Grid /></el-icon>
            <template #title>空间管理</template>
          </el-menu-item>

          <el-menu-item index="/map-editor">
            <el-icon><Location /></el-icon>
            <template #title>地图编辑</template>
          </el-menu-item>

          <el-menu-item index="/user">
            <el-icon><User /></el-icon>
            <template #title>用户管理</template>
          </el-menu-item>

          <el-menu-item index="/route">
            <el-icon><Guide /></el-icon>
            <template #title>路线管理</template>
          </el-menu-item>

          <el-menu-item index="/task">
            <el-icon><List /></el-icon>
            <template #title>任务管理</template>
          </el-menu-item>

          <el-menu-item index="/record">
            <el-icon><Document /></el-icon>
            <template #title>巡检记录</template>
          </el-menu-item>

          <el-menu-item index="/ticket">
            <el-icon><Tickets /></el-icon>
            <template #title>工单管理</template>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container class="main-container">
        <el-header class="header">
          <div class="header-left">
            <el-icon class="collapse-btn" @click="toggleCollapse">
              <Fold v-if="!isCollapse" />
              <Expand v-else />
            </el-icon>
            <span class="page-title">{{ pageTitle }}</span>
          </div>

          <div class="header-right">
            <el-dropdown @command="handleCommand">
              <span class="user-info">
                <el-avatar :size="32" class="avatar">
                  {{ userInfo?.name?.charAt(0) || 'U' }}
                </el-avatar>
                <span class="username">{{ userInfo?.name || '用户' }}</span>
                <el-icon class="arrow"><ArrowDown /></el-icon>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item disabled>
                    <span style="color: #999">
                      角色: {{ getRoleName(userInfo?.role) }}
                    </span>
                  </el-dropdown-item>
                  <el-dropdown-item divided command="logout">
                    <el-icon><SwitchButton /></el-icon>
                    退出登录
                  </el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-header>

        <el-main class="main-content">
          <router-view />
        </el-main>

        <el-footer class="footer">
          <span>© 2026 博巡 - 博物馆智能巡检系统 | 开源守护文化</span>
        </el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Menu,
  DataLine,
  OfficeBuilding,
  Tools,
  Grid,
  Location,
  User,
  Guide,
  List,
  Document,
  Tickets,
  Fold,
  Expand,
  ArrowDown,
  SwitchButton
} from '@element-plus/icons-vue'
import { getUserInfo, logout } from '@/api/auth'

const router = useRouter()
const route = useRoute()

const isCollapse = ref(false)
const userInfo = ref(getUserInfo())

const activeMenu = computed(() => route.path)

const pageTitle = computed(() => {
  const titles: Record<string, string> = {
    '/dashboard': '数据概览',
    '/department': '科室管理',
    '/equipment': '设备台账',
    '/space': '空间管理',
    '/map-editor': '地图编辑',
    '/user': '用户管理',
    '/route': '路线管理',
    '/task': '任务管理',
    '/record': '巡检记录',
    '/ticket': '工单管理'
  }
  return titles[route.path] || ''
})

const toggleCollapse = () => {
  isCollapse.value = !isCollapse.value
}

const getRoleName = (role?: string) => {
  const roleMap: Record<string, string> = {
    'admin': '管理员',
    'inspector': '巡检员',
    'repair': '维修员'
  }
  return roleMap[role || ''] || role
}

const handleCommand = (command: string) => {
  if (command === 'logout') {
    ElMessageBox.confirm('确定要退出登录吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {
      logout()
      ElMessage.success('退出成功')
      router.push('/login')
    })
  }
}
</script>

<style scoped>
.layout-container {
  width: 100%;
  height: 100%;
}

.sidebar {
  background-color: #304156;
  height: 100vh;
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  transition: width 0.3s;
  overflow-x: hidden;
  z-index: 1000;
}

.logo {
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(90deg, #2b3a4a 0%, #1f2d3d 100%);
  color: #fff;
  border-bottom: 1px solid #1f2d3d;
}

.logo-full {
  display: flex;
  align-items: center;
  gap: 8px;
}

.logo-mini {
  display: flex;
  align-items: center;
  justify-content: center;
}

.logo-icon-text {
  font-size: 24px;
}

.logo-text {
  font-size: 20px;
  font-weight: 700;
  color: #fff;
  letter-spacing: 2px;
}

.logo-slogan {
  font-size: 12px;
  color: #909399;
  margin-left: 4px;
}

.el-menu {
  border-right: none;
}

.main-container {
  margin-left: 240px;
  transition: margin-left 0.3s;
  min-height: 100vh;
}

.sidebar.collapsed + .main-container {
  margin-left: 64px;
}

.header {
  background-color: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  height: 60px;
  position: sticky;
  top: 0;
  z-index: 999;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.collapse-btn {
  font-size: 20px;
  cursor: pointer;
  color: #606266;
  transition: color 0.3s;
}

.collapse-btn:hover {
  color: #409eff;
}

.page-title {
  font-size: 16px;
  font-weight: 500;
  color: #303133;
}

.header-right {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  color: #606266;
}

.avatar {
  background: #409eff;
  color: #fff;
  font-weight: 500;
}

.username {
  font-size: 14px;
}

.arrow {
  font-size: 12px;
}

.main-content {
  background-color: #f0f2f5;
  min-height: calc(100vh - 60px - 48px);
  padding: 20px;
}

.footer {
  height: 48px;
  background: #fff;
  border-top: 1px solid #eaeaea;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 13px;
  color: #999;
}
</style>
