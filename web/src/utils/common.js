import { removeToken } from '@/utils/token-util';

/**
 * Log out
 * @param route Whether to use routing jump
 * @param from Login jump address
 * @param push Routing jump method
 */
export function logout(route, from, push) {
  removeToken();
  // 删除
  localStorage.removeItem('userInfo');
  if (route && push) {
    push({
      path: '/login',
      query: from ? { from: encodeURIComponent(from) } : void 0
    });
    return;
  }
  //  This jump avoids logging in again and registering dynamic routes repeatedly. The hash routing mode uses location.reload();
  const BASE_URL = import.meta.env.BASE_URL;
  const url = BASE_URL + 'login';
  location.replace(from ? `${url}?from=${encodeURIComponent(from)}` : url);
}

/**
 * Download file
 * @param data Binary data
 * @param name File name
 * @param type File type
 */
export function download(data, name, type) {
  const blob = new Blob([data], { type: type || 'application/octet-stream' });
  const url = window.URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  a.download = name;
  a.style.display = 'none';
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
  URL.revokeObjectURL(url);
}

/**
 * Convert parameters to URL strings
 * @param params Parameters
 * @param url Address to append parameters
 */
export function toURLSearch(params, url) {
  if (typeof params !== 'object' || params == null) {
    return '';
  }
  const result = transformParams(params)
    .map((d) => `${encodeURIComponent(d[0])}=${encodeURIComponent(d[1])}`)
    .join('&');
  if (!url) {
    return result;
  }
  return (url.includes('?') ? `${url}&` : `${url}?`) + result;
}

/**
 * Convert parameters to form data
 * @param params Parameters
 */
export function toFormData(params) {
  const formData = new FormData();
  if (typeof params !== 'object' || params == null) {
    return formData;
  }
  transformParams(params).forEach((d) => {
    formData.append(d[0], d[1]);
  });
  return formData;
}

/**
 * Handle array and object type parameters for get requests
 * @param params Parameters
 */
export function transformParams(params) {
  const result = [];
  if (params != null && typeof params === 'object') {
    Object.keys(params).forEach((key) => {
      const value = params[key];
      if (value != null && value !== '') {
        if (typeof value === 'object' && !isBlobFile(value)) {
          getObjectParamsArray(value).forEach((item) => {
            result.push([`${key}${item[0]}`, item[1]]);
          });
        } else {
          result.push([key, value]);
        }
      }
    });
  }
  return result;
}

/**
 * Convert object to parameter array
 * @param obj Object
 */
export function getObjectParamsArray(obj) {
  const result = [];
  Object.keys(obj).forEach((key) => {
    const value = obj[key];
    if (value != null && value !== '') {
      const name = `[${key}]`;
      if (typeof value === 'object' && !isBlobFile(value)) {
        getObjectParamsArray(value).forEach((item) => {
          result.push([`${name}${item[0]}`, item[1]]);
        });
      } else {
        result.push([name, value]);
      }
    }
  });
  return result;
}

/**
 * Determine if it is a file
 * @param obj Object
 */
export function isBlobFile(obj) {
  return obj != null && (obj instanceof Blob || obj instanceof File);
}

/**
 * Generate a 32-bit random string
 */
export function getRandom32() {
  // Generate a 32-byte hexadecimal string
  const array = new Uint8Array(16);
  let arr = window.crypto.getRandomValues(array);
  return Array.from(arr)
    .map((b) => b.toString(16).padStart(2, '0'))
    .join('');
}

export function getBeijingHour() {
  const date = new Date();
  const options = { timeZone: 'Asia/Shanghai', hour: '2-digit', hour12: false };
  const dateTimeFormat = new Intl.DateTimeFormat('en-US', options);
  return dateTimeFormat.format(date); // Returns the current Beijing time (24-hour system)
}
