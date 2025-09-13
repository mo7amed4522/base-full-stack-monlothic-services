import request from '@/utils/request';

/**
 * Paginated query user file list
 */
export async function pageUserFiles(params) {
  const res = await request.get('/system/user-file/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query user file list
 */
export async function listUserFiles(params) {
  const res = await request.get('/system/user-file', {
    params
  });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add user file
 */
export async function addUserFile(data) {
  const res = await request.post('/system/user-file', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update user file
 */
export async function updateUserFile(data) {
  const res = await request.put('/system/user-file', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete user file
 */
export async function removeUserFile(id) {
  const res = await request.delete('/system/user-file/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Batch delete user files
 */
export async function removeUserFiles(data) {
  const res = await request.delete('/system/user-file/batch', { data });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
