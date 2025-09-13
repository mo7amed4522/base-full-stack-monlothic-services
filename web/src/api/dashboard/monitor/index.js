import request from '@/utils/request';
const BASE_URL = import.meta.env.BASE_URL;

/**
 * Get China map geo data
 */
export async function getChinaMapData() {
  const res = await request.get(BASE_URL + 'json/china-provinces.geo.json', {
    baseURL: ''
  });
  if (res.data) {
    return res.data;
  }
  return Promise.reject(new Error('Get map data failed'));
}

/**
 * Get user distribution data
 */
export async function getUserCountList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/monitor-user-count.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}

/**
 * Get user browser distribution data
 */
export async function getBrowserCountList() {
  const res = await request.get('https://cdn.eleadmin.com/20200610/monitor-browser-count.json');
  if (res.data.code === 0 && res.data.data) {
    return res.data.data;
  }
  return Promise.reject(new Error(res.data.message));
}
