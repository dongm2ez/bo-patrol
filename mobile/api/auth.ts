import { post } from '../utils/request'

export interface LoginRequest {
  username: string
  password: string
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

export const logout = () => {
  uni.removeStorageSync('token')
  uni.removeStorageSync('userInfo')
}

export const getToken = () => {
  return uni.getStorageSync('token')
}

export const getUserInfo = (): UserInfo | null => {
  const userInfo = uni.getStorageSync('userInfo')
  return userInfo ? JSON.parse(userInfo) : null
}

export const setUserInfo = (user: UserInfo) => {
  uni.setStorageSync('userInfo', JSON.stringify(user))
}
