<template>
  <div class="map-editor-container">
    <div class="page-header">
      <h2>地图编辑</h2>
      <div class="header-actions">
        <el-select v-model="selectedLocationId" placeholder="请选择空间" style="width: 200px" @change="loadMapElements">
          <el-option
            v-for="location in locations"
            :key="location.id"
            :label="location.name"
            :value="location.id"
          />
        </el-select>
        <el-button type="primary" @click="handleAddElement">
          <el-icon><Plus /></el-icon>
          添加元素
        </el-button>
      </div>
    </div>

    <div class="editor-content">
      <div class="map-canvas" ref="canvasRef">
        <div class="canvas-area">
          <div
            v-for="element in mapElements"
            :key="element.id"
            class="map-element"
            :class="getElementClass(element)"
            :style="getElementStyle(element)"
            @click="handleSelectElement(element)"
          >
            <el-icon v-if="element.elementType === 'equipment'" :size="20">
              <Tools />
            </el-icon>
            <el-icon v-else-if="element.elementType === 'patrol_point'" :size="20">
              <Location />
            </el-icon>
            <el-icon v-else-if="element.elementType === 'door'" :size="20">
              <Switch />
            </el-icon>
            <el-icon v-else-if="element.elementType === 'camera'" :size="20">
              <VideoCamera />
            </el-icon>
            <el-icon v-else :size="20">
              <MessageBox />
            </el-icon>
            <span class="element-label">{{ element.name }}</span>
          </div>
        </div>
      </div>

      <div class="properties-panel" v-if="selectedElement">
        <el-card>
          <template #header>
            <div class="panel-header">
              <span>属性编辑</span>
              <el-button type="danger" size="small" @click="handleDeleteElement">
                删除
              </el-button>
            </div>
          </template>

          <el-form :model="selectedElement" label-width="80px">
            <el-form-item label="元素名称">
              <el-input v-model="selectedElement.name" @blur="handleUpdateElement" />
            </el-form-item>
            <el-form-item label="元素类型">
              <el-select
                v-model="selectedElement.elementType"
                style="width: 100%"
                @change="handleUpdateElement"
              >
                <el-option
                  v-for="item in elementTypes"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="X坐标">
              <el-input-number
                v-model="selectedElement.coordsX"
                :min="0"
                :step="10"
                style="width: 100%"
                @change="handleUpdateElement"
              />
            </el-form-item>
            <el-form-item label="Y坐标">
              <el-input-number
                v-model="selectedElement.coordsY"
                :min="0"
                :step="10"
                style="width: 100%"
                @change="handleUpdateElement"
              />
            </el-form-item>
            <el-form-item label="宽度">
              <el-input-number
                v-model="selectedElement.width"
                :min="10"
                :step="10"
                style="width: 100%"
                @change="handleUpdateElement"
              />
            </el-form-item>
            <el-form-item label="高度">
              <el-input-number
                v-model="selectedElement.height"
                :min="10"
                :step="10"
                style="width: 100%"
                @change="handleUpdateElement"
              />
            </el-form-item>
            <el-form-item label="旋转角度">
              <el-input-number
                v-model="selectedElement.rotation"
                :min="0"
                :max="360"
                style="width: 100%"
                @change="handleUpdateElement"
              />
            </el-form-item>
            <el-form-item label="图标">
              <el-input v-model="selectedElement.icon" @blur="handleUpdateElement" />
            </el-form-item>
          </el-form>
        </el-card>
      </div>
    </div>

    <el-dialog
      v-model="dialogVisible"
      title="添加地图元素"
      width="500px"
    >
      <el-form
        ref="formRef"
        :model="form"
        label-width="100px"
        :rules="formRules"
      >
        <el-form-item label="元素名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入元素名称" />
        </el-form-item>
        <el-form-item label="元素类型" prop="elementType">
          <el-select
            v-model="form.elementType"
            placeholder="请选择元素类型"
            style="width: 100%"
          >
            <el-option
              v-for="item in elementTypes"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="X坐标">
          <el-input-number
            v-model="form.coordsX"
            :min="0"
            :step="10"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="Y坐标">
          <el-input-number
            v-model="form.coordsY"
            :min="0"
            :step="10"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="宽度">
          <el-input-number
            v-model="form.width"
            :min="10"
            :step="10"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="高度">
          <el-input-number
            v-model="form.height"
            :min="10"
            :step="10"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="图标">
          <el-input v-model="form.icon" placeholder="请输入图标名称" />
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
import { ref, reactive, onMounted, type FormInstance, type FormRules } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  Plus,
  Tools,
  Switch,
  VideoCamera,
  MessageBox
} from '@element-plus/icons-vue'
import {
  getSpaces,
  getMapElementsByLocation,
  createMapElement,
  updateMapElement,
  deleteMapElement,
  getElementTypes
} from '@/api/space'
import type { MapElement, CreateMapElementRequest } from '@/types'

const canvasRef = ref()
const formRef = ref<FormInstance>()
const dialogVisible = ref(false)
const locations = ref<MapElement[]>([])
const mapElements = ref<MapElement[]>([])
const selectedLocationId = ref<number | null>(null)
const selectedElement = ref<MapElement | null>(null)
const elementTypes = ref<{ value: string; label: string }[]>([])

