<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>任务管理</span>
          <el-button type="primary" @click="openDialog">创建任务</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border stripe style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="任务名称" />
        <el-table-column prop="routeName" label="巡检路线" />
        <el-table-column prop="assigneeName" label="执行人" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag v-if="row.status === 'pending'" type="info">待开始</el-tag>
            <el-tag v-else-if="row.status === 'running'" type="primary">进行中</el-tag>
            <el-tag v-else-if="row.status === 'completed'" type="success">已完成</el-tag>
            <el-tag v-else>{{ row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="planTime" label="计划时间" width="180" />
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button v-if="row.status === 'pending'" size="small" type="primary" @click="startTask(row)">开始</el-button>
            <el-button v-if="row.status === 'running'" size="small" type="success" @click="completeTask(row)">完成</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="loadData"
        @size-change="loadData"
      />
    </el-card>
    
    <el-dialog v-model="dialogVisible" title="创建任务" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="任务名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入任务名称" />
        </el-form-item>
        <el-form-item label="巡检路线" prop="routeId">
          <el-select v-model="form.routeId" placeholder="请选择巡检路线">
            <el-option v-for="route in routes" :key="route.id" :label="route.name" :value="route.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="执行人" prop="assigneeId">
          <el-select v-model="form.assigneeId" placeholder="请选择执行人">
            <el-option v-for="user in users" :key="user.id" :label="user.name" :value="user.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="计划时间" prop="planTime">
          <el-date-picker v-model="form.planTime" type="datetime" placeholder="选择日期时间" style="width: 100%" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入任务描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import request from '@/utils/request'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const dialogVisible = ref(false)
const formRef = ref()
const routes = ref([])
const users = ref([])

const form = reactive({
  name: '',
  routeId: undefined,
  assigneeId: undefined,
  planTime: '',
  description: ''
})

const rules = {
  name: [{ required: true, message: '请输入任务名称', trigger: 'blur' }],
  routeId: [{ required: true, message: '请选择巡检路线', trigger: 'change' }],
  assigneeId: [{ required: true, message: '请选择执行人', trigger: 'change' }],
  planTime: [{ required: true, message: '请选择计划时间', trigger: 'change' }]
}

const loadData = async () => {
  try {
    const res = await request.get('/tasks', { params: { page: page.value, pageSize: pageSize.value } })
    tableData.value = res.data.list
    total.value = res.data.total
  } catch (e) {
    console.error('加载任务列表失败', e)
  }
}

const loadRoutes = async () => {
  try {
    const res = await request.get('/routes')
    routes.value = res.data.list
  } catch (e) {
    console.error('加载路线列表失败', e)
  }
}

const loadUsers = async () => {
  try {
    const res = await request.get('/users')
    users.value = res.data.list
  } catch (e) {
    console.error('加载用户列表失败', e)
  }
}

const openDialog = () => {
  Object.assign(form, {
    name: '',
    routeId: undefined,
    assigneeId: undefined,
    planTime: '',
    description: ''
  })
  dialogVisible.value = true
}

const startTask = async (row: any) => {
  try {
    await request.put(`/tasks/${row.id}/start`)
    ElMessage.success('任务已开始')
    loadData()
  } catch (e) {
    console.error('开始任务失败', e)
  }
}

const completeTask = async (row: any) => {
  try {
    await request.put(`/tasks/${row.id}/complete`)
    ElMessage.success('任务已完成')
    loadData()
  } catch (e) {
    console.error('完成任务失败', e)
  }
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定要删除该任务吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const submitForm = async () => {
  await formRef.value.validate()
  
  try {
    await request.post('/tasks', form)
    ElMessage.success('创建成功')
    dialogVisible.value = false
    loadData()
  } catch (e) {
    console.error('创建任务失败', e)
  }
}

onMounted(() => {
  loadData()
  loadRoutes()
  loadUsers()
})
</script>

<style scoped>
.page-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-pagination {
  margin-top: 20px;
  justify-content: flex-end;
}
</style>
