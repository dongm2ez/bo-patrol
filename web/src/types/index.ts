export interface PageRequest {
  page: number
  pageSize: number
}

export interface PageResponse<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

export interface OptionItem {
  value: string | number
  label: string
}

export interface Department {
  id: number
  name: string
  code: string
  type: string
  description: string
  parentId?: number
  parentName?: string
  sortOrder: number
  managerId?: number
  managerName?: string
  phone: string
  email: string
  status: number
  createdAt: string
}

export interface DepartmentTree extends Department {
  children?: DepartmentTree[]
}

export interface CreateDepartmentRequest {
  name: string
  code: string
  type: string
  description?: string
  parentId?: number
  sortOrder?: number
  managerId?: number
  phone?: string
  email?: string
}

export interface Equipment {
  id: number
  name: string
  code: string
  category: string
  categoryName: string
  brand: string
  model: string
  serialNumber: string
  departmentId: number
  departmentName: string
  locationId?: number
  locationName?: string
  locationDesc: string
  purchaseDate?: string
  purchasePrice: number
  supplierId?: number
  supplierName?: string
  maintenanceVendor: string
  maintenanceStart?: string
  maintenanceEnd?: string
  lastMaintenance?: string
  nextMaintenance?: string
  status: string
  statusName: string
  specs?: string
  remark?: string
  createdAt: string
}

export interface CreateEquipmentRequest {
  name: string
  code: string
  category: string
  brand?: string
  model?: string
  serialNumber?: string
  departmentId: number
  locationId?: number
  locationDesc?: string
  purchaseDate?: string
  purchasePrice?: number
  supplierId?: number
  maintenanceVendor?: string
  maintenanceStart?: string
  maintenanceEnd?: string
  lastMaintenance?: string
  nextMaintenance?: string
  specs?: string
  remark?: string
}

export interface SpaceLocation {
  id: number
  name: string
  code: string
  type: string
  typeName: string
  parentId?: number
  parentName?: string
  galleryId?: number
  galleryName?: string
  floor: number
  description?: string
  mapData?: string
  coordsX: number
  coordsY: number
  sortOrder: number
  status: number
  createdAt: string
}

export interface SpaceTree extends SpaceLocation {
  children?: SpaceTree[]
}

export interface CreateSpaceRequest {
  name: string
  code: string
  type: string
  parentId?: number
  galleryId?: number
  floor?: number
  description?: string
  mapData?: string
  coordsX?: number
  coordsY?: number
  sortOrder?: number
}

export interface MapElement {
  id: number
  locationId: number
  locationName: string
  elementType: string
  elementId?: number
  elementSubType?: string
  name: string
  coordsX: number
  coordsY: number
  width: number
  height: number
  rotation: number
  style?: string
  icon?: string
  sortOrder: number
  status: number
  equipmentInfo?: Equipment
  pointInfo?: any
}

export interface CreateMapElementRequest {
  locationId: number
  elementType: string
  elementId?: number
  elementSubType?: string
  name?: string
  coordsX?: number
  coordsY?: number
  width?: number
  height?: number
  rotation?: number
  style?: string
  icon?: string
}

export interface Supplier {
  id: number
  name: string
  code: string
  contact: string
  phone: string
  email: string
  address: string
  type: string
  status: number
  createdAt: string
}

export interface MaintenanceRecord {
  id: number
  equipmentId: number
  type: string
  date: string
  content: string
  cost: number
  operator: string
  remark: string
  createdAt: string
}
