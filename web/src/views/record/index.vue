<template>
  <div class="page-container">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>巡检记录</span>
        </div>
      </template>
      
      <el-table :data="tableData" border stripe style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="taskName" label="任务名称" />
        <el-table-column prop="pointName" label="巡检点位" />
        <el-table-column prop="inspectorName" label="巡检员" />
        <el-table-column prop="status" label="状态">
          <template #default="{ row }">
            <el-tag v-if="row.status === 'normal'" type="success">正常</el-tag>
            <el-tag v-else-if="row.status === 'abnormal'" type="danger">异常</el-tag>
            <el-tag v-else>{{ row.status }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="checkMethod" label="打卡方式">
          <template #default="{ row }">
            <el-tag v-if="row.checkMethod === 'qr_code'" type="primary">扫码</el-tag>
            <el-tag v-else-if="row.checkMethod === 'nfc'" type="success">NFC</el-tag>
            <el-tag v-else>{{ row.checkMethod }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="checkTime" label="打卡时间" width="180" />
        <el-table-column label="操作" width="120">
          <template #default="{ row }">
            <el-button size="small" @click="viewDetail(row)">查看详情</el-button>
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

    <el-dialog v-model="detailVisible" title="巡检详情" width="600px">
      <el-descriptions :column="2" border v-if="currentRecord">
        <el-descriptions-item label="任务名称">{{ currentRecord.taskName }}</el-descriptions-item>
        <el-descriptions-item label="巡检点位">{{ currentRecord.pointName }}</el-descriptions-item>
        <el-descriptions-item label="巡检员">{{ currentRecord.inspectorName }}</el-descriptions-item>
        <el-descriptions-item label="打卡时间">{{ currentRecord.checkTime }}</el-descriptions-item>
        <el-descriptions-item label="打卡方式">{{ currentRecord.checkMethod }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="currentRecord.status === 'normal'" type="success">正常</el-tag>
          <el-tag v-else type="danger">异常</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="检查结果" :span="2">{{ currentRecord.checkResult || '-' }}</el-descriptions-item>
        <el-descriptions-item label="异常描述" :span="2">{{ currentRecord.abnormalDesc || '-' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import request from '@/utils/request'

const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const detailVisible = ref(false)
const currentRecord = ref()

const loadData = async () => {
  try {
    const res = await request.get('/records', { params: { page: page.value, pageSize: pageSize.value } })
    tableData.value = res.data.list
    total.value = res.data.total
  } catch (e) {
    console.error('加载记录列表失败', e)
  }
}

const viewDetail = async (row: any) => {
  try {
    const res = await request.get(`/records/${row.id}`)
    currentRecord.value = res.data
    detailVisible.value = true
  } catch (e) {
    console.error('加载详情失败', e)
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
