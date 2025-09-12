import request from '@/utils/request';

/**
 * Paginated query organization list
 */
export async function pageOrganizations(params) {
  const res = await request.get('/system/organization/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query organization list
 */
export async function listOrganizations(params) {
  const res = await request.get('/system/organization', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add organization
 */
export async function addOrganization(data) {
  const res = await request.post('/system/organization', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update organization
 */
export async function updateOrganization(data) {
  const res = await request.put('/system/organization', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete organization
 */
export async function removeOrganization(id) {
  const res = await request.delete('/system/organization/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
