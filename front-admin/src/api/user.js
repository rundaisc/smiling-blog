import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/backend/login',
    method: 'post',
    data
  })
}

