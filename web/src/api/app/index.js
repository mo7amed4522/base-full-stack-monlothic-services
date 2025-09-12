import request from '@/utils/request';

// app list
export async function getAppPageListApi(data) {
  const res = await request.post('/appManage/getAppPage', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add app configuration
export async function createAppApi(data) {
  const res = await request.post('/appManage/createApp', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Update app configuration
export async function updateAppApi(data) {
  const res = await request.put('/appManage/updateApp', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete app configuration
export async function deleteAppApi(params) {
  const res = await request.delete('/appManage/deleteAppById', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Domain configuration list
export async function getCDominManageListApi(params) {
  const res = await request.get('/cDominManage/getCDominManageList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add domain configuration
export async function createCDominManageApi(data) {
  const res = await request.post('/cDominManage/createCDominManage', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Update domain configuration status
export async function visibleCDominManageApi(data) {
  const res = await request.put('/cDominManage/visibleCDominManage', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add subdomain configuration
export async function createCSubDominManageApi(data) {
  const res = await request.post('/cSubDominManage/createCSubDominManage', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Create a template
export async function createCPageTemplateApi(data) {
  const res = await request.post('/cPageTemplate/createCPageTemplate', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Update a template
export async function updateCPageTemplateApi(data) {
  const res = await request.put('/cPageTemplate/updateCPageTemplate', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete a template
export async function deleteCPageTemplateApi(params) {
  const res = await request.delete('/cPageTemplate/deleteCPageTemplate', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Template list
export async function getCPageTemplateListApi(params) {
  const res = await request.get('/cPageTemplate/getCPageTemplateList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add promotion link
export async function createCPromotionLinkApi(data) {
  const res = await request.post('/cPromotionLink/createCPromotionLink', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Update promotion link
export async function updateCPromotionLinkApi(data) {
  const res = await request.put('/cPromotionLink/updateCPromotionLink', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete promotion link
export async function deleteCPromotionLinkApi(params) {
  const res = await request.delete('/cPromotionLink/deleteCPromotionLink', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Promotion link list
export async function getCPromotionLinkListApi(params) {
  const res = await request.get('/cPromotionLink/getCPromotionLinkList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
