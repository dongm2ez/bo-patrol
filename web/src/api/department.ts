import { get, post, put, del } from '@/utils/request'
import type { Department, DepartmentTree, CreateDepartmentRequest, OptionItem } from '@/types'

export const getDepartments = (params?: { type?: string }) => {
  return get<{ list: Department[], total: number }>('/departments', params)
}

export const getDepartmentTree = () => {
  return get<DepartmentTree[]>('/departments/tree')
}

export const getDepartment = (id: number) => {
  return get<Department>(`/departments/${id}`)
}

export const createDepartment = (data: CreateDepartmentRequest) => {
  return post<Department>('/departments', data)
}

export const updateDepartment = (id: number, data: CreateDepartmentRequest) => {
  return put<Department>(`/departments/${id}`, data)
}

export const deleteDepartment = (id: number) => {
  return del(`/departments/${id}`)
}

export const getDepartmentTypes = () => {
  return get<OptionItem[]>('/public/department-types')
}
