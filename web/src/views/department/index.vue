<template>
  <div class="department-container">
    <div class="page-header">
      <h2>科室管理</h2>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增科室
      </el-button>
    </div>

    <div class="page-content">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card class="tree-card">
            <template #header>
              <span>科室树形结构</span>
            </template>
            <el-tree
              ref="treeRef"
              :data="treeData"
              :props="{ label: 'name', children: 'children' }"
              node-key="id"
              default-expand-all
              @node-click="handleNodeClick"
              :highlight-current="true"
            >
              <template #default="{ node, data }">
                <span class="custom-tree-node">
                  <el-icon v-if="data.type === 'water'" style="color: #409eff">
                    <Watermelon />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'electric'" style="color: #e6a23c">
                    <Lightning />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'elevator'" style="color: #67c23a">
                    <Tools />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'bas'" style="color: #909399">
                    <Monitor />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'audio'" style="color: #f56c6c">
                    <Microphone />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'civil'" style="color: #606266">
                    <OfficeBuilding />
                  </el-icon>
                  <el-icon v-else-if="data.type === 'finance'" style="color: #409eff">
                    <Money />
                  </el-icon>
                  <el-icon v-else style="color: #409eff">
                    <OfficeBuilding />
                  </el-icon>
                  <span class="node-label">{{ node.label }}</span>
                </span>
              </template>
            </el-tree>
          </el-card>
        </el-col>

        <el-col :span="18">
          <el-card>
            <template #header>
              <div class="card-header">
                <span>科室列表</span>
                <div class="header-actions">
                  <el-select
                    v-model="searchForm.type"
                    placeholder="按类型筛选"
                    style="width: 150px; margin-right: 10px"
                    clearable
                    @change="loadDepartments"
                  >
                    <el-option
                      v-for="item in departmentTypes"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                  </el-select>
                </div>
              </div>
            </template>

            <el-table :data="tableData" v-loading="loading" border stripe>
              <el-table-column prop="id" label="ID" width="80" />
              <el-table-column prop="code" label="科室编号" width="120" />
              <el-table-column prop="name" label="科室名称" width="150" />
              <el-table-column label="科室类型" width="120">
                <template #default="{ row }">
                  <el-tag>{{ getTypeLabel(row.type) }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="phone" label="联系电话" width="130" />
              <el-table-column prop="email" label="邮箱" width="180" />
              <el-table-column label="状态" width="100">
                <template #default="{ row }">
                  <el-tag :type="row.status === 1 ? 'success' : 'danger'">
                    {{ row.status === 1 ? '启用' : '禁用' }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column prop="createdAt" label="创建时间" width="180">
                <template #default="{ row }">
                  {{ formatDate(row.createdAt) }}
                </template>
              </el-table-column>
              <el-table-column label="操作" width="150" fixed="right">
                <template #default="{ row }">
                  <el-button link type="primary" size="small" @click="handleEdit(row)">
                    编辑
                  </el-button>
                  <el-button link type="danger" size="small" @click="handleDelete(row)">
                    删除
                  </el-button>
                </template>
              </el-table-column>
            </el-table>
          </el-card>
        </el-col>
      </el-row>
    </div>

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="form"
        label-width="100px"
        :rules="formRules"
      >
        <el-form-item label="科室编号" prop="code">
          <el-input v-model="form.code" placeholder="请输入科室编号" />
        </el-form-item>
        <el-form-item label="科室名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入科室名称" />
        </el-form-item>
        <el-form-item label="科室类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择科室类型" style="width: 100%">
            <el-option
              v-for="item in departmentTypes"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="上级科室" prop="parentId">
          <el-tree-select
            v-model="form.parentId"
            :data="treeData"
            :props="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择上级科室"
            clearable
            check-strictly
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="联系电话" prop="phone">
          <el-input v-model="form.phone" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="form.sortOrder" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="form.description"
            type="textarea"
            :rows="3"
            placeholder="请输入描述"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import {
  Plus,
  Watermelon,
  Lightning,
  Tools,
  Monitor,
  Microphone,
  OfficeBuilding,
  Money
} from '@element-plus/icons-vue'
import {
  getDepartments,
  getDepartment,
  getDepartmentTree,
  createDepartment,
  updateDepartment,
  deleteDepartment,
  getDepartmentTypes
} from '@/api/department'
import type { Department, CreateDepartmentRequest } from '@/types'

const loading = ref(false)
const dialogVisible = ref(false)
const formRef = ref<FormInstance>()
const treeRef = ref()
const treeData = ref<Department[]>([])
const tableData = ref<Department[]>([])
const departmentTypes = ref<{ value: string; label: string }[]>([])

const searchForm = reactive({
  type: ''
})

const form = reactive<CreateDepartmentRequest & { id?: number }>({
  name: '',
  code: '',
  type: '',
  description: '',
  parentId: undefined,
  sortOrder: 0,
  phone: '',
  email: ''
})

const dialogTitle = computed(() => (form.id ? '编辑科室' : '新增科室'))

const formRules: FormRules = {
  name: [{ required: true, message: '请输入科室名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入科室编号', trigger: 'blur' }],
  type: [{ required: true, message: '请选择科室类型', trigger: 'change' }]
}

const formatDate = (date: string) => {
  return date ? new Date(date).toLocaleString('zh-CN') : '-'
}

const getTypeLabel = (type: string) => {
  const item = departmentTypes.value.find((d) => d.value === type)
  return item?.label || type
}

const loadDepartments = async () => {
  loading.value = true
  try {
    const res = await getDepartments({ type: searchForm.type })
    tableData.value = res.data.list
  } finally {
    loading.value = false
  }
}

const loadTreeData = async () => {
  try {
    const res = await getDepartmentTree()
    treeData.value = res.data
  } catch (error) {
    console.error('加载树数据失败:', error)
  }
}

const loadDepartmentTypes = async () => {
  try {
    const res = await getDepartmentTypes()
    departmentTypes.value = res.data
  } catch (error) {
    console.error('加载科室类型失败:', error)
  }
}

const handleNodeClick = (data: Department) => {
  // 可以在这里处理节点点击事件
  console.log('选中节点:', data)
}

const handleAdd = () => {
  form.id = undefined
  dialogVisible.value = true
}

const handleEdit = (row: Department) => {
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleDelete = (row: Department) => {
  ElMessageBox.confirm(`确定要删除科室【${row.name}】吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteDepartment(row.id)
      ElMessage.success('删除成功')
      loadDepartments()
      loadTreeData()
    } catch (error) {
      console.error('删除失败:', error)
    }
  })
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return

    try {
      if (form.id) {
        await updateDepartment(form.id, form)
        ElMessage.success('更新成功')
      } else {
        await createDepartment(form)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadDepartments()
      loadTreeData()
    } catch (error) {
      console.error('提交失败:', error)
    }
  })
}

const resetForm = () => {
  formRef.value?.resetFields()
  Object.assign(form, {
    name: '',
    code: '',
    type: '',
    description: '',
    parentId: undefined,
    sortOrder: 0,
    phone: '',
    email: ''
  })
}

onMounted(() => {
  loadDepartmentTypes()
  loadDepartments()
  loadTreeData()
})
</script>

<style scoped lang="scss">
.department-container {
  padding: 20px;

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;

    h2 {
      margin: 0;
      font-size: 20px;
      font-weight: 500;
    }
  }

  .tree-card {
    height: calc(100vh - 180px);
    overflow-y: auto;
  }

  .custom-tree-node {
    display: flex;
    align-items: center;

    .node-label {
      margin-left: 8px;
    }
  }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .header-actions {
      display: flex;
      align-items: center;
    }
  }
}
</style>
