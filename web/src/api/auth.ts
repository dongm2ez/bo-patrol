import { post } from '@/utils/request'

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  password: string
  name: string
  phone?: string
  role?: string
}

export interface UserInfo {
  id: number
  username: string
  name: string
  phone: string
  email: string
  role: string
  department: string
  avatar: string
  status: number
}

export interface LoginResponse {
  token: string
  user: UserInfo
}

export const login = (data: LoginRequest) => {
  return post<LoginResponse>('/auth/login', data)
}

export const register = (data: RegisterRequest) => {
  return post<UserInfo>('/auth/register', data)
}

export const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('userInfo')
}

export const getToken = () => {
  return localStorage.getItem('token')
}

export const getUserInfo = (): UserInfo | null => {
  const userInfo = localStorage.getItem('userInfo')
  return userInfo ? JSON.parse(userInfo) : null
}

export const setUserInfo = (user: UserInfo) => {
  localStorage.setItem('userInfo', JSON.stringify(user))
}
