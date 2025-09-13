import request from '@/utils/request';

/**
 * Upload files
 * @param file document
 */
export async function uploadFile(file, config) {
  const formData = new FormData();
  formData.append('file', file);
  const res = await request.post('/fileUploadAndDownload/upload', formData, config);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Upload base64 file
 * @param base64 base64 data
 * @param fileName file name
 */
export async function uploadBase64File(base64, fileName) {
  const formData = new FormData();
  formData.append('base64', base64);
  if (fileName) {
    formData.append('fileName', fileName);
  }
  const res = await request.post('/file/upload/base64', formData);
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Paginated query file upload record list
 */
export async function pageFiles(params) {
  const res = await request.get('/file/page', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Delete file
 */
export async function removeFile(id) {
  const res = await request.delete('/file/remove/' + id);
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Batch delete file
 */
export async function removeFiles(data) {
  const res = await request.delete('/file/remove/batch', {
    data
  });
  if (res.data.code === 0) {
    return res.data.message;
  }
  return Promise.reject(new Error(res.data.message));
}
