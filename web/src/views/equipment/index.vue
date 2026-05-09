<template>
  <div class="equipment-container">
    <div class="page-header">
      <h2>设备台账</h2>
      <el-button type="primary" @click="handleAdd">
        <el-icon><Plus /></el-icon>
        新增设备
      </el-button>
    </div>

    <el-card class="filter-card">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="所属科室">
          <el-select
            v-model="searchForm.departmentId"
            placeholder="请选择科室"
            clearable
            style="width: 200px"
            @change="loadEquipments"
          >
            <el-option
              v-for="dept in departments"
              :key="dept.id"
              :label="dept.name"
              :value="dept.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="设备分类">
          <el-select
            v-model="searchForm.category"
            placeholder="请选择分类"
            clearable
            style="width: 150px"
            @change="loadEquipments"
          >
            <el-option
              v-for="item in equipmentCategories"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="设备状态">
          <el-select
            v-model="searchForm.status"
            placeholder="请选择状态"
            clearable
            style="width: 120px"
            @change="loadEquipments"
          >
            <el-option label="正常" value="normal" />
            <el-option label="故障" value="fault" />
            <el-option label="维护中" value="maintenance" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="loadEquipments">
            <el-icon><Search /></el-icon>
            搜索
          </el-button>
          <el-button @click="resetSearch">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card>
      <el-table :data="tableData" v-loading="loading" border stripe>
        <el-table-column prop="id" label="ID" width="70" />
        <el-table-column prop="code" label="设备编号" width="120" />
        <el-table-column prop="name" label="设备名称" min-width="150" />
        <el-table-column prop="categoryName" label="设备分类" width="100" />
        <el-table-column prop="brand" label="品牌" width="100" />
        <el-table-column prop="model" label="型号" width="120" />
        <el-table-column label="所属科室" width="120">
          <template #default="{ row }">
            {{ row.departmentName }}
          </template>
        </el-table-column>
        <el-table-column prop="locationDesc" label="放置位置" min-width="150" />
        <el-table-column label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">
              {{ getStatusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="purchaseDate" label="采购日期" width="120">
          <template #default="{ row }">
            {{ row.purchaseDate ? row.purchaseDate.substring(0, 10) : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
            <el-button link type="primary" size="small" @click="handleView(row)">
              详情
            </el-button>
            <el-button link type="primary" size="small" @click="handleEdit(row)">
              编辑
            </el-button>
            <el-button link type="danger" size="small" @click="handleDelete(row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="700px"
      @close="resetForm"
    >
      <el-form
        ref="formRef"
        :model="form"
        label-width="100px"
        :rules="formRules"
      >
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="设备编号" prop="code">
              <el-input v-model="form.code" placeholder="请输入设备编号" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="设备名称" prop="name">
              <el-input v-model="form.name" placeholder="请输入设备名称" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="设备分类" prop="category">
              <el-select
                v-model="form.category"
                placeholder="请选择分类"
                style="width: 100%"
              >
                <el-option
                  v-for="item in equipmentCategories"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="所属科室" prop="departmentId">
              <el-select
                v-model="form.departmentId"
                placeholder="请选择科室"
                style="width: 100%"
              >
                <el-option
                  v-for="dept in departments"
                  :key="dept.id"
                  :label="dept.name"
                  :value="dept.id"
                />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="8">
            <el-form-item label="品牌">
              <el-input v-model="form.brand" placeholder="请输入品牌" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="型号">
              <el-input v-model="form.model" placeholder="请输入型号" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="序列号">
              <el-input v-model="form.serialNumber" placeholder="请输入序列号" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="供应商" prop="supplierId">
              <el-select
                v-model="form.supplierId"
                placeholder="请选择供应商"
                style="width: 100%"
                clearable
              >
                <el-option
                  v-for="supplier in suppliers"
                  :key="supplier.id"
                  :label="supplier.name"
                  :value="supplier.id"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="维保厂家">
              <el-input v-model="form.maintenanceVendor" placeholder="请输入维保厂家" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="采购日期">
              <el-date-picker
                v-model="form.purchaseDate"
                type="date"
                placeholder="请选择采购日期"
                style="width: 100%"
                value-format="YYYY-MM-DD"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="采购价格">
              <el-input-number
                v-model="form.purchasePrice"
                :min="0"
                :precision="2"
                style="width: 100%"
                placeholder="请输入采购价格"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="放置位置" prop="locationDesc">
          <el-input
            v-model="form.locationDesc"
            placeholder="请输入放置位置"
          />
        </el-form-item>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="维保开始日期">
              <el-date-picker
                v-model="form.maintenanceStart"
                type="date"
                placeholder="请选择开始日期"
                style="width: 100%"
                value-format="YYYY-MM-DD"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="维保到期日期">
              <el-date-picker
                v-model="form.maintenanceEnd"
                type="date"
                placeholder="请选择到期日期"
                style="width: 100%"
                value-format="YYYY-MM-DD"
              />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="备注" prop="remark">
          <el-input
            v-model="form.remark"
            type="textarea"
            :rows="3"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>

      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确定</el-button>
      </template>
    </el-dialog>

    <el-dialog
      v-model="detailVisible"
      title="设备详情"
      width="800px"
    >
      <el-descriptions :column="2" border v-if="currentEquipment">
        <el-descriptions-item label="设备编号">
          {{ currentEquipment.code }}
        </el-descriptions-item>
        <el-descriptions-item label="设备名称">
          {{ currentEquipment.name }}
        </el-descriptions-item>
        <el-descriptions-item label="设备分类">
          {{ currentEquipment.categoryName }}
        </el-descriptions-item>
        <el-descriptions-item label="品牌">
          {{ currentEquipment.brand }}
        </el-descriptions-item>
        <el-descriptions-item label="型号">
          {{ currentEquipment.model }}
        </el-descriptions-item>
        <el-descriptions-item label="序列号">
          {{ currentEquipment.serialNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="所属科室">
          {{ currentEquipment.departmentName }}
        </el-descriptions-item>
        <el-descriptions-item label="供应商">
          {{ currentEquipment.supplierName || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="放置位置">
          {{ currentEquipment.locationDesc || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="维保厂家">
          {{ currentEquipment.maintenanceVendor || '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="采购日期">
          {{ currentEquipment.purchaseDate ? currentEquipment.purchaseDate.substring(0, 10) : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="采购价格">
          {{ currentEquipment.purchasePrice ? currentEquipment.purchasePrice + ' 元' : '-' }}
        </el-descriptions-item>
        <el-descriptions-item label="备注" :span="2">
          {{ currentEquipment.remark || '-' }}
        </el-descriptions-item>
      </el-descriptions>

      <div class="maintenance-section" v-if="currentEquipment">
        <h4>维保记录</h4>
        <el-table :data="maintenanceRecords" border stripe size="small">
          <el-table-column prop="type" label="类型" width="120" />
          <el-table-column prop="date" label="日期" width="180">
            <template #default="{ row }">
              {{ row.date ? row.date.substring(0, 10) : '-' }}
            </template>
          </el-table-column>
          <el-table-column prop="content" label="内容" />
          <el-table-column prop="operator" label="操作人" width="120" />
          <el-table-column prop="cost" label="费用(元)" width="100" align="right">
            <template #default="{ row }">
              {{ row.cost || 0 }}
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox, type FormInstance, type FormRules } from 'element-plus'
import { Plus, Search } from '@element-plus/icons-vue'
import {
  getEquipments,
  getEquipment,
  createEquipment,
  updateEquipment,
  deleteEquipment,
  getEquipmentCategories,
  getSuppliers
} from '@/api/equipment'
import { getDepartments } from '@/api/department'
import type { Equipment, CreateEquipmentRequest } from '@/types'

const loading = ref(false)
const dialogVisible = ref(false)
const detailVisible = ref(false)
const formRef = ref<FormInstance>()
const tableData = ref<Equipment[]>([])
const equipmentCategories = ref<{ value: string; label: string }[]>([])
const departments = ref<{ id: number; name: string }[]>([])
const suppliers = ref<{ id: number; name: string }[]>([])
const currentEquipment = ref<Equipment | null>(null)

const searchForm = reactive({
  departmentId: undefined as number | undefined,
  category: '',
  status: ''
})

const pagination = reactive({
  page: 1,
  pageSize: 10,
  total: 0
})

const form = reactive<CreateEquipmentRequest & { id?: number }>({
  name: '',
  code: '',
  category: '',
  brand: '',
  model: '',
  serialNumber: '',
  departmentId: 0,
  supplierId: undefined,
  locationDesc: '',
  purchaseDate: '',
  purchasePrice: 0,
  maintenanceVendor: '',
  remark: ''
})

const dialogTitle = computed(() => (form.id ? '编辑设备' : '新增设备'))

const formRules: FormRules = {
  name: [{ required: true, message: '请输入设备名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入设备编号', trigger: 'blur' }],
  category: [{ required: true, message: '请选择设备分类', trigger: 'change' }],
  departmentId: [{ required: true, message: '请选择所属科室', trigger: 'change' }]
}

const getStatusType = (status: string) => {
  const map: Record<string, string> = {
    normal: 'success',
    fault: 'danger',
    maintenance: 'warning'
  }
  return map[status] || 'info'
}

const getStatusText = (status: string) => {
  const map: Record<string, string> = {
    normal: '正常',
    fault: '故障',
    maintenance: '维护中'
  }
  return map[status] || status
}

const loadEquipments = async () => {
  loading.value = true
  try {
    const res = await getEquipments({
      page: pagination.page,
      pageSize: pagination.pageSize,
      departmentId: searchForm.departmentId,
      category: searchForm.category,
      status: searchForm.status
    })
    tableData.value = res.data.list
    pagination.total = res.data.total
  } finally {
    loading.value = false
  }
}

const loadEquipmentCategories = async () => {
  try {
    const res = await getEquipmentCategories()
    equipmentCategories.value = res.data
  } catch (error) {
    console.error('加载设备分类失败:', error)
  }
}

const loadDepartments = async () => {
  try {
    const res = await getDepartments()
    departments.value = res.data.list
  } catch (error) {
    console.error('加载科室失败:', error)
  }
}

const loadSuppliers = async () => {
  try {
    const res = await getSuppliers()
    suppliers.value = res.data.list
  } catch (error) {
    console.error('加载供应商失败:', error)
  }
}

const resetSearch = () => {
  searchForm.departmentId = undefined
  searchForm.category = ''
  searchForm.status = ''
  pagination.page = 1
  loadEquipments()
}

const handleSizeChange = (size: number) => {
  pagination.pageSize = size
  loadEquipments()
}

const handlePageChange = (page: number) => {
  pagination.page = page
  loadEquipments()
}

const handleAdd = () => {
  form.id = undefined
  dialogVisible.value = true
}

const handleEdit = (row: Equipment) => {
  Object.assign(form, row)
  dialogVisible.value = true
}

const handleView = async (row: Equipment) => {
  currentEquipment.value = row
  detailVisible.value = true
}

const handleDelete = (row: Equipment) => {
  ElMessageBox.confirm(`确定要删除设备【${row.name}】吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteEquipment(row.id)
      ElMessage.success('删除成功')
      loadEquipments()
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
        await updateEquipment(form.id, form)
        ElMessage.success('更新成功')
      } else {
        await createEquipment(form)
        ElMessage.success('创建成功')
      }
      dialogVisible.value = false
      loadEquipments()
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
    category: '',
    brand: '',
    model: '',
    serialNumber: '',
    departmentId: 0,
    supplierId: undefined,
    locationDesc: '',
    purchaseDate: '',
    purchasePrice: 0,
    maintenanceVendor: '',
    remark: ''
  })
}

onMounted(() => {
  loadEquipmentCategories()
  loadDepartments()
  loadSuppliers()
  loadEquipments()
})
</script>

<style scoped lang="scss">
.equipment-container {
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

  .filter-card {
    margin-bottom: 20px;
  }

  .pagination-container {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
  }

  .maintenance-section {
    margin-top: 20px;

    h4 {
      margin: 0 0 10px 0;
      font-size: 14px;
      font-weight: 500;
    }
  }
}
</style>
