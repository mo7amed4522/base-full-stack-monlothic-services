<template>
  <ele-page :multi-card="true">
    <ele-card header="Basic usage">
      <div style="display: flex; flex-wrap: wrap">
        <div style="margin: 0 10px 20px 10px">
          <div style="margin-bottom: 8px">canvas Rendering</div>
          <ele-qr-code :value="text" :size="size" :level="level" :margin="margin" :image-settings="image" />
        </div>
        <div style="margin: 0 10px 20px 10px">
          <div style="margin-bottom: 8px">svg Rendering</div>
          <ele-qr-code-svg :value="text" :size="size" :level="level" :margin="margin" :image-settings="image" />
        </div>
      </div>
      <el-form label-width="88px" style="max-width: 420px" @submit.prevent="">
        <el-form-item label="QR Code Content">
          <el-input v-model="text" :maxlength="150" />
        </el-form-item>
        <el-form-item label="Error correction level">
          <el-select v-model="level" class="ele-fluid" :popper-options="{ strategy: 'fixed' }">
            <el-option value="L" label="L" />
            <el-option value="M" label="M" />
            <el-option value="Q" label="Q" />
            <el-option value="H" label="H" />
          </el-select>
        </el-form-item>
        <el-form-item label="Size">
          <el-slider v-model="size" :min="80" :max="280" :step="40" />
        </el-form-item>
        <el-form-item label="Margin">
          <el-slider v-model="margin" :min="0" :max="20" />
        </el-form-item>
        <el-form-item label="Custom image">
          <el-switch v-model="showImage" size="small" @change="onShowImageChange" />
        </el-form-item>
        <template v-if="showImage">
          <el-form-item label="Image address">
            <el-input v-model="image.src" :maxlength="400" />
          </el-form-item>
          <el-form-item label="Image size">
            <div style="display: flex; flex-wrap: wrap">
              <div style="width: 120px; margin-right: 16px">
                <el-input-number v-model="image.width" size="small" :min="0" :max="size" />
              </div>
              <div style="width: 120px">
                <el-input-number v-model="image.height" size="small" :min="0" :max="size" />
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Position">
            <div>
              <el-switch v-model="center" size="small" @change="onCenterChange" />
              <div v-if="!center" style="display: flex; flex-wrap: wrap">
                <div style="display: flex; align-items: center">
                  <div style="padding: 0 12px">x</div>
                  <div style="width: 120px">
                    <el-input-number v-model="image.x" size="small" :min="0" :max="size" />
                  </div>
                </div>
                <div style="display: flex; align-items: center">
                  <div style="padding: 0 12px">y</div>
                  <div style="width: 120px">
                    <el-input-number v-model="image.y" size="small" :min="0" :max="size" />
                  </div>
                </div>
              </div>
            </div>
          </el-form-item>
          <el-form-item label="Background excavate">
            <el-switch v-model="image.excavate" size="small" />
          </el-form-item>
        </template>
      </el-form>
    </ele-card>
  </ele-page>
</template>

<script setup>
import { ref, reactive } from 'vue';
const IMAGE_SRC = 'https://cdn.eleadmin.com/20200610/logo-radius.png';

/** QR Code content */
const text = ref('https://eleadmin.com');

/** Error correction level */
const level = ref('L');

/** Size */
const size = ref(120);

/** Margin */
const margin = ref(0);

/** Whether to add a custom icon */
const showImage = ref(true);

/** Custom icon */
const image = reactive({
  src: IMAGE_SRC,
  width: 28,
  height: 28,
  x: void 0,
  y: void 0,
  excavate: false
});

/** Whether the custom icon is centered */
const center = ref(true);

const onShowImageChange = (checked) => {
  if (checked) {
    image.src = IMAGE_SRC;
  } else {
    image.src = void 0;
  }
};

/** Whether the custom icon is centered */
const onCenterChange = (checked) => {
  if (checked) {
    image.x = void 0;
    image.y = void 0;
  } else {
    image.x = 0;
    image.y = 0;
  }
};
</script>

<script>
export default {
  name: 'ExtensionQrCode'
};
</script>
