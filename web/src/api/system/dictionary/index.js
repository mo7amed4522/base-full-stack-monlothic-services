import request from '@/utils/request';

/**
 * Paginated query dictionary list
 */
export async function pageDictionaries(params) {
  const res = await request.get('/system/dictionary/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Query dictionary list
 */
export async function listDictionaries(params) {
  const res = await request.get('/system/dictionary', {
    params
  });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Add dictionary
 */
export async function addDictionary(data) {
  const res = await request.post('/system/dictionary', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Update dictionary
 */
export async function updateDictionary(data) {
  const res = await request.put('/system/dictionary', data);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete dictionary
 */
export async function removeDictionary(id) {
  const res = await request.delete('/system/dictionary/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
