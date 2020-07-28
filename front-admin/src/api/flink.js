import request from '@/utils/request'

export function saveFlink(data) {
  return request({
    url: '/api/backend/flink/save',
    method: 'post',
    data
  })
}


export function getList(params) {
  return request({
    url: '/api/backend/flink/list',
    method: 'get',
    params
  })
}
export function deleteFlink(data) {
  return request({
    url: '/api/backend/flink/'+data.id,
    method: 'delete',
  })
}
