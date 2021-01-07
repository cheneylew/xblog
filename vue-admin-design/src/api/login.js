import request from '../request'

export function login(data) {
  return request({
    url: '/admin/user/dologin',
    method: 'post',
    data
  })
}

export function getUserInfo(params) {
  return request({
    url: '/user/info',
    method: 'get',
    params
  })
}
