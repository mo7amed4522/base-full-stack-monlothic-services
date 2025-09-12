import request from '@/utils/request';

// Order List
export async function getOrderPayListApi(params) {
  const res = await request.get('/orderPay/getOrderPayList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order operation log list
export async function getOrderPayOperLogListApi(params) {
  const res = await request.get('/orderPayOperLog/getOrderPayOperLogList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - player information
export async function getOrderPayPlayerListApi(params) {
  const res = await request.get('/orderPayPlayer/getOrderPayPlayerList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - request packet
export async function getOrderPayRequestListApi(params) {
  const res = await request.get('/orderPayRequest/getOrderPayRequestList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - reversal record
export async function getOrderPayReversalListApi(params) {
  const res = await request.get('/orderPayReversal/getOrderPayReversalList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - user information
export async function getOrderPayUserListApi(params) {
  const res = await request.get('/orderPayUser/getOrderPayUserList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout
export async function getOrderPayoutListApi(params) {
  const res = await request.get('/orderPayout/getOrderPayoutList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout operation log
export async function getOrderPayoutOperLogListApi(params) {
  const res = await request.get('/orderPayoutOperLog/getOrderPayoutOperLogList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout player information
export async function getOrderPayoutPlayerListApi(params) {
  const res = await request.get('/orderPayoutPlayer/getOrderPayoutPlayerList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout request packet
export async function getOrderPayoutRequestListApi(params) {
  const res = await request.get('/orderPayoutRequest/getOrderPayoutRequestList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout user information
export async function getOrderPayoutUserListApi(params) {
  const res = await request.get('/orderPayoutUser/getOrderPayoutUserList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout - pagination list
export async function getPayoutListApi(params) {
  const res = await request.get('/order/payoutList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Manually push and update third-party interfaces
export async function payoutManualApi(data) {
  const res = await request.post('/order/payout/manual', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout - channel and payment pool dropdown list
export async function getConfigAndPayPoolListApi(params) {
  const res = await request.get('/orderPayReversal/getConfigAndPayPoolList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Get all the following payment method card numbers based on the payment pool ID cPayPoolId
export async function getPayPoolAccountListApi(params) {
  const res = await request.get('/orderPayReversal/GetPayPoolAccountList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Close Order
export async function closeStatusApi(data) {
  const res = await request.put('/orderPayReversal/closeStatus', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout - utr compensation
export async function createOrderPayReversalApi(data) {
  const res = await request.post('/orderPayReversal/createOrderPayReversal', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Withdrawal review
export async function payoutOrderResetApi(data) {
  const res = await request.put('/orderPayout/orderReset', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment order
export async function replenishmentPayoutOrderApi(data) {
  const res = await request.post('/orderPayout/replenishment', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list reset request
export async function orderPayResetApi(data) {
  const res = await request.put('/orderPay/orderPayReset', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list - payout - push order data
export async function pushOrderPayCallbackApi(data) {
  const res = await request.put('/orderPay/pushOrderPayCallback', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Withdrawal re-push order data
export async function pushOrderPayoutCallbackApi(data) {
  const res = await request.put('/orderPayout/pushOrderPayoutCallback', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Simulated order interface
export async function mockPayOrderApi(data) {
  const res = await request.post('/open/api/order/pay', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Simulated withdrawal order interface
export async function mockPayoutOrderApi(data) {
  const res = await request.post('/open/api/order/payout', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Account change record list
export async function getTransactionsListApi(params) {
  const res = await request.get('/transactions/list', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Lock/Unlock Order
export async function orderPayoutLockApi(data) {
  const res = await request.post('/orderPayout/lock', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Get withdrawal account by type
export async function getListByTypeApi(params) {
  const res = await request.get('/cPayPoolAccount/getListByType', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Order list statistics
export async function getPayOrderStaticsApi(params) {
  const res = await request.get('/orderPay/getOrderStatics', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment order statistics
export async function getPayoutOrderStaticsApi(params) {
  const res = await request.get('/orderPayout/getPayoutOrderStatics', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Modification Notes
export async function updateOrderPayoutApi(data) {
  const res = await request.put('/orderPayout/updateOrderPayout', data);
  if (res.data.code === 0) {
    return res.data.msg;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payout order compensation record list
export async function getOrderPayoutLogListApi(params) {
  const res = await request.get('/orderPayout/getOrderPayoutLogList', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Simulated order placement operation version
export async function mockPayOrderTestApi(data) {
  const res = await request.post('/orderPay/payTest', data);
  if (res.data.code === 0) {
    return res.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function getConfigAndPayPoolListByApi(params) {
  const res = await request.get('/orderPayReversal/getConfigAndPayPoolListByApi', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

export async function getTransactionsCardListApi(params) {
  const res = await request.get('/transactions/card/list', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
