<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>路线管理</span>
          <el-button type="primary" @click="openDialog">创建路线</el-button>
        </div>
      </template>
      
      <el-table :data="tableData" border stripe style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="路线名称" />
        <el-table-column prop="zone" label="区域" />
        <el-table-column prop="description" label="描述" />
        <el-table-column label="点位数量">
          <template #default="{ row }">
            {{ row.points?.length || 0 }}
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag v-if="row.status === 1" type="success">启用</el-tag>
            <el-tag v-else type="danger">禁用</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="managePoints(row)">点位管理</el-button>
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
    
    <el-dialog v-model="dialogVisible" title="创建路线" width="500px">
      <el-form :model="form" :rules="rules" ref="formRef" label-width="80px">
        <el-form-item label="路线名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入路线名称" />
        </el-form-item>
        <el-form-item label="区域" prop="zone">
          <el-input v-model="form.zone" placeholder="请输入所属区域" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入路线描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="pointsDialogVisible" title="点位管理" width="600px">
      <div class="points-header">
        <span>路线: {{ currentRoute?.name }}</span>
        <el-button type="primary" size="small" @click="openPointDialog">添加点位</el-button>
      </div>
      <el-table :data="currentPoints" border stripe style="width: 100%; margin-top: 15px;">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="点位名称" />
        <el-table-column prop="qrCode" label="二维码" />
        <el-table-column prop="checkItems" label="检查项" />
        <el-table-column prop="sort" label="排序" />
        <el-table-column label="操作" width="100">
          <template #default="{ row }">
            <el-button size="small" type="danger" @click="deletePoint(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <el-dialog v-model="pointDialogVisible" title="添加点位" width="500px">
        <el-form :model="pointForm" label-width="80px">
          <el-form-item label="点位名称">
            <el-input v-model="pointForm.name" placeholder="请输入点位名称" />
          </el-form-item>
          <el-form-item label="检查项">
            <el-input v-model="pointForm.checkItems" type="textarea" placeholder="请输入检查项" />
          </el-form-item>
          <el-form-item label="排序">
            <el-input-number v-model="pointForm.sort" :min="0" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="pointDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitPoint">确定</el-button>
        </template>
      </el-dialog>
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
const pointsDialogVisible = ref(false)
const pointDialogVisible = ref(false)
const formRef = ref()
const currentRoute = ref()
const currentPoints = ref([])

const form = reactive({
  name: '',
  zone: '',
  description: ''
})

const pointForm = reactive({
  name: '',
  checkItems: '',
  sort: 0
})

const rules = {
  name: [{ required: true, message: '请输入路线名称', trigger: 'blur' }]
}

const loadData = async () => {
  try {
    const res = await request.get('/routes', { params: { page: page.value, pageSize: pageSize.value } })
    tableData.value = res.data.list
    total.value = res.data.total
  } catch (e) {
    console.error('加载路线列表失败', e)
  }
}

const openDialog = () => {
  Object.assign(form, { name: '', zone: '', description: '' })
  dialogVisible.value = true
}

const managePoints = async (row: any) => {
  currentRoute.value = row
  try {
    const res = await request.get(`/routes/${row.id}/points`)
    currentPoints.value = res.data
  } catch (e) {
    console.error('加载点位失败', e)
  }
  pointsDialogVisible.value = true
}

const openPointDialog = () => {
  Object.assign(pointForm, { name: '', checkItems: '', sort: 0 })
  pointDialogVisible.value = true
}

const submitPoint = async () => {
  try {
    await request.post(`/routes/${currentRoute.value.id}/points`, pointForm)
    ElMessage.success('添加点位成功')
    pointDialogVisible.value = false
    managePoints(currentRoute.value)
  } catch (e) {
    console.error('添加点位失败', e)
  }
}

const deletePoint = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定要删除该点位吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    ElMessage.success('删除成功')
    managePoints(currentRoute.value)
  } catch {}
}

const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm('确定要删除该路线吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await request.delete(`/routes/${row.id}`)
    ElMessage.success('删除成功')
    loadData()
  } catch {}
}

const submitForm = async () => {
  await formRef.value.validate()
  
  try {
    await request.post('/routes', form)
    ElMessage.success('创建成功')
    dialogVisible.value = false
    loadData()
  } catch (e) {
    console.error('创建路线失败', e)
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
.page-container {
  padding: 20px;
}

.card-header, .points-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-pagination {
  margin-top: 20px;
  justify-content: flex-end;
}
</style>
