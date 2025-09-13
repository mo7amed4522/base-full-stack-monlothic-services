import { storeToRefs } from 'pinia';
import { useUserStore } from '@/store/modules/user';

/**
 * Permission judgment hook
 */
export function usePermission() {
  const userStore = useUserStore();
  const { authorities, roles } = storeToRefs(userStore);

  /**
   * Whether there are certain permissions
   * @param value Permission character or character array
   */
  const hasPermission = (value) => {
    // If a special identifier is needed to represent all permissions, add the following
    /* if (authorities.value && authorities.value.includes('*:*:*')) {
      return true;
    } */
    return arrayHas(authorities.value, value);
  };

  /**
   * Whether there are any permissions
   * @param value Permission character or character array
   */
  const hasAnyPermission = (value) => {
    // If a special identifier is needed to represent all permissions, add the following
    /* if (authorities.value && authorities.value.includes('*:*:*')) {
      return true;
    } */
    return arrayHasAny(authorities.value, value);
  };

  /**
   * Whether there are certain roles
   * @param value Role character or character array
   */
  const hasRole = (value) => {
    return arrayHas(roles.value, value);
  };

  /**
   * Whether there are any roles
   * @param value Role character or character array
   */
  const hasAnyRole = (value) => {
    return arrayHasAny(roles.value, value);
  };

  return { hasPermission, hasAnyPermission, hasRole, hasAnyRole };
}

/**
 * Whether the array has certain values
 * @param array Array
 * @param value Single value or array
 */
function arrayHas(array, value) {
  if (!value) {
    return true;
  }
  if (!array) {
    return false;
  }
  if (Array.isArray(value)) {
    return value.every((v) => array.includes(v));
  }
  return array.includes(value);
}

/**
 * Whether the array has any values
 * @param array Array
 * @param value Single value or array
 */
function arrayHasAny(array, value) {
  if (!value) {
    return true;
  }
  if (!array) {
    return false;
  }
  if (Array.isArray(value)) {
    return array.some((d) => d && value.includes(d));
  }
  return array.includes(value);
}
