import request from '@/utils/request'

export function saveArticle(data) {
  return request({
    url: '/api/backend/article/save',
    method: 'post',
    data
  })
}
export function updateArticle(data) {
  return request({
    url: '/api/backend/article/save',
    method: 'put',
    data
  })
}

export function getList(params) {
  return request({
    url: '/api/backend/article',
    method: 'get',
    params
  })
}
export function deleteArticle(data) {
  return request({
    url: '/api/backend/article',
    method: 'delete',
    data
  })
}

export function uploadImage(data) {
  return request({
    url: '/api/backend/upload/image',
    method: 'post',
    header:{
     "Content-Type": "multipart/form-data"
    },
    data
  })
}
export function shoeArticle(params) {
  return request({
    url: '/api/backend/article/detail',
    method: 'get',
    params
  })
}
