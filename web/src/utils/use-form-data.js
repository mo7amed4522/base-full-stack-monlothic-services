import { reactive } from 'vue';
import { cloneDeep } from 'lodash-es';

/**
 * Form data hook
 * @param initValue Initial value
 */
export function useFormData(init) {
  let initValue = init == null ? {} : cloneDeep(init);
  /** Form data */
  const form = reactive(init == null ? {} : { ...init });

  /** Reset to initial value */
  const resetFields = (field, excludes) => {
    const keys = Object.keys(form);
    if (typeof field === 'string' && field) {
      if (keys.includes(field)) {
        form[field] = cloneDeep(initValue[field]);
      }
      return;
    }
    keys.forEach((key) => {
      if (!excludes || !excludes.includes(key)) {
        form[key] = cloneDeep(initValue[key]);
      }
    });
  };

  /** Assign values without changing fields */
  const assignFields = (data, excludes) => {
    if (excludes === true) {
      initValue = data == null ? {} : cloneDeep(data);
    }
    Object.keys(form).forEach((key) => {
      if (!excludes || excludes === true || !excludes.includes(key)) {
        form[key] = data?.[key];
      }
    });
  };

  /** Assign value to a field */
  const setFieldValue = (field, value) => {
    form[field] = value;
  };

  const result = [form, resetFields, assignFields, setFieldValue];
  // Support object deconstruction for compatibility with older versions
  Object.assign(result, {
    form,
    resetFields,
    assignFields
  });
  return result;
}
