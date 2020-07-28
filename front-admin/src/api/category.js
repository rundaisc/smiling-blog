import request from '@/utils/request'
import da from "element-ui/src/locale/lang/da";

export function createCategory(data) {
  return request({
    url: '/api/backend/category/save',
    method: 'post',
    data
  })
}
export function updateCategory(data) {
  return request({
    url: '/api/backend/category/save',
    method: 'post',
    data
  })
}

export function getList(params) {
  return request({
    url: '/api/backend/category/list',
    method: 'get',
    params
  })
}
export function deleteCategory(data) {
  return request({
    url: '/api/backend/category/'+data.id,
    method: 'delete',
    data
  })
}
