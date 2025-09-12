import { h } from 'vue';
import ProFormContent from './components/pro-form-content.vue';

/**
 * Built-in component types
 */
export const defaultTypes = [
  'input',
  'textarea',
  'select',
  'multipleSelect',
  'radio',
  'radioButton',
  'checkbox',
  'checkboxButton',
  'date',
  'datetime',
  'daterange',
  'datetimerange',
  'time',
  'timerange',
  'timeSelect',
  'editTag',
  'switch',
  'rate',
  'inputNumber',
  'cascader',
  'treeSelect',
  'treeMultipleSelect',
  'virtualTreeSelect',
  'virtualTreeMultipleSelect',
  'tableSelect',
  'tableMultipleSelect',
  'checkCard',
  'multipleCheckCard',
  'autocomplete',
  'imageUpload',
  'fileUpload',
  'regions',
  'dictRadio',
  'dictCheckbox',
  'dictSelect',
  'dictMultipleSelect',
  'slider',
  'sliderRange',
  'editor',
  'text',
  // Display Type
  'label',
  'divider',
  'button',
  'steps',
  // Container Type
  'card',
  'tabs',
  'table',
  'collapse',
  'collapseItem',
  'row',
  'col',
  'div'
];

/**
 * Display type components
 */
export const viewTypes = ['label', 'divider', 'button', 'steps'];

/**
 * Container type components
 */
export const containerTypes = [
  'card',
  'tabs',
  'table',
  'collapse',
  'collapseItem',
  'row',
  'col',
  'div'
];

/**
 * Select the component type
 */
export const selectTypes = [
  'select',
  'multipleSelect',
  'radio',
  'radioButton',
  'checkbox',
  'checkboxButton',
  'date',
  'datetime',
  'daterange',
  'datetimerange',
  'time',
  'timerange',
  'timeSelect',
  'switch',
  'cascader',
  'rate',
  'slider',
  'sliderRange',
  'treeSelect',
  'tableMultipleSelect',
  'virtualTreeSelect',
  'virtualTreeMultipleSelect',
  'tableSelect',
  'tableMultipleSelect',
  'checkCard',
  'multipleCheckCard',
  'dictRadio',
  'dictSelect',
  'dictCheckbox',
  'dictMultipleSelect',
  'regions'
];

/**
 * Upload type component type
 */
export const uploadTypes = ['imageUpload', 'fileUpload'];

/**
 * Component type that supports options configuration drop-down option data
 */
export const optionsTypes = [
  'select',
  'multipleSelect',
  'radio',
  'radioButton',
  'checkbox',
  'checkboxButton',
  'autocomplete',
  'cascader',
  'treeSelect',
  'treeMultipleSelect',
  'checkCard',
  'multipleCheckCard'
];

/**
 * Component type that triggers form validation rules using blur
 */
export const blurTypes = ['input', 'textarea'];

/**
 * Component type whose form data type is string
 */
export const stringTypes = [
  'input',
  'textarea',
  'date',
  'datetime',
  'time',
  'timeSelect',
  'autocomplete',
  'editor',
  'text'
];

/**
 * Component type whose form data type is numeric
 */
export const numberTypes = [
  'select',
  'radio',
  'radioButton',
  'checkbox',
  'checkboxButton',
  'switch',
  'inputNumber',
  'rate',
  'slider',
  'tableSelect'
];

/**
 * Component type whose form data type is an array
 */
export const arrayTypes = [
  'multipleSelect',
  'checkbox',
  'checkboxButton',
  'daterange',
  'datetimerange',
  'timerange',
  'cascader',
  'sliderRange',
  'treeMultipleSelect',
  'virtualTreeMultipleSelect',
  'tableMultipleSelect',
  'multipleCheckCard',
  'editTag',
  'dictCheckbox',
  'dictMultipleSelect',
  'regions'
];

/**
 * Get the validation rule value type
 * @param type Component Type
 */
export function getRuleType(type) {
  if (type) {
    if (arrayTypes.includes(type)) {
      return 'array';
    }
    if (numberTypes.includes(type)) {
      return 'number';
    }
  }
  return 'string';
}

/**
 * Get the validation rule trigger type
 * @param type Component Type
 */
export function getRuleTrigger(type) {
  return type && blurTypes.includes(type) ? 'blur' : 'change';
}

/**
 * Get the validation rule prompt text
 * @param type Component Type
 * @param label Form item title
 */
export function getRuleMessage(type, label) {
  const text = label ?? '';
  if (type) {
    if (selectTypes.includes(type)) {
      return `Please select${text}`;
    }
    if (uploadTypes.includes(type)) {
      return `Please upload${text}`;
    }
  }
  return `Please enter${text}`;
}

/**
 * Determine whether the form item is displayed
 * @param item Form Items
 * @param form Form Data
 * @param items Form item data
 * @param searchExpand  Search form expanded state
 */
export function isShowItem(item, form, items, searchExpand) {
  if (!item.prop) {
    return false;
  }
  if (item.vIf != null) {
    if (typeof item.vIf === 'function') {
      return item.vIf(form);
    }
    if (typeof item.vIf === 'string' && item.vIf.trim().length) {
      try {
        return new Function(
          'form',
          'items',
          'searchExpand',
          `return (${item.vIf})`
        )(form, items, searchExpand);
      } catch (e) {
        console.error(e);
        return false;
      }
    }
    if (item.vIf === false) {
      return false;
    }
  }
  return true;
}

/**
 * Subcomponent renderer when customizing container components
 * @param props property
 * @param context Context
 */
export function ChildrenRender(props, { emit }) {
  return h(
    ProFormContent,
    {
      model: props.model,
      rules: props.rules,
      items: props.item.children ?? [],
      grid: props.item.grid,
      rowProps: props.item.rowProps,
      formItems: props.formItems,
      searchExpand: props.searchExpand,
      onUpdateItemValue: (prop, value) => {
        emit('updateItemValue', prop, value);
      }
    },
    props.slots
  );
}