const form = reactive<CreateMapElementRequest & { id?: number }>({
  locationId: 0,
  elementType: 'note',
  elementId: undefined,
  elementSubType: '',
  name: '',
  coordsX: 100,
  coordsY: 100,
  width: 60,
  height: 60,
  rotation: 0,
  icon: '',
  style: ''
})

const formRules: FormRules = {
  name: [{ required: true, message: '请输入元素名称', trigger: 'blur' }],
  elementType: [{ required: true, message: '请选择元素类型', trigger: 'change' }]
}

const getElementClass = (element: MapElement) => {
  return {
    'element-equipment': element.elementType === 'equipment',
    'element-patrol-point': element.elementType === 'patrol_point',
    'element-door': element.elementType === 'door',
    'element-camera': element.elementType === 'camera',
    'element-note': element.elementType === 'note',
    'is-selected': selectedElement.value?.id === element.id
  }
}

const getElementStyle = (element: MapElement) => {
  return {
    left: `${element.coordsX}px`,
    top: `${element.coordsY}px`,
    width: `${element.width}px`,
    height: `${element.height}px`,
    transform: `rotate(${element.rotation}deg)`
  }
}

const loadLocations = async () => {
  try {
    const res = await getSpaces()
    locations.value = res.data.list
  } catch (error) {
    console.error('加载空间列表失败:', error)
  }
}

const loadMapElements = async () => {
  if (!selectedLocationId.value) {
    mapElements.value = []
    return
  }

  try {
    const res = await getMapElementsByLocation(selectedLocationId.value)
    mapElements.value = res.data.list
  } catch (error) {
    console.error('加载地图元素失败:', error)
  }
}

const loadElementTypes = async () => {
  try {
    const res = await getElementTypes()
    elementTypes.value = res.data
  } catch (error) {
    console.error('加载元素类型失败:', error)
  }
}

const handleSelectElement = (element: MapElement) => {
  selectedElement.value = element
}

const handleAddElement = () => {
  if (!selectedLocationId.value) {
    ElMessage.warning('请先选择空间')
    return
  }

  form.locationId = selectedLocationId.value
  dialogVisible.value = true
}

const handleSubmit = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid) => {
    if (!valid) return

    try {
      await createMapElement(form)
      ElMessage.success('添加成功')
      dialogVisible.value = false
      loadMapElements()
      formRef.value.resetFields()
    } catch (error) {
      console.error('添加失败:', error)
    }
  })
}

const handleUpdateElement = async () => {
  if (!selectedElement.value) return

  try {
    await updateMapElement(selectedElement.value.id, selectedElement.value)
    ElMessage.success('更新成功')
  } catch (error) {
    console.error('更新失败:', error)
  }
}

const handleDeleteElement = () => {
  if (!selectedElement.value) return

  ElMessageBox.confirm(`确定要删除元素【${selectedElement.value.name}】吗？`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    try {
      await deleteMapElement(selectedElement.value.id)
      ElMessage.success('删除成功')
      selectedElement.value = null
      loadMapElements()
    } catch (error) {
      console.error('删除失败:', error)
    }
  })
}

onMounted(() => {
  loadLocations()
  loadElementTypes()
})
</script>

<style scoped lang="scss">
.map-editor-container {
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

    .header-actions {
      display: flex;
      align-items: center;
      gap: 10px;
    }
  }

  .editor-content {
    display: flex;
    gap: 20px;

    .map-canvas {
      flex: 1;
      border: 2px dashed #dcdfe6;
      border-radius: 4px;
      min-height: 600px;
      position: relative;
      background: linear-gradient(90deg, rgba(0, 0, 0, 0.05) 1px, transparent 1px),
        linear-gradient(rgba(0, 0, 0, 0.05) 1px, transparent 1px);
      background-size: 20px 20px;
    }

    .canvas-area {
      position: relative;
      width: 100%;
      height: 100%;
    }

    .map-element {
      position: absolute;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      border-radius: 4px;
      cursor: pointer;
      transition: all 0.2s;
      z-index: 1;

      &.is-selected {
        outline: 2px solid #409eff;
        outline-offset: 2px;
        z-index: 10;
      }

      &.element-equipment {
        background-color: rgba(103, 194, 58, 0.2);
        border: 1px solid #67c23a;
        color: #67c23a;
      }

      &.element-patrol-point {
        background-color: rgba(64, 158, 255, 0.2);
        border: 1px solid #409eff;
        color: #409eff;
      }

      &.element-door {
        background-color: rgba(245, 108, 108, 0.2);
        border: 1px solid #f56c6c;
        color: #f56c6c;
      }

      &.element-camera {
        background-color: rgba(230, 162, 60, 0.2);
        border: 1px solid #e6a23c;
        color: #e6a23c;
      }

      &.element-note {
        background-color: rgba(144, 147, 153, 0.2);
        border: 1px solid #909399;
        color: #909399;
      }

      .element-label {
        font-size: 12px;
        margin-top: 4px;
        white-space: nowrap;
      }
    }

    .properties-panel {
      width: 300px;

      .panel-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
    }
  }
}
</style>
