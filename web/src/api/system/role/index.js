import request from '@/utils/request';

/**
 * Paginated query role list
 */
export async function pageRoles(params) {
  const res = await request.get('/system/role/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query role list
 */
export async function listRoles(params) {
  const res = await request.get('/system/role', {
    params
  });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add a role
 */
export async function addRole(data) {
  const res = await request.post('/system/role', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update role
 */
export async function updateRole(data) {
  const res = await request.put('/system/role', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete role
 */
export async function removeRole(id) {
  const res = await request.delete('/system/role/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Batch delete role
 */
export async function removeRoles(data) {
  const res = await request.delete('/system/role/batch', {
    data
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get role assigned menu list
 */
export async function listRoleMenus(roleId) {
  const res = await request.get('/system/role-menu/' + roleId);
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update role menu
 */
export async function updateRoleMenus(roleId, data) {
  const res = await request.put('/system/role-menu/' + roleId, data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
