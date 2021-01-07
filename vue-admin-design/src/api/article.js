import request from '../request'

export function articleCreate(data) {
  return request({
    url: '/admin/article/create',
    method: 'post',
    data
  })
}

export function articleList(data) {
  return request({
    url: '/admin/article/list',
    method: 'post',
    data
  })
}

export function getArticle(data) {
  return request({
    url: '/admin/article/one',
    method: 'post',
    data
  })
}

