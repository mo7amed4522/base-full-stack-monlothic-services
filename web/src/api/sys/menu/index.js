import request from '@/utils/request';

/**
 * Pagination query menu
 */
export async function pageMenus(params) {
  const res = await request.get('/system/menu/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get menu list
 */
export async function listMenus(data) {
  const res = await request.post('/menu/getMenuList', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Get menu list
 */
export async function getMenuList(data) {
  const res = await request.post('/menu/getMenuList', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Add base menu
 */
export async function addBaseMenu(data) {
  const res = await request.post('/menu/addBaseMenu', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Get base menu tree
 */
export async function getBaseMenuTree(data) {
  const res = await request.post('/menu/getBaseMenuTree', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Add user menu association relationship
 */
export async function addMenuAuthority(data) {
  const res = await request.post('/menu/addMenuAuthority', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Get user menu association relationship
 */
export async function getMenuAuthority(data) {
  const res = await request.post('/menu/getMenuAuthority', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Update base menu
 */
export async function updateBaseMenu(data) {
  const res = await request.post('/menu/updateBaseMenu', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Get base menu by id
 */
export async function getBaseMenuById(data) {
  const res = await request.post('/menu/getBaseMenuById', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// @Summary Delete menu
// @Produce  application/json
// @Param ID float64
// @Router /menu/deleteBaseMenu [post]
export async function deleteBaseMenu(data) {
  const res = await request.post('/menu/deleteBaseMenu', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Can remove authority btn
 */
export async function canRemoveAuthorityBtnApi(data) {
  const res = await request.post('/authorityBtn/canRemoveAuthorityBtn', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Add menu
 */
export async function addMenu(data) {
  const res = await request.post('/system/menu', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update menu
 */
export async function updateMenu(data) {
  const res = await request.put('/system/menu', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete menu
 */
export async function removeMenu(id) {
  const res = await request.delete('/system/menu/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get all APIs without paging
 */
export async function getAllApis(data) {
  const res = await request.post('/api/getAllApis', data);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
