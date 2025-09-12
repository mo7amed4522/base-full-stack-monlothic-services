import request from '@/utils/request';

// Payment pool daily statistics list
export async function getCPayPoolAccountBillDailyListApi(params) {
  const res = await request.get('/cPayPoolAccountBillDaily/getCPayPoolAccountBillDailyList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment pool hourly statistics list
export async function getCPayPoolAccountBillHourListApi(params) {
  const res = await request.get('/cPayPoolAccountBillHour/getCPayPoolAccountBillHourList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment pool daily statistics list
export async function getCPayPoolBillDailyListApi(params) {
  const res = await request.get('/cPayPoolBillDaily/getCPayPoolBillDailyList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Merchant daily statistics list
export async function getMBillPayDailyListApi(params) {
  const res = await request.get('/mBillPayDaily/getMBillPayDailyList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Operator order statistics
export async function orderPayReversalLogListApi(params) {
  const res = await request.get('/orderPayReversal/logList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Operator withdrawal statistics
export async function orderPayoutLogApi(params) {
  const res = await request.get('/orderPayout/logList', { params });
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment pool daily statistics - statistics
export async function getCPayPoolAccountBillDailyStatisticsApi(params) {
  const res = await request.get('/cPayPoolAccountBillDaily/getCPayPoolAccountBillDailyStatistics', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}

// Payment pool hourly statistics - statistics
export async function getCPayPoolAccountBillHourStatisticsApi(params) {
  const res = await request.get('/cPayPoolAccountBillHour/getCPayPoolAccountBillHourStatistics', { params });
  if (res.data.code === 0) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.msg));
}
