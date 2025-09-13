import request from '@/utils/request';

/**
 * Paginated query login record list
 */
export async function pageLoginRecords(params) {
  const res = await request.get('/system/login-record/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query login record list
 */
export async function listLoginRecords(params) {
  const res = await request.get('/system/login-record', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}
