import request from '@/utils/request'

export function siteInfo() {
  return request({
    url: '/api/backend/site/info',
    method: 'get',
  })
}


export function saveSite(data) {
  return request({
    url: '/api/backend/site/save',
    method: 'post',
    data
  })
}

