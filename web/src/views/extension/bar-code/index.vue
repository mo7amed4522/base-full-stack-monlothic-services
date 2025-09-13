<template>
  <ele-page :multi-card="true">
    <ele-card header="Basic Usage">
      <div>
        <ele-bar-code :value="text" :tag="tag" :options="options" />
      </div>
      <el-form label-width="88px" style="max-width: 340px" @submit.prevent="">
        <el-form-item label="Bar Code Type">
          <el-radio-group :model-value="options.format" @update:modelValue="updateFormat">
            <el-radio value="CODE128" label="CODE128" />
            <el-radio value="EAN13" label="EAN13" />
          </el-radio-group>
        </el-form-item>
        <el-form-item label="Rendering Method">
          <el-radio-group v-model="tag">
            <el-radio value="svg" label="svg" />
            <el-radio value="img" label="img" />
            <el-radio value="canvas" label="canvas" />
          </el-radio-group>
        </el-form-item>
        <el-form-item label="Bar Code Text">
          <el-select v-if="options.format === 'EAN13'" v-model="text" class="ele-fluid"
            :popper-options="{ strategy: 'fixed' }">
            <el-option value="1234567890128" label="1234567890128" />
            <el-option value="6971872201359" label="6971872201359" />
            <el-option value="6954531770199" label="6954531770199" />
            <el-option value="6923644240318" label="6923644240318" />
          </el-select>
          <el-input v-else v-model="text" :maxlength="20" />
        </el-form-item>
        <el-form-item label="Bar Code Height">
          <el-slider v-model="options.height" :min="40" :max="160" :step="10" />
        </el-form-item>
        <el-form-item label="Bar Code Width">
          <el-slider v-model="options.width" :min="1" :max="6" />
        </el-form-item>
        <el-form-item label="Bar Code Margin">
          <el-slider v-model="options.margin" :min="0" :max="40" />
        </el-form-item>
        <el-form-item label="Display Text">
          <el-switch v-model="options.displayValue" size="small" />
        </el-form-item>
        <el-form-item v-if="options.displayValue" label="Text Size">
          <el-slider v-model="options.fontSize" :min="12" :max="36" :step="2" />
        </el-form-item>
        <el-form-item v-if="options.displayValue && options.format === 'CODE128'" label="Text Position">
          <el-radio-group v-model="options.textPosition">
            <el-radio value="bottom" label="bottom" />
            <el-radio value="top" label="top" />
          </el-radio-group>
        </el-form-item>
      </el-form>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, reactive, nextTick } from 'vue';

/** Barcode content */
const text = ref('1234567890');

/** Rendering method */
const tag = ref('svg');

/** Parameter configuration */
const options = reactive({
  height: 60,
  width: 2,
  margin: 2,
  displayValue: true,
  textPosition: 'bottom',
  fontSize: 14,
  format: 'CODE128'
});

/** Switch barcode type */
const updateFormat = (value) => {
  if (value === 'EAN13') {
    text.value = '1234567890128';
    nextTick(() => {
      options.format = value;
    });
  } else {
    options.format = value;
  }
};
</script>

<script>
export default {
  name: 'ExtensionBarCode'
};
</script>
