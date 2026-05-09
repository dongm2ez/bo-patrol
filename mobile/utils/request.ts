const BASE_URL = '/api/v1'

interface RequestConfig {
  url: string
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
  data?: any
  header?: any
}

interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

function request<T = any>(config: RequestConfig): Promise<ApiResponse<T>> {
  const token = uni.getStorageSync('token')
  
  return new Promise((resolve, reject) => {
    uni.request({
      url: BASE_URL + config.url,
      method: config.method || 'GET',
      data: config.data,
      header: {
        'Content-Type': 'application/json',
        'Authorization': token ? `Bearer ${token}` : '',
        ...config.header
      },
      success: (res: any) => {
        const data = res.data as ApiResponse<T>
        
        if (data.code === 200) {
          resolve(data)
        } else if (data.code === 401) {
          uni.removeStorageSync('token')
          uni.removeStorageSync('userInfo')
          
          uni.showToast({
            title: '登录已过期，请重新登录',
            icon: 'none'
          })
          
          setTimeout(() => {
            uni.reLaunch({ url: '/pages/login/login' })
          }, 1500)
          
          reject(new Error(data.message || '未授权'))
        } else {
          uni.showToast({
            title: data.message || '请求失败',
            icon: 'none'
          })
          reject(new Error(data.message || '请求失败'))
        }
      },
      fail: (err) => {
        console.error('请求失败:', err)
        uni.showToast({
          title: '网络连接失败',
          icon: 'none'
        })
        reject(err)
      }
    })
  })
}

export function get<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'GET', data })
}

export function post<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'POST', data })
}

export function put<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'PUT', data })
}

export function del<T = any>(url: string, data?: any): Promise<ApiResponse<T>> {
  return request<T>({ url, method: 'DELETE', data })
}

export default request
