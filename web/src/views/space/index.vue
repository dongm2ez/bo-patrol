<template>
  <div class="space-container">
    <div class="page-header">
      <h2>空间管理</h2>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增空间
      </el-button>
    </div>

    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="tree-card">
          <template #header>
            <span>空间结构</span>
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
                <el-icon v-if="data.type === 'building'" style="color: #409eff">
                  <OfficeBuilding />
                </el-icon>
                <el-icon v-else-if="data.type === 'floor'" style="color: #67c23a">
                  <Grid />
                </el-icon>
                <el-icon v-else-if="data.type === 'zone'" style="color: #e6a23c">
                  <Location />
                </el-icon>
                <el-icon v-else style="color: #909399">
                  <Location />
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
            <span>空间列表</span>
          </template>

          <el-table :data="tableData" v-loading="loading" border stripe>
            <el-table-column prop="id" label="ID" width="70" />
            <el-table-column prop="code" label="空间编号" width="120" />
            <el-table-column prop="name" label="空间名称" min-width="150" />
            <el-table-column label="空间类型" width="100">
              <template #default="{ row }">
                <el-tag>{{ getTypeLabel(row.type) }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="floor" label="楼层" width="80" />
            <el-table-column prop="parentName" label="上级空间" width="120" />
            <el-table-column prop="description" label="描述" min-width="200" />
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
        <el-form-item label="空间编号" prop="code">
          <el-input v-model="form.code" placeholder="请输入空间编号" />
        </el-form-item>
        <el-form-item label="空间名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入空间名称" />
        </el-form-item>
        <el-form-item label="空间类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择空间类型" style="width: 100%">
            <el-option
              v-for="item in spaceTypes"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="上级空间" prop="parentId">
          <el-tree-select
            v-model="form.parentId"
            :data="treeData"
            :props="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择上级空间"
            clearable
            check-strictly
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="所属楼层" prop="floor">
          <el-input-number v-model="form.floor" style="width: 100%" placeholder="请输入楼层" />
        </el-form-item>
        <el-form-item label="经纬度" prop="coordsX">
          <el-input v-model="form.coordsX" placeholder="经度" style="width: 45%; margin-right: 10px" />
          <el-input v-model="form.coordsY" placeholder="纬度" style="width: 45%" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input
            v-model="form.description"
            type="textarea"
            :rows="3"
            placeholder="请输入描述"
          />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="form.sortOrder" :min="0" style="width: 100%" />
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
import { Plus, OfficeBuilding, Grid, Location } from '@element-plus/icons-vue'
import {
  getSpaces,
  getSpaceTree,
  getSpace,
  createSpace,
  updateSpace,
  deleteSpace,
  getSpaceTypes
} from '@/api/space'
import type { SpaceLocation, CreateSpaceRequest } from '@/types'

const loading = ref(false)
const dialogVisible = ref(false)
const formRef = ref<FormInstance>()
const treeRef = ref()
const treeData = ref<SpaceLocation[]>([])
const tableData = ref<SpaceLocation[]>([])
const spaceTypes = ref<{ value: string; label: string }[]>([])

const form = reactive<CreateSpaceRequest & { id?: number }>({
  name: '',
  code: '',
  type: '',
  description: '',
  parentId: undefined,
  floor: 1,
  sortOrder: 0,
  coordsX: 0,
  coordsY: 0
})

const dialogTitle = computed(() => (form.id ? '编辑空间' : '新增空间'))

const formRules: FormRules = {
  name: [{ required: true, message: '请输入空间名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入空间编号', trigger: 'blur' }],
  type: [{ required: true, message: '请选择空间类型', trigger: 'change' }]
}

const formatDate = (date: string) => {
  return date ? new Date(date).toLocaleString('zh-CN') : '-'
}

const getTypeLabel = (type: string) => {
  const item = spaceTypes.value.find((d) => d.value === type)
  return item?.label || type
}

const loadSpaces = async () => {
  loading.value = true
  try {
    const res = await getSpaces()
    tableData.value = res.data.list
  } finally {
    loading.value = false
  }
}

const loadTreeData = async () => {
  try {
    const res = await getSpaceTree()
    treeData.value = res.data
  } catch (error) {
    console.error('加载树数据失败:', error)
  }
}

const loadSpaceTypes = async () => {
  try {
    const res = await getSpaceTypes()
    spaceTypes.value = res.data
  } catch (error) {
    console.error('加载空间类型失败:', error)
  }
}

const handleNodeClick = (data: SpaceLocation) => {
  console.log('选中节点:', data)
}

const handleAdd = () => {
  form.id = undefined
  dialogVisible.value = true
}

const handleEdit = (row: SpaceLocation) => {
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleDelete = (row: SpaceLocation) => {
  ElMessageBox.confirm(`确定要删除空间【${row.name}】吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteSpace(row.id)
      ElMessage.success('删除成功')
      loadSpaces()
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
        await updateSpace(form.id, form)
        ElMessage.success('更新成功')
      } else {
        await createSpace(form)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadSpaces()
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
    floor: 1,
    sortOrder: 0,
    coordsX: 0,
    coordsY: 0
  })
}

onMounted(() => {
  loadSpaceTypes()
  loadSpaces()
  loadTreeData()
})
</script>

<style scoped lang="scss">
.space-container {
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
}
</style>
