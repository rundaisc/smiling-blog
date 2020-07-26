import request from '@/utils/request'

export function createCategory(data) {
  return request({
    url: '/api/backend/category/create',
    method: 'post',
    data
  })
}
export function updateCategory(data) {
  return request({
    url: '/api/backend/category/update',
    method: 'put',
    data
  })
}

export function getList(params) {
  return request({
    url: '/api/backend/category',
    method: 'get',
    params
  })
}
export function deleteCategory(data) {
  return request({
    url: '/api/backend/category',
    method: 'delete',
    data
  })
}
