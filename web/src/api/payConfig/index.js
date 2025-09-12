import request from '@/utils/request';

// Payment configuration label list
export async function getPayLabelListApi(params) {
  const res = await request.get('/payLabel/getPayLabelList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add payment configuration label list
export async function createPayLabelApi(data) {
  const res = await request.post('/payLabel/createPayLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify payment configuration label list
export async function updatePayLabelApi(data) {
  const res = await request.put('/payLabel/updatePayLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete payment configuration label list
export async function deletePayLabel(params) {
  const res = await request.delete('/payLabel/deletePayLabel', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment channel list
export async function getCPayConfigListApi(params) {
  const res = await request.get('/cPayConfig/getCPayConfigList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add payment channel
export async function createCPayConfigApi(data) {
  const res = await request.post('/cPayConfig/createCPayConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify payment channel
export async function updateCPayConfigApi(data) {
  const res = await request.put('/cPayConfig/updateCPayConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete payment channel
export async function deleteCPayConfigApi(params) {
  const res = await request.delete('/cPayConfig/deleteCPayConfig', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment bank dictionary list
export async function getPayBankListApi(params) {
  const res = await request.get('/payBank/getPayBankList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add payment bank dictionary
export async function createPayBankApi(data) {
  const res = await request.post('/payBank/createPayBank', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify payment bank dictionary
export async function updatePayBankApi(data) {
  const res = await request.put('/payBank/updatePayBank', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete payment bank dictionary
export async function deletePayBankApi(params) {
  const res = await request.delete('/payBank/deletePayBank', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Add payment pool label
export async function createCPayPoolLabelApi(data) {
  const res = await request.post('/cPayPoolLabel/createCPayPoolLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete payment pool label
export async function deleteCPayPoolLabelApi(params) {
  const res = await request.delete('/cPayPoolLabel/deleteCPayPoolLabel', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify payment pool label
export async function updateCPayPoolLabelApi(data) {
  const res = await request.put('/cPayPoolLabel/updateCPayPoolLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment pool label list
export async function getCPayPoolLabelListApi(params) {
  const res = await request.get('/cPayPoolLabel/getCPayPoolLabelList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
