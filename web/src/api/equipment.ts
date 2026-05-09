import { get, post, put, del } from '@/utils/request'
import type { Equipment, CreateEquipmentRequest, PageResponse, PageRequest, OptionItem, Supplier, MaintenanceRecord } from '@/types'

export const getEquipments = (params: PageRequest & { departmentId?: number; category?: string; status?: string }) => {
  return get<PageResponse<Equipment>>('/equipments', params)
}

export const getEquipment = (id: number) => {
  return get<Equipment>(`/equipments/${id}`)
}

export const getEquipmentByDepartment = (departmentId: number) => {
  return get<{ list: Equipment[], total: number }>(`/equipments/department/${departmentId}`)
}

export const createEquipment = (data: CreateEquipmentRequest) => {
  return post<Equipment>('/equipments', data)
}

export const updateEquipment = (id: number, data: CreateEquipmentRequest) => {
  return put<Equipment>(`/equipments/${id}`, data)
}

export const deleteEquipment = (id: number) => {
  return del(`/equipments/${id}`)
}

export const getEquipmentCategories = () => {
  return get<OptionItem[]>('/public/equipment-categories')
}

export const getMaintenanceRecords = (equipmentId: number) => {
  return get<{ list: MaintenanceRecord[], total: number }>(`/equipments/${equipmentId}/maintenance`)
}

export const getSuppliers = () => {
  return get<{ list: Supplier[], total: number }>('/suppliers')
}

export const createSupplier = (data: Partial<Supplier>) => {
  return post<Supplier>('/suppliers', data)
}
