import request from '@/utils/request';

/**
 * Paginated query operation record list
 */
export async function pageOperationRecords(params) {
  const res = await request.get('/system/operation-record/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query operation record list
 */
export async function listOperationRecords(params) {
  const res = await request.get('/system/operation-record', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}
