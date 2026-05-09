<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>工单管理</span>
          <el-button type="primary" @click="openDialog">创建工单</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border stripe style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="title" label="工单标题" />
        <el-table-column prop="type" label="类型">
          <template #default="{ row }">
            <el-tag v-if="row.type === 'equipment'" type="info">设备故障</el-tag>
            <el-tag v-else-if="row.type === 'safety'" type="warning">安全隐患</el-tag>
            <el-tag v-else-if="row.type === 'environment'" type="success">环境问题</el-tag>
            <el-tag v-else>其他</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级">
          <template #default="{ row }">
            <el-tag v-if="row.priority === 'urgent'" type="danger">紧急</el-tag>
            <el-tag v-else-if="row.priority === 'high'" type="warning">高</el-tag>
            <el-tag v-else-if="row.priority === 'medium'" type="info">中</el-tag>
            <el-tag v-else type="success">低</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag v-if="row.status === 'pending'" type="info">待处理</el-tag>
            <el-tag v-else-if="row.status === 'processing'" type="primary">处理中</el-tag>
            <el-tag v-else-if="row.status === 'completed'" type="success">已完成</el-tag>
            <el-tag v-else>{{ row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="reporterName" label="上报人" />
        <el-table-column prop="assigneeName" label="处理人" />
        <el-table-column prop="createdAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="250">
          <template #default="{ row }">
            <el-button v-if="row.status === 'pending'" size="small" type="primary" @click="assignTicket(row)">指派</el-button>
            <el-button v-if="row.status === 'processing'" size="small" type="success" @click="completeTicket(row)">完成</el-button>
            <el-button size="small" @click="viewDetail(row)">详情</el-button>
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
    
    <el-dialog v-model="dialogVisible" title="创建工单" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="工单标题" prop="title">
          <el-input v-model="form.title" placeholder="请输入工单标题" />
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择类型">
            <el-option label="设备故障" value="equipment" />
            <el-option label="安全隐患" value="safety" />
            <el-option label="环境问题" value="environment" />
            <el-option label="其他" value="other" />
          </el-select>
        </el-form-item>
        <el-form-item label="优先级" prop="priority">
          <el-select v-model="form.priority" placeholder="请选择优先级">
            <el-option label="紧急" value="urgent" />
            <el-option label="高" value="high" />
            <el-option label="中" value="medium" />
            <el-option label="低" value="low" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="4" placeholder="请详细描述问题" />
        </el-form-item>
        <el-form-item label="位置" prop="location">
          <el-input v-model="form.location" placeholder="请输入问题位置" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="assignVisible" title="指派工单" width="400px">
      <el-form label-width="80px">
        <el-form-item label="处理人">
          <el-select v-model="assignForm.assigneeId" placeholder="请选择处理人" style="width: 100%">
            <el-option v-for="user in users" :key="user.id" :label="user.name" :value="user.id" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="assignVisible = false">取消</el-button>
        <el-button type="primary" @click="confirmAssign">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="detailVisible" title="工单详情" width="600px">
      <el-descriptions :column="2" border v-if="currentTicket">
        <el-descriptions-item label="工单标题">{{ currentTicket.title }}</el-descriptions-item>
        <el-descriptions-item label="类型">{{ currentTicket.type }}</el-descriptions-item>
        <el-descriptions-item label="优先级">{{ currentTicket.priority }}</el-descriptions-item>
        <el-descriptions-item label="状态">{{ currentTicket.status }}</el-descriptions-item>
        <el-descriptions-item label="上报人">{{ currentTicket.reporterName }}</el-descriptions-item>
        <el-descriptions-item label="处理人">{{ currentTicket.assigneeName || '-' }}</el-descriptions-item>
        <el-descriptions-item label="位置" :span="2">{{ currentTicket.location }}</el-descriptions-item>
        <el-descriptions-item label="描述" :span="2">{{ currentTicket.description }}</el-descriptions-item>
      </el-descriptions>
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
const assignVisible = ref(false)
const detailVisible = ref(false)
const formRef = ref()
const users = ref([])
const currentTicket = ref()
const currentTicketId = ref()

const form = reactive({
  title: '',
  type: 'other',
  priority: 'medium',
  description: '',
  location: ''
})

const assignForm = reactive({
  assigneeId: undefined
})

const rules = {
  title: [{ required: true, message: '请输入工单标题', trigger: 'blur' }],
  description: [{ required: true, message: '请输入问题描述', trigger: 'blur' }]
}

const loadData = async () => {
  try {
    const res = await request.get('/tickets', { params: { page: page.value, pageSize: pageSize.value } })
    tableData.value = res.data.list
    total.value = res.data.total
  } catch (e) {
    console.error('加载工单列表失败', e)
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
  Object.assign(form, { title: '', type: 'other', priority: 'medium', description: '', location: '' })
  dialogVisible.value = true
}

const assignTicket = (row: any) => {
  currentTicketId.value = row.id
  assignForm.assigneeId = undefined
  assignVisible.value = true
}

const confirmAssign = async () => {
  try {
    await request.put(`/tickets/${currentTicketId.value}/assign`, assignForm)
    ElMessage.success('指派成功')
    assignVisible.value = false
    loadData()
  } catch (e) {
    console.error('指派失败', e)
  }
}

const completeTicket = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定要完成该工单吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await request.put(`/tickets/${row.id}/complete`)
    ElMessage.success('工单已完成')
    loadData()
  } catch {}
}

const viewDetail = async (row: any) => {
  try {
    const res = await request.get(`/tickets/${row.id}`)
    currentTicket.value = res.data
    detailVisible.value = true
  } catch (e) {
    console.error('加载详情失败', e)
  }
}

const submitForm = async () => {
  await formRef.value.validate()
  
  try {
    await request.post('/tickets', form)
    ElMessage.success('创建成功')
    dialogVisible.value = false
    loadData()
  } catch (e) {
    console.error('创建工单失败', e)
  }
}

onMounted(() => {
  loadData()
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
