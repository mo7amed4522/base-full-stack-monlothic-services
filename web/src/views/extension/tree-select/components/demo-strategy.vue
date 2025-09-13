<template>
  <ele-card header="Checked Strategy">
    <div style="max-width: 240px">
      <ele-tree-select ref="treeSelectRef" clearable multiple placeholder="Select" v-model="selectedValue"
        :popper-width="240" :tree-props="treeProps" :showCheckedStrategy="showCheckedStrategy"
        :checkedValueStrategy="checkedValueStrategy" :popper-options="{ strategy: 'fixed' }" />
    </div>
    <option-item label="Checked Strategy" :label-style="{ paddingTop: '6px' }"
      style="margin-top: 10px; align-items: flex-start">
      <el-radio-group v-model="showCheckedStrategy" style="flex-direction: column; align-items: flex-start">
        <el-radio value="parent" label="parent(When all child nodes are selected, only the parent is displayed)"
          style="margin: 0; display: flex; width: 100%" />
        <el-radio value="child" label="child(Only display selected child nodes)"
          style="margin: 0; display: flex; width: 100%" />
        <el-radio value="all" label="all(Display all selected nodes)" style="margin: 0; display: flex; width: 100%" />
      </el-radio-group>
    </option-item>
    <option-item label="Checked Value Strategy" :responsive="false" style="margin-top: 10px">
      <el-radio-group v-model="checkedValueStrategy">
        <el-radio :value="false" label="All selected values" />
        <el-radio :value="true" label="Same label strategy" />
      </el-radio-group>
    </option-item>
    <option-item label="Checked Values" style="margin-top: 10px">
      {{ JSON.stringify(selectedValue) }}
    </option-item>
  </ele-card>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** 树选择组件 */
const treeSelectRef = ref(null);

/** 选中值 */
const selectedValue = ref([2, 3, 4]);

/** 树属性 */
const treeProps = reactive({
  height: 266,
  data: []
});

/** 标签显示策略 */
const showCheckedStrategy = ref('parent');

/** 值策略 */
const checkedValueStrategy = ref(false);

onMounted(() => {
  setTimeout(() => {
    const keys = [];
    let count = 0;
    const data = Array.from({ length: 4 }).map((_, i) => {
      count++;
      keys.push(count);
      return {
        id: count,
        label: `${String(count).padStart(5, '0')}-node-${i + 1}`,
        disabled: false,
        children: Array.from({ length: 3 }).map((_, j) => {
          count++;
          keys.push(count);
          return {
            id: count,
            label: `${String(count).padStart(5, '0')}-node-${i + 1}-${j + 1}`,
            disabled: false,
            children: Array.from({ length: 2 }).map((_, k) => {
              count++;
              const idStr = String(count).padStart(5, '0');
              return {
                id: count,
                label: `${idStr}-node-${i + 1}-${j + 1}-${k + 1}`,
                disabled: false
              };
            })
          };
        })
      };
    });
    treeProps.data = data;
    if (treeSelectRef.value && treeSelectRef.value.treeRef) {
      treeSelectRef.value.treeRef.setExpandedKeys(keys);
    }
  }, 100);
});
</script>
