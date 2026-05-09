const BASE_URL = '/api/v1'

function request(options) {
  const token = uni.getStorageSync('token')
  
  return new Promise((resolve, reject) => {
    uni.request({
      url: BASE_URL + options.url,
      method: options.method || 'GET',
      data: options.data || {},
      header: {
        'Content-Type': 'application/json',
        'Authorization': token ? `Bearer ${token}` : ''
      },
      success: (res) => {
        if (res.data.code === 200) {
          resolve(res.data)
        } else if (res.data.code === 401) {
          uni.removeStorageSync('token')
          uni.reLaunch({
            url: '/pages/login/login'
          })
          reject(res.data)
        } else {
          uni.showToast({
            title: res.data.message || '请求失败',
            icon: 'none'
          })
          reject(res.data)
        }
      },
      fail: (err) => {
        uni.showToast({
          title: '网络错误',
          icon: 'none'
        })
        reject(err)
      }
    })
  })
}

export default request
