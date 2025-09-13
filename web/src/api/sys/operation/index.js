import request from '@/utils/request';

// Paging query operation record list
export async function getSysOperationRecordListApi(params) {
  const res = await request.get('/sysOperationRecord/getSysOperationRecordList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
