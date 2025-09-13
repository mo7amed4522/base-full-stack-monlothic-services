import request from '@/utils/request';

/**
 * Get the personal information/menu/authority/role of the currently logged-in user
 */
export async function getUserInfo() {
  const res = await request.post('/menu/getMenu');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Change Password
export async function changePasswordApi(data) {
  const res = await request.post('/user/changePassword', data);
  if (res.data.code === 0) {
    return res.data.msg ?? 'Modification successful';
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Modify the password of the currently logged-in user
 */
export async function updatePassword(data) {
  const res = await request.put('/auth/password', data);
  if (res.data.code === 0) {
    return res.data.msg ?? 'Modification successful';
  }
  return Promise.reject(new Error(res.data.msg));
}

/**
 * Modify the personal information of the currently logged-in user
 */
export async function updateUserInfo(data) {
  const res = await request.put('/auth/user', data);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
