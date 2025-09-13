import request from '@/utils/request';

/**
 * Get payment data
 */
export async function getPayNumList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/analysis-pay-num.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get sales data
 */
export async function getSaleroomList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/analysis-saleroom.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get recent 1 hour visit data
 * @returns {Promise<Object>}
 */
export async function getVisitHourList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/analysis-visits.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get word cloud data
 */
export async function getWordCloudList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/analysis-hot-search.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}
