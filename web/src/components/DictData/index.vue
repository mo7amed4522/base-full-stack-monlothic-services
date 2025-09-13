<!-- Dictionary component -->
<template>
  <template v-if="type === 'text'">
    <span v-for="item in valueData" :key="item.dictDataCode">
      {{ item.dictDataName }}
    </span>
  </template>
  <template v-else-if="type === 'tag'">
    <el-tag v-for="item in valueData" :key="item.dictDataCode" :disable-transitions="true" size="small">
      {{ item.dictDataName }}
    </el-tag>
  </template>
  <el-radio-group v-else-if="type === 'radio'" :disabled="disabled" v-model="model">
    <el-radio v-for="item in data" :key="item.dictDataCode" :value="item.dictDataCode" :label="item.dictDataName" />
  </el-radio-group>
  <el-checkbox-group v-else-if="type === 'checkbox'" :disabled="disabled" v-model="model">
    <el-checkbox v-for="item in data" :key="item.dictDataCode" :value="item.dictDataCode" :label="item.dictDataName" />
  </el-checkbox-group>
  <el-select v-else v-model="model" :clearable="true" :disabled="disabled" :placeholder="placeholder"
    :multiple="type === 'multipleSelect'" :teleported="teleported" :filterable="filterable" class="ele-fluid">
    <el-option v-for="item in data" :key="item.dictDataCode" :value="item.dictDataCode" :label="item.dictDataName" />
  </el-select>
</template>

<script setup>
import { computed } from 'vue';
import { useDictData } from '@/utils/use-dict-data';

defineOptions({ name: 'DictData' });

const props = defineProps({
  /** Dictionary type */
  code: String,
  /** Component type */
  type: String,
  /** Whether to disable */
  disabled: Boolean,
  /** Prompt text */
  placeholder: String,
  /** Whether the dropdown of select is inserted into the body */
  teleported: {
    type: Boolean,
    default: true
  },
  /** Whether select is filterable */
  filterable: Boolean
});

/** Dictionary value */
const model = defineModel({
  type: [String, Number, Boolean, Array]
});

/** Dictionary data */
const [data] = useDictData([props.code]);

/** Dictionary data corresponding to the binding value */
const valueData = computed(() => {
  const result = [];
  const val = model.value;
  if (val == null || val === '') {
    return result;
  }
  const values = Array.isArray(val) ? val : [val];
  values.forEach((v) => {
    const temp = data.value.find((d) => d.dictDataCode == v);
    if (temp != null) {
      result.push(temp);
    } else {
      result.push({ dictDataCode: v, dictDataName: v });
    }
  });
  return result;
});
</script>
