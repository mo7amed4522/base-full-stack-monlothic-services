import request from '@/utils/request';

/**
 * Paginated query dictionary data list
 */
export async function pageDictionaryData(params) {
  const res = await request.get('/system/dictionary-data/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query dictionary data list
 */
export async function listDictionaryData(params) {
  const res = await request.get('/system/dictionary-data', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add dictionary data
 */
export async function addDictionaryData(data) {
  const res = await request.post('/system/dictionary-data', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update dictionary data
 */
export async function updateDictionaryData(data) {
  const res = await request.put('/system/dictionary-data', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete dictionary data
 */
export async function removeDictionaryData(id) {
  const res = await request.delete('/system/dictionary-data/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Batch delete dictionary data
 */
export async function removeDictionaryDataBatch(data) {
  const res = await request.delete('/system/dictionary-data/batch', { data });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
