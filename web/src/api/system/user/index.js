import request from '@/utils/request';

/**
 * Paginated query user list
 */
export async function pageUsers(params) {
  const res = await request.get('/system/user/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query user list
 */
export async function listUsers(params) {
  const res = await request.get('/system/user', {
    params
  });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query user by id
 */
export async function getUser(id) {
  const res = await request.get('/system/user/' + id);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add user
 */
export async function addUser(data) {
  const res = await request.post('/system/user', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update user
 */
export async function updateUser(data) {
  const res = await request.put('/system/user', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete user
 */
export async function removeUser(id) {
  const res = await request.delete('/system/user/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Batch delete user
 */
export async function removeUsers(data) {
  const res = await request.delete('/system/user/batch', {
    data
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update user status
 */
export async function updateUserStatus(userId, status) {
  const res = await request.put('/system/user/status', {
    userId,
    status
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Reset user password
 */
export async function updateUserPassword(userId, password = '123456') {
  const res = await request.put('/system/user/password', {
    userId,
    password
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Import Users
 */
export async function importUsers(file) {
  const formData = new FormData();
  formData.append('file', file);
  const res = await request.post('/system/user/import', formData);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Check if the user exists
 */
export async function checkExistence(field, value, id) {
  const res = await request.get('/system/user/existence', {
    params: { field, value, id }
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
