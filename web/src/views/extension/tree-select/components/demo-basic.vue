<template>
  <ele-card header="Basic Usage">
    <div>
      The tree select component is based on the ElTreeV2 virtual scroll tree component encapsulation.
      It has no pressure for a large amount of data.
    </div>
    <div style="max-width: 240px; margin-top: 12px">
      <ele-tree-select ref="treeSelectRef" clearable placeholder="Select" v-model="selectedValue" :popper-width="240"
        :tree-props="treeProps" :disabled="disabled" :popper-options="{ strategy: 'fixed' }" />
    </div>
    <option-item label="Disable" :responsive="false" style="margin-top: 16px">
      <el-radio-group v-model="disabled">
        <el-radio :value="false" label="No" />
        <el-radio :value="true" label="Yes" />
      </el-radio-group>
    </option-item>
    <option-item label="Parent Click" :responsive="false" style="margin-top: 0">
      <el-radio-group v-model="treeProps.expandOnClickNode">
        <el-radio :value="true" label="Expand" />
        <el-radio :value="false" label="Select" />
      </el-radio-group>
    </option-item>
  </ele-card>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import OptionItem from '@/views/extension/avatar/components/option-item.vue';

/** 树选择组件 */
const treeSelectRef = ref(null);

/** 选中值 */
const selectedValue = ref();

/** 树属性 */
const treeProps = reactive({
  height: 266,
  data: [],
  expandOnClickNode: false
});

/** 禁用 */
const disabled = ref(false);

onMounted(() => {
  setTimeout(() => {
    const keys = [];
    let count = 0;
    const data = Array.from({ length: 20 }).map((_, i) => {
      count++;
      keys.push(count);
      return {
        id: count,
        label: `${String(count).padStart(5, '0')}-node-${i + 1}`,
        disabled: false,
        children: Array.from({ length: 20 }).map((_, j) => {
          count++;
          keys.push(count);
          return {
            id: count,
            label: `${String(count).padStart(5, '0')}-node-${i + 1}-${j + 1}`,
            disabled: false,
            children: Array.from({ length: 30 }).map((_, k) => {
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
    data[0].children[0].children[2].disabled = true;
    data[0].children[2].disabled = true;
    data[0].children[2].children.forEach((d) => {
      d.disabled = true;
    });
    treeProps.data = data;
    if (treeSelectRef.value && treeSelectRef.value.treeRef) {
      treeSelectRef.value.treeRef.setExpandedKeys(keys);
    }
  }, 100);
});
</script>
