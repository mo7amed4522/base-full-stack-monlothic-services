import request from '@/utils/request';

// Merchant List
export async function getMUserList(params) {
  const res = await request.get('/mUser/getMUserList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New Merchants
export async function addMUserApi(data) {
  const res = await request.post('/mUser/createMUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify Merchant
export async function updateMUserApi(data) {
  const res = await request.put('/mUser/updateMUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete Merchant
export async function deleteMUserApi(params) {
  const res = await request.delete('/mUser/deleteMUser', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant Label List
export async function getMUserLabelListApi(params) {
  const res = await request.get('/mUserLabel/getMUserLabelList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New Merchant Label
export async function createMUserLabelApi(data) {
  const res = await request.post('/mUserLabel/createMUserLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify Merchant Label
export async function updateMUserLabelApi(data) {
  const res = await request.put('/mUserLabel/updateMUserLabel', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete Merchant Label
export async function deleteMUserLabelApi(params) {
  const res = await request.delete('/mUserLabel/deleteMUserLabel', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant Wallet List
export async function getMUserWalletListApi(params) {
  const res = await request.get('/mUserWallet/getMUserWalletList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New Merchant Wallet
export async function createMUserWalletApi(data) {
  const res = await request.post('/mUserWallet/createMUserWallet', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify Merchant Wallet
export async function updateMUserWalletApi(data) {
  const res = await request.put('/mUserWallet/updateMUserWallet', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete Merchant Wallet
export async function deleteMUserWalletApi(params) {
  const res = await request.delete('/mUserWallet/deleteMUserWallet', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment method list
export async function getCPayPoolAccountListApi(params) {
  const res = await request.get('/cPayPoolAccount/getCPayPoolAccountList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New payment method
export async function createCPayPoolAccountApi(data) {
  const res = await request.post('/cPayPoolAccount/createCPayPoolAccount', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify payment method
export async function updateCPayPoolAccountApi(data) {
  const res = await request.put('/cPayPoolAccount/updateCPayPoolAccount', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete payment method
export async function deleteCPayPoolAccountApi(params) {
  const res = await request.delete('/cPayPoolAccount/deleteCPayPoolAccount', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant payment pool list
export async function getCPayPoolListApi(params) {
  const res = await request.get('/cPayPool/getCPayPoolList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New merchant payment pool
export async function createCPayPoolApi(data) {
  const res = await request.post('/cPayPool/createCPayPool', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify merchant payment pool
export async function updateCPayPoolApi(data) {
  const res = await request.put('/cPayPool/updateCPayPool', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete merchant payment pool
export async function deleteCPayPoolApi(params) {
  const res = await request.delete('/cPayPool/deleteCPayPool', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant payment pool configuration list
export async function getCPayPoolUserListApi(params) {
  const res = await request.get('/cPayPoolUser/getCPayPoolUserList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Added merchant payment pool configuration
export async function createCPayPoolUserApi(data) {
  const res = await request.post('/cPayPoolUser/createCPayPoolUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify merchant payment pool configuration
export async function updateCPayPoolUserApi(data) {
  const res = await request.put('/cPayPoolUser/updateCPayPoolUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete merchant payment pool configuration
export async function deleteCPayPoolUserApi(params) {
  const res = await request.delete('/cPayPoolUser/deleteCPayPoolUser', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant payment pool channel configuration list
export async function getMPayChannelConfigListApi(params) {
  const res = await request.get('/mPayChannelConfig/getMPayChannelConfigList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New merchant payment pool channel configuration
export async function createMPayChannelConfigApi(data) {
  const res = await request.post('/mPayChannelConfig/createMPayChannelConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify merchant payment pool channel configuration
export async function updateMPayChannelConfigApi(data) {
  const res = await request.put('/mPayChannelConfig/updateMPayChannelConfig', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete merchant payment pool channel configuration
export async function deleteMPayChannelConfigApi(params) {
  const res = await request.delete('/mPayChannelConfig/deleteMPayChannelConfig', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant player list
export async function getOrderPayUserListApi(params) {
  const res = await request.get('/orderPayUser/getOrderPayUserList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// New merchant player
export async function createOrderPayUserApi(data) {
  const res = await request.post('/orderPayUser/createOrderPayUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify merchant player
export async function updateOrderPayUserApi(data) {
  const res = await request.put('/orderPayUser/updateOrderPayUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete merchant player
export async function deleteOrderPayUserApi(params) {
  const res = await request.delete('/orderPayUser/deleteOrderPayUser', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Update merchant status
export async function updateMUserStatusApi(data) {
  const res = await request.put('/mUser/updateMUserStatus', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Configure the gift method
export async function configMUserApi(data) {
  const res = await request.post('/mUser/configMUser', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Get current login merchant information
export async function currentMUserApi(params) {
  const res = await request.get('/mUser/currentMUser', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Get gift method configuration
export async function getConfigMUserApi(params) {
  const res = await request.get('/mUser/getConfigMUser', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Account change operation
export async function updateAmountApi(data) {
  const res = await request.post('/cPayPoolAccount/updateAmount', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Transfer payment pool
export async function transferCPayPoolApi(data) {
  const res = await request.put('/cPayPoolAccount/transferCPayPool', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Receiving/withdrawing account list
export async function getCPayPoolAccountListByDelList(params) {
  const res = await request.get('/cPayPoolAccount/getCPayPoolAccountListByDel', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Recover deleted accounts
export async function recoverCPayPoolAccountApi(params) {
  const res = await request.delete('/cPayPoolAccount/recoverCPayPoolAccount', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Receiving account list
export async function getCPayPoolCardsListApi(params) {
  const res = await request.get('/cPayPoolCards/getCPayPoolCardsList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Delete receiving account
export async function deleteCPayPoolCardsApi(params) {
  const res = await request.delete('cPayPoolCards/deleteCPayPoolCards', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Create a receiving account
export async function createCPayPoolCardsApi(data) {
  const res = await request.post('/cPayPoolCards/createCPayPoolCards', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modify receiving account
export async function updateCPayPoolCardsApi(data) {
  const res = await request.put('/cPayPoolCards/updateCPayPoolCards', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Receiving account - deleted list
export async function getCPayPoolCardsByDelApi(params) {
  const res = await request.get('/cPayPoolCards/getCPayPoolCardsByDel', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Recover receiving account
export async function recoverCPayPoolCardsApi(params) {
  const res = await request.delete('/cPayPoolCards/recoverCPayPoolCards', { params });
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
