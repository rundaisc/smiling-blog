import request from '@/utils/request'

export function saveNav(data) {
  return request({
    url: '/api/backend/nav/save',
    method: 'post',
    data
  })
}


export function getList(params) {
  return request({
    url: '/api/backend/nav/list',
    method: 'get',
    params
  })
}
export function deleteNav(data) {
  return request({
    url: '/api/backend/nav/'+data.id,
    method: 'delete',
  })
}
