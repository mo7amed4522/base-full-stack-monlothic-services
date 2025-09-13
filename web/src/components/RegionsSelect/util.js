import { defineStore, storeToRefs } from 'pinia';
import { EleMessage } from 'ele-admin-plus/es';
import request from '@/utils/request';
const BASE_URL = import.meta.env.BASE_URL;

/**
 * Customize provincial, municipal and district data status management
 */
export const useRegionsStore = defineStore('regions', {
  state: () => ({
    regionsData: null
  }),
  actions: {
    setRegionsData(data) {
      this.regionsData = data;
    }
  }
});

/**
 * Customize provincial, municipal and district data hook
 */
export function useRegionsData() {
  const regionsStore = useRegionsStore();
  const { regionsData } = storeToRefs(regionsStore);
  // If it has not been cached yet, get the province, city and district data
  if (regionsData.value == null) {
    regionsStore.setRegionsData([]);
    getRegionsData()
      .then((data) => {
        regionsStore.setRegionsData(data);
      })
      .catch((e) => {
        EleMessage.error(e.message);
      });
  }
  return regionsData;
}

/**
 * Customize provincial, municipal and district data
 */
export async function getRegionsData() {
  const res = await request.get(BASE_URL + 'json/regions-data.json', {
    baseURL: ''
  });
  return res.data ?? [];
}

/**
 * Province, city, district data value processing
 * @param data Provincial, municipal and district data
 * @param field value field name
 */
export function formatData(data, field) {
  if (field === 'label') {
    return data.map((d) => {
      const item = { label: d.label, value: d.label };
      if (d.children) {
        item.children = d.children.map((c) => {
          const cItem = { label: c.label, value: c.label };
          if (c.children) {
            cItem.children = c.children.map((cc) => {
              return { label: cc.label, value: cc.label };
            });
          }
          return cItem;
        });
      }
      return item;
    });
  }
  return data;
}

/**
 * Provincial, municipal and district data screening
 * @param data Provincial, municipal and district data
 * @param type Only provincial and municipal data or only provincial data
 * @param field value field name
 */
export function filterData(data, type, field) {
  if (type === 'provinceCity') {
    return formatData(
      data.map((d) => {
        const item = { label: d.label, value: d.value };
        if (d.children) {
          item.children = d.children.map((c) => {
            return { label: c.label, value: c.value };
          });
        }
        return item;
      }),
      field
    );
  } else if (type === 'province') {
    return formatData(
      data.map((d) => ({ label: d.label, value: d.value })),
      field
    );
  }
  return formatData(data, field);
}

/**
 * Customize provincial, municipal and district data selected value corresponding text
 * @param values Selected value
 * @param data Provincial, municipal and district data
 */
export function getValueLabel(values, data) {
  const labels = [];
  if (values && values.length && data) {
    data.forEach((d) => {
      if (d.value === values[0]) {
        labels.push(d.label);
        if (d.children && d.children.length) {
          d.children.forEach((c) => {
            if (c.value === values[1]) {
              labels.push(c.label);
            }
            if (c.children && c.children.length) {
              c.children.forEach((cc) => {
                if (cc.value === values[2]) {
                  labels.push(cc.label);
                }
              });
            }
          });
        }
      }
    });
  }
  return labels;
}
