import md5 from 'js-md5';
import { ElMessage } from 'element-plus';

/**
 * Processing time format
 * @param {*} isoString
 * @returns
 */
export const _timeFormat = (isoString) => {
  // Use regular expressions to remove the time zone information and get the date and time part
  const match = isoString.match(/^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})/);

  if (match) {
    // Replace 'T' with ' ' to format the output
    return match[1].replace('T', ' ');
  }

  return null; // If the match fails, return null.
};

// Generate sign
export const generateSign = (formData, signKey) => {
  // Filter out the sign field because sign is the last one to be generated
  const filteredData = { ...formData };
  delete filteredData.sign;

  // Sort the parameter names in ASCII dictionary order
  const sortedKeys = Object.keys(filteredData).sort();

  // Concatenate in the format of key=value&key=value
  const stringA = sortedKeys.map((key) => `${key}=${filteredData[key]}`).join('&');

  // Concatenate signKey
  const stringSignTemp = `${stringA}&signKey=${signKey}`;
  return md5(stringSignTemp);
  // Calculate MD5 and return
  // return md5(stringSignTemp).toUpperCase() // Calculate MD5 and return
};

export const formatToTwoDecimals = (value) => {
  const num = parseFloat(value);
  return isNaN(num) ? '0.00' : num.toFixed(2);
};

// Your copyText function
export const copyText = async (text, t) => {
  // Check if the page is in HTTPS environment or file protocol
  if (window.location.protocol !== 'https:' && window.location.protocol !== 'file:') {
    // Using execCommand as a fallback
    try {
      const textArea = document.createElement('textarea');
      textArea.value = text.replace(/\n/g, '\r\n'); // Replace line breaks (for different systems)
      document.body.appendChild(textArea);
      textArea.select(); // Select text
      const success = document.execCommand('copy'); // Use execCommand for copying
      document.body.removeChild(textArea);

      if (success) {
        ElMessage.success(t('common.copySuccess')); // Call t to get internationalized text
      } else {
        console.error('execCommand Copy failed');
        ElMessage.error(t('common.copyFail'));
      }
    } catch (error) {
      console.error('execCommand Copy failed:', error);
      ElMessage.error(t('common.copyFail'));
    }
  } else {
    // Use Clipboard API for copying (requires HTTPS environment)
    try {
      await navigator.clipboard.writeText(text);
      ElMessage.success(t('common.copySuccess'));
    } catch (error) {
      console.error('Clipboard API Copy failed:', error);
      ElMessage.error(t('common.copyFail'));
    }
  }
};

// Returns hours, minutes, and seconds
export const formatTime = (seconds) => {
  // If it is a negative number, return 0
  if (seconds < 0) return 0;

  // Calculate hours, minutes, and seconds
  const hours = Math.floor(seconds / 3600); // Calculate hours
  const minutes = Math.floor((seconds % 3600) / 60); // Calculate minutes
  const secs = seconds % 60; // Calculate seconds

  // Return time in HH:MM:SS format, ensuring each part is two digits
  return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
};

export const maskMiddle8 = (input) => {
  if (typeof input !== 'string' || input.length !== 12) {
    console.warn('Input must be a 12-digit string');
    return input;
  }

  const prefix = input.slice(0, 2);
  const suffix = input.slice(-2);
  const masked = '*'.repeat(8);

  return `${prefix}${masked}${suffix}`;
};
