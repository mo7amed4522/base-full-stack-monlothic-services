import request from '@/utils/request';

// Automatic Receipt List
export async function getReceiptAccountListApi(params) {
  const res = await request.get('/receiptAccount/getReceiptAccountList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify automatic receipts
export async function updateReceiptAccountApi(data) {
  const res = await request.put('/receiptAccount/updateReceiptAccount', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete automatic receipts
export async function deleteReceiptAccountApi(params) {
  const res = await request.delete('/receiptAccount/deleteReceiptAccount', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add automatic receipts
export async function createReceiptAccountApi(data) {
  const res = await request.post('/receiptAccount/createReceiptAccount', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Receipt configuration list
export async function getReceiptConfigListApi(params) {
  const res = await request.get('/receiptConfig/getReceiptConfigList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify receipt configuration
export async function updateReceiptConfigApi(data) {
  const res = await request.put('/receiptConfig/updateReceiptConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete receipt configuration
export async function deleteReceiptConfigApi(params) {
  const res = await request.delete('/receiptConfig/deleteReceiptConfig', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add receipt configuration
export async function createReceiptConfigApi(data) {
  const res = await request.post('/receiptConfig/createReceiptConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Receipt order record list
export async function getReceiptOrderReqListApi(params) {
  const res = await request.get('/receiptOrderReq/getReceiptOrderReqList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add receipt order record
export async function createReceiptOrderReqApi(data) {
  const res = await request.post('/receiptOrderReq/createReceiptOrderReq', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Proxy Configuration
// Proxy Configuration List Page
export async function getReceiptProxyListApi(params) {
  const res = await request.get('/receiptProxy/getReceiptProxyList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add proxy configuration
export async function createReceiptProxyApi(data) {
  const res = await request.post('/receiptProxy/createReceiptProxy', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete proxy configuration
export async function deleteReceiptProxyApi(params) {
  const res = await request.delete('/receiptProxy/deleteReceiptProxy', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify proxy configuration
export async function updateReceiptProxyApi(data) {
  const res = await request.put('/receiptProxy/updateReceiptProxy', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Proxy Configuration Log List
export async function getReceiptLogsListApi(params) {
  const res = await request.get('/receiptLogs/getReceiptLogsList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// View agent details
export async function findReceiptProxyApi(params) {
  const res = await request.get('/receiptProxy/findReceiptProxy', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
